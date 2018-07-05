package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-chi/chi"
)

func init() {
	loadConfiguration()
	loadTemplates()
}

var templates map[string][]string

// TemplateConfig ...
type TemplateConfig struct {
	TemplateLayoutPath  string
	TemplateIncludePath string
}

var templateConfig TemplateConfig

func loadConfiguration() {
	templateConfig.TemplateLayoutPath = "templates/layouts/"
	templateConfig.TemplateIncludePath = "templates/"
}

func loadTemplates() {
	if templates == nil {
		templates = make(map[string][]string)
	}
	layoutFiles, err := filepath.Glob(templateConfig.TemplateLayoutPath + "*.html")
	if err != nil {
		log.Fatal(err)
	}
	includeFiles, err := filepath.Glob(templateConfig.TemplateIncludePath + "*.html")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName] = files
	}
}

// RenderTemplate ...
func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	// fmt.Println("RenderTemplate HTML")
	// fmt.Println(templates[name])

	tmpl, err := template.ParseFiles(templates[name]...)
	// tmpl, err := template.ParseFiles("templates/test.html")

	if err != nil {
		http.Error(w, fmt.Sprintf("The template %s does not exist.", name),
			http.StatusInternalServerError)
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, 5*5); err != nil {
		panic(err)
	}
	tpl.WriteTo(w)
}

func main() {
	router := chi.NewRouter()
	// router.Use(middleware.RequestID)
	// router.Use(middleware.RealIP)
	// router.Use(middleware.Logger)
	// router.Use(middleware.Recoverer)
	// router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/", Index)
	router.Get("/home", Home)
	http.ListenAndServe(":8080", router)

}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/", Expires: expire, MaxAge: 0}
	http.SetCookie(w, &cookie)
	RenderTemplate(w, "receipt_log.html", nil)
}

// Home ...
func Home(w http.ResponseWriter, r *http.Request) {
	// read cookie
	var cookie, err = r.Cookie("testcookiename")
	if err == nil {
		var cookievalue = cookie.Value
		io.WriteString(w, "<b>get cookie value is "+cookievalue+"</b>\n")
	}
	fmt.Println(cookie.Value)

	RenderTemplate(w, "home.html", nil)
}
