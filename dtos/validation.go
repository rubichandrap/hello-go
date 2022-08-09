package dtos

type ValidationResponse struct {
	Success     bool                `json:"success"`
	Message     string              `json:"message"`
	Data        interface{}         `json:"data"`
	Validations map[string][]string `json:"validations"`
}

func (v *ValidationResponse) Send(validations map[string][]string) *ValidationResponse {
	return &ValidationResponse{
		Success:     false,
		Message:     "Please review your input",
		Data:        nil,
		Validations: validations,
	}
}
