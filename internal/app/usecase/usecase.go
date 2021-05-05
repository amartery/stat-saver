package usecase

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amartery/statSaver/internal/app"
	"github.com/amartery/statSaver/internal/app/models"
)

type StatUsecase struct {
	repository app.Repository
}

func NewStatUsecase(statRep app.Repository) app.Usecase {
	return &StatUsecase{
		repository: statRep,
	}
}

func (usecase *StatUsecase) Add(req *models.RequestForSave) error {
	res := &models.StatisticsShow{}

	res.Date = req.Date

	if req.Views != "" {
		viewsInt, err := strconv.ParseInt(req.Views, 10, 64)
		if err != nil {
			return fmt.Errorf("'views' is not a valid")
		}
		if viewsInt < 0 {
			return fmt.Errorf("'views' must be > 0")
		}
		res.Views = viewsInt
	}
	if req.Clicks != "" {
		clicksInt, err := strconv.ParseInt(req.Clicks, 10, 64)
		if err != nil {
			return fmt.Errorf("'clicks' is not a valid")
		}
		if clicksInt < 0 {
			return fmt.Errorf("'clicks' must be > 0")
		}
		res.Clicks = clicksInt
	}
	if req.Cost != "" {
		if !validCostFormat(req.Cost) {
			return fmt.Errorf("'cost' must have two decimal places")
		}
		costFloat, err := strconv.ParseFloat(req.Cost, 64)
		if err != nil {
			return fmt.Errorf("'cost' is not a valid")
		}
		if costFloat < 0 {
			return fmt.Errorf("'cost' must be > 0")
		}
		res.Cost = costFloat
	}

	if req.Cost != "" && req.Clicks != "" && res.Clicks != 0 {
		res.Cpc = res.Cost / float64(res.Clicks)
	}
	if req.Cost != "" && req.Views != "" && res.Views != 0 {
		res.Cpm = res.Cost / float64(res.Views) * 1000
	}
	err := usecase.repository.Add(res)
	return err

}

func (usecase *StatUsecase) Show(model *models.RequestForShow) (*[]models.StatisticsShow, error) {
	if model.SortField == "" {
		model.SortField = "event_date"
	}

	arrayStat, err := usecase.repository.ShowOrdered(model)
	if err != nil {
		return nil, err
	}
	return arrayStat, nil
}

func (usecase *StatUsecase) ClearStatistics() error {
	err := usecase.repository.ClearStatistics()
	return err
}

func validCostFormat(value string) bool {
	s := strings.Split(value, ".")
	if len(s) != 2 {
		return false
	}
	if len(s[1]) != 2 {
		return false
	}
	return true
}
