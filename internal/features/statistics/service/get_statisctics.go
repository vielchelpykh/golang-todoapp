package statistics_service

import (
	"context"
	"fmt"
	"time"

	domain "github.com/vielchelpykh/golang-todoapp/internal/core/domains"
	core_errors "github.com/vielchelpykh/golang-todoapp/internal/core/errors"
)

func (s *StaticticsService) GetStatistics(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Statistics{}, fmt.Errorf(
				"'to' must be after 'from': %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	tasks, err := s.staticticsRepository.GetTasks(ctx, userID, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("get tasks from repository: %w", err)
	}

	statictics := calcStatistics(tasks)

	return statictics, nil
}

func calcStatistics(tasks []domain.Task) domain.Statistics {
	if len(tasks) == 0 {
		return domain.Statistics{}
	}

	tasksCreated := len(tasks)

	taskCompleted := 0
	var totalCaompletedDuration time.Duration
	for _, task := range tasks {
		if task.Completed {
			taskCompleted++
		}

		completionDuration := task.CompletionDuration()
		if completionDuration != nil {
			totalCaompletedDuration += *completionDuration
		}
	}

	taskCompletedRate := float64(taskCompleted) / float64(tasksCreated) * 100

	var tasksAverageCompletionTime *time.Duration
	if taskCompleted > 0 && totalCaompletedDuration != 0 {
		avg := totalCaompletedDuration / time.Duration(taskCompleted)

		tasksAverageCompletionTime = &avg
	}

	return domain.NewStatistics(tasksCreated, taskCompleted, &taskCompletedRate, tasksAverageCompletionTime)
}
