package data_base

import (
	"github.com/alexperezortuno/go-auth/common/environment"
	"github.com/alexperezortuno/go-auth/internal/platform/storage/data_base/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Connection *gorm.DB

func Init(params environment.ServerValues) {
	Connection = ORM(params, StrConn(params))
	log.Println("Connection has been successfully")
}

func ORM(params environment.ServerValues, strConn string) *gorm.DB {
	//var connection *gorm.DB

	if params.EngineSql == "postgres" {
		Connection, _ = gorm.Open(postgres.Open(strConn), &gorm.Config{})
	}

	return Connection
}

func CloseConnection() {
	sqlDB, err := Connection.DB()

	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection close")
}

func StrConn(params environment.ServerValues) string {
	if params.EngineSql == "postgres" {
		return "host=" + params.DbHost +
			" port=" + params.DbPort +
			" user=" + params.DbUser +
			" dbname=" + params.DbName +
			" password=" + params.DbPass +
			" sslmode=disable" +
			" TimeZone=" + params.DbTimeZone
	}

	return ""
}

func Migrate() {
	migration.UserMigrate(Connection)
}

func Instance() *gorm.DB {
	return Connection
}
