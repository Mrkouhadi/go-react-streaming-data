package main

import (
	"fmt"
	"net/http"
)

func main() {
	//  handler for events streaming...
	http.HandleFunc("/events", handleEventsStream)
	http.HandleFunc("/newevents", eventsHandler)
	// Start the server
	fmt.Println("SSE server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
