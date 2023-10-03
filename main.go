package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested %s, token is %s\n", r.URL.Path, r.URL.Query().Get("token"))
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Web server has started")
	log.Fatal(http.ListenAndServe(":3080", nil))
}

// package main

// import (
// 	"fmt"

// 	geo "github.com/iamengg/microsvc/geometry"
// 	"rsc.io/quote"
// )

// func main() {
// 	fmt.Println("Main package")
// 	fmt.Println(quote.Go())
// 	a, p := rectProps(7, 4.1)
// 	fmt.Printf("Area of rectangle is %v, %v\n", a, p)

// 	var dayOfTheMonth = map[string]int{"Jan": 31, "Feb": 22}
// 	fmt.Println(dayOfTheMonth)

// 	fmt.Println(geo.Area(1, 44))

// }

// func rectProps(length, width float64) (area, perimeter float64) {
// 	area = length * width
// 	perimeter = (length + width) * 2
// 	return
// }
