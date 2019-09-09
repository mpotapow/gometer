package responses

import "gometer/modules/http/contracts"

type authReponse struct {
	UserID int `json:"user_id"`
}

// NewAuthReponse ...
func NewAuthReponse(userID int) contracts.Response {
	return NewMainResponse(authReponse{
		UserID: userID,
	})
}
