package handler

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details"`
}
