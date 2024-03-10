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


// new handler to stream data using channels
func eventsHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers for Server-Sent Events
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a channel to send data to the client
	dataChannel := make(chan string)

	// Close the channel when the request is closed
	defer close(dataChannel)

	// Simulate continuous data generation
	go func() {
		for i := 0; i++ {
			time.Sleep(1 * time.Second)
			dataChannel <- fmt.Sprintf("Message %d", i)
		}
	}()

	// Continuously send data to the client
	for {
		select {
		case data := <-dataChannel:
			fmt.Fprintf(w, "data: %s\n\n", data)
			w.(http.Flusher).Flush()
		case <-r.Context().Done():
			return
		}
	}
}