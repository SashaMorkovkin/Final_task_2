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

	// Стартуем горутины для обработки задач (агенты)
	computingPower := os.Getenv("COMPUTING_POWER")
	if computingPower == "" {
		computingPower = "4" // Если не задано, используем 4 агента
	}

	numWorkers := 4 // Можем взять из переменной окружения или задать по умолчанию
	if computingPower != "" {
		// Пытаемся преобразовать переменную в число
		parsedWorkers, err := fmt.Sscanf(computingPower, "%d", &numWorkers)
		if err != nil || parsedWorkers == 0 {
			numWorkers = 4
		}
	}

	// Запускаем агентов
	worker.StartWorkers(numWorkers)

	// Настроим сервер
	serverAddr := ":8080"
	fmt.Println("server has start on", serverAddr)

	// Запускаем сервер
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
