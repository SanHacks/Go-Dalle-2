package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func runDb() {
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
    defer db.Close()

    // Execute SELECT statement
    rows, err := db.Query("SELECT * FROM generatedProducts")
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()

    // Print results
    for rows.Next() {
        var id int
        var name string
        var price float64
        if err := rows.Scan(&id, &name, &price); err != nil {
            panic(err.Error())
        }
        fmt.Printf("Product %d: %s ($%.2f)\n", id, name, price)
    }
    if err := rows.Err(); err != nil {
        panic(err.Error())
    }
}
