package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	db "github.com/AkezhanOb1/business-company/repository"
)



//GetBusinessCompanyOperationHours is
func (*Server) GetBusinessCompanyOperationHours(ctx context.Context, request *pb.GetBusinessCompanyOperationHoursRequest) (*pb.GetBusinessCompanyOperationHoursResponse, error) {
	companyID := request.GetCompanyID()
	operationHours, err := db.GetBusinessCompanyOperationHoursRepository(ctx, companyID)
	if err != nil {
		return nil, err
	}

	return operationHours, nil
}
