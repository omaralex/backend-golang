package vaccination

import (
	"backend-kata/internal/delivery/rest"
	"backend-kata/internal/domain/vaccination"
	"context"
	"fmt"
	"net/http"
)

type GetAllVaccinationsService interface {
	GetAll(ctx context.Context, limit int, offset int) (*[]vaccination.Vaccination, error)
}

type GetAllVaccinationHandler struct {
	getAllVaccinationsService GetAllVaccinationsService
}

func NewGetAllVaccinationsHandler(getAllVaccinationsService GetAllVaccinationsService) GetAllVaccinationHandler {
	return GetAllVaccinationHandler{
		getAllVaccinationsService: getAllVaccinationsService,
	}
}
func (handler GetAllVaccinationHandler) Handle(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	limit := 10
	offset := 0
	vaccinationList, errorGetAllVaccination := handler.getAllVaccinationsService.GetAll(ctx, limit, offset)

	if errorGetAllVaccination != nil {
		rest.InternalError(response, request, nil, fmt.Sprint("Unexpected error"))
		return
	}

	rest.OK(response, request, nil, vaccinationList)
}
