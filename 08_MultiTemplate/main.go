package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/oxtoacart/bpool"
)

var templates map[string]*template.Template
var bufpool *bpool.BufferPool

// UserData ...
type UserData struct {
	Name        string
	City        string
	Nationality string
}

// SkillSet ...
type SkillSet struct {
	Language string
	Level    string
}

// TemplateConfig ...
type TemplateConfig struct {
	TemplateLayoutPath  string
	TemplateIncludePath string
}

// SkillSets ...
type SkillSets []*SkillSet

var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

var templateConfig TemplateConfig

func loadConfiguration() {
	templateConfig.TemplateLayoutPath = "templates/layouts/"
	templateConfig.TemplateIncludePath = "templates/"
}

func loadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layoutFiles, err := filepath.Glob(templateConfig.TemplateLayoutPath + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	includeFiles, err := filepath.Glob(templateConfig.TemplateIncludePath + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	mainTemplate := template.New("main")

	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			log.Fatal(err)
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}

	log.Println("templates loading successful")

	bufpool = bpool.NewBufferPool(64)
	log.Println("buffer allocation successful")
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("The template %s does not exist.", name),
			http.StatusInternalServerError)
	}

	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	userData := &UserData{Name: "Asit Dhal", City: "Bhubaneswar", Nationality: "Indian"}
	renderTemplate(w, "aboutme.html", userData)
}

func skillSet(w http.ResponseWriter, r *http.Request) {
	skillSets := SkillSets{&SkillSet{Language: "Golang", Level: "Beginner"},
		&SkillSet{Language: "C++", Level: "Advanced"},
		&SkillSet{Language: "Python", Level: "Advanced"}}
	renderTemplate(w, "skillset.html", skillSets)
}

func main() {
	loadConfiguration()
	loadTemplates()
	// server := http.Server{
	// 	Addr: "127.0.0.1:8080",
	// }

	// http.HandleFunc("/", index)
	// http.HandleFunc("/aboutme", aboutMe)
	// http.HandleFunc("/skillset", skillSet)
	// server.ListenAndServe()

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/", index)
	router.Get("/aboutme", aboutMe)
	router.Get("/skillset", skillSet)

	http.ListenAndServe(":8080", router)
}
