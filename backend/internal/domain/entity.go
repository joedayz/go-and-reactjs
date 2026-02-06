package domain

import "time"

// Task representa la entidad de dominio (DDD) para una tarea.
// En arquitectura hexagonal vive en el núcleo del dominio.
type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "PENDING"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusDone      TaskStatus = "DONE"
)

// TaskRepository es el puerto (interfaz) que define las operaciones de persistencia.
// La infraestructura implementa este puerto (inversión de dependencias).
type TaskRepository interface {
	FindByID(id string) (*Task, error)
	FindAll() ([]*Task, error)
	Save(task *Task) error
	Update(task *Task) error
	Delete(id string) error
}
