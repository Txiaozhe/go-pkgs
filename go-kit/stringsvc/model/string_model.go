package model

// StringService
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// UppercaseRequest
type UppercaseRequest struct {
	S string `json:"s"`
}

// UppercaseResponse
type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err"`
}

// CountRequest
type CountRequest struct {
	S string `json:"s"`
}

// CountRequest
type CountResponse struct {
	V int `json:"v"`
}
