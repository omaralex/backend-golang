package auth

import (
	"backend-kata/internal/delivery/rest"
	"backend-kata/internal/delivery/rest/auth/dtos"
	"backend-kata/internal/domain/auth"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SignUpService interface {
	SignUp(ctx context.Context, user auth.User) error
}
type SignUpHandler struct {
	signUpService SignUpService
}

func NewSignUpHandler(signUpService SignUpService) SignUpHandler {
	return SignUpHandler{
		signUpService: signUpService,
	}
}
func (handler SignUpHandler) Handle(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var signUpRequest dtos.SignUpRequest
	errorDecode := json.NewDecoder(request.Body).Decode(&signUpRequest)
	if errorDecode != nil {
		rest.BadRequest(response, request, nil, fmt.Sprint("Error: Bad Request. Please check your request parameters and try again."))
		return
	}

	user := signUpRequest.ToUser()

	errorSignUp := handler.signUpService.SignUp(ctx, user)
	if errorSignUp != nil {
		rest.InternalError(response, request, nil, fmt.Sprint("Unexpected error"))
		return
	}

	rest.OK(response, request, nil, fmt.Sprint("Registration successful!"))
}
