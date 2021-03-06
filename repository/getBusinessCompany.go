package repository

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	"github.com/AkezhanOb1/business-company/config"
	"github.com/jackc/pgx/v4"
)

//GetBusinessCompanyRepository is a repository that responsible to all the requests to DB
//about business categories
func  GetBusinessCompanyRepository(ctx context.Context, companyID int64) (*pb.GetBusinessCompanyResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	sqlQuery := `SELECT id, name, category_id, address FROM business_company WHERE id=$1;`

	row := conn.QueryRow(ctx, sqlQuery, companyID)

	var company pb.BusinessCompany

	err = row.Scan(
		&company.BusinessCompanyID,
		&company.BusinessCompanyName,
		&company.BusinessCompanyCategoryID,
		&company.BusinessCompanyAddress,
	)

	if err != nil {
		return nil, err
	}

	return &pb.GetBusinessCompanyResponse{
		BusinessCompany: &company,
	}, nil
}
