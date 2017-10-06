package main

import (
	"net/http"
	"fmt"
	"time"
)

func testRoute(rw http.ResponseWriter, req *http.Request){
	fmt.Fprintf(rw,"<h1>Testing routing!</h1>")
}

func helloUser(rw http.ResponseWriter, req *http.Request){

}

type message struct{
	text string
}

/**
	This is a Custom handler that we can define with a Signature that in this case is message
 */
func(msg message) ServeHTTP(resp http.ResponseWriter, req *http.Request){
	fmt.Fprintf(resp,msg.text)
}

func main(){

	routerMux:= http.NewServeMux()

	staticFiles:= http.FileServer(http.Dir("public"))

	routerMux.Handle("/",staticFiles)

	routerMux.HandleFunc("/test",testRoute)

	routerMux.HandleFunc("/user",helloUser)

	someMessage:=message{text:"Hello There"}

	/*
		And since we have the default handler defined, our RouterMux recognize it and use it
	 */
	routerMux.Handle("/hello",someMessage)

	server := http.Server{
		Addr: ":8080",
		Handler:routerMux,
		ReadTimeout:10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
