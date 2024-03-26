package auth

import (
	"backend-kata/internal/delivery/rest"
	"backend-kata/internal/delivery/rest/auth/dtos"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginService interface {
	Authenticate(ctx context.Context, email string, password string) (*string, error)
}
type LoginHandler struct {
	loginService LoginService
}

func NewLoginHandler(loginService LoginService) LoginHandler {
	return LoginHandler{
		loginService: loginService,
	}
}
func (handler LoginHandler) Handle(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var loginRequest dtos.LoginRequest
	errorDecode := json.NewDecoder(request.Body).Decode(&loginRequest)
	if errorDecode != nil {
		rest.BadRequest(response, request, nil, fmt.Sprint("Error: Bad Request. Please check your request parameters and try again."))
		return
	}

	token, errorLogin := handler.loginService.Authenticate(ctx, loginRequest.Email, loginRequest.Password)
	if errorLogin != nil {
		rest.InternalError(response, request, nil, errorLogin.Error())
		return
	}

	rest.OK(response, request, nil, dtos.LoginResponse{Token: token})
}
