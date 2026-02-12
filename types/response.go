package types

type ResponseType string

const (
	Error   ResponseType = "error"
	Success ResponseType = "success"
	Warning ResponseType = "warning"
)

type ResponseMessage struct {
	Type    ResponseType `json:"type" example:"error"`
	Message string       `json:"message" example:"Message"`
}
