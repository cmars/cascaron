package cascaron

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Service struct {
	*mux.Router
}

func NewService() *Service {
	s := &Service{Router: mux.NewRouter()}
	s.HandleFunc("/login", s.Login)
	s.HandleFunc("/logout", s.Logout)
	s.HandleFunc("/validate", s.Validate)
	s.HandleFunc("/serviceValidate", s.ServiceValidate)
	s.HandleFunc("/proxyValidate", s.ProxyValidate)
	s.HandleFunc("/proxy", s.Proxy)
	s.HandleFunc("/samlValidate", s.SamlValidate)
	return s
}

func (s *Service) Login(w http.ResponseWriter, r *http.Request) {
	notImplemented(w)
}

func (s *Service) Logout(w http.ResponseWriter, r *http.Request) {
	notImplemented(w)
}

func (s *Service) Validate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w)
}

func (s *Service) ServiceValidate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w)
}

func (s *Service) ProxyValidate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w)
}

func (s *Service) Proxy(w http.ResponseWriter, r *http.Request) {
	notImplemented(w)
}

func (s *Service) SamlValidate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w)
}
