package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const port = "8080"

var homeTemplate *template.Template
var message string

func init() {
	homeTemplate = template.Must(template.ParseFiles("static/index.html"))
	message = "Hi!"
}

func main() {
	r := mux.NewRouter()

	r.Path("/").HandlerFunc(homeHandler)
	r.Path("/image.png").HandlerFunc(imageHandler)
	r.Path("/message").HandlerFunc(messageHandler).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Printf("Listening on :%s", port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln("Could not start origin server:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("homeHandler")
  w.WriteHeader(http.StatusOK)
	homeTemplate.Execute(w, struct{Message string}{message})
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./image.png")
}


func messageHandler(w http.ResponseWriter, r *http.Request) {
	message = r.FormValue("message")
	log.Printf("messageHandler: %s", message)
	Draw(message)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
