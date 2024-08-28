package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"protobuf/models"
	"strconv"

	"google.golang.org/protobuf/proto"
)

var persons = make(map[int32]*models.Person)

func main() {
	http.HandleFunc("/add", addPersonHandler)
	http.HandleFunc("/get", getPersonHandler)

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addPersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed.", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	person := models.Person{}

	err = proto.Unmarshal(body, &person)

	if err != nil {
		http.Error(w, "Failed to unmarshal protobuf", http.StatusBadRequest)
		return
	}

	persons[person.Id] = &person

	fmt.Println("Person added successfully: %v", &person)
}

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed.", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
	}

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error when converting to int: ", err)
	}

	id := int32(idInt)

	person := persons[id]

	if person == nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	data, err := proto.Marshal(person)
	if err != nil {
		http.Error(w, "Failed to marshal protobuf", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octed-stream")
	w.Write(data)
}
