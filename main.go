package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	flag.Parse()

	http.HandleFunc("/", home)
	http.HandleFunc("/events", events)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Errorf("error starting server: %v", err)
	}

}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	tokens := []string{"this", "is", "testing"}

	for _, token := range tokens {
		content := fmt.Sprintf("data: %s\n\n", token)
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		time.Sleep(1 * time.Second)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
