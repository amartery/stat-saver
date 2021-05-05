package postgresDB

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/amartery/statSaver/internal/app/mocks"
	"github.com/amartery/statSaver/internal/app/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Cant create mock: %v", err)
	}
	statRep := NewStatRepository(db)
	defer db.Close()

	statModel := &models.StatisticsShow{
		Date:   "2000-01-01",
		Views:  2,
		Clicks: 5,
		Cost:   10.00,
		Cpc:    10.00 / float64(5),
		Cpm:    10.00 / float64(2) * 1000,
	}

	rows := sqlmock.NewRows([]string{"stat_id"}).AddRow(1)

	query := "INSERT INTO"
	mock.ExpectQuery(query).WithArgs(statModel.Date,
		statModel.Views,
		statModel.Clicks,
		statModel.Cost,
		statModel.Cpc,
		statModel.Cpm).WillReturnRows(rows)

	if err = statRep.Add(statModel); err != nil {
		t.Errorf("err: %v", err)
	}
	require.NoError(t, err)
}

func TestShow(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Cant create mock: %v", err)
	}
	statRep := NewStatRepository(db)
	defer db.Close()

	model1 := models.StatisticsShow{
		StatID: 1,
		Date:   "2001-01-01",
		Views:  22,
		Clicks: 41,
		Cost:   10.00,
		Cpc:    10.00 / float64(41),
		Cpm:    10.00 / float64(22) * 1000,
	}
	model2 := models.StatisticsShow{
		StatID: 2,
		Date:   "2005-01-01",
		Views:  22,
		Clicks: 41,
		Cost:   10.00,
		Cpc:    10.00 / float64(41),
		Cpm:    10.00 / float64(22) * 1000,
	}
	wantRes := []models.StatisticsShow{model1, model2}

	rows := sqlmock.NewRows([]string{"stat_id", "event_date", "views", "clicks", "cost", "cpc", "cpm"})
	rows.AddRow(model1.StatID, model1.Date, model1.Views, model1.Clicks, model1.Cost, model1.Cpc, model1.Cpm)
	rows.AddRow(model2.StatID, model2.Date, model2.Views, model2.Clicks, model2.Cost, model2.Cpc, model2.Cpm)

	req := &models.RequestForShow{
		From:      "2000-01-01",
		To:        "2010-01-01",
		SortField: "event_date",
	}

	query := "SELECT"
	mock.ExpectQuery(query).WithArgs(req.From, req.To).WillReturnRows(rows)

	reciveRes, err := statRep.ShowOrdered(req)

	require.NoError(t, err)
	require.Equal(t, &wantRes, reciveRes)
}

func TestShowOrdered(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Cant create mock: %v", err)
	}
	statRep := NewStatRepository(db)
	defer db.Close()
	model1 := models.StatisticsShow{
		StatID: 1,
		Date:   "2001-01-01",
		Views:  100,
		Clicks: 41,
		Cost:   10.00,
		Cpc:    10.00 / float64(41),
		Cpm:    10.00 / float64(100) * 1000,
	}
	model2 := models.StatisticsShow{
		StatID: 2,
		Date:   "2005-01-01",
		Views:  10,
		Clicks: 41,
		Cost:   10.00,
		Cpc:    10.00 / float64(41),
		Cpm:    10.00 / float64(10) * 1000,
	}
	model3 := models.StatisticsShow{
		StatID: 3,
		Date:   "2008-01-01",
		Views:  50,
		Clicks: 41,
		Cost:   10.00,
		Cpc:    10.00 / float64(41),
		Cpm:    10.00 / float64(50) * 1000,
	}
	wantRes := []models.StatisticsShow{model1, model2, model3}

	rows := sqlmock.NewRows([]string{"stat_id", "event_date", "views", "clicks", "cost", "cpc", "cpm"})
	rows.AddRow(model1.StatID, model1.Date, model1.Views, model1.Clicks, model1.Cost, model1.Cpc, model1.Cpm)
	rows.AddRow(model2.StatID, model2.Date, model2.Views, model2.Clicks, model2.Cost, model2.Cpc, model2.Cpm)
	rows.AddRow(model3.StatID, model3.Date, model3.Views, model3.Clicks, model3.Cost, model3.Cpc, model3.Cpm)

	req := &models.RequestForShow{
		From:      "2000-01-01",
		To:        "2010-01-01",
		SortField: "views",
	}

	query := "SELECT"
	mock.ExpectQuery(query).WithArgs(req.From, req.To).WillReturnRows(rows)

	reciveRes, err := statRep.ShowOrdered(req)

	require.NoError(t, err)
	require.Equal(t, &wantRes, reciveRes)
}

func TestClearStatistics(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mocks.NewMockRepository(ctrl)

	mock.EXPECT().ClearStatistics().Times(1).Return(nil)

	err := mock.ClearStatistics()
	require.NoError(t, err)
}
