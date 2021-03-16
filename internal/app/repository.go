package app

import "github.com/amartery/statSaver/internal/app/models"

type Repository interface {
	Add(s *models.StatisticsShow) error
	Show(d *models.DateLimit) ([]models.StatisticsShow, error)
	ShowOrdered(d *models.DateLimit, category string) ([]models.StatisticsShow, error)
	ClearStatistics() error
}
