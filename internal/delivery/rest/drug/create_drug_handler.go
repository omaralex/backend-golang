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

type CreateDrugService interface {
	Create(ctx context.Context, drug drug.Drug) (*uint32, error)
}

type CreateDrugHandler struct {
	createDrugService CreateDrugService
}

func NewCreateDrugHandler(createDrugService CreateDrugService) CreateDrugHandler {
	return CreateDrugHandler{
		createDrugService: createDrugService,
	}
}
func (handler CreateDrugHandler) Handle(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var drugRequest dtos.DrugRequest
	errorDecode := json.NewDecoder(request.Body).Decode(&drugRequest)
	if errorDecode != nil {
		rest.BadRequest(response,
			request, nil,
			fmt.Sprint("Error: Bad Request. Please check your request parameters and try again."),
		)
		return
	}

	drugFromRequest := drugRequest.ToDrug()

	drugId, errorCreation := handler.createDrugService.Create(ctx, drugFromRequest)

	if errorCreation != nil {
		rest.InternalError(response, request, nil, fmt.Sprint("Unexpected error"))
		return
	}

	rest.OK(response, request, nil, dtos.DrugResponse{
		ID:          drugId,
		Name:        drugFromRequest.Name,
		Approved:    drugFromRequest.Approved,
		MiniDose:    drugFromRequest.MaxDose,
		MaxDose:     drugFromRequest.MaxDose,
		AvailableAt: drugFromRequest.AvailableAt,
	})
}
