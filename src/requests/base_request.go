package requests

import (
	"encoding/json"
	"gometer/src/contracts"
	"net/http"
)

type baseRequest struct {
}

// ParseJson ...
func (b *baseRequest) ParseJson(req *http.Request, container contracts.Request) error {
	defer req.Body.Close()

	err := json.NewDecoder(req.Body).Decode(container)
	if err != nil {
		return err
	}
	return nil
}
