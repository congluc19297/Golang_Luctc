package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

type fetchIP struct {
	Origin string
}

func main() {
	fmt.Println("Before Get data")
	response, err := http.Get("https://httpbin.org/ip")
	fmt.Println("After Get data")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// log.Fatal(err)
		panic(err)
	}

	fmt.Println(string(responseData))
	var fetchIps fetchIP
	err = json.Unmarshal(responseData, &fetchIps)
	if err != nil {
		panic(err)
	}
	fmt.Println(fetchIps)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(responseData))
	})
	http.ListenAndServe(":3000", r)
}
