package main

import "time"

type search struct {
	Time    time.Time
	QueryIn string
}

// Gen
type GenerateImages struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}
type Products struct {
	id          int
	name        string
	description string
	price       string
	image       string
	category    string
	subcategory string
	created_at  string
	updated_at  string
}

type MetaData struct {
	ParentSku string `json:"parent_sku"`
}

type GeneratedProducts struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
	Subcategory string  `json:"subcategory"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
