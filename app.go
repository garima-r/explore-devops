package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/garima-r/explore-devops/src/blog"

	"github.com/gorilla/mux"
)

func main() {
	//Postgres DB initialization
	blog.InitDB()

	router := mux.NewRouter()

	// Route handles & endpoints
	router.HandleFunc("/display/articles", blog.DisplayArticlesHandler).Methods("GET")
	router.HandleFunc("/publish/article", blog.PublishArticleHandler).Methods("POST")

	// serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	log.Println("test github action workflow")
}
