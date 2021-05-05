package usecase

import (
	"testing"

	"github.com/amartery/statSaver/internal/app/mocks"
	"github.com/amartery/statSaver/internal/app/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	beforeUsecase := &models.RequestForSave{
		Date:   "2000-01-01",
		Views:  "22",
		Clicks: "41",
		Cost:   "9.22",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	mockUse := NewStatUsecase(mockRepo)

	afterUsecase := &models.StatisticsShow{
		StatID: 0,
		Date:   "2000-01-01",
		Views:  22,
		Clicks: 41,
		Cost:   9.22,
		Cpc:    9.22 / float64(41),
		Cpm:    9.22 / float64(22) * 1000,
	}

	mockRepo.EXPECT().Add(afterUsecase).Times(1).Return(nil)

	err := mockUse.Add(beforeUsecase)
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

	mockRepo := mocks.NewMockRepository(ctrl)
	mockUse := NewStatUsecase(mockRepo)

	timeLimit := &models.RequestForShow{
		From:      "2000-01-01",
		To:        "2010-01-01",
		SortField: "event_date",
	}

	mockRepo.EXPECT().ShowOrdered(timeLimit).Times(1).Return(&wantRes, nil)

	reciveRes, err := mockUse.Show(timeLimit)
	require.NoError(t, err)
	require.Equal(t, &wantRes, reciveRes)
}

func TestShowOrdered(t *testing.T) {
	statModel1 := models.StatisticsShow{
		Date:   "2001-01-01",
		Views:  22,
		Clicks: 41,
		Cost:   99.22,
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

	mockRepo := mocks.NewMockRepository(ctrl)
	mockUse := NewStatUsecase(mockRepo)

	timeLimit := &models.RequestForShow{
		From:      "2000-01-01",
		To:        "2010-01-01",
		SortField: "cost",
	}

	mockRepo.EXPECT().ShowOrdered(timeLimit).Times(1).Return(&wantRes, nil)

	reciveRes, err := mockUse.Show(timeLimit)
	require.NoError(t, err)
	require.Equal(t, &wantRes, reciveRes)
}

func TestClearStatistics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	mockUse := NewStatUsecase(mockRepo)

	mockRepo.EXPECT().ClearStatistics().Times(1).Return(nil)

	err := mockUse.ClearStatistics()
	require.NoError(t, err)
}
