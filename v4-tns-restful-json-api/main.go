package main

import (
	"net/http"
	"log"
)





func main() {

	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", Index)
	//router.HandleFunc("/todos", TodoIndex)
	//router.HandleFunc("/todos/{todoId}", TodoShow)

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//})

	//log.Fatal(http.ListenAndServe(":8080", router))

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}

