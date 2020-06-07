package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	db "github.com/AkezhanOb1/business-company/repository"
)


//CreateBusinessCompany is
func (*Server) CreateBusinessCompany(ctx context.Context, request *pb.CreateBusinessCompanyRequest) (*pb.CreateBusinessCompanyResponse, error) {
	businessCompany, err := db.CreateCompanyRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	return businessCompany, nil
}

