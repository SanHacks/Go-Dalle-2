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
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
)

func main() {

	//Initiate Router
	router := mux.NewRouter()

	//Main Backend Routes
	platformRouter(router)

	//TEMPLATES
	platform, inventory, product := templates()

	templateHandler(platform, inventory, product)

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
func templateHandler(platform, inventory, product *template.Template) {

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

}

func storeImageDB(Image string, Prompt string) int {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/StorePlatform")
	if err != nil {
		log.Println("Error in Connecting to Database")
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
	return 0
}

func saveImageLocally(imageURL string) string {

	// Download the image
	resp, err := http.Get(imageURL)
	if err != nil {
		fmt.Println("Error downloading image:", err)
		return ""
	}
	defer resp.Body.Close()

	// Generate a random 5-letter string
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 5)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	randomString := string(b)

	// Create the file and write the image data to it
	fname := path.Base(imageURL)
	fname = "generatedProducts/AIGEN_" + randomString + ".jpg"
	f, err := os.Create(fname)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return ""
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return ""
	}

	fmt.Println("Image downloaded and saved as", fname)
	return fname
}
