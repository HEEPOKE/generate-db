package examples

type SuccessStatusMessage struct {
	Code        int    `json:"code" example:"0000"`
	Message     string `json:"message" example:"Success"`
	Service     string `json:"service"`
	Description string `json:"description"`
}

type FailedStatusMessage struct {
	Code        int    `json:"code" example:"0001"`
	Message     string `json:"message" example:"Fail"`
	Service     string `json:"service"`
	Description string `json:"description"`
}

type NotFoundStatusMessage struct {
	Code        int    `json:"code" example:"0002"`
	Message     string `json:"message" example:"Not Found"`
	Service     string `json:"service"`
	Description string `json:"description"`
}

type DuplicateStatusMessage struct {
	Code        int    `json:"code" example:"0003"`
	Message     string `json:"message" example:"Duplicate"`
	Service     string `json:"service"`
	Description string `json:"description"`
}

type InvalidStatusMessage struct {
	Code        int    `json:"code" example:"0004"`
	Message     string `json:"message" example:"Invalid"`
	Service     string `json:"service"`
	Description string `json:"description"`
}
