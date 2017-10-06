package main

import (
	"net/http"
	"fmt"
)

func helloWorld(rw http.ResponseWriter, req *http.Request){
	fmt.Fprintf(rw,"<h1>Hello World GoLan</h1>")
}

func testRoute(rw http.ResponseWriter, req *http.Request){
	fmt.Fprintf(rw,"<h1>Testing routing!</h1>")
}

func helloUser(rw http.ResponseWriter, req *http.Request){
	fmt.Fprintf(rw,"<h1>Some User route</h1>")
}

func main(){


	routerMux:= http.NewServeMux();

	routerMux.HandleFunc("/",helloWorld)

	routerMux.HandleFunc("/test",testRoute)

	routerMux.HandleFunc("/user",helloUser)

	http.ListenAndServe(":8080",routerMux)
}
