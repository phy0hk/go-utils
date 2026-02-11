package types

type ResponseMessage struct {
	Status  ResponseStatus
	Message string `json:"message" example:"Message"`
}

type ResponseStatus struct {
	Code    int    `json:"code" example:"Status Code"`
	Message string `json:"message" example:"Status Name"`
}
