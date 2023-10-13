package korecontract

import (
	"time"

	"kore_chaincode/core/utils"
)

// Reference data fields
type Reference struct {
	ReferenceID   string `json:"reference_id"`
	ReferenceName string `json:"reference_name"`
}

// Jurisdiction data fields
type Jurisdiction struct {
	Country string `json:"country"`
	State   string `json:"state"`
}

// Document data fields
type Document struct {
	EntityID         string `json:"entity_id"`
	DocumentSourceID string `json:"document_source_id"`
	DocumentHash     string `json:"document_hash"`
	HashAlgorithm    string `json:"hash_algorithm"`
	Status           int32  `json:"status"`
	utils.MetaData
}

// DocumentFilter data fields
type DocumentFilter struct {
	EntityID string `json:"entity_id"`
}

// CompanyFilter data fields
type CompanyFilter struct {
	CompanyID string `json:"company_id"`
}

// OfferingMemorandum data fields
type OfferingMemorandum struct {
	CompanyID              string         `json:"company_id"`
	BrokerDealers          []string       `json:"broker_dealers"`
	CompanySymbol          string         `json:"company_symbol"`
	DocumentID             string         `json:"document_id"`
	Exemptions             []string       `json:"exemptions"`
	IssuanceJurisdictions  []Jurisdiction `json:"issuance_jurisdictions"`
	OfferingDetail         OfferingDetail `json:"offering_detail"`
	PublicReportingCompany bool           `json:"public_reporting_company"`
	Status                 int32          `json:"status"`
	utils.MetaData
}

// OfferingDetail data fields
type OfferingDetail struct {
	OfferingAmount           float64     `json:"offering_amount"`
	OfferingCurrency         string      `json:"offering_currency"`
	OfferingEndDate          time.Time   `json:"offering_end_date"`
	OfferingPricePerSecurity float64     `json:"offering_price_per_security"`
	OfferingSecuritiesClass  string      `json:"offering_securities_class"`
	OfferingSecuritiesNumber float64     `json:"offering_securities_number"`
	OfferingStartDate        time.Time   `json:"offering_start_date"`
	OfferingType             string      `json:"offering_type"`
	References               []Reference `json:"references"`
}

// BrokerDealersDetail data fields
type BrokerDealersDetail struct {
	BrokerDealerID           string       `json:"broker_dealer_id"`
	BrokerDealerJurisdiction Jurisdiction `json:"broker_dealer_jurisdiction"`
}

// OfferingMemorandumFilter data fields
type OfferingMemorandumFilter struct {
	CompanyID string `json:"company_id"`
}

// PaymentMethod data fields
type PaymentMethod struct {
	PaymentType     string    `json:"payment_type"`
	TransactionDate time.Time `json:"transaction_date"`
	IPAddress       string    `json:"ip_address"`
	PayerID         string    `json:"payer_id"`
	Amount          float64   `json:"amount"`
	Currency        string    `json:"currency"`
	// CreditCard      CreditCard     `json:"credit_card"`
	// ACH             ACH            `json:"ach"`
	// Cryptocurrency  Cryptocurrency `json:"cryptocurrency"`
	// Wire            Wire           `json:"wire"`
	// IRA             IRA            `json:"ira"`
	PaymentMethod string `json:"payment_method"`
	utils.MetaData
}

// CreditCard data fields
type CreditCard struct {
	CardNumber string `json:"card_number"`
	NameOnCard string `json:"name_on_card"`
	ExpiryDate string `json:"expiry_date"`
}

// ACH data fields
type ACH struct {
	BankName      string `json:"expiry_date"`
	NameOnAccount string `json:"name_on_account"`
	TransitNumber string `json:"transit_number"`
	AccountNumber string `json:"account_number"`
}

// Cryptocurrency data fields
type Cryptocurrency struct {
	CryptoType              string `json:"crypto_type"`
	Amount                  string `json:"amount"`
	USDEquivalent           string `json:"usd_equivalent"`
	CryptoExchangeCustodian string `json:"crypto_exchange_custodian"`
}

// Wire data fields
type Wire struct {
	BankName      string `json:"expiry_date"`
	NameOnAccount string `json:"name_on_account"`
	SwiftCode     string `json:"swift_code"`
	AccountNumber string `json:"account_number"`
}

// IRA data fields
type IRA struct {
	TrustName     string `json:"trust_name"`
	IRAType       string `json:"ira_type"`
	NameOnAccount string `json:"name_on_account"`
	AccountNumber string `json:"account_number"`
	TrustID       string `json:"trust_id"`
}

// ShareHolderAgreement data fields
type ShareHolderAgreement struct {
	Beneficiaries          []Beneficiary          `json:"beneficiaries"`
	CompanyID              string                 `json:"company_id"`
	DocumentID             string                 `json:"document_id"`
	Disclosures            []Disclosure           `json:"disclosures"`
	Exits                  Exits                  `json:"exits"`
	FinancialParticipation FinancialParticipation `json:"financial_participation"`
	Notifications          []Notification         `json:"notifications"`
	Reports                []Report               `json:"reports"`
	Rights                 Rights                 `json:"rights"`
	Trading                Trading                `json:"trading"`
	Transfer               Transfer               `json:"transfer"`
	Voting                 Voting                 `json:"voting"`
	Status                 int32                  `json:"status"`
	utils.MetaData
}

// Beneficiary data fields
type Beneficiary struct {
	BeneficiaryID   string `json:"beneficiary_id"`
	BeneficiaryType int    `json:"beneficiary_type"`
}

// Disclosure data fields
type Disclosure struct {
	DisclosureFrequency     int                   `json:"disclosure_frequency"`
	DisclosureFrequencyUnit int                   `json:"disclosure_frequency_unit"`
	DisclosureID            string                `json:"disclosure_id"`
	DisclosureRecipients    []DisclosureRecipient `json:"disclosure_recipients"`
	DisclosureTitle         string                `json:"disclosure_title"`
}

// DisclosureRecipient data fields
type DisclosureRecipient struct {
	DisclosureRecipientID   string `json:"disclosure_recipient_id"`
	DisclosureRecipientName string `json:"disclosure_recipient_name"`
	DisclosureRecipientType string `json:"disclosure_recipient_type"`
}

// SubExists data fields
type SubExists struct {
	Exists     bool                   `json:"exists"`
	References []ShareHolderReference `json:"references"`
}

// ShareHolderReference data fields
type ShareHolderReference struct {
	Reference
	ReferenceClause string `json:"reference_clause"`
}

// Exits data fields
type Exits struct {
	IPO        SubExists `json:"IPO"`
	MA         SubExists `json:"M&A"`
	RTO        SubExists `json:"RTO"`
	Bankruptcy SubExists `json:"bankruptcy"`
}

// SubFinancialParticipation data fields
type SubFinancialParticipation struct {
	Description string                 `json:"description"`
	Exists      bool                   `json:"exists"`
	References  []ShareHolderReference `json:"references"`
}

// FinancialParticipation data fields
type FinancialParticipation struct {
	Dividends    SubFinancialParticipation `json:"dividends"`
	Preferreds   SubFinancialParticipation `json:"preferreds"`
	RevenueShare SubFinancialParticipation `json:"revenue_share"`
	Warrants     SubExists                 `json:"warrants"`
}

// Notification data fields
type Notification struct {
	NotificationMandatory bool                   `json:"notification_mandatory"`
	NotificationTitle     string                 `json:"notification_title"`
	References            []ShareHolderReference `json:"references"`
	Trigger               string                 `json:"trigger"`
}

// Report data fields
type Report struct {
	ReportConsumers []ReportConsumer `json:"report_consumers"`
	ReportFrequency string           `json:"report_frequency"`
	ReportID        string           `json:"report_id"`
	ReportTitle     string           `json:"report_title"`
}

// ReportConsumer data fields
type ReportConsumer struct {
	ReportConsumerID   string `json:"report_consumer_id"`
	ReportConsumerName string `json:"report_consumer_name"`
	ReportConsumerType string `json:"report_consumer_type"`
}

// Exists data fields
type Exists struct {
	Exists bool `json:"exists"`
}

// Rights data fields
type Rights struct {
	OfferOfWarrants     Exists `json:"offer_of_warrants"`
	RightOfDragAlong    Exists `json:"right_of_drag_along"`
	RightOfFirstRefusal Exists `json:"right_of_first_refusal"`
	RightOfPreemption   Exists `json:"right_of_preemption"`
	RightOfTagAlong     Exists `json:"right_of_tag_along"`
}

// Trading data fields
type Trading struct {
	Ats                 []TradingATS       `json:"ats"`
	TradingRestrictions TradingRestriction `json:"trading_restrictions"`
}

// TradingATS data fields
type TradingATS struct {
	AtsID            string       `json:"ats_id"`
	AtsJurisdictions Jurisdiction `json:"ats_jurisdictions"`
}

// SalesToJuridiction data fields
type SalesToJuridiction struct {
	Country         string `json:"country"`
	RestrictionType string `json:"restriction_type"`
	State           string `json:"state"`
}

// TradingHoldPeriod data fields
type TradingHoldPeriod struct {
	References                    []ShareHolderReference `json:"references"`
	TradingHoldPeriodDuration     int                    `json:"trading_hold_period_duration"`
	TradingHoldPeriodDurationUnit int                    `json:"trading_hold_period_duration_unit"`
	TradingHoldPeriodStart        time.Time              `json:"trading_hold_period_start"`
}

// TradingRestriction data fields
type TradingRestriction struct {
	SaleToJurisdictions []SalesToJuridiction `json:"sale_to_jurisdictions"`
	TradingFrequency    TradingFrequency     `json:"trading_frequency"`
	TradingHoldPeriod   TradingHoldPeriod    `json:"trading_hold_period"`
	TradingQuantity     TradingQuantity      `json:"trading_quantity"`
}

// TradingFrequency data fields
type TradingFrequency struct {
	NumberOfTrades      int `json:"number_of_trades"`
	TradeFrequencyEvery int `json:"trade_frequency_every"`
	TradeFrequencyUnit  int `json:"trade_frequency_unit"`
}

// TradingQuantity data fields
type TradingQuantity struct {
	TradingQuantityBasis             string `json:"trading_quantity_basis"`
	MaximumNumberOfSecuritiesToTrade int    `json:"maximum_number_of_securities_to_trade"`
}

// Transfer data fields
type Transfer struct {
	TransferRestriction       bool      `json:"transfer_restriction"`
	TransferRestrictionNumber int       `json:"transfer_restriction_number"`
	TransferRestrictionUnit   int       `json:"transfer_restriction_unit"`
	TransferRestrictionStart  time.Time `json:"transfer_restriction_start"`
}

// Voting data fields
type Voting struct {
	Exists        bool           `json:"exists"`
	VotingActions []VotingAction `json:"voting_actions"`
	VotingBasis   VotingBasis    `json:"voting_basis"`
}

// VotingAction data fields
type VotingAction struct {
	CorporateAction string      `json:"corporate_action"`
	References      []Reference `json:"references"`
}

// VotingBasis data fields
type VotingBasis struct {
	References      []ShareHolderReference `json:"references"`
	VotingBasisName string                 `json:"voting_basis_name"`
}

// ShareHolderAgreementFilter data fields
type ShareHolderAgreementFilter struct {
	CompanyID string `json:"company_id"`
}

// SubscriptionAgreement data fields
type SubscriptionAgreement struct {
	CompanyID     string    `json:"company_id"`
	DocumentID    string    `json:"document_id"`
	AgreementText string    `json:"agreement_text"`
	Date          time.Time `json:"date"`
	Status        int32     `json:"status"`
	utils.MetaData
}
