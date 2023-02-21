package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// ProductsGenerator API Generator for products
func ProductsGenerator(w http.ResponseWriter) any {

	db, err := sql.Open("mysql", "ndiGundoSan:@Sifhufhi2024@tcp(aigen.mysql.database.azure.com:3306)/aigen")
	if err != nil {
		panic(err.Error())
	}

	productsData, err := db.Query("SELECT id, name, description, price, image, category, subcategory FROM generatedProducts ORDER BY id DESC")
	//Display all the products in the database to .html page for inventory management
	var products []GeneratedProducts

	//Loop through products and get variants
	for productsData.Next() {
		var product GeneratedProducts

		err = productsData.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Image, &product.Category, &product.Subcategory)
		products = append(products, product)

	}
	//add headers to the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusCreated)
	productsReadyPush := json.NewEncoder(w).Encode(products)

	log.Println(productsReadyPush)
	return productsReadyPush
}
