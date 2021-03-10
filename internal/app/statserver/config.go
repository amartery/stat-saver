package statserver

import "github.com/amartery/statSaver/internal/app/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}

// psql -h localhost statserver_db statserver
// insert into Statistic (event_date, views, clicks, cost, cpc, cpm) values ('2021-03-09', 25, 27, 1.67, 0.06185185, 66.8)
