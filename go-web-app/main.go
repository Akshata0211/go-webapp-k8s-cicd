package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./website"))

	// Custom handler for root `/`
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./website/index.html")
			return
		}
		fs.ServeHTTP(w, r) // fallback to file server
	})

	log.Println("Serving on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
