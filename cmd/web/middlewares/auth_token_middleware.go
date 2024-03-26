package middlewares

import (
	"backend-kata/internal/delivery/rest"
	"backend-kata/internal/infrastructure/security"
	"backend-kata/pkg"
	"fmt"
	"net/http"
)

type AuthMiddleware struct {
	security security.Security
}

func NewAuthMiddleware(security security.Security) *AuthMiddleware {
	return &AuthMiddleware{
		security: security,
	}
}

func (a *AuthMiddleware) HandleMandatoryToken(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(pkg.AuthorizationHeader)
		if token == "" {
			rest.BadRequest(w, r, nil, fmt.Sprint("header Authorization is mandatory"))
			return
		}

		isValid := a.security.IsValidateJWT(token)
		if !isValid {
			rest.BadRequest(w, r, nil, fmt.Sprint("token invalid"))
			return
		}
		h(w, r)
	}
}
