package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	ctx := r.Context() // context da request
	log.Println("Request Started")
	defer log.Println("Request Finished")

	select {
		case <- time.After(5 * time.Second):
			log.Println("Request finished with Successful") // Imprimi no comand line stdout
			w.Write([]byte("Request finished with Success")) // Imprimir no browser
		case <-ctx.Done():
			log.Println("Request canceled for the client") // Imprimi no comand line stdout
			http.Error(w, "Request canceled for the client", http.StatusRequestTimeout)
	}
}