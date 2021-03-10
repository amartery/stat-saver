package validation

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/amartery/statSaver/internal/app/model"
)

// DateValidate ...
func DateValidate(from, to string) (*model.DateLimit, error) {
	if from == "" || to == "" {
		return nil, fmt.Errorf("from, to are required")
	}
	dateFrom, err := time.Parse("2006-01-02", from)
	if err != nil {
		return nil, fmt.Errorf("from is not a valid date")
	}
	dateTo, err := time.Parse("2006-01-02", to)
	if err != nil {
		return nil, fmt.Errorf("to is not a valid date")
	}
	if dateTo.Year() > time.Now().Year() {
		return nil, fmt.Errorf("dateTo.Year must be < Now().Year")
	}

	if dateFrom.Year() > dateTo.Year() {
		return nil, fmt.Errorf("from must be > to")
	} else if dateFrom.Year() == dateTo.Year() && dateFrom.YearDay() >= dateTo.YearDay() {
		return nil, fmt.Errorf("from must be > to")
	}
	return &model.DateLimit{
		From: dateFrom.Format("2006-01-02"),
		To:   dateTo.Format("2006-01-02"),
	}, nil
}

// FieldSortValid ...
func FieldSortValid(category string) bool {
	var fieldsToSort = [6]string{"event_date", "views", "clicks", "cost", "cpc", "cpm"}
	for _, field := range fieldsToSort {
		if field == category {
			return true
		}
	}
	return false
}

// RequestValidate ...
func RequestValidate(r *model.Request) (*model.StatisticsShow, error) {
	res := &model.StatisticsShow{}

	date, err := time.Parse("2006-01-02", r.Date)
	if err != nil {
		return nil, fmt.Errorf("is not a valid date")
	}
	res.Date = date.Format("2006-01-02")

	if r.Views != "" {
		viewsInt, err := strconv.ParseInt(r.Views, 10, 64)
		if err != nil {
			return nil, err
		}
		if viewsInt < 0 {
			return nil, fmt.Errorf("views must be > 0")
		}
		res.Views = viewsInt
	}
	if r.Clicks != "" {
		clicksInt, err := strconv.ParseInt(r.Clicks, 10, 64)
		if err != nil {
			return nil, err
		}
		if clicksInt < 0 {
			return nil, fmt.Errorf("clicks must be > 0")
		}
		res.Clicks = clicksInt
	}
	if r.Cost != "" {
		if len(r.Cost) != 4 {
			return nil, fmt.Errorf("cost must have two decimal places")
		}
		costFloat, err := strconv.ParseFloat(r.Cost, 64)
		if err != nil {
			return nil, err
		}
		if costFloat < 0 {
			return nil, fmt.Errorf("cost must be > 0")
		}
		res.Cost = costFloat
	}

	if r.Cost != "" && r.Clicks != "" {
		res.Cpc = res.Cost / float64(res.Clicks)
	}
	if r.Cost != "" && r.Views != "" {
		res.Cpm = res.Cost / float64(res.Views) * 1000
	}
	return res, nil
}

func checkDecimalPlaces(i int, value float64) bool {
	valueFloat := value * float64(math.Pow(10.0, float64(i)))
	extra := valueFloat - float64(int(valueFloat))
	fmt.Println(extra)
	return extra == 0
}
