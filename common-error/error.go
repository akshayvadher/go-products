package common_error

type CommonError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (m *CommonError) Error() string {
	return m.Code + "," + m.Message
}
