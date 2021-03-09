package store_test

import (
	"testing"

	"github.com/amartery/statSaver/internal/app/model"
	"github.com/amartery/statSaver/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestStatRepository_Add(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("statistics")

	stat, err := s.Stat().Add(&model.StatisticsShow{
		Date:   "2020-07-22",
		Views:  33,
		Clicks: 45,
		Cost:   1.76,
		Cpc:    1.76 / 45,
		Cpm:    1.76 / 33 * 1000,
	})
	assert.NoError(t, err)
	assert.NotNil(t, stat)
}

func TestStatRepository_Show(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("statistics")

	s.Stat().Add(&model.StatisticsShow{
		Date:   "2020-07-22",
		Views:  33,
		Clicks: 45,
		Cost:   1.76,
		Cpc:    1.76 / 45,
		Cpm:    1.76 / 33 * 1000,
	})

	s.Stat().Add(&model.StatisticsShow{
		Date:   "2020-07-22",
		Views:  45,
		Clicks: 456,
		Cost:   1.76,
		Cpc:    1.76 / 45,
		Cpm:    1.76 / 33 * 1000,
	})

	testDate := &model.DateLimit{
		From: "2000-01-01",
		To:   "2025-01-01",
	}
	stat, err := s.Stat().Show(testDate)
	assert.NoError(t, err)
	assert.NotNil(t, stat)
}
