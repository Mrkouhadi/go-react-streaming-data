package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleEventsStream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Send events to client
	for i := 0; ; i++ {
		fmt.Fprintf(w, "data: Message %d\n", i)
		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second) // Simulate periodic updates
	}
}
