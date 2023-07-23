package request

type GenerateRequest struct {
	Table    string      `json:"table"`
	Column   interface{} `json:"column"`
	Quantity int64       `json:"quantity"`
}
