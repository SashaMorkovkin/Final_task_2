package main

import (
	"Final_task_2/api"
	"log"
	"net/http"
)

func main() {
	// Регистрация обработчиков API
	http.HandleFunc("/api/v1/calculate", api.AddCalculation)
	http.HandleFunc("/api/v1/expressions", api.GetAllExpressions)
	http.HandleFunc("/api/v1/expressions/", api.GetExpressionByID) // Обработчик с параметром ID

	// Запуск HTTP сервера
	log.Fatal(http.ListenAndServe(":8080", nil))
}
