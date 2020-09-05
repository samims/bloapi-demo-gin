package config

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/orm"
)

var db orm.Ormer

// GetDb : It will returns orm object
func GetDB() orm.Ormer {
	if db == nil {
		dbHost := os.Getenv("POSTGRES_HOST")
		dbName := os.Getenv("POSTGRES_DB")
		dbUser := os.Getenv("POSTGRES_USER")
		dbPassword := os.Getenv("POSTGRES_PASSWORD")
		dbPort := os.Getenv("POSTGRES_PORT")

		if dbHost == "" {
			fmt.Println("Environment variable DB_HOST is null.")
			return nil
		}
		if dbName == "" {
			fmt.Println("Environment variable DB_NAME is null.")
			return nil
		}
		if dbUser == "" {
			fmt.Println("Environment variable DB_USERNAME is null.")
			return nil
		}
		if dbPassword == "" {
			fmt.Println("Environment variable DB_PASSWORD is null.")
			return nil
		}

		if dbPort == "" {
			dbPort = "5432"
		}

		orm.RegisterDriver("postgres", orm.DRPostgres)
		orm.RegisterDataBase("default", "postgres", "postgres://"+dbUser+":"+dbPassword+"@"+dbHost+":"+dbPort+"/"+dbName+"?sslmode=disable")

		// This is for auto generating tables
		orm.RunSyncdb("default", false, true)

		db = orm.NewOrm()
		db.Using("default")

	}
	return db
}
