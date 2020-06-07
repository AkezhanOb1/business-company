package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	db "github.com/AkezhanOb1/business-company/repository"
)


//GetBusinessCompanyOperationHourByDay is
func (*Server) GetBusinessCompanyOperationHourByDay(ctx context.Context, request *pb.GetBusinessCompanyOperationHourByDayRequest) (*pb.GetBusinessCompanyOperationHourByDayResponse, error) {
	dayOfWeek := request.GetDayOfWeek()
	companyID := request.GetCompanyID()
	operationHour, err := db.GetBusinessCompanyOperationHourByDayRepository(ctx, companyID, dayOfWeek)
	if err != nil {
		return nil, err
	}

	return operationHour, nil
}
