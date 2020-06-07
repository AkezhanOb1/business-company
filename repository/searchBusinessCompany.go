package repository

import (
	pb "github.com/AkezhanOb1/business-company/api"
	"github.com/AkezhanOb1/business-company/config"
	"github.com/jackc/pgx/v4"
	"context"

)


//SearchBusinessCompanyRepository is a repository that responsible to all the requests to DB
//about business categories
func SearchBusinessCompanyRepository(ctx context.Context, companyName string) (*pb.SearchBusinessCompanyResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)


	companyName = companyName + "%"
	sqlQuery := `SELECT id, name, category_id, address 
					FROM business_company
				 WHERE LOWER(name) LIKE LOWER($1);`

	rows, err := conn.Query(ctx, sqlQuery, companyName)
	if err != nil {
		return nil, err
	}

	var companies pb.SearchBusinessCompanyResponse

	for rows.Next() {
		var company pb.BusinessCompany

		err = rows.Scan(
			&company.BusinessCompanyID,
			&company.BusinessCompanyName,
			&company.BusinessCompanyCategoryID,
			&company.BusinessCompanyAddress,
		)

		if err != nil {
			return nil, err
		}

		companies.BusinessCompanies = append(companies.BusinessCompanies, &company)
	}

	return &companies, nil
}
