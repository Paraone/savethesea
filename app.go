package main

import (
	// "fmt"
	"log"
	"net/http"
)

// func myHandler(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Hi %s", r.URL.Path[1:])

// }

func main(){
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	// http.HandleFunc("/", myHandler )

	log.Println("Listening... ")
	http.ListenAndServe(":3000", nil)
}
