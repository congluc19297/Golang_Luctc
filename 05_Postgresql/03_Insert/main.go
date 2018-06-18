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

	fmt.Println("Successfully Connected!")

	empl := employee{}
	empl.Name = "Tran Cong Luc"
	empl.Score = 44
	empl.Salary = 23000.55

	// insert values
	_, err = db.Exec("INSERT INTO employees (NAME, SCORE, SALARY) VALUES ($1, $2, $3)", empl.Name, empl.Score, empl.Salary)

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Inserted!")
}
