// Entry point of the Golang application
package main 

import (
	"log"
	"net/http"
	"receipt-processor/handlers"
)

func main() {
	// setting up HTTP routes for API requests
	http.HandleFunc("/receipts/process", handlers.ProcessReceiptHandler)
	http.HandleFunc("/receipts/", handlers.GetPointsHandler)
	
	// using default 8080 port
	log.Fatal(http.ListenAndServe(":8080", nil))
}
