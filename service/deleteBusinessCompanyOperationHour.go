package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	db "github.com/AkezhanOb1/business-company/repository"
)


//CreateBusinessCompanyOperationHour is
func (*Server) DeleteBusinessCompanyOperationHour(ctx context.Context, request *pb.DeleteBusinessCompanyOperationHourRequest) (*pb.DeleteBusinessCompanyOperationHourResponse, error) {
	var operationHourID = request.OperationHourID
	deletedOperationHour, err := db.DeleteBusinessCompanyOperationHourRepository(ctx, operationHourID)
	if err != nil {
		return nil, err
	}

	return deletedOperationHour, nil
}

