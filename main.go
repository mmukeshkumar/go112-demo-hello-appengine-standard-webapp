package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)

//based of https://github.com/GoogleCloudPlatform/golang-samples/tree/master/appengine/go11x/helloworld
//http://localhost:8080/?name=Mukesh
func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("Starting Server")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	log.Printf("Received request for %s\n", name)
	name = fmt.Sprintf("Hello, %s\n", name)
	fmt.Fprintln(w, name)
}