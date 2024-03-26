package web

import (
	"backend-kata/cmd/web/middlewares"
	"backend-kata/internal/delivery/rest/health"
	"net/http"
)

type Route struct {
	Path        string
	Method      string
	Handler     func(w http.ResponseWriter, r *http.Request)
	Middlewares []middleware
}

func initRoutes(
	restAuthHandler *AuthHandler,
	restDrugHandler *DrugHandler,
	restVaccinationHandler *VaccinationHandler,
	authMiddleware *middlewares.AuthMiddleware,
) []Route {
	return []Route{
		{
			Path:    "/health",
			Method:  http.MethodGet,
			Handler: health.HandleHealth,
		},
		{
			Path:    "/v1/signup",
			Method:  http.MethodPost,
			Handler: restAuthHandler.signUpHandler.Handle,
		},
		{
			Path:    "/v1/login",
			Method:  http.MethodPost,
			Handler: restAuthHandler.loginHandler.Handle,
		},
		{
			Path:    "/v1/drugs",
			Method:  http.MethodPost,
			Handler: restDrugHandler.createDrugHandler.Handle,
			Middlewares: []middleware{
				authMiddleware.HandleMandatoryToken,
			},
		},
		{
			Path:    "/v1/drugs/:id",
			Method:  http.MethodPut,
			Handler: restDrugHandler.updateDrugHandler.Handle,
			Middlewares: []middleware{
				authMiddleware.HandleMandatoryToken,
			},
		},
		{
			Path:    "/v1/drugs",
			Method:  http.MethodGet,
			Handler: restDrugHandler.getAllDrugsHandler.Handle,
			Middlewares: []middleware{
				authMiddleware.HandleMandatoryToken,
			},
		},
		{
			Path:    "/v1/drugs/:id",
			Method:  http.MethodDelete,
			Handler: restDrugHandler.deleteDrugHandler.Handle,
			Middlewares: []middleware{
				authMiddleware.HandleMandatoryToken,
			},
		},
		{
			Path:    "/v1/vaccination",
			Method:  http.MethodPost,
			Handler: restVaccinationHandler.createVaccinationHandler.Handle,
			Middlewares: []middleware{
				authMiddleware.HandleMandatoryToken,
			},
		},
		{
			Path:    "/v1/vaccination/:id",
			Method:  http.MethodPut,
			Handler: restVaccinationHandler.updateVaccinationHandler.Handle,
			Middlewares: []middleware{
				authMiddleware.HandleMandatoryToken,
			},
		},
		{
			Path:    "/v1/vaccination",
			Method:  http.MethodGet,
			Handler: restVaccinationHandler.getAllVaccinationsHandler.Handle,
			Middlewares: []middleware{
				authMiddleware.HandleMandatoryToken,
			},
		},
		{
			Path:    "/v1/vaccination/:id",
			Method:  http.MethodDelete,
			Handler: restVaccinationHandler.deleteVaccinationHandler.Handle,
			Middlewares: []middleware{
				authMiddleware.HandleMandatoryToken,
			},
		},
	}
}
