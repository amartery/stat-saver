package statserver

import (
	"encoding/json"
	"fmt"
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
		w.Header().Set("Content-Type", "application/json")

		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")

		fromto, err := validation.DateValidate(from, to)
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("error data validation"))
			return
		}

		var arrayStat []model.StatisticsShow
		fieldSort := r.URL.Query().Get("sort")

		if fieldSort != "" {
			if !validation.FieldSortValid(fieldSort) {
				msg := fieldSort + "field doesn`t exist, available fields [event_date, views, clicks, cost, cpc, cpm]"
				s.error(w, r, http.StatusBadRequest, fmt.Errorf(msg))
				return
			}
			arrayStat, err = s.store.Stat().ShowOrdered(fromto, fieldSort)
			if err != nil {
				fmt.Println(err)
				s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error on the server"))
				return
			}
		} else {
			arrayStat, err = s.store.Stat().Show(fromto)
			if err != nil {
				fmt.Println(err)
				s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error on the server"))
				return
			}
		}
		result := make(map[string][]model.StatisticsShow)
		result["statistics"] = arrayStat
		s.respond(w, r, http.StatusOK, result)
	}
}

func (s *StatServer) handleAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleAdd")
		w.Header().Set("Content-Type", "application/json")

		req := &model.Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error on the server"))
			return
		}
		statForBD, err := validation.RequestValidate(req)
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("error request validate"))
			return
		}
		_, err = s.store.Stat().Add(statForBD)
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error on the server"))
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *StatServer) handleDel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleDel")

		err := s.store.Stat().ClearingStatistics()
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error on the server"))
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *StatServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *StatServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		resp, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error on the server"))
			return
		}
		io.WriteString(w, string(resp))
	}
}
