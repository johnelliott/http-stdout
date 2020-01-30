package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			io.Copy(os.Stdout, r.Body)
			io.WriteString(os.Stdout, "\n")
		} else {
			log.Println(r.Method)
		}
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "ok")
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
