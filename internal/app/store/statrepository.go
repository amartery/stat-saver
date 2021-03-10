package store

import (
	"fmt"

	"github.com/amartery/statSaver/internal/app/model"
)

// StatRepository ...
type StatRepository struct {
	store *Store
}

// Add ...
func (r *StatRepository) Add(s *model.StatisticsShow) (*model.StatisticsShow, error) {
	return s, r.store.db.QueryRow("insert into Statistic "+
		"(event_date, views, clicks, cost, cpc, cpm) "+
		"values ($1, $2, $3, $4, $5, $6) returning stat_id",
		s.Date,
		s.Views,
		s.Clicks,
		s.Cost,
		s.Cpc,
		s.Cpm).Scan(&s.StatID)
}

// Show ...
func (r *StatRepository) Show(d *model.DateLimit) ([]model.StatisticsShow, error) {
	var result = []model.StatisticsShow{}

	queryReq := fmt.Sprintf("select * from Statistic "+
		"where event_date >= '%s' and event_date <= '%s' order by %s;",
		d.From,
		d.To,
		"event_date")
	rows, err := r.store.db.Query(queryReq)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		currentStat := model.StatisticsShow{}
		if err := rows.Scan(
			&currentStat.StatID,
			&currentStat.Date,
			&currentStat.Views,
			&currentStat.Clicks,
			&currentStat.Cost,
			&currentStat.Cpc,
			&currentStat.Cpm,
		); err != nil {
			return nil, err
		}
		currentStat.Date = currentStat.Date[:10]
		result = append(result, currentStat)
	}
	return result, nil
}

// ShowOrdered ...
func (r *StatRepository) ShowOrdered(d *model.DateLimit, category string) ([]model.StatisticsShow, error) {
	var result = []model.StatisticsShow{}
	queryReq := fmt.Sprintf("select * from Statistic "+
		"where event_date >= '%s' and event_date <= '%s' order by %s;",
		d.From,
		d.To,
		category)

	rows, err := r.store.db.Query(queryReq)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		currentStat := model.StatisticsShow{}
		if err := rows.Scan(
			&currentStat.StatID,
			&currentStat.Date,
			&currentStat.Views,
			&currentStat.Clicks,
			&currentStat.Cost,
			&currentStat.Cpc,
			&currentStat.Cpm,
		); err != nil {
			return nil, err
		}
		currentStat.Date = currentStat.Date[:10]
		result = append(result, currentStat)
	}
	return result, nil
}

// ClearingStatistics ...
func (r *StatRepository) ClearingStatistics() error {
	_, err := r.store.db.Exec("truncate Statistic restart identity;")
	if err != nil {
		return err
	}
	return nil
}
