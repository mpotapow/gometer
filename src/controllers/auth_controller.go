package controllers

import (
	"gometer/modules/core"
	httpContracts "gometer/modules/http/contracts"
	"gometer/src/contracts"
	"gometer/src/requests"
	"gometer/src/responses"
	"gometer/src/tools"
	"net/http"
)

// AuthController ...
func AuthController(w httpContracts.ResponseWriter, r *http.Request) {

	loginRequest := requests.LoginRequest{}

	if err := loginRequest.ParseJson(r, &loginRequest); err != nil {
		w.ToJson(responses.NewMainUnprocessableResponse("Invalid request"))
		return
	}
	if err := loginRequest.Validate(); err != nil {
		w.ToJson(responses.NewMainUnprocessableResponse(err.Error()))
		return
	}

	authService, _ := core.GetApplicationInstance().Get("auth-service")
	service := authService.(contracts.AuthService)

	user, err := service.Authorize(loginRequest.Login, loginRequest.Password)
	if err != nil {
		w.ToJson(responses.NewMainUnprocessableResponse("Wrong login or password"))
		return
	}

	session := tools.GetSessionInstance(r)
	session.GetStorage().Set("user_id", user.ID)

	w.ToJson(responses.NewAuthReponse(user.ID))
}
