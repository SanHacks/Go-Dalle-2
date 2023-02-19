package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "go/ast"
	_ "go/types"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	//Initiate Router
	router := mux.NewRouter()

	//Main Backend Routes
	platformRouter(router)

	//TEMPLATES
	platform, inventory, product, order := templates()

	templateHandler(platform, inventory, product, order)

	http.Handle("/", router)
	port := openPort()

	log.Printf("Listening on port %s", port)
	log.Printf("🚀🚀🚀🚀AIGEN🚀🚀🚀🚀")
	log.Printf("Open http://localhost:%s/platform in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}

// Opens Up The Port 8085, although it can get changed by the PORT env variable
func openPort() string {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8095"
		log.Printf("Enojy! %s", port)
	}

	return port
}

// Front End routes
func templateHandler(platform, inventory, product, order *template.Template) {

	//Landing Page
	http.HandleFunc("/platform", func(w http.ResponseWriter, r *http.Request) {
		//IF THE REQUEST IS NOT A POST
		if r.Method != http.MethodPost {
			//Render the Home Page
			err := platform.Execute(w, nil)
			if err != nil {
				log.Println("Error in Rendering the Platform Page")
			}
			return
		}

		//Handle the POST Request
		details := search{
			Time:    time.Time{},
			QueryIn: r.FormValue("search"),
		}
		// Details Collected From Request
		_ = details
		//Get Prompt from the Form
		var imageURL = GenerateImage(details.QueryIn)
		//Log entry of all the search queries
		log.Println("Prompt typed in: ", details.QueryIn)
		//display the image on the home page
		var imageOut GenerateImages
		getJsonValues := json.Unmarshal([]byte(imageURL), &imageOut)
		if getJsonValues != nil {
			log.Println("Error in Unmarshalling JSON")
		}
		//Store the Image in the Database
		var ImageID = storeImageDB(imageOut.Data[0].Url, details.QueryIn)

		err := platform.Execute(w, struct {
			Success  bool
			ImageURL any
			Search   string
			ID       int
		}{Success: true, ImageURL: imageOut.Data[0].Url, Search: details.QueryIn, ID: ImageID})
		if err != nil {
			return
		}
		log.Println("Image URL: ", imageURL)

	})

	//Get Product from proudct id and display in template product.html page for inventory management
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		//Get Product ID from URL
		id := r.URL.Query().Get("id")

		//Get Product Data from Database

		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/StorePlatform")
		if err != nil {
			panic(err.Error())
		}

		productsData, err := db.Query("SELECT id, name, description, price, image, category, subcategory FROM generatedProducts WHERE id = ?", id)
		//Display all the products in the database to .html page for inventory management
		var products []GeneratedProducts

		//Loop through products and get variants
		for productsData.Next() {
			var product GeneratedProducts

			err = productsData.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Image, &product.Category, &product.Subcategory)
			products = append(products, product)

		}
		log.Println("Product ID: ", id)
		log.Println("Product Data: ", products)
		product.Execute(w, struct {
			Success     bool
			ChannelData any
		}{Success: false, ChannelData: products})
		return
	})

	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("sku")
		/*		if r.Method == http.MethodPost && id == "" {
				//Get Values from the Form and store in the database to create an order
				//Get Product ID from URL
				//id := r.URL.Query().Get("id")

				sku := r.FormValue("sku")
				name := r.FormValue("name")
				email := r.FormValue("email")
				phone := r.FormValue("phone")
				address1 := r.FormValue("address1")
				//address2 := r.FormValue("address2")
				city := r.FormValue("city")
				state := r.FormValue("state")
				zipcode := r.FormValue("zipcode")
				country := r.FormValue("country")
				payment := r.FormValue("payment")

				//Check if all the fields are filled
				if sku == "" || name == "" || email == "" || phone == "" || address1 == "" || city == "" || state == "" || zipcode == "" || country == "" || payment == "" {
					log.Println("All fields are required")
					order.Execute(w, struct {
						Success     bool
						ChannelData any
					}{Success: false, ChannelData: "All fields are required"})

					//Reirect to the order Thank you page
					http.Redirect(w, r, "/order", http.StatusSeeOther)
					return
				}

				if r.Method == http.MethodPost {
					//Get Values from the Form and store in the database to create an order
					//Get Product ID from URL
					id := r.URL.Query().Get("id")

					sku := r.FormValue("sku")
					name := r.FormValue("name")
					email := r.FormValue("email")
					phone := r.FormValue("phone")
					address1 := r.FormValue("address1")
					address2 := r.FormValue("address2")
					city := r.FormValue("city")
					state := r.FormValue("state")
					zipcode := r.FormValue("zipcode")
					country := r.FormValue("country")
					payment := r.FormValue("payment")

					//Check if all the fields are filled
					if sku == "" || name == "" || email == "" || phone == "" || address1 == "" || city == "" || state == "" || zipcode == "" || country == "" || payment == "" {
						log.Println("All fields are required")
						order.Execute(w, struct {
							Success     bool
							ChannelData any
						}{Success: false, ChannelData: "All fields are required"})
						return
					}

					//Store Order in the database
					var orderID = storeOrderDB(sku, name, email, phone, address1, address2, city, state, zipcode, country, payment)
					log.Println("Order ID: ", orderID)
					//Send email to customer
					sendEmail(name, email, orderID, id)
					//Reirect to the order Thank you page
					http.Redirect(w, r, "/CompleteOrder?id="+orderID, http.StatusSeeOther)
					return

				}
			}*/

		//If GET Request is made to the order page
		if r.Method == http.MethodPost {
			db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/StorePlatform")
			if err != nil {
				panic(err.Error())
			}

			productsData, err := db.Query("SELECT id, name, description, price, image, category, subcategory FROM generatedProducts WHERE id = ?", id)
			//Display all the products in the database to .html page for inventory management
			var products []GeneratedProducts

			//Loop through products and get variants
			for productsData.Next() {
				var product GeneratedProducts

				err = productsData.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Image, &product.Category, &product.Subcategory)
				products = append(products, product)

			}
			log.Println("Product ID: ", id)
			log.Println("Product Data: ", products)
			order.Execute(w, struct {
				Success     bool
				ChannelData any
			}{Success: false, ChannelData: products})
			return
		}

	})
}
