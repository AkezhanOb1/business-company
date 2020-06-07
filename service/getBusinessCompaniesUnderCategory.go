package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	db "github.com/AkezhanOb1/business-company/repository"
)


//GetBusinessCompaniesUnderCategory is
func (*Server) GetBusinessCompaniesUnderCategory(ctx context.Context, request *pb.GetBusinessCompaniesUnderCategoryRequest) (*pb.GetBusinessCompaniesUnderCategoryResponse, error) {
	companies, err := db.GetBusinessCompaniesUnderCategoryRepository(ctx, request.GetCategoryID())
	if err != nil {
		return nil, err
	}

	for i, company := range companies.BusinessCompanies {
		businessCompanyImages, err := db.GetBusinessCompanyImagesRepository(context.Background(), company.GetBusinessCompanyID())
		if err != nil {
			return nil, err
		}

		companies.BusinessCompanies[i].BusinessCompanyImages = businessCompanyImages
	}

	return companies, nil
}
