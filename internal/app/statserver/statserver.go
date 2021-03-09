package statserver

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/amartery/statSaver/internal/app/model"
	"github.com/amartery/statSaver/internal/app/store"
	"github.com/amartery/statSaver/internal/app/validation"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// StatServer ...
type StatServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New ...
func New(config *Config) *StatServer {
	return &StatServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *StatServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting statistics server" + s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *StatServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *StatServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *StatServer) configureRouter() {
	s.router.HandleFunc("/stat/show", s.handleShow()).Methods("GET")
	s.router.HandleFunc("/stat/add", s.handleAdd()).Methods("POST")
	s.router.HandleFunc("/stat/del", s.handleDel()).Methods("DELETE")
}

func (s *StatServer) handleShow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleShow")
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")

		// валидация даты
		fromto, err := validation.DateValidate(from, to)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		var arrayStat []model.StatisticsShow
		fieldSort := r.URL.Query().Get("sort")

		if fieldSort != "" {
			if !validation.FieldSortValid(fieldSort) { // валидация метода сортировки
				io.WriteString(w, "error: field \"sort\" must be one of the(event_date, views, clicks, cost, cpc, cpm)")
				return
			}
			arrayStat, err = s.store.Stat().ShowOrdered(fromto, fieldSort)
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}
		} else {
			arrayStat, err = s.store.Stat().Show(fromto)
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}
		}
		arrayStatJSON, err := json.Marshal(arrayStat)
		io.WriteString(w, string(arrayStatJSON))
	}
}

func (s *StatServer) handleAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleAdd")

		req := &model.Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			io.WriteString(w, err.Error())
			return
		}
		statForBD, err := validation.RequestValidate(req)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		_, err = s.store.Stat().Add(statForBD)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		io.WriteString(w, "success")
	}
}

func (s *StatServer) handleDel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleDel")
		io.WriteString(w, "handlehandleDel")
	}
}
