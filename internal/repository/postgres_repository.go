package repository

import (
	"database/sql"
	"errors"

	"go-task-api/internal/domain"
)

type postgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository creates a new PostgreSQL task repository.
func NewPostgresRepository(db *sql.DB) TaskRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(t domain.Task) (domain.Task, error) {
	query := `
		INSERT INTO tasks (title, description, status) 
		VALUES ($1, $2, $3) 
		RETURNING id
	`
	err := r.db.QueryRow(query, t.Title, t.Description, t.Status).Scan(&t.ID)
	if err != nil {
		return domain.Task{}, err
	}
	return t, nil
}

func (r *postgresRepository) List() ([]domain.Task, error) {
	query := `SELECT id, title, description, status FROM tasks`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var t domain.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *postgresRepository) GetByID(id int) (domain.Task, error) {
	query := `SELECT id, title, description, status FROM tasks WHERE id = $1`
	var t domain.Task
	err := r.db.QueryRow(query, id).Scan(&t.ID, &t.Title, &t.Description, &t.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Task{}, errors.New("task not found") // simplify mapping later
		}
		return domain.Task{}, err
	}
	return t, nil
}

func (r *postgresRepository) Update(id int, t domain.Task) (domain.Task, error) {
	query := `
		UPDATE tasks 
		SET title = $1, description = $2, status = $3 
		WHERE id = $4 
		RETURNING id, title, description, status
	`
	var updated domain.Task
	err := r.db.QueryRow(query, t.Title, t.Description, t.Status, id).Scan(
		&updated.ID, &updated.Title, &updated.Description, &updated.Status,
	)
	if err != nil {
		return domain.Task{}, err
	}
	return updated, nil
}

func (r *postgresRepository) Delete(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}
