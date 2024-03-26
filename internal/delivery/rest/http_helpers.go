package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type errorResponse struct {
	Messages []string `json:"messages"`
}

func (e *errorResponse) Error() string {
	return fmt.Sprintf("%s", strings.Join(e.Messages, ","))
}

func GetStringParam(r *http.Request, key, defaultValue string) string {
	keys, ok := r.URL.Query()[key]

	if !ok || len(keys[0]) < 1 {
		return defaultValue
	}

	return keys[0]
}

func GetUInt32Param(r *http.Request, key string, defaultValue uint32) (uint32, error) {
	keys, ok := r.URL.Query()[key]

	if !ok || len(keys[0]) < 1 {
		return defaultValue, nil
	}

	value, err := strconv.ParseInt(keys[0], 10, 0)
	if err != nil {
		return defaultValue,
			fmt.Errorf("server: %s is not a valid int value for %s", keys[0], key)
	}

	return uint32(value), nil
}

func InternalError(w http.ResponseWriter, r *http.Request, headers map[string]string, messages ...string) {
	err := &errorResponse{
		Messages: messages,
	}
	Render(w, r, err, headers, http.StatusInternalServerError)
}

func BadRequest(w http.ResponseWriter, r *http.Request, headers map[string]string, messages ...string) {
	err := &errorResponse{
		Messages: messages,
	}
	Render(w, r, err, headers, http.StatusBadRequest)
}

func OK(w http.ResponseWriter, r *http.Request, headers map[string]string, obj interface{}) {
	Render(w, r, obj, headers, http.StatusOK)
}

func Render(w http.ResponseWriter, _ *http.Request, obj interface{}, headers map[string]string, status int) {
	js, err := json.Marshal(obj)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	addHeaders := func() {
		if headers == nil {
			w.Header().Set("Content-Type", "application/json")
			return
		}

		for key, value := range headers {
			w.Header().Set(key, value)
		}
		value, exists := headers["Content-Type"]
		if exists {
			w.Header().Set("Content-Type", value)
		} else {
			w.Header().Set("Content-Type", "application/json")
		}
		return
	}

	addHeaders()
	w.WriteHeader(status)
	w.Write(js)
}
