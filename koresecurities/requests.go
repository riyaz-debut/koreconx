package koresecurities

import "time"

// AssociateATSWithSecurityRequest data fields
type AssociateATSWithSecurityRequest struct {
	CompanyID        string    `json:"company_id"`
	KoresecuritiesID string    `json:"koresecurities_id"`
	AtsOperatorID    string    `json:"ats_operator_id"`
	CreatedAt        time.Time `json:"created_at"`
}

// AssociateBrokerWithSecurityRequest data fields
type AssociateBrokerWithSecurityRequest struct {
	CompanyID        string    `json:"company_id"`
	KoresecuritiesID string    `json:"koresecurities_id"`
	BrokerDealerID   string    `json:"broker_dealer_id"`
	CreatedAt        time.Time `json:"created_at"`
}

// AllHoldingsByAllSecuritiesRequest data fields
type AllHoldingsByAllSecuritiesRequest struct {
	CompanyID   string `json:"company_id"`
	RequestorID string `json:"requestor_id"`
}

// GetAllHoldingsbySecuritiesIDRequest data fields
type GetAllHoldingsbySecuritiesIDRequest struct {
	CompanyID        string `json:"company_id"`
	RequestorID      string `json:"requestor_id"`
	KoreSecuritiesID string `json:"koresecurities_id"`
}

// AllTradableHoldingsRequest data fields
type AllTradableHoldingsRequest struct {
	SecuritiesHolderID string `json:"securities_holder_id"`
	RequestorID        string `json:"requestor_id"`
}
