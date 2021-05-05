package app

import "github.com/amartery/statSaver/internal/app/models"

// go:generate mockgen -destination=./internal/app/mocks/mock_repo.go -package=mocks github.com/amartery/statSaver/internal/app Repository
type Repository interface {
	Add(s *models.StatisticsShow) error
	ShowOrdered(model *models.RequestForShow) (*[]models.StatisticsShow, error)
	ClearStatistics() error
}
