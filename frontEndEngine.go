package main

import (
	"html/template"
)

// Point To The Front End Templates
func templates() (*template.Template, *template.Template) {
	platform := template.Must(template.ParseFiles("frontend/index.html"))
	inventory := template.Must(template.ParseFiles("frontend/inventory.html"))
	return platform, inventory
}
