package main

import (
	"kore_chaincode/core/utils"
	"kore_chaincode/industry"
	"kore_chaincode/korecontract"
	"kore_chaincode/koresecurities"
	"kore_chaincode/person"
	"kore_chaincode/trade"
	"kore_chaincode/user"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddKorecontract adds a new person in world state
func (s *KoreChainCode) AddKorecontract(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return korecontract.AddKorecontract(ctx, []byte(data))
}

// ExecuteKorecontract adds a new person in world state
func (s *KoreChainCode) ExecuteKorecontract(ctx contractapi.TransactionContextInterface, data string) (*korecontract.KorecontractResponse, error) {
	return korecontract.ExecuteKorecontract(ctx, []byte(data))
}

// TestKorecontract adds a new person in world state
func (s *KoreChainCode) TestKorecontract(ctx contractapi.TransactionContextInterface, data string) (*korecontract.KorecontractResponse, error) {
	return korecontract.TestKorecontract(ctx, []byte(data))
}

//---------- Person

// AddPerson adds a new person in world state
func (s *KoreChainCode) AddPerson(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return person.AddPerson(ctx, []byte(data))
}

// ImportPerson adds a new person in world state
func (s *KoreChainCode) ImportPerson(ctx contractapi.TransactionContextInterface, data string) ([]person.PersonWithID, error) {
	return person.ImportPerson(ctx, []byte(data))
}

// GetAllPersons adds a new person in world state
// func (s *KoreChainCode) GetAllPersons(ctx contractapi.TransactionContextInterface) ([]person.PersonDoc, error) {
// 	return person.GetAllPersons(ctx)
// }

// GetPerson fetches a person from world state
func (s *KoreChainCode) GetPerson(ctx contractapi.TransactionContextInterface, data string) (*person.Person, error) {
	return person.GetPerson(ctx, []byte(data))
}

// GetCompany fetches a person from world state
func (s *KoreChainCode) GetCompany(ctx contractapi.TransactionContextInterface, data string) (*user.Company, error) {
	return user.GetCompany(ctx, []byte(data))
}

// GetServiceProviderByID fetches a person from world state
func (s *KoreChainCode) GetServiceProviderByID(ctx contractapi.TransactionContextInterface, data string) (*user.ServiceProvider, error) {
	return user.GetServiceProvider(ctx, []byte(data))
}

// GetATSOperator fetches a person from world state
func (s *KoreChainCode) GetATSOperator(ctx contractapi.TransactionContextInterface, data string) (*user.ATSOperator, error) {
	return user.GetATSOperator(ctx, []byte(data))
}

// GetBrokerDealer fetches a person from world state
func (s *KoreChainCode) GetBrokerDealer(ctx contractapi.TransactionContextInterface, data string) (*user.BrokerDealer, error) {
	return user.GetBrokerDealer(ctx, []byte(data))
}

// GetTransferAgent fetches a person from world state
func (s *KoreChainCode) GetTransferAgent(ctx contractapi.TransactionContextInterface, data string) (*user.TransferAgent, error) {
	return user.GetTransferAgent(ctx, []byte(data))
}

// GetSecurities fetches a person from world state
func (s *KoreChainCode) GetSecurities(ctx contractapi.TransactionContextInterface, data string) (*koresecurities.Securities, error) {
	return koresecurities.GetSecurities(ctx, []byte(data))
}

//---------- Company

// AddCompany adds a new company in world state
func (s *KoreChainCode) AddCompany(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return user.AddCompany(ctx, []byte(data))
}

// ImportCompanies adds a new company in world state
func (s *KoreChainCode) ImportCompanies(ctx contractapi.TransactionContextInterface, data string) ([]user.CompanyWithID, error) {
	return user.ImportCompanies(ctx, []byte(data))
}

// AssociateNotificationURLWithCompany adds a new company in world state
func (s *KoreChainCode) AssociateNotificationURLWithCompany(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
	return user.AssociateNotificationURLWithCompany(ctx, []byte(data))
}

// GetAllCompanies returns all persons found in world state
// func (s *KoreChainCode) GetAllCompanies(ctx contractapi.TransactionContextInterface, data string) ([]user.CompanyDoc, error) {
// 	return user.GetAllCompanies(ctx, []byte(data))
// }

// GetAllCompaniesByRequestorID returns all persons found in world state
func (s *KoreChainCode) GetAllCompaniesByRequestorID(ctx contractapi.TransactionContextInterface, data string) ([]user.CompanyList, error) {
	return user.GetAllCompaniesByRequestorID(ctx, []byte(data))
}

// UpdateCompany updates the company information in world state
func (s *KoreChainCode) UpdateCompany(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return user.UpdateCompany(ctx, []byte(data))
}

// CompanyBankruptcyProceedingStatus updates the Company Bankruptcy Proceeding status in Korechian
// func (s *KoreChainCode) CompanyBankruptcyProceedingStatus(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
// 	return user.CompanyBankruptcyProceedingStatus(ctx, []byte(data))
// }

// // CompanyRegulatoryInjunctionStatus updates the Company Regulatory Injunction status in world state
// func (s *KoreChainCode) CompanyRegulatoryInjunctionStatus(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
// 	return user.CompanyRegulatoryInjunctionStatus(ctx, []byte(data))
// }

// // CompanyHoldByATSOperatorStatus updates the Company Hold By ATS Operator status in Korechian
// func (s *KoreChainCode) CompanyHoldByATSOperatorStatus(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
// 	return user.CompanyHoldByATSOperatorStatus(ctx, []byte(data))
// }

// AssociateATSOperator Associates the ATS Operators with the Company
func (s *KoreChainCode) AssociateATSOperator(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
	return user.AssociateATSOperator(ctx, []byte(data))
}

// AssociateBrokerDealer Associates the Brokerdealer with the Company
func (s *KoreChainCode) AssociateBrokerDealer(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
	return user.AssociateBrokerDealer(ctx, []byte(data))
}

// AssociateTransferAgent Associates the TransferAgent with the Company
func (s *KoreChainCode) AssociateTransferAgent(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
	return user.AssociateTransferAgent(ctx, []byte(data))
}

// AssociateServiceProvider Associates the ATS Operators with the Company
func (s *KoreChainCode) AssociateServiceProvider(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
	return user.AssociateServiceProvider(ctx, []byte(data))
}

// AssignManagementPeople assigns the director to company
// func (s *KoreChainCode) AssignManagementPeople(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
// 	return user.AssignManagementPeople(ctx, []byte(data))
// }

// // RemoveManagementPeople removes the director to company
// func (s *KoreChainCode) RemoveManagementPeople(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
// 	return user.RemoveManagementPeople(ctx, []byte(data))
// }

// // GetManagementPeople fetch management people
// func (s *KoreChainCode) GetManagementPeople(ctx contractapi.TransactionContextInterface, data string) (*user.Management, error) {
// 	return user.GetManagementPeople(ctx, []byte(data))
// }

//---------- ATS Operator

// AddATSOperator saves the ATS Operator in world state
func (s *KoreChainCode) AddATSOperator(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return user.AddATSOperator(ctx, []byte(data))
}

//---------- Broker Dealer

// AddBrokerDealer saves the Broker Dealer in world state
func (s *KoreChainCode) AddBrokerDealer(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return user.AddBrokerDealer(ctx, []byte(data))
}

//---------- Transfer Agent

// AddTransferAgent saves the transfer agent in world state
func (s *KoreChainCode) AddTransferAgent(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return user.AddTransferAgent(ctx, []byte(data))
}

//---------- Service Provider

// AddServiceProvider saves the ServiceProvider information into the world state
func (s *KoreChainCode) AddServiceProvider(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return user.AddServiceProvider(ctx, []byte(data))
}

// GetAllServiceProviders returns all service providers found in world state
// func (s *KoreChainCode) GetAllServiceProviders(ctx contractapi.TransactionContextInterface, data string) ([]user.ServiceProviderDoc, error) {
// 	return user.GetAllServiceProviders(ctx, []byte(data))
// }

//---------- Industry

// AddIndustry saves the Industry information in world state
func (s *KoreChainCode) AddIndustry(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return industry.AddIndustry(ctx, []byte(data))
}

// GetAllIndustries returns all industries found in world state
func (s *KoreChainCode) GetAllIndustries(ctx contractapi.TransactionContextInterface, data string) ([]industry.IndustryDoc, error) {
	return industry.GetAllIndustries(ctx, []byte(data))
}

//---------- Document

// AddDocument saves the document hash in world state
func (s *KoreChainCode) AddDocument(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return korecontract.AddDocument(ctx, []byte(data))
}

// GetAllDocuments returns all documents found in world state
// func (s *KoreChainCode) GetAllDocuments(ctx contractapi.TransactionContextInterface, data string) ([]korecontract.DocumentDoc, error) {
// 	return korecontract.GetAllDocuments(ctx, []byte(data))
// }

//---------- Offering Memorandum

// AddOfferingMemorandum saves the Offering Memorandum in world state
func (s *KoreChainCode) AddOfferingMemorandum(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return korecontract.AddOfferingMemorandum(ctx, []byte(data))
}

// SaveKoreContract saves the Offering Memorandum in world state
func (s *KoreChainCode) SaveKoreContract(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return korecontract.SaveKoreContract(ctx, []byte(data))
}

// GetAllOfferingMemorandums returns all offeringMemorandums found in world state
// func (s *KoreChainCode) GetAllOfferingMemorandums(ctx contractapi.TransactionContextInterface, data string) ([]korecontract.OfferingMemorandumDoc, error) {
// 	return korecontract.GetAllOfferingMemorandums(ctx, []byte(data))
// }

//---------- Shareholder Agreement

// AddShareHolderAgreement saves the shareHolderAgreement hash in world state
func (s *KoreChainCode) AddShareHolderAgreement(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return korecontract.AddShareHolderAgreement(ctx, []byte(data))
}

// GetAllShareHolderAgreements returns all shareHolderAgreements found in world state
// func (s *KoreChainCode) GetAllShareHolderAgreements(ctx contractapi.TransactionContextInterface, data string) ([]korecontract.ShareHolderAgreementDoc, error) {
// 	return korecontract.GetAllShareHolderAgreements(ctx, []byte(data))
// }

// AddSubscriptionAgreement saves the  hash in world state
func (s *KoreChainCode) AddSubscriptionAgreement(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return korecontract.AddSubscriptionAgreement(ctx, []byte(data))
}

//---------- Payment Method

// AddPaymentMethod saves the document hash in world state
func (s *KoreChainCode) AddPaymentMethod(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return korecontract.AddPaymentMethod(ctx, []byte(data))
}

//---------- Certificate text

// AddSecuritiesCertificateText saves the securitiesCertificateText hash in world state
func (s *KoreChainCode) AddSecuritiesCertificateText(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return koresecurities.AddSecuritiesCertificateText(ctx, []byte(data))
}

// GetAllSecuritiesCertificateTexts returns all securitiesCertificateTexts found in world state
// func (s *KoreChainCode) GetAllSecuritiesCertificateTexts(ctx contractapi.TransactionContextInterface, data string) ([]koresecurities.SecuritiesCertificateTextDoc, error) {
// 	return koresecurities.GetAllSecuritiesCertificateTexts(ctx, []byte(data))
// }

//---------- Certificate

// AddCertificate saves the certificate hash in world state
func (s *KoreChainCode) AddCertificate(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return koresecurities.AddCertificate(ctx, []byte(data))
}

// GetAllCertificates returns all certificates found in world state
// func (s *KoreChainCode) GetAllCertificates(ctx contractapi.TransactionContextInterface, data string) ([]koresecurities.CertificateDoc, error) {
// 	return koresecurities.GetAllCertificates(ctx, []byte(data))
// }

// UpdateCertificate updates the Securities certificate in the world state
func (s *KoreChainCode) UpdateCertificate(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return koresecurities.UpdateCertificate(ctx, []byte(data))
}

//---------- SecuritiesExchangePrice

// AddSecuritiesExchangePrice saves the certificate hash in world state
func (s *KoreChainCode) AddSecuritiesExchangePrice(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return koresecurities.AddSecuritiesExchangePrice(ctx, []byte(data))
}

//---------- Holding

// // GetInvestorHoldings returns all holdings found in world state
// func (s *KoreChainCode) GetInvestorHoldings(ctx contractapi.TransactionContextInterface, data string) ([]*koresecurities.Holding, error) {
// 	return koresecurities.GetInvestorHoldings(ctx, []byte(data))
// }

// InvestorHoldingExists checks if the investor has a particular holding or not
func (s *KoreChainCode) InvestorHoldingExists(ctx contractapi.TransactionContextInterface, data string) (*koresecurities.ResponseInvestorHoldingExists, error) {
	return koresecurities.InvestorHoldingExists(ctx, []byte(data))
}

// GetAvailableShares checks if the investor has a particular holding or not
func (s *KoreChainCode) GetAvailableShares(ctx contractapi.TransactionContextInterface, data string) (*koresecurities.ResponseInvestorHoldingExists, error) {
	return koresecurities.GetAvailableShares(ctx, []byte(data))
}

// PlaceHoldOnShares checks if the investor has a particular holding or not
func (s *KoreChainCode) PlaceHoldOnShares(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return koresecurities.PlaceHoldOnShares(ctx, []byte(data))
}

// ReleaseHoldOnShares checks if the investor has a particular holding or not
func (s *KoreChainCode) ReleaseHoldOnShares(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return koresecurities.ReleaseHoldOnShares(ctx, []byte(data))
}

// UpdateHolding checks if the investor has a particular holding or not
func (s *KoreChainCode) UpdateHolding(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return koresecurities.UpdateHolding(ctx, []byte(data))
}

// AddHolding saves the holding hash in world state
func (s *KoreChainCode) AddHolding(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return koresecurities.AddHolding(ctx, []byte(data))
}

// GetNumberOfSharesInHolding returns all holdings found in world state
func (s *KoreChainCode) GetNumberOfSharesInHolding(ctx contractapi.TransactionContextInterface, data string) ([]koresecurities.HoldingDoc, error) {
	return koresecurities.GetNumberOfSharesInHolding(ctx, []byte(data))
}

// GetAllHoldingsbyAllSecurities returns all holdings found in world state
func (s *KoreChainCode) GetAllHoldingsbyAllSecurities(ctx contractapi.TransactionContextInterface, data string) ([]koresecurities.AllHoldingAllSecuritiesResponse, error) {
	return koresecurities.GetAllHoldingsbyAllSecurities(ctx, []byte(data))
}

// GetAllHoldingsbySecuritiesID returns all holdings found in world state
func (s *KoreChainCode) GetAllHoldingsbySecuritiesID(ctx contractapi.TransactionContextInterface, data string) ([]koresecurities.AllHoldingBySecuritiesResponse, error) {
	return koresecurities.GetAllHoldingsbySecuritiesID(ctx, []byte(data))
}

// GetAllShareHoldersByComapny returns all holdings found in world state
func (s *KoreChainCode) GetAllShareHoldersByComapny(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseIDArray, error) {
	return koresecurities.GetAllShareHoldersByComapny(ctx, []byte(data))
}

// GetAllTradableHoldings returns all holdings found in world state
func (s *KoreChainCode) GetAllTradableHoldings(ctx contractapi.TransactionContextInterface, data string) ([]koresecurities.AllTradableHoldingsResponse, error) {
	return koresecurities.GetAllTradableHoldings(ctx, []byte(data))
}

// GetAllShareHolders returns all holdings found in world state
func (s *KoreChainCode) GetAllShareHolders(ctx contractapi.TransactionContextInterface, data string) (*koresecurities.ShareHoldersIDResponse, error) {
	return koresecurities.GetAllShareHolders(ctx, []byte(data))
}

// GetAllHoldingsByCompany returns all holdings found in world state for company
func (s *KoreChainCode) GetAllHoldingsByCompany(ctx contractapi.TransactionContextInterface, data string) ([]koresecurities.HoldingByCompanyResponse, error) {
	return koresecurities.GetAllHoldingsByCompany(ctx, []byte(data))
}

// UpdateSecurities checks if the investor has a particular holding or not
func (s *KoreChainCode) UpdateSecurities(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
	return koresecurities.UpdateSecurities(ctx, []byte(data))
}

// AssociateATSWithSecurity checks if the investor has a particular holding or not
func (s *KoreChainCode) AssociateATSWithSecurity(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
	return koresecurities.AssociateATSWithSecurity(ctx, []byte(data))
}

// AssociateBrokerWithSecurity checks if the investor has a particular holding or not
func (s *KoreChainCode) AssociateBrokerWithSecurity(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
	return koresecurities.AssociateBrokerWithSecurity(ctx, []byte(data))
}

//---------- SecuritiesInstrument

// AddSecuritiesInstrument saves the document hash in world state
func (s *KoreChainCode) AddSecuritiesInstrument(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return koresecurities.AddSecuritiesInstrument(ctx, []byte(data))
}

//--------- Securities

// IssueSecurities saves the Securities information into the world state
func (s *KoreChainCode) IssueSecurities(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return koresecurities.IssueSecurities(ctx, []byte(data))
}

// GetAllSecurities returns all securities found in world state
func (s *KoreChainCode) GetAllSecurities(ctx contractapi.TransactionContextInterface, data string) ([]koresecurities.SecuritiesDoc, error) {
	return koresecurities.GetAllSecurities(ctx, []byte(data))
}

// TransferSecurities saves the document hash in world state
func (s *KoreChainCode) TransferSecurities(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return koresecurities.TransferSecurities(ctx, []byte(data))
}

// UpdateTransferSecurities saves the document hash in world state
func (s *KoreChainCode) UpdateTransferSecurities(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseMessage, error) {
	return koresecurities.UpdateTransferSecurities(ctx, []byte(data))
}

// Trade

// AddTradeRequest adds a new person in world state
func (s *KoreChainCode) AddTradeRequest(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return trade.AddTradeRequest(ctx, []byte(data))
}

// AddATSTrade adds a new person in world state
func (s *KoreChainCode) AddATSTrade(ctx contractapi.TransactionContextInterface, data string) (*utils.ResponseID, error) {
	return trade.AddATSTrade(ctx, []byte(data))
}

// GetAllTradeRequests adds a new person in world state
func (s *KoreChainCode) GetAllTradeRequests(ctx contractapi.TransactionContextInterface, data string) ([]trade.TradeRequestDoc, error) {
	return trade.GetAllTradeRequests(ctx, []byte(data))
}
