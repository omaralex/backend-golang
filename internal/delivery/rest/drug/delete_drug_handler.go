package drug

import (
	"backend-kata/internal/delivery/rest"
	"context"
	"fmt"
	"net/http"
)

type DeleteDrugService interface {
	Delete(ctx context.Context, id uint32) error
}

type DeleteDrugHandler struct {
	deleteDrugService DeleteDrugService
}

func NewDeleteDrugHandler(deleteDrugService DeleteDrugService) DeleteDrugHandler {
	return DeleteDrugHandler{
		deleteDrugService: deleteDrugService,
	}
}
func (handler DeleteDrugHandler) Handle(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	drugId, errorParam := rest.GetUInt32Param(request, "id", 0)
	if errorParam != nil {
		rest.BadRequest(response,
			request, nil,
			fmt.Sprint("Error: Bad Request. Please check your request parameters and try again."))
		return
	}

	errorDeleteDrug := handler.deleteDrugService.Delete(ctx, drugId)

	if errorDeleteDrug != nil {
		rest.InternalError(response, request, nil, fmt.Sprint("Unexpected error"))
		return
	}

	rest.OK(response, request, nil, fmt.Sprint("Record deleted successful"))
}
