package vaccination

import (
	"backend-kata/internal/delivery/rest"
	"backend-kata/internal/delivery/rest/vaccination/dtos"
	"backend-kata/internal/domain/vaccination"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type UpdateVaccinationService interface {
	Update(ctx context.Context, vaccinationId uint32, vaccination vaccination.Vaccination) error
}

type UpdateVaccinationHandler struct {
	updateVaccinationService UpdateVaccinationService
}

func NewUpdateVaccinationHandler(updateVaccinationService UpdateVaccinationService) UpdateVaccinationHandler {
	return UpdateVaccinationHandler{
		updateVaccinationService: updateVaccinationService,
	}
}
func (handler UpdateVaccinationHandler) Handle(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	vaccinationId, errorParam := rest.GetUInt32Param(request, "id", 0)
	if errorParam != nil {
		rest.BadRequest(response, request, nil, fmt.Sprint("Error: Bad Request. Please check your request parameters and try again."))
		return
	}

	var vaccinationRequest dtos.VaccinationRequest
	errorDecode := json.NewDecoder(request.Body).Decode(&vaccinationRequest)
	if errorDecode != nil {
		rest.BadRequest(response, request, nil, fmt.Sprint("Error: Bad Request. Please check your request parameters and try again."))
		return
	}

	vaccinationFromRequest := vaccinationRequest.ToVaccination()

	errorUpdateVaccination := handler.updateVaccinationService.Update(ctx, vaccinationId, vaccinationFromRequest)

	if errorUpdateVaccination != nil {
		rest.InternalError(response, request, nil, fmt.Sprint("Unexpected error"))
		return
	}

	rest.OK(response, request, nil, fmt.Sprint("Record updated successful"))
}
