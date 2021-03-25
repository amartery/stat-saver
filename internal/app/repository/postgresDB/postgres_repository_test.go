package postgresDB

import (
	"testing"

	"github.com/amartery/statSaver/internal/app/mocks"
	"github.com/amartery/statSaver/internal/app/models"
	"github.com/amartery/statSaver/internal/app/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	statModel := &models.StatisticsShow{
		Date:   "2000-01-01",
		Views:  22,
		Clicks: 41,
		Cost:   9.22,
		Cpc:    9.22 / float64(41),
		Cpm:    9.22 / float64(22) * 1000,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mocks.NewMockRepository(ctrl)

	mock.EXPECT().Add(statModel).Times(1).Return(nil)

	u := usecase.NewStatUsecase(mock)
	err := u.Add(statModel)
	require.NoError(t, err)
}

func TestShow(t *testing.T) {
	statModel1 := models.StatisticsShow{
		Date:   "2001-01-01",
		Views:  22,
		Clicks: 41,
		Cost:   9.22,
		Cpc:    9.22 / float64(41),
		Cpm:    9.22 / float64(22) * 1000,
	}
	statModel2 := models.StatisticsShow{
		Date:   "2005-01-01",
		Views:  22,
		Clicks: 41,
		Cost:   9.22,
		Cpc:    9.22 / float64(41),
		Cpm:    9.22 / float64(22) * 1000,
	}
	wantRes := []models.StatisticsShow{statModel1, statModel2}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mocks.NewMockRepository(ctrl)

	timeLimit := &models.DateLimit{
		From: "2000-01-01",
		To:   "2010-01-01",
	}

	mock.EXPECT().Show(timeLimit).Times(1).Return(wantRes, nil)

	u := usecase.NewStatUsecase(mock)
	reciveRes, err := u.Show(timeLimit)
	require.NoError(t, err)
	require.Equal(t, wantRes, reciveRes)
}

func TestShowOrdered(t *testing.T) {
	statModel1 := models.StatisticsShow{
		Date:   "2001-01-01",
		Views:  224,
		Clicks: 12,
		Cost:   9.22,
		Cpc:    9.22 / float64(41),
		Cpm:    9.22 / float64(22) * 1000,
	}
	statModel2 := models.StatisticsShow{
		Date:   "2005-01-01",
		Views:  22,
		Clicks: 1000,
		Cost:   9.22,
		Cpc:    9.22 / float64(41),
		Cpm:    9.22 / float64(22) * 1000,
	}
	wantRes := []models.StatisticsShow{statModel1, statModel2}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mocks.NewMockRepository(ctrl)

	timeLimit := &models.DateLimit{
		From: "2000-01-01",
		To:   "2010-01-01",
	}
	sortCategory := "clicks"
	mock.EXPECT().ShowOrdered(timeLimit, sortCategory).Times(1).Return(wantRes, nil)

	u := usecase.NewStatUsecase(mock)
	reciveRes, err := u.ShowOrdered(timeLimit, sortCategory)
	require.NoError(t, err)
	require.Equal(t, wantRes, reciveRes)
}

func TestClearStatistics(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mocks.NewMockRepository(ctrl)

	mock.EXPECT().ClearStatistics().Times(1).Return(nil)

	u := usecase.NewStatUsecase(mock)
	err := u.ClearStatistics()
	require.NoError(t, err)
}
