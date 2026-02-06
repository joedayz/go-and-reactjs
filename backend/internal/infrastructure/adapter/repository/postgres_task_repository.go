package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/josediaz/go-and-reactjs/backend/internal/domain"
)

// PostgresTaskRepository implementa el puerto domain.TaskRepository (adaptador de infraestructura).
type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) FindByID(id string) (*domain.Task, error) {
	var t domain.Task
	var createdAt, updatedAt time.Time
	err := r.db.QueryRow(
		`SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = $1`,
		id,
	).Scan(&t.ID, &t.Title, &t.Description, &t.Status, &createdAt, &updatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t.CreatedAt = createdAt
	t.UpdatedAt = updatedAt
	return &t, nil
}

func (r *PostgresTaskRepository) FindAll() ([]*domain.Task, error) {
	rows, err := r.db.Query(`SELECT id, title, description, status, created_at, updated_at FROM tasks ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []*domain.Task
	for rows.Next() {
		var t domain.Task
		var createdAt, updatedAt time.Time
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &createdAt, &updatedAt); err != nil {
			return nil, err
		}
		t.CreatedAt = createdAt
		t.UpdatedAt = updatedAt
		tasks = append(tasks, &t)
	}
	return tasks, rows.Err()
}

func (r *PostgresTaskRepository) Save(task *domain.Task) error {
	if task.ID == "" {
		task.ID = uuid.New().String()
	}
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now
	_, err := r.db.Exec(
		`INSERT INTO tasks (id, title, description, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`,
		task.ID, task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt,
	)
	return err
}

func (r *PostgresTaskRepository) Update(task *domain.Task) error {
	task.UpdatedAt = time.Now()
	_, err := r.db.Exec(
		`UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5`,
		task.Title, task.Description, task.Status, task.UpdatedAt, task.ID,
	)
	return err
}

func (r *PostgresTaskRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	return err
}
