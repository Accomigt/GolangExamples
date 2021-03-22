package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/about", aboutHandler)
	//Server
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("contact page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about yes"))
}
