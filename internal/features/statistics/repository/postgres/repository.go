package statistics_postgres_repository

import core_postgres_pool "github.com/dyingvoid/todoapp/internal/core/repository/postgres/conn"

type StatisticsRepository struct {
	pool core_postgres_pool.Pool
}

func NewStatisticsRepository(
	pool core_postgres_pool.Pool,
) *StatisticsRepository {
	return &StatisticsRepository{
		pool: pool,
	}
}
