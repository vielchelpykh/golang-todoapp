package statistics_transport_http

import (
	"context"
	"net/http"
	"time"

	domain "github.com/vielchelpykh/golang-todoapp/internal/core/domains"
	core_http_server "github.com/vielchelpykh/golang-todoapp/internal/core/transport/http/server"
)

type StatisticsHTTPHandler struct {
	statiscticsService StaticticsService
}

type StaticticsService interface {
	GetStatistics(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(
	statiscticsService StaticticsService,
) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		statiscticsService: statiscticsService,
	}
}

func (h *StatisticsHTTPHandler) Routers() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statisctics",
			Handler: h.GetStatistics,
		},
	}
}
