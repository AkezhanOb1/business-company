package config

import "fmt"

//ServerAddress is
var ServerAddress = "0.0.0.0:50057"
//postgresAddress is the address of the postgres
//const postgresAddress = "127.0.0.1"
const postgresAddress = "46.101.138.224"

//postgresPort is the port of the postgres
const postgresPort = "5432"

//postgresDataBase is the name of the database
const postgresDataBase = "diploma"

//postgresUsername is the name of the user inside DBA
const postgresUsername = "postgres"

//postgresPassword is the password of the user
const postgresPassword = "postgres"

//PostgresConnection is the connection string to the database
var PostgresConnection = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	postgresAddress, postgresPort, postgresUsername, postgresPassword, postgresDataBase)

//AwsAccessKey is
var AwsAccessKey = "NLYUPR3DZFAJ62SPHTPQ"
//AwsSecretKey is
var AwsSecretKey = "p5xvehVAaR+Nmxkc5hxMvgHYJCmukN3eKpk+quwnZPA"
//AwsEndPoint is
var AwsEndPoint = "ams3.digitaloceanspaces.com"
//ImageSpaceName is
var ImageSpaceName = "qaqtus-images"
//DigitalOceanSpaceURL is
var DigitalOceanSpaceURL = "https://qaqtus-images.ams3.digitaloceanspaces.com/"

