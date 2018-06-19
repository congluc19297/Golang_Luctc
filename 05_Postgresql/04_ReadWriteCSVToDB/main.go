package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"time"

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

	// _, err := db.Exec(`CREATE TABLE IF NOT EXISTS quote (
	// 	id 		SERIAL PRIMARY KEY NOT NULL,
	// 	quote	TEXT NOT NULL,
	// 	auth 	TEXT NOT NULL,
	// 	kind 	TEXT
	// )`)
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS quote (
		quote	TEXT NOT NULL,
		auth 	TEXT PRIMARY KEY NOT NULL,
		kind 	TEXT NOT NULL
	)`)

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully created table QUOTE!")
}

func insertDataToQuote(csvData [][]string, fragNum int) {
	lenData := len(csvData)
	fmt.Println("Inserting...")
	var i int
	var fragment int
	start := time.Now()
	fmt.Println(start)

	wg := sync.WaitGroup{}

	wg.Add(fragNum)
	fragment = lenData / fragNum
	for i = 0; i < fragNum; i++ {
		go func(idx int) {
			defer wg.Done()
			start := fragment * idx
			end := fragment * (idx + 1)
			if idx == (fragNum - 1) {
				end = lenData
			}
			for j := start; j < end; j++ {
				_, _ = db.Exec("INSERT INTO quote (QUOTE, AUTH, KIND) VALUES ($1, $2, $3)", csvData[j][0], csvData[j][1], csvData[j][2])
			}

		}(i)
	}
	// fmt.Println(fragment)
	/*for i = 0; i < 100; i++ {
		_, err := db.Exec("INSERT INTO quote (QUOTE, AUTH, KIND) VALUES ($1, $2, $3)", csvData[i][0], csvData[i][1], csvData[i][2])
		if err != nil {
			panic(err)
		}
	}*/
	wg.Wait()
	end := time.Now()
	fmt.Println(end)

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
	insertDataToQuote(csvData, 90)

	/*50 goroutines
	2018-06-19 14:32:28.539460984 +0700 +07 m=+0.135548243
	2018-06-19 14:33:18.980734256 +0700 +07 m=+50.576821457*/

	/*60 goroutines
	2018-06-19 14:35:34.017387773 +0700 +07 m=+0.037819768
	2018-06-19 14:36:10.819939037 +0700 +07 m=+36.840371200*/

	/*70 goroutines
	2018-06-19 14:36:40.21244279 +0700 +07 m=+0.040360130
	2018-06-19 14:37:15.471842468 +0700 +07 m=+35.299759986*/

	/*80 goroutines
	2018-06-19 14:41:40.523446589 +0700 +07 m=+0.127882018
	2018-06-19 14:42:07.619207725 +0700 +07 m=+27.223643151*/

	/*90 goroutines
	2018-06-19 14:43:33.279647591 +0700 +07 m=+0.038205097
	2018-06-19 14:43:55.139216944 +0700 +07 m=+21.897774611*/
	db.Close()
}
