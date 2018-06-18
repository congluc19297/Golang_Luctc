package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test123"
	dbname   = "quotes"
)

type quoteInfo struct {
	Quote string
	Auth  string
	Kind  string
}

var db *sql.DB

func init() {
	// Creating the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Opening a connection to our database
	var err error
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

func readCSVFile(fileName string, setComma rune) [][]string {
	// read data from CSV file

	csvFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Comma = setComma

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return csvData
}

func createTableQuote() {

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS quote (
		id 		SERIAL PRIMARY KEY NOT NULL,
		quote	TEXT NOT NULL,
		auth 	TEXT NOT NULL,
		kind 	TEXT
	)`)

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully created table QUOTE!")
}

func insertDataToQuote(csvData [][]string) {
	lenData := len(csvData)
	fmt.Println("Inserting...")
	var i int
	for i = 0; i < lenData; i++ {
		_, err := db.Exec("INSERT INTO quote (QUOTE, AUTH, KIND) VALUES ($1, $2, $3)", csvData[i][0], csvData[i][1], csvData[i][2])
		if err != nil {
			panic(err)
		}
	}
	// for _, each := range csvData {
	// 	_, err := db.Exec("INSERT INTO quote (QUOTE, AUTH, KIND) VALUES ($1, $2, $3)", each[0], each[1], each[2])

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
}

func main() {
	csvData := readCSVFile("./quotes_all.csv", ';')

	createTableQuote()
	insertDataToQuote(csvData)
	db.Close()
}
