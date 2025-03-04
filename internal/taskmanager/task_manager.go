package taskmanager

import (
	"fmt"
	"sync"
)

var (
	mu          sync.Mutex
	tasks       = make(map[string]*Task)
	taskCounter = 0
)

type Task struct {
	ID         string
	Status     string
	Result     float64
	Expression string
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

type TaskManager struct{}

func generateTaskID() string {
	taskCounter++
	return fmt.Sprintf("%d", taskCounter)
}

func AddTask(expression string) string {
	mu.Lock()
	defer mu.Unlock()

	taskID := generateTaskID()
	task := &Task{
		ID:         taskID,
		Status:     "pending",
		Expression: expression,
	}
	tasks[taskID] = task
	return taskID
}

func GetTasks() []*Task {
	mu.Lock()
	defer mu.Unlock()

	var taskList []*Task
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	return taskList
}

func GetTaskByID(taskID string) (*Task, bool) {
	mu.Lock()
	defer mu.Unlock()
	task, exists := tasks[taskID]
	return task, exists
}

func UpdateTaskStatus(taskID, status string, result float64) {
	mu.Lock()
	defer mu.Unlock()

	if task, exists := tasks[taskID]; exists {
		task.Status = status
		task.Result = result
	}
}
