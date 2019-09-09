package controllers

import (
	"gometer/modules/core"
	httpContracts "gometer/modules/http/contracts"
	"gometer/src/contracts"
	"gometer/src/requests"
	"gometer/src/responses"
	"net/http"
)

// AuthController ...
func AuthController(w httpContracts.ResponseWriter, r *http.Request) {

	loginRequest := requests.LoginRequest{}
	err := loginRequest.ParseJson(r, &loginRequest)
	if err != nil {
		w.ToJson(responses.NewMainUnprocessableResponse("Некорректный запрос"))
		return
	}
	if len(loginRequest.Login) <= 0 || len(loginRequest.Password) <= 0 {
		w.ToJson(responses.NewMainUnprocessableResponse("Некорректный запрос"))
		return
	}

	authService, _ := core.GetApplicationInstance().Get("auth-service")
	service := authService.(contracts.AuthService)

	user, err := service.Authorize(loginRequest.Login, loginRequest.Password)
	if err != nil {
		w.ToJson(responses.NewMainUnprocessableResponse("Неверный логин или пароль"))
		return
	}

	w.ToJson(responses.NewAuthReponse(user.ID))
}
