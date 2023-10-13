package user

// AssociateTransferAgentRequest data fields
type AssociateTransferAgentRequest struct {
	CompanyID       string `json:"company_id"`
	TransferAgentID string `json:"transfer_agent_id"`
}

// CompaniesByRequestorID data fields
type CompaniesByRequestorID struct {
	RequestorID string `json:"requestor_id"`
}

// AssociateNotificationURLRequest data fields
type AssociateNotificationURLRequest struct {
	CompanyID string `json:"company_id"`
	URL       string `json:"url"`
}
