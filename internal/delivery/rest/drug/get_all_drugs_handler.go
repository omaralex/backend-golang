package drug

import (
	"backend-kata/internal/delivery/rest"
	"backend-kata/internal/domain/drug"
	"context"
	"fmt"
	"net/http"
)

type GetAllDrugService interface {
	GetAll(ctx context.Context, limit int, offset int) (*[]drug.Drug, error)
}

type GetAllDrugsHandler struct {
	getAllDrugService GetAllDrugService
}

func NewGetAllDrugsHandler(getAllDrugService GetAllDrugService) GetAllDrugsHandler {
	return GetAllDrugsHandler{
		getAllDrugService: getAllDrugService,
	}
}
func (handler GetAllDrugsHandler) Handle(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	limit := 10
	offset := 0
	drugsList, errorGetAllDrugs := handler.getAllDrugService.GetAll(ctx, limit, offset)

	if errorGetAllDrugs != nil {
		rest.InternalError(response, request, nil, fmt.Sprint("Unexpected error"))
		return
	}

	rest.OK(response, request, nil, drugsList)
}
