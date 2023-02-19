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
	"path"
	"time"
)

func main() {

	//Initiate Router
	router := mux.NewRouter()

	//Main Backend Routes
	platformRouter(router)

	//TEMPLATES
	platform, inventory := templates()

	templateHandler(platform, inventory)

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
func templateHandler(platform, inventory *template.Template) {

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

}

func storeImageDB(Image string, Prompt string) int {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/StorePlatform")
	if err != nil {
		log.Println("Error in Connecting to Database")
	} else {
		log.Println("Connected to Database")
		//Save Image To Local Database for future use
		/*var imagePaths =*/
		saveImageLocally(Image)

		prepare := "INSERT INTO generatedProducts (name, description, price, image, category, subcategory) VALUES (?, ?, ?, ?, ?, ?)"
		_, err := db.Exec(prepare, "Test", Prompt, 0.00, Image, "Design", "Shirts")
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

func saveImageLocally(image string) string {

	r, err := http.Get(image)

	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	fname := path.Base(image)
	fname = time.Now().Format("20060102150405") + fname
	fname = fname + ".jpg"
	//give the image a name
	fname = "generatedProducts/" + fname

	f, err := os.Create(fname)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.ReadFrom(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("image downloaded")
	return fname
}
