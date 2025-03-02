package internal

import (
	"log"
	"time"
)

type Expression struct {
	ID        string  `json:"id"`
	Status    string  `json:"status"`
	Result    float64 `json:"result,omitempty"`
	Arg1      float64 `json:"arg1"`
	Arg2      float64 `json:"arg2"`
	Operation string  `json:"operation"`
}

// Простой парсер выражений
func ParseExpression(expression string) Expression {
	// Например, выражение "2+2*2" -> {Arg1: 2, Arg2: 2, Operation: "addition"}
	return Expression{
		Arg1:      2,
		Arg2:      2,
		Operation: "addition", // Упрощаем, оставляем только "addition"
	}
}

// Выполнение вычислений
func ComputeExpression(expr Expression) float64 {
	var result float64
	switch expr.Operation {
	case "addition":
		time.Sleep(time.Millisecond * 100) // Имитация задержки
		result = expr.Arg1 + expr.Arg2
	case "multiplication":
		time.Sleep(time.Millisecond * 200)
		result = expr.Arg1 * expr.Arg2
	case "subtraction":
		time.Sleep(time.Millisecond * 150)
		result = expr.Arg1 - expr.Arg2
	case "division":
		time.Sleep(time.Millisecond * 250)
		result = expr.Arg1 / expr.Arg2
	default:
		log.Println("Unknown operation")
	}
	return result
}
