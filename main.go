package main

import (
	"encoding/json"
	"fmt"
	_ "go/ast"
	_ "go/types"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	//Initiate Router
	router := mux.NewRouter()

	//Map the routes to the handlers in the backend api handler and
	//Handle 404 errors and redirect to /platform
	//Open Up routes for the frontend engine to use the backend api handler
	//Such as /backend/api/products (api) and /generatedProducts/ directory
	platformRouter(router)

	//Templates allocation to variables for the frontend engine
	platform, inventory, product, order, errorPage, orderSuccess := templates()

	templateHandler(platform, inventory, product, order, errorPage, orderSuccess)
	http.Handle("/", router)

	port := openPort()

	log.Printf("Listening on port %s", port)
	log.Printf("🚀🚀🚀🚀AIGEN🚀🚀🚀🚀")
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}

func routeNotFoundError() http.Handler {
	//Redirect to /platform
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//if is to / then redirect to /platform else redirect if is route not found just go to 404
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/platform", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/404", http.StatusSeeOther)
		}
	})
}

// Opens Up The Port 8080, although it can get changed by the PORT env variable
func openPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Enojy! %s", port)
	}
	return port
}

// Front End routes
func templateHandler(platform, inventory, product, order, errorPage, orderSuccess *template.Template) {

	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		err := errorPage.Execute(w, nil)
		if err != nil {
			log.Println("Error in Rendering the 404 Page")
		}
		return
	})

	//HandleFunc() the Platform Page
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

		details := search{
			Time:    time.Time{},
			QueryIn: r.FormValue("search"),
		}

		_ = details

		var imageURL = GenerateImage(details.QueryIn)
		log.Println("Prompt typed in: ", details.QueryIn)
		var imageOut GenerateImages
		getJsonValues := json.Unmarshal([]byte(imageURL), &imageOut)
		if getJsonValues != nil {
			log.Println("Error in Unmarshalling JSON")
		}

		var imageUrls []string
		for i, image := range imageOut.Data {
			fmt.Printf("Image %d: %s\n", i+1, image.Url)
			imageUrls = append(imageUrls, image.Url)
			// Do something with the image URL, such as download or display the image
		}

		//Store the Images in the Database
		var ImageIDs []int
		for _, url := range imageUrls {
			var ImageID = storeImageDB(url, details.QueryIn)
			ImageIDs = append(ImageIDs, ImageID)
		}

		err := platform.Execute(w, struct {
			Success  bool
			ImageURL []string
			Search   string
			IDs      []int
		}{Success: true, ImageURL: imageUrls, Search: details.QueryIn, IDs: ImageIDs})

		if err != nil {
			return
		}
		log.Println("Image URLs: ", imageUrls)
	})

	//Get Product from product id and display in template product.html page for inventory management
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		//Get Product ID from URL
		id := r.URL.Query().Get("id")

		db, err := dbPass()
		if err != nil {
			panic(err.Error())
		}

		productsData, err := db.Query("SELECT id, name, description, price, image, category, subcategory FROM generatedProducts WHERE id = ?", id)
		//Display all the products in the database to .html page for inventory management
		if err != nil {
			panic(err.Error())
		}
		var products []GeneratedProducts

		//Loop through products and get variants
		for productsData.Next() {
			var product GeneratedProducts

			err = productsData.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Image, &product.Category, &product.Subcategory)
			if err != nil {
				panic(err.Error())
			}
			products = append(products, product)

		}
		//Display the product in the template
		log.Println("Product ID: ", id)
		log.Println("Product Data: ", products)
		product.Execute(w, struct {
			Success     bool
			ChannelData any
		}{Success: false, ChannelData: products})

	})

	//Get Product from product id and display in template product.html page for inventory management
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
			db, fail := dbPass()
			if fail != nil {
				panic(fail.Error())
			}

			productsData, err := db.Query("SELECT id, name, description, price, image, category, subcategory FROM generatedProducts WHERE id = ?", id)
			//Display all the products in the database to .html page for inventory management
			if err != nil {
				log.Panic(err.Error())
			}
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

	http.HandleFunc("/CompleteOrder", func(w http.ResponseWriter, r *http.Request) {
		//Get Order ID from URL
		id := r.URL.Query().Get("id")
		log.Println("Order ID: ", id)
		orderSuccess.Execute(w, struct {
			Success     bool
			ChannelData any
		}{Success: true, ChannelData: id})

	})

}
