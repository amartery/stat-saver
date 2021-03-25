package app

import "github.com/amartery/statSaver/internal/app/models"

// go:generate mockgen -destination=./internal/app/mocks/mock_repo.go -package=mocks github.com/amartery/statSaver/internal/app Repository
type Repository interface {
	Add(s *models.StatisticsShow) error
	Show(d *models.DateLimit) ([]models.StatisticsShow, error)
	ShowOrdered(d *models.DateLimit, category string) ([]models.StatisticsShow, error)
	ClearStatistics() error
}
