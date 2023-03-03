package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"io/ioutil"
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

// Store Image in Database Function

func storeImageDB(Image string, Prompt string) int {

	db, err := dbPass()
	if err != nil {
		log.Println("Error in Connecting to Database")
		panic(err.Error())
	} else {
		log.Println("Connected to Database")
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
	dbPassword := "@Sifhufhi2024"
	dbHost := "aigen.mysql.database.azure.com"
	dbName := "aigen"

	rootCertPool := x509.NewCertPool()
	pem, _ := ioutil.ReadFile(CaCertPath)
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}
	mysql.RegisterTLSConfig("custom", &tls.Config{RootCAs: rootCertPool})
	var connectionString string
	connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true&tls=custom", dbUser, dbPassword, dbHost, dbName)
	//dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=true&tlsca=%s", dbUser, dbPassword, dbHost, dbPort, dbName, CaCertPath)

	db, err := sql.Open("mysql", connectionString)
	return db, err
}

//db, err := sql.Open("mysql", "ndiGundoSan:{your_password}@tcp
//(aigen.mysql.database.azure.com:3306)/{your_database}?tls=custom&tls-ca={ca-cert filename}")
