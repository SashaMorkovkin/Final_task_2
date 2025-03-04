package worker

import (
	"log"
	"time"

	"github.com/SashaMorkovkin/Final_task_2/internal/calculator"
	"github.com/SashaMorkovkin/Final_task_2/internal/taskmanager"
)

func StartWorkers(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go workerLoop()
	}
}

func workerLoop() {
	for {
		task := getTaskFromOrchestrator()
		if task == nil {
			time.Sleep(1 * time.Second)
			continue
		}

		// Выполнение вычислений
		result, err := calculator.Calculate(task.Expression)
		if err != nil {
			log.Printf("Error calculating task %s: %v", task.ID, err)
			continue
		}

		// Обновление статуса задачи с результатом
		taskmanager.UpdateTaskStatus(task.ID, "completed", result)
	}
}

func getTaskFromOrchestrator() *taskmanager.Task {
	tasks := taskmanager.GetTasks()
	if len(tasks) > 0 {
		return tasks[0]
	}
	return nil
}
