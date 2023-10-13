package trade

import (
	"kore_chaincode/core/utils"
	"time"
)

// TradeRequest data fields
type TradeRequest struct {
	ShareholderID    string    `json:"shareholder_id"`
	OrderDate        time.Time `json:"order_date"`
	Direction        string    `json:"direction"`
	CompanyID        string    `json:"company_id"`
	KoresecuritiesID string    `json:"koresecurities_id"`
	NumberToTrade    float64   `json:"number_to_trade"`
	TradePrice       float64   `json:"trade_price"`
	OrderType        string    `json:"order_type"`
	Limit            float64   `json:"limit"`
	Stop             float64   `json:"stop"`
	OrderDuration    string    `json:"order_duration"`
	OrderExpiry      time.Time `json:"order_expiry"`
	Misc             string    `json:"misc"`
	utils.MetaData
}

// AtsTradeRequest data fields
type AtsTradeRequest struct {
	CompanyID        string    `json:"company_id"`
	KoresecuritiesID string    `json:"koresecurities_id"`
	RequestorID      string    `json:"requestor_id"`
	ATSTransactionID string    `json:"ats_transaction_id"`
	OwnerID          string    `json:"owner_id"`
	TransferredToID  string    `json:"transferred_to_id"`
	AuthorizationID  string    `json:"transfer_authorization_transaction_id"`
	TotalSecurities  float64   `json:"total_securities"`
	TradePrice       float64   `json:"trade_price"`
	EffectiveDate    time.Time `json:"effective_date"`
	TransactionID    string    `json:"transaction_id"`
	CreatedAt        time.Time `json:"created_at"`
}

// AtsTrade data fields
type AtsTrade struct {
	CompanyID        string    `json:"company_id"`
	KoresecuritiesID string    `json:"koresecurities_id"`
	RequestorID      string    `json:"requestor_id"`
	ATSTransactionID string    `json:"ats_transaction_id"`
	OwnerID          string    `json:"owner_id"`
	TransferredToID  string    `json:"transferred_to_id"`
	AuthorizationID  string    `json:"transfer_authorization_transaction_id"`
	TotalSecurities  float64   `json:"total_securities"`
	TradePrice       float64   `json:"trade_price"`
	EffectiveDate    time.Time `json:"effective_date"`
	utils.MetaData
}

// AllTradeRequest data fields
type AllTradeRequest struct {
	RequestorID   string `json:"requestor_id"`
	ShareholderID string `json:"shareholder_id"`
	CompanyID     string `json:"company_id"`
}
