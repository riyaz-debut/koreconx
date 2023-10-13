package korecontract

import (
	"encoding/json"
	"fmt"
	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"
	"kore_chaincode/user"
	"strings"
	"time"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ShareHolderAgreement struct
type ShareHolderAgreementForm struct {
	Title        string         `json:"title"`
	PreambleText string         `json:"preamble_text"`
	Clauses      map[int]Clause `json:"clauses"`
	utils.MetaData
}

// Clause struct
type Clause struct {
	Title      string            `json:"title"`
	Text       string            `json:"text"`
	References []string          `json:"references"`
	Data       map[string]string `json:"data"`
	Clauses    map[int]Clause    `json:"clauses"`
}

type Korecontract struct {
	ID    string           `json:"id"`
	Meta  MetaKorecontract `json:"meta"`
	Rules []string         `json:"rules"`
	utils.MetaData
}

type MetaKorecontract struct {
	Version      string    `json:"version"`
	Author       string    `json:"author"`
	Company      string    `json:"company"`
	DateCreated  time.Time `json:"date_created"`
	DateUntil    time.Time `json:"date_until"`
	Description  string    `json:"description"`
	ContractText string    `json:"contract_text"`
	Category     string    `json:"category"`
	Subcategory  string    `json:"subcategory"`
	Keywords     []struct {
		Keyword string `json:"keyword"`
	} `json:"keywords"`
}

type ExecuteKorecontractRequest struct {
	Version         string    `json:"version"`
	ID              string    `json:"id"`
	Company         string    `json:"company"`
	TransactionID   string    `json:"transaction_id"`
	Variables       string    `json:"variables"`
	ReturnVariables []string  `json:"return_variables"`
	CreatedAt       time.Time `json:"created_at"`
}

type TestKorecontractRequest struct {
	Rules           []string `json:"rules"`
	Variables       string   `json:"variables"`
	ReturnVariables []string `json:"return_variables"`
}

type KorecontractResponse struct {
	MyJSON string `json:"myJSON"`
}

// SaveKoreContract Data
func SaveKoreContract(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(ShareHolderAgreementForm)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeKoreContract
	request.UpdatedAt = request.CreatedAt

	// change strcut to json
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(ctx.GetStub().GetTxID(), jsonData)

	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseID)
	response.ID = ctx.GetStub().GetTxID()
	return response, nil
}

// AddKorecontract adds a new person in world state
func AddKorecontract(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(Korecontract)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// company
	_, err = user.GetCompanyByID(ctx, request.Meta.Company)
	if err != nil {
		return nil, err
	}

	// korecontract
	contract, _ := GetKoreContractByID(ctx, request.ID)
	if contract.ID == request.ID {
		return nil, status.ErrBadRequest.WithMessage("Korecontract Already exists!")
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeKoreContractRules
	request.UpdatedAt = request.CreatedAt

	// change strcut to json
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(ctx.GetStub().GetTxID(), jsonData)

	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseID)
	response.ID = ctx.GetStub().GetTxID()
	return response, nil
}

// ExecuteKorecontract adds a new person in world state
func ExecuteKorecontract(ctx contractapi.TransactionContextInterface, data []byte) (*KorecontractResponse, error) {
	request := new(ExecuteKorecontractRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\",\"meta.company\":\"%s\",\"meta.version\":\"%s\",\"id\":\"%s\" }}", utils.DocTypeKoreContractRules, request.TransactionID, request.Company, request.Version, request.ID)
	korecontractBA, _, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("Korecontract with ID: %s does not exists!", request.TransactionID))
	if err != nil {
		return nil, err
	}

	// ummarshal the byte array to structure
	korecontract := new(Korecontract)
	err = json.Unmarshal(korecontractBA, korecontract)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// Check the date here
	if !korecontract.Meta.DateUntil.IsZero() && !korecontract.Meta.DateUntil.After(request.CreatedAt) {
		return nil, status.ErrBadRequest.WithMessage("Korecontract can not be executed as the date untill is of past")
	}

	myJSON := []byte(request.Variables)

	dataContext := ast.NewDataContext()
	err = dataContext.AddJSON("KCH", myJSON)
	if err != nil {
		return nil, status.ErrBadRequest.WithMessage("Error in adding JSON to Grule.")
	}

	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)
	bs := pkg.NewBytesResource([]byte(strings.Join(korecontract.Rules, " ")))
	err = ruleBuilder.BuildRuleFromResource(korecontract.ID, korecontract.Meta.Version, bs)
	if err != nil {
		return nil, status.ErrBadRequest.WithMessage("Error in building the Rule Library.")
	}

	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance(korecontract.ID, korecontract.Meta.Version)
	engine := engine.NewGruleEngine()
	err = engine.Execute(dataContext, knowledgeBase)
	if err != nil {
		return nil, status.ErrBadRequest.WithMessage("Error in executing the rules against library.")
	}

	response := new(KorecontractResponse)
	iter := dataContext.Get("KCH").Value().MapRange()
	jsonString := "{"
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()

		key := fmt.Sprintf("%s", k)
		if stringInSlice(key, request.ReturnVariables) {
			jsonString += fmt.Sprintf("\"%s\":\"%v\",", k, v)
		}
	}

	jsonString = strings.TrimRight(jsonString, ",")
	jsonString += "}"

	fmt.Println("JSON Stirng ", jsonString)

	response.MyJSON = jsonString
	return response, nil
}

// ExecuteKorecontract adds a new person in world state
func TestKorecontract(ctx contractapi.TransactionContextInterface, data []byte) (*KorecontractResponse, error) {
	request := new(TestKorecontractRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	myJSON := []byte(request.Variables)

	dataContext := ast.NewDataContext()
	err = dataContext.AddJSON("KCH", myJSON)
	if err != nil {
		return nil, status.ErrBadRequest.WithMessage("Error in adding JSON to Grule.")
	}

	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)
	bs := pkg.NewBytesResource([]byte(strings.Join(request.Rules, " ")))
	err = ruleBuilder.BuildRuleFromResource("Test", "1.0", bs)
	if err != nil {
		return nil, status.ErrBadRequest.WithMessage("Error in building the Rule Library.")
	}

	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("Test", "1.0")
	engine := engine.NewGruleEngine()
	err = engine.Execute(dataContext, knowledgeBase)
	if err != nil {
		return nil, status.ErrBadRequest.WithMessage("Error in executing the rules against library.")
	}

	response := new(KorecontractResponse)
	iter := dataContext.Get("KCH").Value().MapRange()
	jsonString := "{"
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()

		key := fmt.Sprintf("%s", k)
		if stringInSlice(key, request.ReturnVariables) {
			jsonString += fmt.Sprintf("\"%s\":\"%v\",", k, v)
		}
	}

	jsonString = strings.TrimRight(jsonString, ",")
	jsonString += "}"

	fmt.Println("JSON Stirng ", jsonString)

	response.MyJSON = jsonString
	return response, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// GetKorecontractByID returns the koreContract with given ID from world state
func GetKoreContractByID(ctx contractapi.TransactionContextInterface, ID string) (Korecontract, error) {
	koreContractData := Korecontract{}

	koreContractBA, err := KorecontractExists(ctx, ID)
	if err != nil {
		return koreContractData, err
	}

	// ummarshal the byte array to structure
	err = json.Unmarshal(koreContractBA, &koreContractData)
	if err != nil {
		return koreContractData, status.ErrInternal.WithError(err)
	}

	return koreContractData, nil
}

// KorecontractExists checks whether the koreContract with given ID exists or not in world state
func KorecontractExists(ctx contractapi.TransactionContextInterface, ID string) ([]byte, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypeKoreContractRules, ID)
	koreContractBA, _, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("Korecontract with ID: %s does not exists!", ID))

	return koreContractBA, err
}
