package main

import (
	"fmt"
	"log"
	"net/http"
)

func initRouter(s *service) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /memes", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		tags := params["tags"]

		res, err := s.findMeme(tags...)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(string("something went wrong!")))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})

	fmt.Print("Starting router...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
