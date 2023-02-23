package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func runDb() (*sql.DB) {
    // Set up database connection string
    dbUser := "ndiGundoSan"
    dbPass := "@Sifhufhi2024"
    dbHost := "aigen.mysql.database.azure.com"
    dbPort := "3306"
    dbName := "aigen"
    dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=true", dbUser, dbPass, dbHost, dbPort, dbName)

    // Connect to database
    db, err := sql.Open("mysql", dbURI)
    if err != nil {
        panic(err.Error())
    }
	// Test connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}else{
		fmt.Println("Connected to database")

        //Query 
        rows, err := db.Query("SELECT * FROM generatedProducts")
        if err != nil {
            panic(err.Error())
        }
        defer rows.Close()

        //Loop through rows
        for rows.Next() {
            var id int
            var name string
            var description string

            err = rows.Scan(&id, &name, &description)
            if err != nil {
                panic(err.Error())
            }
            fmt.Println(id, name, description)
        }

        //Get error
        err = rows.Err()
        if err != nil {
            panic(err.Error())
        }


	}

	return db
}
