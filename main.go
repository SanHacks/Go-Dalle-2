package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "go/ast"
	_ "go/types"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	//Initiate Router
	router := mux.NewRouter()
	//router.Use(checkLogin)
	//Map the routes to the handlers in the backend api handler and
	//Handle 404 errors and redirect to /platform
	//Open Up routes for the frontend engine to use the backend api handler
	//Such as /backend/api/products (api) and /generatedProducts/ directory
	platformRouter(router)

	//Templates allocation to variables for the frontend engine
	platform, inventory, product, order, errorPage, orderSuccess,joinPlatform := templates()

	templateHandler(platform, inventory, product, order, errorPage, orderSuccess, joinPlatform)
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
func templateHandler(platform, inventory, product, order, errorPage, orderSuccess, jointPlatform *template.Template) {

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
	http.HandleFunc("/signup", signUpHandler)
	//http.HandleFunc("/login", loginHandler)

	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		//IF THE REQUEST IS NOT A POST
		if r.Method != http.MethodPost {
			//Render the Home Page
			err := jointPlatform.Execute(w, nil)
			if err != nil {
				log.Println("Error in Rendering the Join Page")
			}
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
		}{Success: false, ChannelData: id})

	})

}


func signUpHandler(w http.ResponseWriter, r *http.Request) {
	// Get the form values
	name := r.FormValue("name")
	email := r.FormValue("email")
	username := r.FormValue("email")
	password := r.FormValue("password")

	// Hash the password for security
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	log.Println(hashedPassword)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Save the user to the database
	err = saveUser(name, email, username, hashedPassword)
	if err != nil {
		http.Error(w, "Error saving user to database", http.StatusInternalServerError)
		return
	}

	// Redirect the user to the login page
	//user credentials are saved to the database to log user in automatically
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

