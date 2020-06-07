package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	db "github.com/AkezhanOb1/business-company/repository"
)


//CreateBusinessCompany is
func (*Server) BusinessCompanyImageDelete (ctx context.Context, request *pb.BusinessCompanyImageDeleteRequest) (*pb.BusinessCompanyImageDeleteResponse, error) {
	image, err := db.BusinessCompanyImageDeleteRepository(context.Background(), request.GetImageID())
	if err != nil {
		return nil, err
	}

	return image, err
}

