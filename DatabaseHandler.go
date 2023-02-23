package main

import (
	"database/sql"
	"fmt"
	"log"
)

// Store Order in Database Function
func storeOrderDB(sku string, name string, email string, phone string, address1 string, address2 string, city string, state string, zipcode string, country string, payment string) string {
	db, fail := dbPass()

	if fail != nil {
		panic(fail.Error())
	}
	StoreOrder, _ := db.Query("INSERT INTO orders (sku, name, email, phone, address1, address2, city, state, zipcode, country, payment) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", sku, name, email, phone, address1, address2, city, state, zipcode, country, payment)
	if StoreOrder != nil {
		log.Println("Error in Storing Order in Database")
	} else {
		log.Println("Stored Order in Database")
	}
	//Get ID of Order that was just created
	var orderID string
	connector := db.QueryRow("SELECT id FROM orders ORDER BY id DESC LIMIT 1").Scan(&orderID)
	if connector != nil {
		log.Println("Error in Getting Last Order ID")
	} else {
		log.Println("Got Last Order ID")
	}
	return orderID
}

func storeImageDB(Image string, Prompt string) int {

   // Set up database connection string
    // Connect to database
    db, err := dbPass()

    if err != nil {
		log.Println("Error in Connecting to Database")
        panic(err.Error())
	} else {
		log.Println("Connected to Database")
		//Save Image To Local Database for future use
		/*var imagePaths =*/
		var localImage = saveImageLocally(Image)

		prepare := "INSERT INTO generatedProducts (name, description, price, image, category, subcategory) VALUES (?, ?, ?, ?, ?, ?)"
		_, err := db.Exec(prepare, "Test", Prompt, 0.00, localImage, "Design", "Shirts")
		if err != nil {
			log.Println("Error in Storing Image in Database")
		} else {
			log.Println("Stored Image in Database")
		}
		//Get Last Inserted ID
		var id int

		err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)
		
		if err != nil {
			log.Println("Error in Getting Last Insert ID")
		} else {
			log.Println("Last Insert ID: ", id)
		}
		return id
	}
}

func dbPass() (*sql.DB, error) {
	dbUser := "ndiGundoSan"
	dbPass := "@Sifhufhi2024"
	dbHost := "aigen.mysql.database.azure.com"
	dbPort := "3306"
	dbName := "aigen"
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=true", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbURI)
	return db, err
}
