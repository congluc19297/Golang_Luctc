package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"sync"

	"github.com/lib/pq"
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
	ID       int
	Quote    string
	Auth     string
	Category string
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
		author 	TEXT UNIQUE NOT NULL,
		category 	TEXT
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
	var fragmentVal int

	wg := sync.WaitGroup{}

	wg.Add(fragNum)
	fragmentVal = lenData / fragNum
	for i = 0; i < fragNum; i++ {
		go devideFragment(i, fragmentVal, fragNum, csvData, lenData, &wg)
	}
	wg.Wait()
}

func devideFragment(idx int, fragmentVal int, fragNum int, csvData [][]string, lenData int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := fragmentVal * idx
	end := fragmentVal * (idx + 1)
	if idx == (fragNum - 1) {
		end = lenData
	}
	for j := start; j < end; j++ {
		_, err := db.Exec("INSERT INTO quote (quote, author, category) VALUES ($1, $2, $3)", csvData[j][0], csvData[j][1], csvData[j][2])
		if err != nil {
			pqErr, ok := err.(*pq.Error)
			if ok {
				// fmt.Println(pqErr.Code.Name())
				if pqErr.Code.Name() == "unique_violation" {
					continue
				} else {
					panic(err)
				}
			}
		}
	}
}

func getAllQuote() []quoteInfo {
	rows, err := db.Query("SELECT * FROM quote order by id;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	quotes := make([]quoteInfo, 0)
	for rows.Next() {
		quote := quoteInfo{}
		err := rows.Scan(&quote.ID, &quote.Quote, &quote.Auth, &quote.Category)
		if err != nil {
			panic(err)
		}
		quotes = append(quotes, quote)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return quotes
}

func main() {
	csvData := readCSVFile("./quotes_all.csv", ';')
	createTableQuote()
	insertDataToQuote(csvData, 90)
	db.Close()
}
