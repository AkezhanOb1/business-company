package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	db "github.com/AkezhanOb1/business-company/repository"
	"github.com/golang/protobuf/ptypes/empty"
)

//GetBusinessCompanies is
func (*Server) GetBusinessCompanies(ctx context.Context, emp *empty.Empty) (*pb.GetBusinessCompaniesResponse, error) {
	businessCompanies, err := db.GetBusinessCompaniesRepository(ctx)
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
