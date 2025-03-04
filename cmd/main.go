package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SashaMorkovkin/Final_task_2/internal/api"
	"github.com/SashaMorkovkin/Final_task_2/internal/worker"
)

func main() {
	// Создание серверных обработчиков
	http.HandleFunc("/api/v1/calculate", api.CalculateExpression)
	http.HandleFunc("/api/v1/expressions", api.GetExpressionsList)
	http.HandleFunc("/api/v1/expressions/", api.GetExpressionStatus) // Статус задачи по ID
	computingPower := os.Getenv("COMPUTING_POWER")
	if computingPower == "" {
		computingPower = "4" // Если не задано, используем 4 агента
	}
	numWorkers := 4
	if computingPower != "" {
		parsedWorkers, err := fmt.Sscanf(computingPower, "%d", &numWorkers)
		if err != nil || parsedWorkers == 0 {
			numWorkers = 4
		}
	}
	worker.StartWorkers(numWorkers)

	serverAddr := ":8080"
	fmt.Println("server has start on", serverAddr)

	// Запускаем сервер
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
