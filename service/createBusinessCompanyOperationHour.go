package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	db "github.com/AkezhanOb1/business-company/repository"
)

//CreateBusinessCompanyOperationHour is
func (*Server) CreateBusinessCompanyOperationHour(ctx context.Context, request *pb.CreateBusinessCompanyOperationHourRequest) (*pb.CreateBusinessCompanyOperationHourResponse, error) {
	businessCompany, err := db.CreateBusinessCompanyOperationHourRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	return businessCompany, nil
}


