package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
)

type employee struct {
	ID     int
	Name   string
	Score  int
	Salary float32
}
type contextKey string

func (c contextKey) String() string {
	return string(c)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test123"
	dbname   = "employees"
)

var db *sql.DB
var tpl *template.Template

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

	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	r := chi.NewRouter()

	// r.Use(middleware.RequestID)
	// r.Use(middleware.Logger)
	// r.Use(middleware.RealIP)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/", func(r chi.Router) {
		r.Get("/", index)
	})
	r.Route("/show-empl", func(r chi.Router) {
		r.Route("/{emplID}", func(r chi.Router) {
			// r.Use(showEmplCtx)
			r.Get("/", showEmpl)
		})
	})
	r.Route("/add-empl.htm", func(r chi.Router) {
		r.Get("/", addEmployee)
		r.Post("/process", addEmployeeProcess)
	})

	http.ListenAndServe(":3000", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	employees, err := getAllEmployees()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	// for _, empl := range employees {
	// 	// fmt.Println(bk.isbn, bk.title, bk.author, bk.price)
	// 	fmt.Printf("%d, %s, %d, $%.2f\n", empl.ID, empl.Name, empl.Score, empl.Salary)
	// }
	tpl.ExecuteTemplate(w, "index.html", employees)
}

func showEmpl(w http.ResponseWriter, r *http.Request) {
	temp := chi.URLParam(r, "emplID")
	emplID, err := strconv.Atoi(temp)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	employee, err := getEmployee(emplID)
	if err != nil {
		http.Error(w, http.StatusText(422)+" Can't access database on my server", 422)
		return
	}
	if len(employee) == 0 {
		// http.Error(w, http.StatusText(404), 404)
		http.Error(w, http.StatusText(404)+" ID employee invalid", http.StatusNotAcceptable)
		return
	}
	// fmt.Println(len(employee))
	tpl.ExecuteTemplate(w, "show-empl.html", employee)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "add-empl.html", nil)
}

func addEmployeeProcess(w http.ResponseWriter, r *http.Request) {
	// get form values
	empl := employee{}
	var err error
	var temp1 int
	var temp2 float64
	empl.Name = r.FormValue("Name")
	temp1, err = strconv.Atoi(r.FormValue("Score"))
	temp2, err = strconv.ParseFloat(r.FormValue("Salary"), 32)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please hit back and enter a number for the score and salary", http.StatusNotAcceptable)
		return
	}
	empl.Score = temp1
	empl.Salary = float32(temp2)

	// Insert values
	_, err = db.Exec("INSERT INTO employees (name, score, salary) VALUES ($1, $2, $3)", empl.Name, empl.Score, empl.Salary)

	http.Redirect(w, r, "/", http.StatusSeeOther)
	// tpl.ExecuteTemplate(w, "index.html", empl)
}

func getAllEmployees() ([]employee, error) {
	employees := make([]employee, 0)

	rows, err := db.Query("SELECT * FROM employees;")
	if err != nil {
		return employees, err
	}
	defer rows.Close()

	for rows.Next() {
		empl := employee{}
		err := rows.Scan(&empl.ID, &empl.Name, &empl.Score, &empl.Salary) // order matters
		if err != nil {
			panic(err)
		}
		employees = append(employees, empl)
	}
	if err = rows.Err(); err != nil {
		return employees, err
	}
	return employees, nil
}

func getEmployee(ID int) ([]employee, error) {
	employees := make([]employee, 0)

	rows, err := db.Query("SELECT * FROM employees where ID = $1;", ID)
	if err != nil {
		return employees, err
	}
	defer rows.Close()

	for rows.Next() {
		empl := employee{}
		err := rows.Scan(&empl.ID, &empl.Name, &empl.Score, &empl.Salary) // order matters
		if err != nil {
			panic(err)
		}
		employees = append(employees, empl)
	}
	if err = rows.Err(); err != nil {
		return employees, err
	}
	return employees, nil
}
