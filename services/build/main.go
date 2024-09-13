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

type Build struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time,omitempty"`
}

var builds = make(map[string]Build)
var buildMutex sync.RWMutex

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/builds/trigger", triggerBuild).Methods("POST")
	r.HandleFunc("/builds/{id}", getBuild).Methods("GET")
	r.HandleFunc("/builds", listBuilds).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func triggerBuild(w http.ResponseWriter, r *http.Request) {
	buildID := fmt.Sprintf("build-%d", time.Now().UnixNano())
	build := Build{
		ID:        buildID,
		Status:    "running",
		StartTime: time.Now(),
	}

	buildMutex.Lock()
	builds[buildID] = build
	buildMutex.Unlock()

	go func() {
		time.Sleep(30 * time.Second) // Simulate build process
		buildMutex.Lock()
		build.Status = "completed"
		build.EndTime = time.Now()
		builds[buildID] = build
		buildMutex.Unlock()
	}()

	json.NewEncoder(w).Encode(build)
}

func getBuild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	buildID := vars["id"]

	buildMutex.RLock()
	build, exists := builds[buildID]
	buildMutex.RUnlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(build)
}

func listBuilds(w http.ResponseWriter, r *http.Request) {
	buildMutex.RLock()
	buildList := make([]Build, 0, len(builds))
	for _, build := range builds {
		buildList = append(buildList, build)
	}
	buildMutex.RUnlock()

	json.NewEncoder(w).Encode(buildList)
}
