package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/SashaMorkovkin/Final_task_2/internal"
)

var expressions = make(map[string]internal.Expression)

// Добавление нового вычисления
func AddCalculation(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Expression string `json:"expression"`
	}

	// Парсим входящие данные
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Преобразуем выражение
	parsedExpr := internal.ParseExpression(req.Expression)

	// Генерируем уникальный ID
	id := fmt.Sprintf("%d", len(expressions)+1)
	parsedExpr.ID = id
	parsedExpr.Status = "in_progress"

	// Добавляем выражение в список
	expressions[id] = parsedExpr

	// Выполнение вычисления в горутине
	go func() {
		result := internal.ComputeExpression(parsedExpr)
		parsedExpr.Status = "completed"
		parsedExpr.Result = result
		expressions[id] = parsedExpr
	}()

	// Отправляем ответ с ID нового вычисления
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"id": id,
	})
}

// Получение всех выражений
func GetAllExpressions(w http.ResponseWriter, r *http.Request) {
	var result []internal.Expression
	for _, expr := range expressions {
		result = append(result, expr)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"expressions": result,
	})
}

// Получение выражения по ID
func GetExpressionByID(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID из URL
	id := strings.TrimPrefix(r.URL.Path, "/api/v1/expressions/")

	expression, exists := expressions[id]

	if !exists {
		http.Error(w, "Expression not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"expression": expression,
	})
}
