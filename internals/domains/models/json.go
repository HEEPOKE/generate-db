package models

type JsonStructure struct {
	Table    string                   `json:"table" example:"users"`
	DataKey  string                   `json:"key" example:"123456"`
	Quantity int64                    `json:"quantity" example:"1000"`
	Datas    []map[string]interface{} `json:"datas"`
}
