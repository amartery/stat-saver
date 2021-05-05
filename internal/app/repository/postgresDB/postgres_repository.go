package postgresDB

import (
	"database/sql"

	"github.com/amartery/statSaver/internal/app"
	"github.com/amartery/statSaver/internal/app/models"
)

// StatRepository ...
type StatRepository struct {
	con *sql.DB
}

func NewStatRepository(con *sql.DB) app.Repository {
	return &StatRepository{
		con: con,
	}
}

// Add ...
func (r *StatRepository) Add(s *models.StatisticsShow) error {
	query := `INSERT INTO Statistic (event_date, views, clicks, cost, cpc, cpm) 
			  VALUES ($1, $2, $3, $4, $5, $6) 
	          RETURNING stat_id;`
	return r.con.QueryRow(query,
		s.Date,
		s.Views,
		s.Clicks,
		s.Cost,
		s.Cpc,
		s.Cpm).Scan(&s.StatID)
}

// ShowOrdered ...
func (r *StatRepository) ShowOrdered(model *models.RequestForShow) (*[]models.StatisticsShow, error) {
	query := `SELECT * FROM Statistic 
			  WHERE event_date >= $1 AND event_date <= $2 
			  ORDER BY ` + model.SortField

	rows, err := r.con.Query(query, model.From, model.To)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]models.StatisticsShow, 0)
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
	return &result, nil
}

// ClearingStatistics ...
func (r *StatRepository) ClearStatistics() error {
	_, err := r.con.Exec("TRUNCATE Statistic RESTART IDENTITY;")
	if err != nil {
		return err
	}
	return nil
}
