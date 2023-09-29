package dto

type ApiResponse[T any] struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
	Data       T      `json:"data"`
}
