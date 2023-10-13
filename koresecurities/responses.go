package koresecurities

// CertificateDoc is used to return the response which contains only one message field
type CertificateDoc struct {
	Key string       `json:"key"`
	Doc *Certificate `json:"doc"`
}

// SecuritiesCertificateTextDoc is used to return the response which contains only one message field
type SecuritiesCertificateTextDoc struct {
	Key string                     `json:"key"`
	Doc *SecuritiesCertificateText `json:"doc"`
}

// HoldingDoc is used to return the response which contains only one message field
type HoldingDoc struct {
	HoldingID       string  `json:"holding_id"`
	HoldingAmount   float64 `json:"holding_amount"`
	AvailableShares float64 `json:"available_shares"`
}

// SecuritiesDoc is used to return the response which contains only one message field
type SecuritiesDoc struct {
	Key string      `json:"key"`
	Doc *Securities `json:"doc"`
}

// ShareHoldersIDResponse is used to return the response which contains only one message field
type ShareHoldersIDResponse struct {
	Data []string `json:"data"`
}

// AllHoldingAllSecuritiesResponse data fields
type AllHoldingAllSecuritiesResponse struct {
	KoreSecuritiesID   string `json:"koresecurities_id"`
	HoldingID          string `json:"holding_id"`
	SecuritiesHolderID string `json:"securities_holder_id"`
}

// AllHoldingBySecuritiesResponse data fields
type AllHoldingBySecuritiesResponse struct {
	HoldingID          string `json:"holding_id"`
	SecuritiesHolderID string `json:"securities_holder_id"`
}

// HoldingByCompanyResponse data fields
type HoldingByCompanyResponse struct {
	HoldingAmount      float64 `json:"holding_amount"`
	AvailableShares    float64 `json:"available_shares"`
	SecuritiesHolderID string  `json:"securities_holder_id"`
}

// AllTradableHoldingsResponse data fields
type AllTradableHoldingsResponse struct {
	CompanyLegalName        string  `json:"company_legal_name"`
	CompanySymbol           string  `json:"company_symbol"`
	OfferingSecuritiesClass string  `json:"offering_securities_class"`
	HoldingAmount           float64 `json:"holding_amount"`
	AvailableShares         float64 `json:"available_shares"`
	SecuritiesHolderID      string  `json:"securities_holder_id"`
	KoresecuritiesID        string  `json:"koresecurities_id"`
	HoldingID               string  `json:"holding_id"`
}
