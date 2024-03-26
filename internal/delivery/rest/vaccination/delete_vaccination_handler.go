package vaccination

import (
	"backend-kata/internal/delivery/rest"
	"context"
	"fmt"
	"net/http"
)

type DeleteVaccinationService interface {
	Delete(ctx context.Context, id uint32) error
}

type DeleteVaccinationHandler struct {
	deleteVaccinationService DeleteVaccinationService
}

func NewDeleteVaccinationHandler(deleteVaccinationService DeleteVaccinationService) DeleteVaccinationHandler {
	return DeleteVaccinationHandler{
		deleteVaccinationService: deleteVaccinationService,
	}
}
func (handler DeleteVaccinationHandler) Handle(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	vaccinationId, errorParam := rest.GetUInt32Param(request, "id", 0)
	if errorParam != nil {
		rest.BadRequest(response,
			request, nil,
			fmt.Sprint("Error: Bad Request. Please check your request parameters and try again."))
		return
	}

	errorDeleteVaccination := handler.deleteVaccinationService.Delete(ctx, vaccinationId)

	if errorDeleteVaccination != nil {
		rest.InternalError(response, request, nil, fmt.Sprint("Unexpected error"))
		return
	}

	rest.OK(response, request, nil, fmt.Sprint("Record deleted successful"))
}
