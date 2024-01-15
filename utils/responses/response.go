package responses

type MapResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WebResponse(message string, data interface{}) MapResponse {
	return MapResponse{
		Message: message,
		Data:    data,
	}
}