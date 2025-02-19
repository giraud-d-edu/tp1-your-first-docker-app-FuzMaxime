package groupieTracker

import (
	"fmt"
	"log"
	"net/http"
)

// Server func
func Server() {
	/* ------ import css and img ---- */
	fs := http.FileServer(http.Dir("../../assets/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	/* ------------------------------ */
	http.HandleFunc("/", Home)
	http.HandleFunc("/details", Details)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting server to test (listen 8080 - link => http://localhost:8080/)\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
