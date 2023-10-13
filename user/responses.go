package user

// CompanyDoc is used to return the response which contains only one message field
type CompanyDoc struct {
	Key string   `json:"key"`
	Doc *Company `json:"doc"`
}

// ServiceProviderDoc is used to return the response which contains only one message field
type ServiceProviderDoc struct {
	Key string           `json:"key"`
	Doc *ServiceProvider `json:"doc"`
}

// CompanyList is used to return the response which contains only one message field
type CompanyList struct {
	ID               string `json:"id"`
	CompanyLegalName string `json:"company_legal_name"`
}
