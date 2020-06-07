package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	db "github.com/AkezhanOb1/business-company/repository"
)

//CreateBusinessCompanyOperationHour is
func (*Server) UpdateBusinessCompanyOperationHour(ctx context.Context, request *pb.UpdateBusinessCompanyOperationHourRequest) (*pb.UpdateBusinessCompanyOperationHourResponse, error) {
	updatedOperationHour, err := db.UpdateBusinessCompanyOperationHourRepository(ctx, request.BusinessCompanyOperationHour)
	if err != nil {
		return nil, err
	}

	return updatedOperationHour, nil
}

