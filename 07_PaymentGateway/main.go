package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	psqlInfo := os.Getenv("PSQLINFOR")

	// Opening a connection to our database
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func main() {

	pattern := os.Getenv("REGEXVCB")
	sms := os.Getenv("SMSVCB")

	mapMatch := GetMatchRegex(pattern, sms)
	fmt.Println("map: ", mapMatch)
	// greeting()
}

//GetMatchRegex ...
func GetMatchRegex(pattern string, sms string) map[string]string {

	r := regexp.MustCompile(pattern)
	valueMap := r.FindStringSubmatch(sms)
	keyMap := r.SubexpNames()

	mapMatch := make(map[string]string)
	for k, v := range keyMap {
		if v != "" {
			mapMatch[v] = valueMap[k]
		}
	}
	return mapMatch
}
