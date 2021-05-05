package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amartery/statSaver/internal/app"
	"github.com/amartery/statSaver/internal/app/middleware"
	"github.com/amartery/statSaver/internal/app/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// StatServer ...
type StatServer struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	usecase app.Usecase
}

// New ...
func New(config *Config, statUsecase app.Usecase) *StatServer {
	return &StatServer{
		config:  config,
		logger:  logrus.New(),
		router:  mux.NewRouter(),
		usecase: statUsecase,
	}
}

// Start ...
func (s *StatServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

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

func (s *StatServer) configureRouter() {
	s.router.Use(middleware.PanicMiddleware)
	s.router.Use(middleware.ContentTypeJson)
	s.router.HandleFunc("/stat/show", s.handleShow).Methods("GET")
	s.router.HandleFunc("/stat/add", s.handleAdd).Methods("POST")
	s.router.HandleFunc("/stat/del", s.handleDel).Methods("DELETE")
}

func (s *StatServer) handleShow(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("starting handleShow")

	forShow := &models.RequestForShow{
		From:      r.URL.Query().Get("from"),
		To:        r.URL.Query().Get("to"),
		SortField: r.URL.Query().Get("sort"),
	}

	err := forShow.Validate()
	if err != nil {
		fmt.Println(err)
		s.error(w, r, http.StatusBadRequest, fmt.Errorf("validation: "+err.Error()))
		return
	}

	arrayStat, err := s.usecase.Show(forShow)
	if err != nil {
		fmt.Println(err)
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error on the server"))
		return
	}
	s.respond(w, r, http.StatusOK, map[string]*[]models.StatisticsShow{"statistics": arrayStat})
}

func (s *StatServer) handleAdd(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("starting handleAdd")

	req := &models.RequestForSave{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Println(err)
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error on the server"))
		return
	}
	err := req.Validate()
	if err != nil {
		fmt.Println(err)
		s.error(w, r, http.StatusBadRequest, fmt.Errorf("validation: "+err.Error()))
		return
	}
	err = s.usecase.Add(req)
	if err != nil {
		fmt.Println(err)
		s.error(w, r, http.StatusBadRequest, fmt.Errorf("validation: "+err.Error()))
		return
	}
	s.respond(w, r, http.StatusOK, nil)
}

func (s *StatServer) handleDel(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("starting handleDel")

	err := s.usecase.ClearStatistics()
	if err != nil {
		fmt.Println(err)
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error on the server"))
		return
	}
	s.respond(w, r, http.StatusOK, nil)

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
