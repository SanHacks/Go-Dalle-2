package main

import (
	"html/template"
)

// Point To The Front End Templates
func templates() *template.Template {
	platform := template.Must(template.ParseFiles("frontend/index.html"))
	return platform
}
