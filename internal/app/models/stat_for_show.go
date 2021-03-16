package models

// StatisticsShow ...
type StatisticsShow struct {
	StatID int64   `json:"-"`
	Date   string  `json:"date"`
	Views  int64   `json:"views"`
	Clicks int64   `json:"clicks"`
	Cost   float64 `json:"cost"`
	Cpc    float64 `json:"cpc"`
	Cpm    float64 `json:"cpm"`
}
