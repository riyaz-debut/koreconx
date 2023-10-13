package person

// PersonDoc is used to return the response which contains only one message field
type PersonDoc struct {
	Key string  `json:"key"`
	Doc *Person `json:"doc"`
}
