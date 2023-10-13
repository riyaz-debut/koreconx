package koresecurities

import (
	"time"

	"kore_chaincode/core/utils"
)

// SecuritiesCertificateText data fields
type SecuritiesCertificateText struct {
	CompanyID        string `json:"company_id"`
	KoreSecuritiesID string `json:"koresecurities_id"`
	CertificateText  string `json:"certificate_text"`
	utils.MetaData
}

// CertificateTextFilter data fields
type CertificateTextFilter struct {
	KoreSecuritiesID string `json:"koresecurities_id"`
}

// Certificate data fields
type Certificate struct {
	CompanyID          string    `json:"company_id"`
	SecuritiesHolderID string    `json:"securities_holder_id"`
	KoresecuritiesID   string    `json:"koresecurities_id"`
	KoretransactionID  string    `json:"koretransaction_id"`
	CertificateNumber  string    `json:"certificate_number"`
	HoldingAmount      float64   `json:"holding_amount"`
	AveragePrice       float64   `json:"average_price"`
	DateAcquired       time.Time `json:"date_acquired"`
	Status             string    `json:"status"`
	utils.MetaData
}

// UpdateCertificateRequest data fields
type UpdateCertificateRequest struct {
	CertificateID string    `json:"certificate_id"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

// SecuritiesExchangePrice data fields
type SecuritiesExchangePrice struct {
	CompanyID        string    `json:"company_id"`
	KoresecuritiesID string    `json:"koresecurities_id"`
	QuoteProviderID  string    `json:"quote_provider_id"`
	ExchangePrice    float64   `json:"exchange_price"`
	ExchangeUnit     string    `json:"exchange_unit"`
	ExchangeUnitType string    `json:"exchange_unit_type"`
	QuoteDate        time.Time `json:"quote_date"`
	utils.MetaData
}

// InvestorHoldingRequest data fields
type InvestorHoldingRequest struct {
	SecuritiesHolderID string `json:"securities_holder_id"`
	KoresecuritiesID   string `json:"koresecurities_id"`
	RequestorID        string `json:"requestor_id"`
}

// ResponseInvestorHoldingExists data fields
type ResponseInvestorHoldingExists struct {
	Exists bool `json:"exists"`
}

// SecuritiesInstrument data fields
type SecuritiesInstrument struct {
	SecuritiesType string `json:"securities_type"`
	ClassType      string `json:"class_type"`
	Name           string `json:"name"`
	utils.MetaData
}

// Securities data fields
type Securities struct {
	OfferingMemorandumID    string                `json:"offering_memorandum_id"`
	SecuritiesCertificate   SecuritiesCertificate `json:"securities_certificate"`
	AvailableSecurities     float64               `json:"available_securities"`
	ShareholderAgreementID  string                `json:"shareholder_agreement_id"`
	SubscriptionAgreementID string                `json:"subscription_agreement_id"`
	AtsOperatorID           string                `json:"ats_operator_id"`
	BrokerDealerID          string                `json:"broker_dealer_id"`
	Status                  int32                 `json:"status"`
	utils.MetaData
}

// IssueSecuritiesInputRequest data fields
type IssueSecuritiesInputRequest struct {
	CompanyID               string    `json:"company_id"`
	OfferingMemorandumID    string    `json:"offering_memorandum_id"`
	ShareholderAgreementID  string    `json:"shareholder_agreement_id"`
	SubscriptionAgreementID string    `json:"subscription_agreement_id"`
	CertificateNumber       string    `json:"certificate_number"`
	SecuritiesType          string    `json:"securities_type"`
	Status                  int32     `json:"status"`
	CreatedAt               time.Time `json:"created_at"`
}

// SecuritiesFilter data fields
type SecuritiesFilter struct {
	CompanyID   string `json:"company_id"`
	RequestorID string `json:"requestor_id"`
}

// TransactionRequest data fields
type TransactionRequest struct {
	CompanyID          string    `json:"company_id"`
	SecuritiesHolderID string    `json:"securities_holder_id"`
	KoresecuritiesID   string    `json:"koresecurities_id"`
	TransactionID      string    `json:"transaction_id"`
	CertificateNumber  string    `json:"certificate_number"`
	HoldingAmount      float64   `json:"holding_amount"`
	AveragePrice       float64   `json:"average_price"`
	DateAcquired       time.Time `json:"date_acquired"`
	SourceSystemID     string    `json:"source_system_id"`
	CreatedAt          time.Time `json:"created_at"`
}

// Transaction data fields
type Transaction struct {
	CompanyID          string    `json:"company_id"`
	SecuritiesHolderID string    `json:"securities_holder_id"`
	KoresecuritiesID   string    `json:"koresecurities_id"`
	CertificateNumber  string    `json:"certificate_number"`
	HoldingAmount      float64   `json:"holding_amount"`
	AveragePrice       float64   `json:"average_price"`
	DateAcquired       time.Time `json:"date_acquired"`
	SourceSystemID     string    `json:"source_system_id"`
	utils.MetaData
}

// Holding data fields
type Holding struct {
	CompanyID          string    `json:"company_id"`
	SecuritiesHolderID string    `json:"securities_holder_id"`
	KoresecuritiesID   string    `json:"koresecurities_id"`
	HoldingAmount      float64   `json:"holding_amount"`
	AvailableShares    float64   `json:"available_shares"`
	LastUpdatedAt      time.Time `json:"last_updated_at"`
	ReasonCode         string    `json:"reason_code"`
	utils.MetaData
}

// HoldingFilter data fields
type HoldingFilter struct {
	CompanyID          string `json:"company_id"`
	KoresecuritiesID   string `json:"koresecurities_id"`
	SecuritiesHolderID string `json:"securities_holder_id"`
	RequestorID        string `json:"requestor_id"`
}

// ShareholderFilter data fields
type ShareholderFilter struct {
	CompanyID        string `json:"company_id"`
	KoresecuritiesID string `json:"koresecurities_id"`
	RequestorID      string `json:"requestor_id"`
}

// GetAvailableSharesRequest data fields
type GetAvailableSharesRequest struct {
	SecuritiesHolderID string  `json:"securities_holder_id"`
	KoresecuritiesID   string  `json:"koresecurities_id"`
	NumberOfShares     float64 `json:"number_of_shares"`
	RequestorID        string  `json:"requestor_id"`
}

// HoldSharesTransaction data fields
type HoldSharesTransaction struct {
	SecuritiesHolderID      string    `json:"securities_holder_id"`
	KoresecuritiesID        string    `json:"koresecurities_id"`
	NumberOfShares          float64   `json:"number_of_shares"`
	AvailableNumberOfShares float64   `json:"available_number_of_shares"`
	AtsID                   string    `json:"ats_id"`
	ReasonCode              string    `json:"reason_code"`
	ATSTransactionID        string    `json:"ats_transaction_id"`
	LastUpdatedAt           time.Time `json:"last_updated_at"`
	Status                  string    `json:"status"`
	utils.MetaData
}

// ReleaseSharesTransaction data fields
type ReleaseSharesTransaction struct {
	SecuritiesHolderID string    `json:"securities_holder_id"`
	KoresecuritiesID   string    `json:"koresecurities_id"`
	NumberOfShares     float64   `json:"number_of_shares"`
	AtsID              string    `json:"ats_id"`
	ReasonCode         string    `json:"reason_code"`
	ATSTransactionID   string    `json:"ats_transaction_id"`
	LastUpdatedAt      time.Time `json:"last_updated_at"`
	utils.MetaData
}

// UpdateHoldingRequest data fields
type UpdateHoldingRequest struct {
	SecuritiesHolderID string    `json:"securities_holder_id"`
	KoresecuritiesID   string    `json:"koresecurities_id"`
	NumberOfShares     float64   `json:"number_of_shares"`
	ReasonCode         string    `json:"reason_code"`
	LastUpdatedAt      time.Time `json:"last_updated_at"`
	CreatedAt          time.Time `json:"created_at"`
}

// SecuritiesCertificate data fields
type SecuritiesCertificate struct {
	CompanyID                string    `json:"company_id"`
	CertificateNumber        string    `json:"certificate_number"`
	ClassOfSecurities        string    `json:"class_of_securities"`
	CompanySymbol            string    `json:"company_symbol"`
	Currency                 string    `json:"currency"`
	DateIssued               time.Time `json:"date_issued"`
	NumberOfSecurities       float64   `json:"number_of_securities"`
	OfferingPricePerSecurity float64   `json:"offering_price_per_security"`
	SecuritiesholderID       string    `json:"securitiesholder_id"`
	SecuritiesholderName     string    `json:"securitiesholder_name"`
	SecuritiesType           string    `json:"securities_type"`
	Status                   string    `json:"status"`
	CertificateParent        []string  `json:"certificate_parent"`
	CertificateSucessor      []string  `json:"certificate_successor"`
}

// SecuritiesTransferRequest data fields
type SecuritiesTransferRequest struct {
	CompanyID                          string    `json:"company_id"`
	OwnerID                            string    `json:"owner_id"`
	KoresecuritiesID                   string    `json:"koresecurities_id"`
	TransferredToID                    string    `json:"transferred_to_id"`
	TotalSecurities                    float64   `json:"total_securities"`
	EffectiveDate                      time.Time `json:"effective_date"`
	TransferAuthorizationTransactionID string    `json:"transfer_authorization_transaction_id"`
	TransferType                       string    `json:"transfer_type"`
	TransferRequestor                  string    `json:"transfer_requestor"`
	TransferApprover                   string    `json:"transfer_approver"`
	TransactionID                      string    `json:"transaction_id"`
	Status                             int64     `json:"status"`
	Reason                             Reason    `json:"reason"`
	utils.MetaData
}

// UpdatingSecuritiesRequest data fields
type UpdatingSecuritiesRequest struct {
	KoresecuritiesID    string    `json:"koresecurities_id"`
	AvailableSecurities float64   `json:"available_securities"`
	CreatedAt           time.Time `json:"created_at"`
}

// UpdateSecuritiesTransferRequest data fields
type UpdateSecuritiesTransferRequest struct {
	TransferRequestID string    `json:"transfer_request_id"`
	Status            int64     `json:"status"`
	TransactionID     string    `json:"transaction_id"`
	Reason            Reason    `json:"reason"`
	CreatedAt         time.Time `json:"created_at"`
}

// Reason data fields
type Reason struct {
	ReasonCode string `json:"reason_code"`
	ReasonText string `json:"reason_text"`
}
