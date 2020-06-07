package repository

import (
	"context"
	pb "github.com/AkezhanOb1/business-company/api"
	"github.com/AkezhanOb1/business-company/config"

	"github.com/jackc/pgx/v4"
)

//DeleteBusinessCompanyOperationHourRepository is a repository that responsible for inserting data into the company
//table inside the database
func BusinessCompanyImageDeleteRepository(ctx context.Context, imageID int64) (*pb.BusinessCompanyImageDeleteResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())


	sqlQuery := `DELETE FROM business_company_image WHERE id=$1 RETURNING *;`

	row := conn.QueryRow(ctx, sqlQuery, imageID)

	var image pb.BusinessCompanyImageDeleteResponse

	err = row.Scan(
		&image.ImageID,
		&image.ImagePath,
	)


	return &image, nil
}
