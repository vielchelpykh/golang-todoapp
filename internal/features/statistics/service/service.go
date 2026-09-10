package statistics_service

import (
	"context"
	"time"

	domain "github.com/vielchelpykh/golang-todoapp/internal/core/domains"
)

type StaticticsService struct {
	staticticsRepository StaticticsRepository
}

type StaticticsRepository interface {
	GetTasks(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) ([]domain.Task, error)
}

func NewStatisticsService(staticticsRepository StaticticsRepository) *StaticticsService {
	return &StaticticsService{
		staticticsRepository: staticticsRepository,
	}
}
