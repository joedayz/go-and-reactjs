package application

import (
	"github.com/josediaz/go-and-reactjs/backend/internal/domain"
)

// TaskService aplica los casos de uso (Application Layer / DDD).
// Depende del puerto TaskRepository, no de implementaciones concretas.
type TaskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetByID(id string) (*domain.Task, error) {
	return s.repo.FindByID(id)
}

func (s *TaskService) ListAll() ([]*domain.Task, error) {
	return s.repo.FindAll()
}

func (s *TaskService) Create(title, description string) (*domain.Task, error) {
	task := &domain.Task{
		Title:       title,
		Description: description,
		Status:      domain.TaskStatusPending,
	}
	if err := s.repo.Save(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) UpdateStatus(id string, status domain.TaskStatus) (*domain.Task, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	task.Status = status
	if err := s.repo.Update(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) Delete(id string) error {
	return s.repo.Delete(id)
}
