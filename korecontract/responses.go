package korecontract

// DocumentDoc is used to return the response which contains only one message field
type DocumentDoc struct {
	Key string    `json:"key"`
	Doc *Document `json:"doc"`
}

// OfferingMemorandumDoc is used to return the response which contains only one message field
type OfferingMemorandumDoc struct {
	Key string              `json:"key"`
	Doc *OfferingMemorandum `json:"doc"`
}

// ShareHolderAgreementDoc is used to return the response which contains only one message field
type ShareHolderAgreementDoc struct {
	Key string                `json:"key"`
	Doc *ShareHolderAgreement `json:"doc"`
}
