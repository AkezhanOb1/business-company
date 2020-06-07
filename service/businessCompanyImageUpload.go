package service

import (
	"bytes"
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	"github.com/AkezhanOb1/business-company/config"
	db "github.com/AkezhanOb1/business-company/repository"
	"github.com/minio/minio-go"
	"log"
	"strings"
)


//CreateBusinessCompany is
func (*Server) BusinessCompanyImageUpload(ctx context.Context, request *pb.BusinessCompanyImageUploadRequest) (*pb.BusinessCompanyImageUploadResponse, error) {
	company, err := db.GetBusinessCompanyRepository(context.Background(), request.BusinessCompanyID)
	if err != nil {
		return nil, err
	}

	request.Image.FileName = company.BusinessCompany.GetBusinessCompanyName() + "-" + request.Image.FileName
	request.Image.FileName = strings.Replace(request.Image.FileName, " ", "", -1)

	client, err := minio.New(config.AwsEndPoint, config.AwsAccessKey, config.AwsSecretKey, true)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.PutObject(
		config.ImageSpaceName,
		request.Image.FileName,
		bytes.NewReader(request.Image.File),
		request.Image.Size,
		minio.PutObjectOptions{
			UserMetadata: 			 map[string]string{"x-amz-acl": "public-read"},
			ContentType:             request.Image.ContentType,
		},
	)


	if err != nil {
		return nil, err
	}


	actualPath := config.DigitalOceanSpaceURL + request.Image.FileName
	id, err := db.UploadBusinessCompanyImageRepository(ctx, actualPath, request.BusinessCompanyID)
	if err != nil {
		return nil, err
	}


	return &pb.BusinessCompanyImageUploadResponse{
		ImageID: *id,
	}, err
}

