package rest

import (
	"net/http"
	"vk-film-library/uimport"

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
	s.mux.HandleFunc("/actor/create", s.CreateActor)
	s.mux.HandleFunc("/actor/update", s.UpdateActor)
	s.mux.HandleFunc("/actor/delete", s.DeleteActor)

	s.mux.HandleFunc("/movie/create", s.CreateMovie)
	s.mux.HandleFunc("/movie/update", s.UpdateMovie)
	s.mux.HandleFunc("/movie/delete", s.DeleteMovie)
	s.mux.HandleFunc("/movies", s.GetMovieList)

	s.log.Infoln("сервер успешно запущен на порту :9000")
	if err := http.ListenAndServe(":9000", s.mux); err != nil {
		s.log.Fatalln("не удалось начать прослушивание, ошибка:", err)
	}
}
