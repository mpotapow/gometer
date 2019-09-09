package http

import (
	"encoding/json"
	"gometer/modules/http/contracts"
	"net/http"
)

// Response ...
type Response struct {
	http.ResponseWriter
}

// NewResponse ...
func NewResponse(w http.ResponseWriter) Response {
	return Response{w}
}

// ToJson ...
func (r Response) ToJson(v contracts.Response) {

	r.ResponseWriter.WriteHeader(v.GetStatus())
	r.ResponseWriter.Header().Add("Content-Type", "application/json")

	json.NewEncoder(r.ResponseWriter).Encode(v)
}
