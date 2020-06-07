package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	db "github.com/AkezhanOb1/business-company/repository"
)


//GetBusinessCompanyServices is
func (*Server) GetBusinessCompanyServices(ctx context.Context, request *pb.GetBusinessCompanyServicesRequest) (*pb.GetBusinessCompanyServicesResponse, error) {
	businessCompanyID := request.GetBusinessCompanyID()
	businessCompany, err := db.GetBusinessCompanyServicesRepository(ctx, businessCompanyID)
	if err != nil {
		return nil, err
	}

	return businessCompany, nil
}