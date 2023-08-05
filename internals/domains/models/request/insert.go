package request

import "github.com/HEEPOKE/generate-db/pkg/enums"

type InsertRequest struct {
	Key    string       `json:"key"`
	Insert enums.Insert `json:"insert"`
}
