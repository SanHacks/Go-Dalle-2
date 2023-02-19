package main

import "log"

// Logs Products Information to the Console for Debugging Purposes Only::API
func PlatformLogger(product Products) {
	log.Println("Name: ", product.name)
	log.Println("Price: ", product.price)
	log.Println("Image: ", product.image)
	log.Println("Description: ", product.description)
	log.Println("Created At: ", product.created_at)
	log.Println("Updated At: ", product.updated_at)
	log.Println("Category: ", product.category)
	log.Println("Subcategory: ", product.subcategory)
}
