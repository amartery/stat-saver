package app

import "github.com/amartery/statSaver/internal/app/models"

// go:generate mockgen -destination=./internal/app/mocks/mock_usecase.go -package=mocks github.com/amartery/statSaver/internal/app Usecase
type Usecase interface {
	Add(req *models.RequestForSave) error
	Show(model *models.RequestForShow) (*[]models.StatisticsShow, error)
	ClearStatistics() error
}
