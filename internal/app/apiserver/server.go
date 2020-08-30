package apiserver

import (
	"mitsoСhat/internal/app/model"
	. "encoding/json"
	"errors"

	"mitsoСhat/internal/app/store"
	"mitsoСhat/internal/app/store/sqlstore"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  sqlstore.Store
}

var (
	errorNotAuthorize           = errors.New("no authorization token")
	errIncorrectEmailOrPassword = errors.New("email or password not valid")
	TOKEN_SECRET_RULE           = "mitsoChatSecret"
)


func newServer(store sqlstore.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/api/get_user", s.getUser()).Methods("GET")
	s.router.HandleFunc("/api/get_installment_plan", s.getInstallmentPlan()).Methods("GET")
}

func (s *server) logRequest(w http.ResponseWriter, r *http.Request) {
	logger := s.logger
	logger.Infof("started %s %s %s", r.Method, r.RequestURI, r.Header)
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{
		"error": err.Error(),
	})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		_ = NewEncoder(w).Encode(data)
	}
}

func (s *server) getUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type response struct {
			Login string `json:"login"`
			Pass  string `json:"pass"`
		}
		logins, ok := r.URL.Query()["login"]
		if !ok || len(logins[0]) < 1 {
			s.error(w, r, http.StatusInternalServerError, store.ErrWrongRequset)
			return
		}
		login := logins[0]

		passwords, ok := r.URL.Query()["pass"]
		if !ok || len(passwords[0]) < 1 {
			s.error(w, r, http.StatusInternalServerError, store.ErrWrongRequset)
			return
		}
		pass := passwords[0]

		err, student := s.store.GetStudent(login, pass)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, student)
	}
}

func (s *server) getInstallmentPlan() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type response struct {
			Plans []*model.InstallmentPlan `json:"plans"`
		}
		logins, ok := r.URL.Query()["login"]
		if !ok || len(logins[0]) < 1 {
			s.error(w, r, http.StatusInternalServerError, store.ErrWrongRequset)
			return
		}
		login := logins[0]

		passwords, ok := r.URL.Query()["pass"]
		if !ok || len(passwords[0]) < 1 {
			s.error(w, r, http.StatusInternalServerError, store.ErrWrongRequset)
			return
		}
		pass := passwords[0]

		err, plans := s.store.GetStudentInstallmentPlan(login, pass)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		resp := response{
			Plans: plans,
		}

		s.respond(w, r, http.StatusOK, resp)
	}
}
