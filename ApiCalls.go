package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)
func GenerateImage(PromptIn string) string {
	url := "https://api.openai.com/v1/images/generations"
	method := "POST"

	//Fine Tune
	if PromptIn == "" {
		PromptIn = "T-Shirt Real 3D"
	} else {
		if strings.Contains(PromptIn, "T-Shirt") || strings.Contains(PromptIn, "t-shirt") || strings.Contains(PromptIn, "shirt") || strings.Contains(PromptIn, "Shirt") {
			PromptIn = PromptIn + " Real 3D"
		} else {
			PromptIn = PromptIn + " T-Shirt Real 3D"
		}
		log.Println("Prompt Sent: ", PromptIn)

		payload := strings.NewReader(`{
    "prompt": "` + PromptIn + `",
    "n": 3,
    "size": "512x512"
  }`)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return ""
		}
		req.Header.Add("Content-Type", "application/json")
		OpenAIkey := os.Getenv("OPENAIKEY")
		req.Header.Add("Authorization", "Bearer "+OpenAIkey)

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
		panic(err)
			}
		}(res.Body)

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		fmt.Println(string(body))

		return string(body)
	}
	return "OK"
}
