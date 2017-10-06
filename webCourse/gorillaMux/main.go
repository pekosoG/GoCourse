package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"fmt"
)

func GetUsers(resp http.ResponseWriter, req *http.Request){
	fmt.Fprintf(resp,"Get Users")
}

func PostUsers(resp http.ResponseWriter, req *http.Request){
	fmt.Fprintf(resp,"Post Users")
}

func PutUsers(resp http.ResponseWriter, req *http.Request){
	fmt.Fprintf(resp,"Put Users")
}

func DeleteUsers(resp http.ResponseWriter, req *http.Request){
	fmt.Fprintf(resp,"Delete Users")
}

func main(){
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/api/users",GetUsers).Methods("GET")
	router.HandleFunc("/api/users",PostUsers).Methods("POST")
	router.HandleFunc("/api/users",PutUsers).Methods("PUT")
	router.HandleFunc("/api/users",DeleteUsers).Methods("DELETE")

	server := http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}