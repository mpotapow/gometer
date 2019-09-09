package responses

import (
	"gometer/modules/http/contracts"
	"net/http"
)

// MainResponse ...
type MainResponse struct {
	Status  int         `json:"-"`
	Error   bool        `json:"error"`
	Content interface{} `json:"content"`
}

// NewMainResponse ...
func NewMainResponse(content interface{}) contracts.Response {

	return &MainResponse{
		Error:   false,
		Content: content,
		Status:  http.StatusOK,
	}
}

// NewMainUnprocessableResponse ...
func NewMainUnprocessableResponse(content interface{}) contracts.Response {

	return &MainResponse{
		Error:   true,
		Content: content,
		Status:  http.StatusUnprocessableEntity,
	}
}

// GetStatus ...
func (m *MainResponse) GetStatus() int {
	return m.Status
}

// SetStatus ...
func (m *MainResponse) SetStatus(status int) contracts.Response {
	m.Status = status
	return m
}

// HasError ...
func (m *MainResponse) HasError() bool {
	return m.Error
}

// SetHasError ...
func (m *MainResponse) SetHasError(flag bool) contracts.Response {
	m.Error = flag
	return m
}
