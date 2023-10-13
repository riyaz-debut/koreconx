package user

import (
	"time"

	"kore_chaincode/core/utils"
)

// Company Data fields
type Company struct {
	Cd               string             `json:"cd"`
	CompanyID        []RegistrationData `json:"company_id"`
	Source           SourceData         `json:"source"`
	Verifications    []Verification     `json:"verifications"`
	OtherReferences  []OtherReference   `json:"other_references"`
	ATSOperators     []string           `json:"ats_operators"`
	BrokerDealers    []string           `json:"broker_dealers"`
	ServiceProviders []string           `json:"service_providers"`
	TransferAgentID  string             `json:"transfer_agent_id"`
	Status           int32              `json:"status"`
	// Management           *Management           `json:"management"`
	// BankruptcyProceeding *BankruptcyProceeding `json:"bankruptcy_proceeding"`
	// RegulatoryInjunction *BankruptcyProceeding `json:"regulatory_injunction"`
	// HoldByATSOperator    *BankruptcyProceeding `json:"hold_by_ats_operator"`
	utils.MetaData
}

// RegistrationData data fields
type RegistrationData struct {
	RegistrationID        string `json:"registration_id"`
	RegistrationDomicile  string `json:"registration_domicile"`
	RegistrationAuthority string `json:"registration_authority"`
	RegistrationRecordURL string `json:"registration_record_url"`
}

// SourceData data fields
type SourceData struct {
	SourceSystemID   string `json:"source_system_id"`
	SourcePlatformID string `json:"source_platform_id"`
}

// Verification data fields
type Verification struct {
	VerificationType string    `json:"verification_type"`
	VerifyingOrg     string    `json:"verifying_org"`
	VerificationID   string    `json:"verification_id"`
	VerificationDate time.Time `json:"verification_date"`
	VerificationURL  string    `json:"verification_url"`
}

// OtherReference data fields
type OtherReference struct {
	OtherReferenceID string `json:"other_reference_id"`
	OtherPlatformID  string `json:"other_platform_id"`
}

// Company data fields
// type  struct {
// 	CompanyLegalName      string                `json:"company_legal_name"`
// 	LegalExtension        string                `json:"legal_extension"`
// 	DateOfIncorporation   time.Time             `json:"date_of_incorporation"`
// 	AddressLine1          string                `json:"address_line_1"`
// 	AddressLine2          string                `json:"address_line_2"`
// 	Website               string                `json:"website"`
// 	City                  string                `json:"city"`
// 	Region                string                `json:"region"`
// 	Postal                string                `json:"postal"`
// 	Country               string                `json:"country"`
// 	CikNumber             string                `json:"cik_number"`
// 	IndustryID            string                `json:"industry_id"`
// 	SubIndustryID         string                `json:"sub_industry_id"`
// 	RegistrationDomicile  string                `json:"registration_domicile"`
// 	RegistrationRecordURL string                `json:"registration_record_url"`
// 	BankruptcyProceeding  *BankruptcyProceeding `json:"bankruptcy_proceeding"`
// 	RegulatoryInjunction  *BankruptcyProceeding `json:"regulatory_injunction"`
// 	HoldByATSOperator     *BankruptcyProceeding `json:"hold_by_ats_operator"`
// 	SourceSystemID        string                `json:"source_system_id"`
// 	SourcePlatform        string                `json:"source_platform"`
// 	SourcePlatformID      string                `json:"source_platform_id"`
// 	ATSOperators          []string              `json:"ats_operators"`
// 	BrokerDealers         []string              `json:"broker_dealers"`
// 	ServiceProviders      []string              `json:"service_providers"`
// 	TransferAgentID       string                `json:"transfer_agent_id"`
// 	NotificationURLs      []string              `json:"notification_urls"`
// 	Management            *Management           `json:"management"`
// NotificationURLs      []string              `json:"notification_urls"`
// 	utils.MetaData
// 	// TransferAgents        []string              `json:"transfer_agents"`
// }

// Management data fields
type Management struct {
	CEO      map[string]ManagementPeople `json:"ceo"`
	CFO      map[string]ManagementPeople `json:"cfo"`
	Director map[string]ManagementPeople `json:"director"`
	Officer  map[string]ManagementPeople `json:"officer"`
}

// ManagementPeople data fields
type ManagementPeople struct {
	PersonID                     string    `json:"person_id"`
	StartDate                    time.Time `json:"start_date"`
	EndDate                      time.Time `json:"end_date"`
	Status                       string    `json:"status"`
	ShareholderVoteTransactionID string    `json:"shareholder_vote_transaction_id"`
}

// BankruptcyProceeding data fields
type BankruptcyProceeding struct {
	Status           bool                        `json:"status"`
	ReferenceDetails map[string]ReferenceDetails `json:"reference_details"`
}

// CompanyStatusRequest data fields
type CompanyStatusRequest struct {
	CompanyID        string             `json:"company_id"`
	Status           bool               `json:"status"`
	ReferenceDetails []ReferenceDetails `json:"reference_details"`
	CreatedAt        time.Time          `json:"created_at"`
}

// CompanyManagementRequest data fields
type CompanyManagementRequest struct {
	PersonID                     string    `json:"person_id"`
	ShareholderVoteTransactionID string    `json:"shareholder_vote_transaction_id"`
	Role                         string    `json:"role"`
	CompanyID                    string    `json:"company_id"`
	StartDate                    time.Time `json:"start_date"`
	CreatedAt                    time.Time `json:"created_at"`
}

// CompanyManagementRemoveRequest data fields
type CompanyManagementRemoveRequest struct {
	PersonID  string    `json:"person_id"`
	Role      string    `json:"role"`
	CompanyID string    `json:"company_id"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
}

// ReferenceDetails data fields
type ReferenceDetails struct {
	AuthorizationReference string    `json:"authorization_reference"`
	EffectiveStartDate     time.Time `json:"effective_start_date"`
	EffectiveEndDate       time.Time `json:"effective_end_date"`
	AuthorizationDate      time.Time `json:"authorization_date"`
	AuthorizingEntityID    string    `json:"authorizing_entity_id"`
}

// AssociateIDsRequest data fields
type AssociateIDsRequest struct {
	CompanyID       string    `json:"company_id"`
	Data            []string  `json:"data"`
	AssociationDate time.Time `json:"association_date"`
}

// CompanyFilter data fields
type CompanyFilter struct {
	BrokerDealerID string `json:"broker_dealer_id"`
	// Country          string `json:"country"`
	// IndustryID       string `json:"industry_id"`
	// SubIndustryID    string `json:"sub_industry_id"`
	// CompanyLegalName string `json:"company_legal_name"`
}

// UpdateCompanyRequest data fields
type UpdateCompanyRequest struct {
	ID   string  `json:"id"`
	Data Company `json:"data"`
}

// GetManagementPeopleRequest data fields
type GetManagementPeopleRequest struct {
	CompanyID string `json:"company_id"`
}

// ATSOperator data fields
type ATSOperator struct {
	ATSOperatorID          string     `json:"ats_operator_id"`
	CorporateName          string     `json:"corporate_name"`
	RegistrationNumber     string     `json:"registration_number"`
	RegistrationAuthority  string     `json:"registration_authority"`
	TypeOfLicense          string     `json:"type_of_license"`
	LicenseGrantType       string     `json:"license_grant_type"`
	RegistrationDate       time.Time  `json:"registration_date"`
	RegistrationExpiryDate time.Time  `json:"registration_expiry_date"`
	Domicile               []Domicile `json:"domicile"`
	SourceSystemID         string     `json:"source_system_id"`
	Status                 int32      `json:"status"`
	utils.MetaData
}

// BrokerDealer data fields
type BrokerDealer struct {
	BrokerDealerID        string     `json:"broker_dealer_id"`
	CorporateName         string     `json:"corporate_name"`
	RegistrationNumber    string     `json:"registration_number"`
	RegistrationAuthority string     `json:"registration_authority"`
	Domicile              []Domicile `json:"domicile"`
	SourceSystemID        string     `json:"source_system_id"`
	Status                int32      `json:"status"`
	utils.MetaData
}

// Domicile data fields
type Domicile struct {
	Country string   `json:"country"`
	State   []string `json:"state"`
}

// TransferAgent data fields
type TransferAgent struct {
	TransferAgentID       string     `json:"transfer_agent_id"`
	CorporateName         string     `json:"corporate_name"`
	RegistrationNumber    string     `json:"registration_number"`
	RegistrationAuthority string     `json:"registration_authority"`
	Domicile              []Domicile `json:"domicile"`
	SourceSystemID        string     `json:"source_system_id"`
	Status                int32      `json:"status"`
	utils.MetaData
}

// ServiceProvider data fields
type ServiceProvider struct {
	LegalName             string `json:"legal_name"`
	ServiceName           string `json:"service_name"`
	RegistrationAuthority string `json:"registration_authority"`
	RegistrationNumber    string `json:"registration_number"`
	Country               string `json:"country"`
	SourceSystemID        string `json:"source_system_id"`
	Status                int32  `json:"status"`
	utils.MetaData
}

// ImportCompaniesRequest data fields
type ImportCompaniesRequest struct {
	Data []CompanyWithID `json:"data"`
}

// CompanyWithID data fields
type CompanyWithID struct {
	Data Company `json:"data"`
	ID   string  `json:"id"`
}
