package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
)

//// sendEmail To Customer with Order Confirmation
//func sendEmail(name string, email string, id any, id2 string) {
//	//Set up authentication information.
//	from := "sifhufhisg@gmail.com"
//	pass := "_ihfuhfis96_"
//	to := email
//	body := `
//		<!DOCTYPE html>
//		<html>
//		<head>
//		<style>
//		* {
//		  box-sizing: border-box;
//		}
//
//		.column {
//		  float: left;
//		  width: 50%;
//		  padding: 10px;
//		  height: 300px; /* Should be removed. Only for demonstration */
//		}
//
//		.row:after {
//		  content: "";
//		  display: table;
//		  clear: both;
//		}
//		</style>
//		</head>
//		<body>
//
//		<h2>Order Confirmation</h2>
//
//		<div class="row">
//		  <div class="column" style="background-color:#aaa;">
//			<h2>Order Details</h2>
//		  </div>
//		  <div class="column" style="background-color:#bbb;">
//			<h2>Shipping Details</h2>
//			<p>Address 1: 1234 Main St</p>
//			<p>Address 2: Apt 1</p>
//			<p>City: San Francisco</p>
//			<p>State: CA</p>
//			<p>Zipcode: 94111</p>
//			<p>Country: USA</p>
//		  </div>
//		</div>
//
//		</body>
//		</html>
//		`
//	msg := "From: " + from + "\n" +
//		"To: " + to + "\n" +
//		"Subject: Order Confirmation\n\n" +
//		body
//
//	err := smtp.SendMail("smtp.gmail.com:587",
//		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
//		from, []string{to}, []byte(msg))
//
//	if err != nil {
//		log.Printf("smtp error: %s", err)
//		return
//	}
//
//}

// Save Image Locally
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
