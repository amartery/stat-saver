package store

import (
	"fmt"
	"io/ioutil"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	config         *Config
	db             *sqlx.DB
	statRepository *StatRepository
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// InitDatabase ...
func (s *Store) InitDatabase() error {
	sqlInitFile, err := ioutil.ReadFile("scripts/statistics.sql")
	if err != nil {
		return err
	}
	_, err = s.db.Exec(string(sqlInitFile))
	if err != nil {
		return err
	}
	return nil
}

// Open ...
func (s *Store) Open() error {
	db, err := sqlx.Connect("postgres", s.config.DataBaseURL)
	if err != nil {
		fmt.Println("open db")
		return err
	}
	if err := db.Ping(); err != nil {
		fmt.Println("ping db")
		return err
	}
	s.db = db
	if err := s.InitDatabase(); err != nil {
		fmt.Println("init db")
		return err
	}
	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// Stat ...
func (s *Store) Stat() *StatRepository {
	if s.statRepository != nil {
		return s.statRepository
	}

	s.statRepository = &StatRepository{
		store: s,
	}

	return s.statRepository
}
