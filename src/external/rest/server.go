package rest

import (
	"grpc-test/uimport"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Server struct {
	log *logrus.Logger
	mux *http.ServeMux
	uimport.UsecaseImports
}

func NewServer(log *logrus.Logger, ui uimport.UsecaseImports) *Server {
	return &Server{
		log:            log,
		UsecaseImports: ui,
		mux:            http.NewServeMux(),
	}
}

func (s *Server) Run() {
	s.mux.HandleFunc("/actor/create", s.CreateUser)

	s.log.Infoln("сервер успешно запущен на порту :9000")
	if err := http.ListenAndServe(":9000", s.mux); err != nil {
		s.log.Fatalln("не удалось начать прослушивание, ошибка:", err)
	}
}
