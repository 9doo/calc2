package main

import (
	"log"
	"net/http"

	"calc2/internal/agent"
	"calc2/internal/handlers"
	"calc2/internal/orchestrator"
)

func main() {
	orchestratorInstance := orchestrator.NewOrchestrator()
	orchestratorInstance.Start()
	agent := agent.NewAgent(orchestratorInstance, 1)
	go agent.Start()
	http.HandleFunc("/api/v1/calculate", handlers.CalculateHandler)
	http.HandleFunc("/api/v1/expressions", handlers.ExpressionsHandler)
	http.HandleFunc("/api/v1/expressions/", handlers.ExpressionByIDHandler)
	log.Println("Starting calculator service on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
