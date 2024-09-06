package main

import (
	"api-testing/controllers"
	"api-testing/models"
	"log"
	"net/http"
)

func Checkhealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's working"))
}

func main() {
	PORT := ":8080"

	models.ConnectDB()
	models.DBMigrate()

	mux := http.NewServeMux()
	mux.HandleFunc("/posts", controllers.GetAllPosts)
	mux.HandleFunc("/checkhealth", Checkhealth)

	log.Printf("The server is running at %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
