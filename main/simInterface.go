package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",defaultGateWay)
	http.ListenAndServe(":8080",nil)
}
func rootGarWay(w http.ResponseWriter, r *http.Request) {
	println("Welcome to Chris's homepage!")
}
func defaultGateWay(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello jingjing")
	println("Welcome to Chris's homepage! ")
}