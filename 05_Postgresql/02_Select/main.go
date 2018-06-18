package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type employee struct {
	ID     int
	Name   string
	Score  int
	Salary float32
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test123"
	dbname   = "employees"
)

func main() {
	// Creating the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Opening a connection to our database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	rows, err := db.Query("SELECT * FROM employees;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	employees := make([]employee, 0)
	for rows.Next() {
		empl := employee{}
		err := rows.Scan(&empl.ID, &empl.Name, &empl.Score, &empl.Salary) // order matters
		if err != nil {
			panic(err)
		}
		employees = append(employees, empl)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, empl := range employees {
		// fmt.Println(bk.isbn, bk.title, bk.author, bk.price)
		fmt.Printf("%d, %s, %d, $%.2f\n", empl.ID, empl.Name, empl.Score, empl.Salary)
	}
}
