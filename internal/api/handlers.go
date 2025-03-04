package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/SashaMorkovkin/Final_task_2/internal/calculator"
	"github.com/SashaMorkovkin/Final_task_2/internal/taskmanager"
)

// Функция для расчета выражения
func CalculateExpression(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Expression string `json:"expression"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		log.Printf("Invalid data: %v", err)
		return
	}

	// Выполняем расчет
	result, err := calculator.Calc(req.Expression)
	if err != nil {
		http.Error(w, "Error calculating expression", http.StatusInternalServerError)
		log.Printf("Error calculating expression: %v", err)
		return
	}

	taskID := taskmanager.AddTask(req.Expression)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":     taskID,
		"result": result,
	})
}

func GetExpressionsList(w http.ResponseWriter, r *http.Request) {
	tasks := taskmanager.GetTasks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"expressions": tasks,
	})
}

// Получение статуса выражения по ID
func GetExpressionStatus(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID из пути URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	taskID := parts[4]
	task, exists := taskmanager.GetTaskByID(taskID)
	if !exists {
		http.Error(w, "Expression not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
