package industry

// IndustryDoc is used to return the response which contains only one message field
type IndustryDoc struct {
	Key string    `json:"key"`
	Doc *Industry `json:"doc"`
}
