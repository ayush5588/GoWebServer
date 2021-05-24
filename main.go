package main

import (
	"fmt"
	"net/http")


func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello world")
} 

func main() {
	http.HandleFunc("/",helloWorld)
	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080",nil)

}