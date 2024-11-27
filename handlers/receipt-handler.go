// creating ProcessReceiptHandler and GetPointsHandler for API endpoints
package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor/models"
	"receipt-processor/services"
	"strings"
	"sync"
)


var (
	receiptStore = make(map[string]models.Receipt)
	storeMutex   sync.Mutex
)

// ProcessReceiptHandler handles POST /receipts/process endpoint
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	
	// Generate a unique ID for the receipt
	receiptID := strings.ReplaceAll(receipt.Retailer, " ", "") + "-" + receipt.PurchaseDate
	points := services.CalculatePoints(receipt)	

	// storing the receipt in in-memory map
	storeMutex.Lock()
	receiptStore[receiptID] = receipt
	storeMutex.Unlock()
	
	// returning the id value for the receipt
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": receiptID})
}

// GetPointsHandler handles /receipts/{id}/points
func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// extractind the id value from the URL
	id := strings.TrimPrefix(r.URL.Path, "/receipts/")
	if id == "" {
		http.Error(w, "Missing receipt", http.StatusBadRequest)
		return
	}
	
	// fetching the corresponding receipt from nthe in-memory map
	storeMutex.Lock()
	receipt, exists := receiptStore[id]
	storeMutex.Unlock()

	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}
	
	// calculating points for the receipt using CalculatePoints method. view method implementation in point_calculator.go
	points := services.CalculatePoints(receipt)

	// getting the number of points
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
