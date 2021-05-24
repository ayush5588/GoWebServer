package main

import (
	"fmt"
	"log"
	"net/http"
)


func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello world")
	log.Println("Hello World")
} 

func main() {
	http.HandleFunc("/",helloWorld)
	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080",nil)

}