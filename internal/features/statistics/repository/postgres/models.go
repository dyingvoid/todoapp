package statistics_postgres_repository

import (
	"time"

	"github.com/dyingvoid/todoapp/internal/core/domain"
)

type taskModel struct {
	ID           int64      `db:"id"`
	Version      int64      `db:"version"`
	Title        string     `db:"title"`
	Description  *string    `db:"description"`
	Completed    bool       `db:"completed"`
	CreatedAt    time.Time  `db:"created_at"`
	CompletedAt  *time.Time `db:"completed_at"`
	AuthorUserID int64      `db:"user_id"`
}

func (t *taskModel) ToDomain() domain.Task {
	return domain.NewTask(
		t.ID,
		t.Version,
		t.Title,
		t.Description,
		t.Completed,
		t.CreatedAt,
		t.CompletedAt,
		t.AuthorUserID,
	)
}

func modelsToDomains(models []taskModel) []domain.Task {
	domains := make([]domain.Task, len(models))
	for i, model := range models {
		domains[i] = model.ToDomain()
	}
	return domains
}
