package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	db "github.com/AkezhanOb1/business-company/repository"
)


//GetBusinessCompany is
func (*Server) GetBusinessCompany(ctx context.Context, request *pb.GetBusinessCompanyRequest) (*pb.GetBusinessCompanyResponse, error) {
	businessCompany, err := db.GetBusinessCompanyRepository(ctx, request.BusinessCompanyID)
	if err != nil {
		return nil, err
	}

	businessCompanyImages, err := db.GetBusinessCompanyImagesRepository(context.Background(), request.GetBusinessCompanyID())
	if err != nil {
		return nil, err
	}

	businessCompany.BusinessCompany.BusinessCompanyImages = businessCompanyImages


	return businessCompany, nil
}
