package usecase

import (
	"github.com/amartery/statSaver/internal/app"
	"github.com/amartery/statSaver/internal/app/models"
)

type StatUsecase struct {
	statRep app.Repository
}

func NewStatUsecase(statRep app.Repository) app.Usecase {
	return &StatUsecase{
		statRep: statRep,
	}
}

func (usecase *StatUsecase) Add(s *models.StatisticsShow) error {
	err := usecase.statRep.Add(s)
	return err

}

func (usecase *StatUsecase) Show(d *models.DateLimit) ([]models.StatisticsShow, error) {
	arrayStat, err := usecase.statRep.Show(d)
	if err != nil {
		return nil, err
	}
	return arrayStat, nil
}

func (usecase *StatUsecase) ShowOrdered(d *models.DateLimit, category string) ([]models.StatisticsShow, error) {
	arrayStatOrder, err := usecase.statRep.ShowOrdered(d, category)
	if err != nil {
		return nil, err
	}
	return arrayStatOrder, nil
}

func (usecase *StatUsecase) ClearStatistics() error {
	err := usecase.statRep.ClearStatistics()
	return err
}
