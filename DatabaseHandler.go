package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"strings"
)

// Store Order in Database Function
func storeOrderDB(sku string, name string, email string, phone string, address1 string, address2 string, city string, state string, zipcode string, country string, payment string) string {
	db, fail := dbPass()

	if fail != nil {
		panic(fail.Error())
	}
	StoreOrder, _ := db.Query("INSERT INTO orders (sku, name, email, phone, address_line_1, address_line_2, city, province, zipcode, country, payment_method) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", sku, name, email, phone, address1, address2, city, state, zipcode, country, payment)
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
		//If words in prompt are less than 3, use the first 3 words as the product name
		var productName = Prompt
		if len(strings.Split(Prompt, " ")) < 3 {
			//Prompt = strings.Join(strings.Split(Prompt, " ")[:3], " ")
		}

		//Determine price by number of words in prompt (1 word = $5, 2 words = $10, 3 words = $15, 4 words = $20, 5 words = $25)
		var price float64
		var wordCount = len(strings.Split(Prompt, " "))
		if wordCount <= 5 {
			price = 42.69
		} else if wordCount == 7 {
			price = 69.42
		} else if wordCount == 10 {
			price = 159.69
		} else if wordCount == 15 {
			price = 169.42
		} else if wordCount == 20 {
			price = 249.69
		} else {
			price = 99.42
		}

		prepare := "INSERT INTO generatedProducts (name, description, price, image, category, subcategory) VALUES (?, ?, ?, ?, ?, ?)"
		_, err := db.Exec(prepare, productName, Prompt, price, localImage, "Design", "Shirts")
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
	dbPassword := "Philemon70"
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
	db, err := sql.Open("mysql", connectionString)

	return db, err
}

func savePurchase(name string, address string, number string, address2 string, city string, state string, zip string, country string, payment string, sku string) int {
	db, err := dbPass()
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	prepare := "INSERT INTO orders (name, product_id, address_line_1, address_line_2, province, city, zipcode, country, payment_method, sku) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(prepare)
	if err != nil {
		log.Fatal(err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)
	_, err = stmt.Exec(name, sku, address, address2, state, city, zip, country, payment, sku)
	if err != nil {
		log.Fatal(err)
	}
	//Get Last Inserted ID
	var id int

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)

	if err != nil {
		log.Fatal(err)
	}
	return id
}


func saveUser(name string, email string, username string, password interface{}) error {
	db, err := dbPass()
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	prepare := "INSERT INTO storecustomers (name, email, username, password) VALUES (?, ?, ?, ?)"
	stmt, err := db.Prepare(prepare)
	if err != nil {
		log.Fatal(err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)
	_, err = stmt.Exec(name, email, username, password)
	if err != nil {
		log.Fatal(err)
	}
	//Get Last Inserted ID
	var id int

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)

	if err != nil {
		log.Fatal(err)
	}
	return nil

}

