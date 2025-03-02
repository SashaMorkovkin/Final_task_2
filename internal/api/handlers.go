package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SashaMorkovkin/Final_task_2/internal/calculator"
	"github.com/SashaMorkovkin/Final_task_2/internal/taskmanager"
)

// Функция для расчета выражения
func CalculateExpression(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Expression string `json:"expression"`
	}

	// Декодируем тело запроса в структуру
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Используем вашу функцию Calc для вычисления результата
	result, err := calculator.Calc(req.Expression)
	if err != nil {
		http.Error(w, "Error calculating expression", http.StatusInternalServerError)
		log.Printf("Error calculating expression: %v", err)
		return
	}

	// Создаем задачу с результатом (ID задачи можно генерировать здесь или где-то еще)
	taskID := taskmanager.AddTask(req.Expression)

	// Отправляем результат обратно в клиентский ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":     taskID,
		"result": result,
	})
}

// Получение списка выражений
func GetExpressionsList(w http.ResponseWriter, r *http.Request) {
	tasks := taskmanager.GetTasks()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"expressions": tasks,
	})
}

// Получение статуса выражения по ID
func GetExpressionStatus(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Query().Get("id") // Получаем ID из URL

	task, exists := taskmanager.GetTaskByID(taskID)
	if !exists {
		http.Error(w, "Expression not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
