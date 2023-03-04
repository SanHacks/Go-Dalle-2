package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Map the routes to the handlers in the backend api handler
// Handle 404 errors and redirect to /platform
func platformRouter(router *mux.Router) {
	router.NotFoundHandler = routeNotFoundError()
	//Backend API Routes Definition (API) (ie: open route and map to handler)
	router.HandleFunc("/backend/api/products", ProductsHandler).Methods("GET")

	//Submission Handlers (ie: open route and map to handler)
	router.HandleFunc("/checkout", CheckoutHandler).Methods("POST")
	router.HandleFunc("/postsignup", SignupHandler).Methods("POST")
	router.HandleFunc("/postlogin", LoginHandler).Methods("POST")
	router.HandleFunc("/postlogout", LogoutHandler).Methods("POST")
	router.HandleFunc("/postAdminLogin", AdminLoginHandler).Methods("POST")
	router.HandleFunc("/postAdminLogout", AdminLogoutHandler).Methods("POST")
	router.HandleFunc("/postAdminSignup", AdminSignupHandler).Methods("POST")

	//File Server for Platform CDN and Frontend Engine Resources (CSS, JS, Images, Fonts and stuff)
	router.PathPrefix("/generatedProducts/").Handler(http.StripPrefix("/generatedProducts/", http.FileServer(http.Dir("./generatedProducts/"))))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./frontend/assets/css/"))))
	router.PathPrefix("/assets/css/").Handler(http.StripPrefix("/assets/css/", http.FileServer(http.Dir("./frontend/assets/css/"))))
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./frontend/assets/img/"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/frontend/assets/js/", http.FileServer(http.Dir("./frontend/assets/js/"))))
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images /", http.FileServer(http.Dir("./frontend/assets/img/"))))
	router.PathPrefix("/frontend/assets/fonts/").Handler(http.StripPrefix("/frontend/assets/fonts/", http.FileServer(http.Dir("./frontend/assets/fonts/"))))
	router.PathPrefix("/frontend/assets/").Handler(http.StripPrefix("/frontend/assets/", http.FileServer(http.Dir("./frontend/assets/"))))
	router.PathPrefix("/frontend/").Handler(http.StripPrefix("/frontend/", http.FileServer(http.Dir("./frontend/"))))
}

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		emailAddress := r.FormValue("email")
		phoneNumber := r.FormValue("phone")
		address := r.FormValue("address1")
		address2 := r.FormValue("address2")
		city := r.FormValue("city")
		state := r.FormValue("state")
		zip := r.FormValue("zipcode")
		country := r.FormValue("country")
		payment := r.FormValue("payment")
		sku := r.FormValue("sku")

		if sku == "" {
			sku = name
		}

		if name == "" || emailAddress == "" || phoneNumber == "" || address == "" || city == "" || state == "" || zip ==
			"" || country == "" || payment == "" || sku == "" {
			log.Println("Missing Fields")
			http.Redirect(w, r, "/product?id="+sku+"&error=1", http.StatusSeeOther)

		} else {

			saveOrder := savePurchase(name, address, phoneNumber, address2, city, state, zip, country, payment, sku)

			if saveOrder == true {
				log.Println("Order Saved")
				http.Redirect(w, r, "/succesfulOrder", http.StatusSeeOther)

			} else {

				log.Println("Order Not Saved")
				http.Redirect(w, r, "/product?id="+sku+"&error=2", http.StatusSeeOther)
			}
		}
	}
}

func AdminSignupHandler(w http.ResponseWriter, r *http.Request) {

}

func AdminLogoutHandler(w http.ResponseWriter, r *http.Request) {

}

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {

}
