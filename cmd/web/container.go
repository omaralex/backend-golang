package web

import (
	"backend-kata/cmd/web/middlewares"
	"backend-kata/config"
	authRest "backend-kata/internal/delivery/rest/auth"
	drugRest "backend-kata/internal/delivery/rest/drug"
	vaccinationRest "backend-kata/internal/delivery/rest/vaccination"
	userRepository "backend-kata/internal/domain/auth/repository"
	authService "backend-kata/internal/domain/auth/service"
	drugRepository "backend-kata/internal/domain/drug/repository"
	drugService "backend-kata/internal/domain/drug/service"
	vaccinationRepository "backend-kata/internal/domain/vaccination/repository"
	vaccinationService "backend-kata/internal/domain/vaccination/service"
	"backend-kata/internal/infrastructure/datasource"
	"backend-kata/internal/infrastructure/security"
	"net/http"
)

type RestHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
	signUpHandler RestHandler
	loginHandler  RestHandler
}

type DrugHandler struct {
	createDrugHandler  RestHandler
	updateDrugHandler  RestHandler
	getAllDrugsHandler RestHandler
	deleteDrugHandler  RestHandler
}
type VaccinationHandler struct {
	createVaccinationHandler  RestHandler
	updateVaccinationHandler  RestHandler
	getAllVaccinationsHandler RestHandler
	deleteVaccinationHandler  RestHandler
}
type Container struct {
	cfg                *config.Config
	authMiddleware     *middlewares.AuthMiddleware
	authHandler        *AuthHandler
	drugHandler        *DrugHandler
	vaccinationHandler *VaccinationHandler
}

func NewContainer(cfg *config.Config) Container {
	return Container{cfg: cfg}
}

func (container *Container) GetAuthMiddleware() *middlewares.AuthMiddleware {
	if container.authMiddleware == nil {
		security := security.NewSecurity(container.cfg.Security)
		container.authMiddleware = middlewares.NewAuthMiddleware(security)
	}

	return container.authMiddleware
}
func (container *Container) GetAuthHandler() *AuthHandler {
	if container.authHandler == nil {
		dataSourcePostgreSQL := datasource.New(container.cfg.PostgreSql)
		repository := userRepository.NewUserRepository(dataSourcePostgreSQL)

		security := security.NewSecurity(container.cfg.Security)
		authService := authService.NewAuthService(repository, security)

		container.authHandler = &AuthHandler{
			loginHandler:  authRest.NewLoginHandler(authService),
			signUpHandler: authRest.NewSignUpHandler(authService),
		}
	}

	return container.authHandler
}

func (container *Container) GetDrugHandler() *DrugHandler {
	if container.drugHandler == nil {
		dataSourcePostgreSQL := datasource.New(container.cfg.PostgreSql)
		repository := drugRepository.NewDrugRepository(dataSourcePostgreSQL)

		drugService := drugService.NewDrugService(repository)

		container.drugHandler = &DrugHandler{
			createDrugHandler:  drugRest.NewCreateDrugHandler(drugService),
			updateDrugHandler:  drugRest.NewUpdateDrugHandler(drugService),
			getAllDrugsHandler: drugRest.NewGetAllDrugsHandler(drugService),
			deleteDrugHandler:  drugRest.NewDeleteDrugHandler(drugService),
		}
	}

	return container.drugHandler
}

func (container *Container) GetVaccinationHandler() *VaccinationHandler {
	if container.vaccinationHandler == nil {
		dataSourcePostgreSQL := datasource.New(container.cfg.PostgreSql)
		repository := vaccinationRepository.NewVaccinationRepository(dataSourcePostgreSQL)

		vaccinationService := vaccinationService.NewVaccinationService(repository)

		container.vaccinationHandler = &VaccinationHandler{
			createVaccinationHandler:  vaccinationRest.NewCreateVaccinationHandler(vaccinationService),
			updateVaccinationHandler:  vaccinationRest.NewUpdateVaccinationHandler(vaccinationService),
			getAllVaccinationsHandler: vaccinationRest.NewGetAllVaccinationsHandler(vaccinationService),
			deleteVaccinationHandler:  vaccinationRest.NewDeleteVaccinationHandler(vaccinationService),
		}
	}

	return container.vaccinationHandler
}
