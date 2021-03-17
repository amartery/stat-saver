package postgresDB

import (
	"fmt"

	"github.com/amartery/statSaver/internal/app"
	"github.com/amartery/statSaver/internal/app/models"
	"github.com/jmoiron/sqlx"
)

// StatRepository ...
type StatRepository struct {
	con *sqlx.DB
}

func NewStatRepository(con *sqlx.DB) app.Repository {
	return &StatRepository{
		con: con,
	}
}

// Add ...
func (r *StatRepository) Add(s *models.StatisticsShow) error {
	return r.con.QueryRow("insert into Statistic "+
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
func (r *StatRepository) Show(d *models.DateLimit) ([]models.StatisticsShow, error) {
	var result = []models.StatisticsShow{}

	queryReq := fmt.Sprintf("select * from Statistic "+
		"where event_date >= '%s' and event_date <= '%s' order by %s;",
		d.From,
		d.To,
		"event_date")
	rows, err := r.con.Query(queryReq)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		currentStat := models.StatisticsShow{}
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
func (r *StatRepository) ShowOrdered(d *models.DateLimit, category string) ([]models.StatisticsShow, error) {
	var result = []models.StatisticsShow{}
	queryReq := fmt.Sprintf("select * from Statistic "+
		"where event_date >= '%s' and event_date <= '%s' order by %s;",
		d.From,
		d.To,
		category)

	rows, err := r.con.Query(queryReq)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		currentStat := models.StatisticsShow{}
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
func (r *StatRepository) ClearStatistics() error {
	_, err := r.con.Exec("truncate Statistic restart identity;")
	if err != nil {
		return err
	}
	return nil
}
