package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// "github.com/aws/aws-sdk-go/aws"
// "github.com/aws/aws-sdk-go/aws/credentials"
// "github.com/aws/aws-sdk-go/aws/session"
// "github.com/aws/aws-sdk-go/service/dynamodb"
// "github.com/vsouza/go-gin-boilerplate/config"

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/golang_poc")
	// defer db.Close()
	// c := config.GetConfig()
	// db = dynamodb.New(session.New(&aws.Config{
	// 	Region:      aws.String(c.GetString("db.region")),
	// 	Credentials: credentials.NewEnvCredentials(),
	// 	Endpoint:    aws.String(c.GetString("db.endpoint")),
	// 	DisableSSL:  aws.Bool(c.GetBool("db.disable_ssl")),

	// }))

}

func GetDB() *gorm.DB {
	return db
}
