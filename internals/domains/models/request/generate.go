package request

type GenerateRequest struct {
	Table    string      `json:"table"`
	Columns  interface{} `json:"columns"`
	Quantity int64       `json:"quantity"`
}
