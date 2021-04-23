package main

import (
	"esb/api"
	"esb/config"
	"esb/securitylog"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var SecurityLog *log.Logger

func init() {

	// Setup of database and tables if do not exist
	log.Println("Initializing Database")

	// Make Sql server connection
	db, err := config.GetMySQLDB("mysql")
	if err != nil {
		panic(err.Error())
	}

	// Close DB after the function has finished executing
	defer db.Close()

	res, err := db.Query("SHOW DATABASES")
	if err != nil {
		panic(err.Error())
	}

	var nameDb string
	found := false

	for res.Next() {
		res.Scan(&nameDb)
		if nameDb == config.DbName {
			found = true
			break
		}
	}

	if !found {

		log.Println("Creating goSchoolDb")
		_, err := db.Query("Create database goSchoolDb")
		if err != nil {
			panic(err.Error())
		}

		db.Close()

		db, err := config.GetMySQLDB(config.DbName)
		if err != nil {
			panic(err.Error())
		}

		// Close DB after the function has finished executing
		defer db.Close()

		_, err = db.Query("Create Table course (coursePk int(6) unsigned auto_increment primary key, title varchar(50) not null, insDateTime char(30) not null, updDateTime char(30) , insByFk int(6) not null, updByFk int(6)  , unique key title (title))")
		if err != nil {
			panic(err.Error())
		}
		log.Println("Courses Table Created in goSchoolDb")

		_, err = db.Query("Create Table user (userPk int(6) unsigned auto_increment primary key, userId char(100) not null, pswdHash  char(64) not null , UNIQUE KEY userId (userId))")
		if err != nil {
			panic(err.Error())
		}
		log.Println("User Table Created in goSchoolDb")

		

	}

	// Initiate Security Log
	securityLog := securitylog.Instance()

	// Log Initiating
	log.Println("Initiating Default Log")
	
	securityLog.Println("Initating Security Log")

}

// The Enterprise Service Bus (ESB) provides API services.
func main() {

	if len(os.Args) == 2 {
		api.Port = os.Args[1]
	} else {
		api.Port = "9001"
	}

	fmt.Println("Initializing System ....... please wait")

	api.HandleReq()

}
