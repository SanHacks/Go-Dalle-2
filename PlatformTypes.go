package main

import "time"

type search struct {
	Time    time.Time
	QueryIn string
}

type GenerateImages struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}
