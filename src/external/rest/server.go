package rest

import (
	"grpc-test/uimport"

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
