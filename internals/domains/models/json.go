package models

type JsonStructure struct {
	Table    string                   `json:"table"`
	DataKey  string                   `json:"key"`
	Quantity int64                    `json:"quantity"`
	Datas    []map[string]interface{} `json:"datas"`
}
