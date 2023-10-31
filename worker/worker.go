package worker

import (
	"fmt"
	"monazzim/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// Worker struct represents a system component responsible for running tasks as Docker containers,
// receiving task assignments from a manager, providing statistics for scheduling purposes,
// and maintaining task state information.
type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]task.Task
	TaskCount int
}

func (w *Worker) RunTask() {
	fmt.Println("I will start or stop a task")
}

func (w *Worker) StartTask() {
	fmt.Println("I will start a task")
}

func (w *Worker) StopTask() {
	fmt.Println("I will stop a task")
}

func (w *Worker) CollectStats() {
	fmt.Println("I will collect stats")
}
