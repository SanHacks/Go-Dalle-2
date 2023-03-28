package main

import (
	"log"
	"os"
)



//CONTAINS SENSITIVE INFORMATION - DO NOT SHARE WITH ANYONE ELSE
// TODO: REMOVE THIS FILE FROM GIT REPO OR MAKE IT PRIVATE
// COMMITTING THIS FILE IS CAREER SUICIDE, YOU DON'T NEED THAT STRESS.
func setup()  {
	err := os.Setenv("OPENAIKEY", "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")
	if err != nil {
		log.Printf("Error setting environment variable: %v", err)
	}
	
	os.Setenv("DB_USER", "ndiGundoSan")
	
	os.Setenv("DB_PASS", "Philemon70")
	os.Setenv("DB_HOST", "aigen.mysql.database.azure.com")
	os.Setenv("DB_NAME", "aigen")

}