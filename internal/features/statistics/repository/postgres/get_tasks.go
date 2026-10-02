package statistics_postgres_repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dyingvoid/todoapp/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

func (r *StatisticsRepository) GetTasks(
	ctx context.Context,
	userID *int64,
	from *time.Time,
	to *time.Time,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	args := pgx.NamedArgs{
		"user_id": userID,
		"from":    from,
		"to":      to,
	}

	var sb strings.Builder
	sb.WriteString("SELECT * FROM todoapp.task")

	conditions := []string{}

	if userID != nil {
		conditions = append(conditions, "user_id = @user_id")
		args["user_id"] = *userID
	}
	if from != nil {
		conditions = append(conditions, "created_at >= @from")
		args["from"] = *from
	}
	if to != nil {
		conditions = append(conditions, "created_at < @to")
		args["to"] = *to
	}

	if len(conditions) > 0 {
		sb.WriteString(" WHERE ")
		sb.WriteString(strings.Join(conditions, " AND "))
	}

	sb.WriteString(" ORDER BY id ASC;")

	rows, err := r.pool.Query(ctx, sb.String(), args)
	if err != nil {
		return nil, fmt.Errorf("statistics tasks query: %w", err)
	}

	tasksModels, err := pgx.CollectRows(rows, pgx.RowToStructByName[taskModel])
	if err != nil {
		return nil, fmt.Errorf("statistics tasks collect: %w", err)
	}

	return modelsToDomains(tasksModels), nil
}
