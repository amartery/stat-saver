package store

import (
	"fmt"

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

// Open ...
func (s *Store) Open() error {
	db, err := sqlx.Connect("postgres", s.config.DataBaseURL) // "host=localhost dbname=profiles_db user=server password=password sslmode=disable"
	if err != nil {
		fmt.Println("open db")
		return err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("ping db")
		return err
	}
	s.db = db
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
