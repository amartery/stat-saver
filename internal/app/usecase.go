package app

import "github.com/amartery/statSaver/internal/app/models"

// go:generate mockgen -destination=./internal/app/mocks/mock_usecase.go -package=mocks github.com/amartery/statSaver/internal/app Usecase
type Usecase interface {
	Add(s *models.StatisticsShow) error
	Show(d *models.DateLimit) ([]models.StatisticsShow, error)
	ShowOrdered(d *models.DateLimit, category string) ([]models.StatisticsShow, error)
	ClearStatistics() error
}
