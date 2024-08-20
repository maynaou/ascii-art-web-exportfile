package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"bonus_1/Handler"
	utils "bonus_1/helpers"
)

var tmpl *template.Template

func main() {
	blue := "\033[34m"
	yellow := "\033[33m"
	reset := "\033[0m"

	tmpl = template.Must(template.ParseGlob("Templates/*.html"))
	Handler.VetTemplate(tmpl)
	utils.SetTemplate(tmpl)
	fs := http.FileServer(http.Dir("Static"))
	http.Handle("/Static/", http.StripPrefix("/Static/", fs))
	http.HandleFunc("/", Handler.HomeHandler)
	http.HandleFunc("/ascii-art", Handler.FormHandler)
	http.HandleFunc("/download", Handler.DownHandler)

	fmt.Println(blue + "/-------------------------------------------------------------------\\" + reset)
	fmt.Println(blue + "|" + yellow + "                Server started on http://localhost:8088            " + blue + " |" + reset)
	fmt.Println(blue + "\\-------------------------------------------------------------------/" + reset)

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
