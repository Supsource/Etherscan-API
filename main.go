package main

// Importing
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Define a struct to hold the transaction data
type Transaction struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
	Block string `json:"blockNumber"`
	Time  string `json:"timeStamp"`
}

// Define a struct to hold the API response
type ApiResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []Transaction `json:"result"`
}

// Endpoint to fetch the data from the Etherscan API
func fetchData() ([]Transaction, error) {
	// The API key goes here
	apiKey := "XE5JUF29JI5711WR6WKACV2PHVAIRUYFYI"
	url := "https://api.etherscan.io/api?module=account&action=tokentx&contractaddress=0x9355372396e3F6daF13359B7b607a3374cc638e0&page=1&sort=asc&apikey=" + apiKey + "&limit=100"

	// Making the API call
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var apiResponse ApiResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}
	return apiResponse.Result, nil
}

var database []Transaction // In-memory database to store the transactions

// Initialize the in-memory database
func initDatabase() error {
	data, err := fetchData()
	if err != nil {
		return err
	}
	database = make([]Transaction, len(data))
	copy(database, data)

	return nil
}

// Handler function to return all transacrions
func getAllTransactions(w http.ResponseWriter, r *http.Request) {
	var limit int = 10 // Slice the database to return a maximum of 10 transactions
	if len(database) > limit {
		database = database[:limit]
	}
	response, err := json.Marshal(database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Handler function to return paginated transactions
func getPaginatedTransactions(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		http.Error(w, "Invalid offset parameter", http.StatusBadRequest)
		return
	}

	start := offset
	end := offset + limit
	if end > len(database) {
		end = len(database)
	}

	response, err := json.Marshal(database[start:end])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Handler function to return transactions based on from and to addresses
func getTransactionsByAddress(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	// Create a slice to hold the matching transactions
	matches := make([]Transaction, 0)

	for _, tx := range database {
		if tx.From == from && tx.To == to {
			matches = append(matches, tx)
		}
	}

	response, err := json.Marshal(matches)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	err := initDatabase()
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter() // Create a new router

	router.HandleFunc("/transactions", getAllTransactions).Methods("GET")
	router.HandleFunc("/transactions/paginated", getPaginatedTransactions).Methods("GET")
	router.HandleFunc("/transactions/address", getTransactionsByAddress).Methods("GET")

	// Start the HTTP server:
	fmt.Println("Server listening on port 8001")
	err = http.ListenAndServe(":8001", router)
	if err != nil {
		log.Fatal(err)
	}
}
