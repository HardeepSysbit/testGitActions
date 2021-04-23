package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Container2Container = false

const (
	DbDriver = "mysql"
	DbUser   = "root"
	DbPass   = "mypassword"
	DbName   = "goSchoolDb"
	//DbHost2Container = "@tcp(192.168.2.1:3306)"
	DbHost2Container = "@tcp(goSchoolSql:3306)"

	DbContainer2Container = "@tcp(172.28.5.0:3307)"
)

func GetMySQLDB(dbName string) (db *sql.DB, err error) {

	if !Container2Container {
		log.Printf("Host 2 Container: %s ", DbHost2Container)
		db, err = sql.Open(DbDriver, DbUser+":"+DbPass+DbHost2Container+"/"+dbName)
	} else {
		log.Printf("Host Container 2 Container: %s ", DbContainer2Container)
		db, err = sql.Open(DbDriver, DbUser+":"+DbContainer2Container+"/"+dbName)
	}

	return
	
}
