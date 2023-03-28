package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/mssola/user_agent"
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

	//
	setup()
	//Initiate Router
	router := mux.NewRouter()

	//router.Use(checkLogin)
	//Map the routes to the handlers in the backend api handler and
	//Handle 404 errors and redirect to /platform
	//Open Up routes for the frontend engine to use the backend api handler
	//Such as /backend/api/products (api) and /generatedProducts/ directory
	platformRouter(router)

	//Templates allocation to variables for the frontend engine
	platform, inventory, product, order, errorPage, orderSuccess,joinPlatform, loginPlatform := templates()

	templateHandler(platform, inventory, product, order, errorPage, orderSuccess, joinPlatform, loginPlatform)
	http.Handle("/", router)


	port := openPort()

	log.Printf("Listening on port %s", port)
	log.Printf("🚀🚀🚀🚀AIGEN🚀🚀🚀🚀")
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}

func isAuthenticated(r *http.Request) bool {
	// Get the session token from the cookie
	session, err := store.Get(r, "aigenID")
	if err != nil {
		log.Println("Error getting session:", err)
		return false
	}

	// Check if the "username" value is set in the session
	username, ok := session.Values["username"].(string)
	if !ok {
		log.Println("Username is not set in the session", username)
		return false
	}
	log.Println("Username:", session.Values["username"])
	log.Println("Session ID:", session.Values["token"])

	// Check if the session token is valid
	valid, err := isValidSessionToken(session.Values["username"].(string), session.Values["token"])
	if err != nil {
		log.Println("Error validating session token:", err)
		return false
	}

	return valid
}

func isValidSessionToken(username string, id interface{}) (bool, error) {
	// Get the session token from the database
	db, fail := dbPass()
	if fail != nil {
		return false, fail
	}
	defer db.Close()
	var token string
	err := db.QueryRow("SELECT token FROM sessions WHERE username = ?", username).Scan(&token)
	if err != nil {
		return false, err
	}

	// Check if the session token matches the one in the database
	if token == id {
		return true, nil
	}

	return false, nil

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
func templateHandler(platform, inventory, product, order, errorPage, orderSuccess, jointPlatform, loginPlatform *template.Template) {

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
		//Log Vistor Device
		ua := user_agent.New(r.UserAgent())
		
		fmt.Println(ua.OS())      // Output: Linux
		fmt.Println(ua.Browser()) // Output: Chrome
		
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
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		//IF THE REQUEST IS NOT A POST
		authCheckIn(w, r)

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

	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		authCheckIn(w, r)
		if r.Method != http.MethodPost {
			//Render the Home Page
			err := loginPlatform.Execute(w, nil)
			if err != nil {
				log.Println("Error in Rendering the Join Page")
			}
			return
		}

	})

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//Get the Form Values
	email := r.FormValue("email")
	password := r.FormValue("password")

	//Check if the User Exists
	db, err := dbPass()
	if err != nil {
		log.Println("Error in Connecting to the Database")
	}else {
		log.Println("Connected to the Database")
	}

	//Check if the User Exists
	userData, err := db.Query("SELECT id, email, password FROM storecustomers WHERE email = ?", email)
	if err != nil {
		log.Println("Error in User the Database, User Does Not Exist With the Email: ", email)
	}

	var users []Users
	for userData.Next() {
		var user Users

		err = userData.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			log.Println("Error in Scanning the Database, User Does Not Exist With the Email: ", email)
		}
		users = append(users, user)
		log.Println(users)

	}

	//Check if the User Exists
	if len(users) == 0 {
		log.Println("User Does Not Exist With the Email: ", email)
	}
	//decrypt and compare the password

	//Check if the Password is Correct
	if len(users) > 0 {
		err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(password))
		log.Println(err)
		if err != nil {
			log.Println("User Exists With the Email: ", email)
			log.Println("Password is Correct")
			//Create a Session
			session, _ := store.Get(r, "aiGenID")
			session.Values["email"] = email
			session.Save(r, w)
			//Redirect to the Home Page
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			log.Println("User Exists With the Email: ", email)
			log.Println("Password is Incorrect")
			http.Redirect(w, r, "/join?error=1", http.StatusSeeOther)
		}
	}

}

func logoutHandler(writer http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "username")
	session.Options.MaxAge = -1
	session.Save(request, writer)
	http.Redirect(writer, request, "/", http.StatusSeeOther)

}

func authCheckIn(w http.ResponseWriter, r *http.Request) {
	sessionError := isAuthenticated(r)
	log.Println("Session Error: ", sessionError)
	if sessionError != true {
		log.Println("Error in Authenticating Session")
	} else {
		log.Println("Authenticated Session")
		//Redirect to the Home Page
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

var store = sessions.NewCookieStore([]byte("aigenID"))

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	// Get the form values
	authCheckIn(w, r)
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

	// Generate a random session token
	randomBytes := make([]byte, 64)
	_, err = rand.Read(randomBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate a random session token
	tokenBytes := make([]byte, 64)
	_, err = rand.Read(tokenBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token := "aigen-" + base64.URLEncoding.EncodeToString(tokenBytes)

	// Save the session token to the database
	err = saveSessionToken(username, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new session for the user
	session, err := store.New(r, "aigenID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["username"] = username

	// Save the session token in a cookie
	session.Values["token"] = token
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Redirect the user to the homepage or some other page
	http.Redirect(w, r, "/join", http.StatusSeeOther)
}

func saveSessionToken(username string, token string) error {

	// Open the database
	db, err := dbPass()
	if err != nil {
		return err
	}

	// Save the session token to the database
	_, err = db.Exec("INSERT INTO sessions (username, token) VALUES (?, ?)", username, token)
	if err != nil {
		return err
	}

	return nil
}

