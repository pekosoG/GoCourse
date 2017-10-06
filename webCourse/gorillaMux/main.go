package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
	"encoding/json"
	"strconv"
)

type Note struct{
	title string `json: "title"`
	description string `json: "description"`
	createdAt time.Time `json: "created_at"`
}

var noteStore = make(map[string]Note)

var id int = 0

func getNotesHandler(resp http.ResponseWriter, req *http.Request){

	var notes[]Note
	for _,note := range noteStore{
		notes = append(notes,note)
	}

	resp.Header().Set("Content-Type","application/json")
	jsonResp, error := json.Marshal(notes)

	if error!=nil {
		panic(error)
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(jsonResp)
}

func createNotesHandler(resp http.ResponseWriter, req *http.Request){

	log.Println("Got a Get Call")
	var note Note
	error := json.NewDecoder(req.Body).Decode(&note)

	if error != nil{
		panic(error)
	}

	note.createdAt=time.Now()

	log.Println("Before adding: ",len(noteStore))

	id++
	noteStore[strconv.Itoa(id)]=note

	log.Println("Note Added: ",len(noteStore))
	resp.Header().Set("Content-Type","application/json")
	resp.WriteHeader(http.StatusCreated)

	var jsonResp,_ = json.Marshal(note)

	resp.Write(jsonResp)
}

func updateNotesHandler(resp http.ResponseWriter, req *http.Request){

}

func deleteNotesHandler(resp http.ResponseWriter, req *http.Request){

}


func main(){
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/api/notes", getNotesHandler).Methods("GET")
	router.HandleFunc("/api/notes", createNotesHandler).Methods("POST")
	router.HandleFunc("/api/notes/{id}", updateNotesHandler).Methods("PUT")
	router.HandleFunc("/api/notes/{id}", deleteNotesHandler).Methods("DELETE")

	server := http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening localhots:8080")
	server.ListenAndServe()
}