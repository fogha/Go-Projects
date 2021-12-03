package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Fogha/Go-Projects/youtube/youtube"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(": 8080", nil))
}

func main() {
	fmt.Println("youtube subscriber monitor")

	item, err := youtube.GetSubscribers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", item)
	//setupRoutes()
}
