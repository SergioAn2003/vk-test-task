package rest

import (
	"grpc-test/uimport"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Server struct {
	log *logrus.Logger
	uimport.UsecaseImports
}

func NewServer(log *logrus.Logger, ui uimport.UsecaseImports) *Server {
	return &Server{
		log:            log,
		UsecaseImports: ui,
	}
}

func (s *Server) Run() {
	router := http.NewServeMux()

	if err := http.ListenAndServe(":8080", router); err != nil {
		s.log.Errorln("не удалось начать прослушивание, ошибка:", err)
		return
	}
}
