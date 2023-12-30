package main

//docker build -t go-app:latest .
//go distroless image
//googlecontainer tools
//To run url
//http://localhost/books/hi/page/123

import (
	"fmt"
	"log"
	"net/http"

	geo "github.com/iamengg/microsvc/devops-made-easy/geometry"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func handlerArea(w http.ResponseWriter, r *http.Request) {
	res := geo.Area(3, 5)
	fmt.Fprintf(w, "area is %f\n", res)
}

func msg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my website!")
}

func main() {
	fmt.Println("Start of microservice architecture")
	r := mux.NewRouter()

	//handlers
	r.HandleFunc("/books/{title}/page/{page}", handler)
	r.HandleFunc("/v1/area", handlerArea)
	r.HandleFunc("/v1/msg", msg)
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
