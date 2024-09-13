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

type Deployment struct {
	ID        string    `json:"id"`
	BuildID   string    `json:"build_id"`
	Environment string  `json:"environment"`
	Status    string    `json:"status"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time,omitempty"`
}

var deployments = make(map[string]Deployment)
var deployMutex sync.RWMutex

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/deployments/create", createDeployment).Methods("POST")
	r.HandleFunc("/deployments/{id}", getDeployment).Methods("GET")
	r.HandleFunc("/deployments", listDeployments).Methods("GET")

	log.Fatal(http.ListenAndServe(":8082", r))
}

func createDeployment(w http.ResponseWriter, r *http.Request) {
	var deployRequest struct {
		BuildID     string `json:"build_id"`
		Environment string `json:"environment"`
	}
	json.NewDecoder(r.Body).Decode(&deployRequest)

	deployID := fmt.Sprintf("deploy-%d", time.Now().UnixNano())
	deploy := Deployment{
		ID:          deployID,
		BuildID:     deployRequest.BuildID,
		Environment: deployRequest.Environment,
		Status:      "in_progress",
		StartTime:   time.Now(),
	}

	deployMutex.Lock()
	deployments[deployID] = deploy
	deployMutex.Unlock()

	go func() {
		time.Sleep(45 * time.Second) // Simulate deployment process
		deployMutex.Lock()
		deploy.Status = "completed"
		deploy.EndTime = time.Now()
		deployments[deployID] = deploy
		deployMutex.Unlock()
	}()

	json.NewEncoder(w).Encode(deploy)
}

func getDeployment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deployID := vars["id"]

	deployMutex.RLock()
	deploy, exists := deployments[deployID]
	deployMutex.RUnlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(deploy)
}

func listDeployments(w http.ResponseWriter, r *http.Request) {
	deployMutex.RLock()
	deployList := make([]Deployment, 0, len(deployments))
	for _, deploy := range deployments {
		deployList = append(deployList, deploy)
	}
	deployMutex.RUnlock()

	json.NewEncoder(w).Encode(deployList)
}
