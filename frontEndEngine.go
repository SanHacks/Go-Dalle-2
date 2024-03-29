package main

import (
	"html/template"
)

// Point To The Front End Templates
func templates() (*template.Template, *template.Template, *template.Template, *template.Template, *template.Template,
	*template.Template, *template.Template, *template.Template) {
	platform := template.Must(template.ParseFiles("frontend/index.html"))
	inventory := template.Must(template.ParseFiles("frontend/inventory.html"))
	product := template.Must(template.ParseFiles("frontend/product.html"))
	order := template.Must(template.ParseFiles("frontend/purchase.html"))
	errorPage := template.Must(template.ParseFiles("frontend/404.html"))
	orderSuccess := template.Must(template.ParseFiles("frontend/orderSuccess.html"))
	joinPlatform := template.Must(template.ParseFiles("frontend/signup.html"))
	loginPlatform := template.Must(template.ParseFiles("frontend/login.html"))
	return platform, inventory, product, order, errorPage, orderSuccess, joinPlatform, loginPlatform
}
