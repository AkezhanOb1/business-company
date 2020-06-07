package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	db "github.com/AkezhanOb1/business-company/repository"
)

//GetBusinessCompanies is
func (*Server) SearchBusinessCompany(ctx context.Context, request *pb.SearchBusinessCompanyRequest) (*pb.SearchBusinessCompanyResponse, error) {
	businessCompanies, err := db.SearchBusinessCompanyRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	for i, company := range businessCompanies.BusinessCompanies {
		businessCompanyImages, err := db.GetBusinessCompanyImagesRepository(context.Background(), company.GetBusinessCompanyID())
		if err != nil {
			return nil, err
		}

		businessCompanies.BusinessCompanies[i].BusinessCompanyImages = businessCompanyImages
	}

	return businessCompanies, nil
}
