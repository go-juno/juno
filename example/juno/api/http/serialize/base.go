package serialize

type List struct {
	Items interface{} `json:"items"`
	Total int64       `json:"total"`
}
