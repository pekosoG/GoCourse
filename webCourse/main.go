package main

import (
	"net/http"
	"fmt"
)

func main(){

	http.HandleFunc("/",func(rw http.ResponseWriter, req *http.Request){
		fmt.Fprintf(rw,"<h1>Hello World GoLan</h1>")
	})

	http.ListenAndServe(":8080",nil)
}
