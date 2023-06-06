package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)
 
func ConnectDB() (*sql.DB,error) {
	dbDriver := "mysql"
    dbUser := "root"
    dbPass := "ganesh"
    dbName := "project"
	
    // Open the database connection
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+ "?parseTime=true")
    if err != nil {
        log.Fatal(err)
    }
   

    // Ping the database to ensure the connection is established
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    // Connection established
    fmt.Println("Connected to the database!")
    return db,nil
    
}

// GetDB returns a database object
func GetDB() (*sql.DB) {
	db,err := ConnectDB()
	if err!=nil{
		log.Fatal(err)
	}
	return db
  }