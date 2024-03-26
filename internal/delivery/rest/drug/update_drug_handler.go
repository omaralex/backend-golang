package drug

import (
	"backend-kata/internal/delivery/rest"
	"backend-kata/internal/delivery/rest/drug/dtos"
	"backend-kata/internal/domain/drug"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type UpdateDrugService interface {
	Update(ctx context.Context, id uint32, drug drug.Drug) error
}

type UpdateDrugHandler struct {
	updateDrugService UpdateDrugService
}

func NewUpdateDrugHandler(updateDrugService UpdateDrugService) UpdateDrugHandler {
	return UpdateDrugHandler{
		updateDrugService: updateDrugService,
	}
}
func (handler UpdateDrugHandler) Handle(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	drugId, errorParam := rest.GetUInt32Param(request, "id", 0)
	if errorParam != nil {
		rest.BadRequest(response, request, nil, fmt.Sprint("Error: Bad Request. Please check your request parameters and try again."))
		return
	}

	var drugRequest dtos.DrugRequest
	errorDecode := json.NewDecoder(request.Body).Decode(&drugRequest)
	if errorDecode != nil {
		rest.BadRequest(response, request, nil, fmt.Sprint("Error: Bad Request. Please check your request parameters and try again."))
		return
	}

	drugFromRequest := drugRequest.ToDrug()

	errorUpdateDrug := handler.updateDrugService.Update(ctx, drugId, drugFromRequest)

	if errorUpdateDrug != nil {
		rest.InternalError(response, request, nil, fmt.Sprint("Unexpected error"))
		return
	}

	rest.OK(response, request, nil, fmt.Sprint("Record updated successful"))
}
