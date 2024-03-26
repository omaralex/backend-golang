package health

import (
	"backend-kata/internal/delivery/rest"
	"net/http"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	rest.OK(w, r, nil, map[string]string{
		"status":  "ok",
		"version": "1.0",
	})
}
