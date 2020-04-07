package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}
type Covid struct {
	Country string
	Cases   string
}

func main() {
	welcome := Welcome{"EZ", time.Now().Format(time.Stamp)}

	//pk := Covid{"Pakistan", "3000"}
	//sk := Covid{"South Korea", "3000"}
	//fr := Covid{"France", "3000"}

	templates := template.Must(template.ParseFiles("templates/welcome.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "welcome.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Println("Listening")
	log.Println(http.ListenAndServe(":8080", nil))
}
