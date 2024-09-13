package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type TestRun struct {
	ID        string    `json:"id"`
	BuildID   string    `json:"build_id"`
	Status    string    `json:"status"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time,omitempty"`
	Results   string    `json:"results,omitempty"`
}

var testRuns = make(map[string]TestRun)
var testMutex sync.RWMutex

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tests/run", runTest).Methods("POST")
	r.HandleFunc("/tests/{id}", getTest).Methods("GET")
	r.HandleFunc("/tests", listTests).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func runTest(w http.ResponseWriter, r *http.Request) {
	var testRequest struct {
		BuildID string `json:"build_id"`
	}
	json.NewDecoder(r.Body).Decode(&testRequest)

	testID := fmt.Sprintf("test-%d", time.Now().UnixNano())
	test := TestRun{
		ID:        testID,
		BuildID:   testRequest.BuildID,
		Status:    "running",
		StartTime: time.Now(),
	}

	testMutex.Lock()
	testRuns[testID] = test
	testMutex.Unlock()

	go func() {
		time.Sleep(15 * time.Second) // Simulate test process
		testMutex.Lock()
		test.Status = "completed"
		test.EndTime = time.Now()
		test.Results = "All tests passed"
		testRuns[testID] = test
		testMutex.Unlock()
	}()

	json.NewEncoder(w).Encode(test)
}

func getTest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testID := vars["id"]

	testMutex.RLock()
	test, exists := testRuns[testID]
	testMutex.RUnlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(test)
}

func listTests(w http.ResponseWriter, r *http.Request) {
	testMutex.RLock()
	testList := make([]TestRun, 0, len(testRuns))
	for _, test := range testRuns {
		testList = append(testList, test)
	}
	testMutex.RUnlock()

	json.NewEncoder(w).Encode(testList)
}
