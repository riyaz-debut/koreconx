package industry

import "kore_chaincode/core/utils"

// Industry data fields
type Industry struct {
	Title    string `json:"title"`
	ParentID string `json:"parent_id"`
	Industry string `json:"industry"`
	utils.MetaData
}

// RequestAllIndustries data fields
type RequestAllIndustries struct {
	IndustryID string `json:"industry_id"`
}
