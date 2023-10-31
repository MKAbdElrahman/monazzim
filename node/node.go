package node

import (
	"fmt"
	"monazzim/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string]task.Task
	EventDb       map[string]task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

// responsible for looking at the requirements specified in a Task
// and evaluating the resources available in the
// pool of workers to see which worker is best suited to run the task.
func (m *Manager) SelectWorker() {
	fmt.Println("I will select an appropriate worker")
}

// Keep track of tasks, their states, and the machine on which they run.
func (m *Manager) UpdateTasks() {
	fmt.Println("I will update tasks")
}

// send tasks to workers
func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")
}
