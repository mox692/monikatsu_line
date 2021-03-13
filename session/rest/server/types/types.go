package types

type SessionRequest struct {
	StatusID string `json:"status_id"`
	UserID string	`json:"user_id"`
}
type SetSessionResponse struct { 
	StatusCode int32 `json:"status_code"`
	ErrMessage string`json:"error_message"`
}

type GetSessionResponse struct {
	StatusCode int32 `json:"status_code"`
	Data string `json:"data"`
	ErrMessage string`json:"error_message"`
}