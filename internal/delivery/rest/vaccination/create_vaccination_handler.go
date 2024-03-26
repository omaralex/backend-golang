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

type CreateVaccinationService interface {
	Create(ctx context.Context, vaccination vaccination.Vaccination) (*uint32, error)
}

type CreateVaccinationHandler struct {
	createVaccinationService CreateVaccinationService
}

func NewCreateVaccinationHandler(createVaccinationService CreateVaccinationService) CreateVaccinationHandler {
	return CreateVaccinationHandler{
		createVaccinationService: createVaccinationService,
	}
}
func (handler CreateVaccinationHandler) Handle(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var vaccinationRequest dtos.VaccinationRequest
	errorDecode := json.NewDecoder(request.Body).Decode(&vaccinationRequest)
	if errorDecode != nil {
		rest.BadRequest(response, request, nil, fmt.Sprint("Error: Bad Request. Please check your request parameters and try again."))
		return
	}

	vaccinationFromRequest := vaccinationRequest.ToVaccination()

	vaccinationId, errorCreateVaccunation := handler.createVaccinationService.Create(ctx, vaccinationFromRequest)

	if errorCreateVaccunation != nil {
		rest.InternalError(response, request, nil, fmt.Sprint("Unexpected error"))
		return
	}

	rest.OK(response, request, nil, dtos.VaccinationResponse{
		ID:     vaccinationId,
		Name:   vaccinationFromRequest.Name,
		DrugId: vaccinationFromRequest.DrugId,
		Dose:   vaccinationFromRequest.Dose,
		Date:   vaccinationFromRequest.Date,
	})
}
