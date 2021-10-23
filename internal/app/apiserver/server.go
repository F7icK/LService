package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/F7icK/LService/internal/app/model"
	"github.com/F7icK/LService/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store: store,
	}

	s.configureRouter()

	return s
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.logRequest)
	s.router.HandleFunc("/create", s.handleUsersCreate())
	s.router.HandleFunc("/users", s.handleUsersAllSelect())
	s.router.HandleFunc("/users/{id:[0-9]+}", s.handleUsersDelete())
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type  request struct {
		Name string `json:"name"`
		Surname string `json:"surname"`
		Telephone string `json:"telephone"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Name: req.Name,
			Surname: req.Surname,
			Telephone: req.Telephone,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleUsersAllSelect() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		telephone := r.URL.Query().Get("telephone")

		if telephone != "" {
			s.handleFindByTelephone()(w, r)
			return
		}

		us, err := s.store.User().AllSelect()
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, us)
	}
}

func (s *server) handleFindByTelephone() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		telephone := r.URL.Query().Get("telephone")

		fmt.Println(telephone)
		telephoneRune := []rune(telephone)
		if telephoneRune[0] == ' ' {
			telephoneRune[0] = '+'
		}
		telephone = string(telephoneRune)

		us, err := s.store.User().FindByTelephone(telephone)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w, r, http.StatusOK, us)
	}
}

func (s *server) handleUsersDelete() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		us, err := s.store.User().DeleteFromID(id)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w, r, http.StatusOK, us)
	}
}

func (s *server) error (w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
