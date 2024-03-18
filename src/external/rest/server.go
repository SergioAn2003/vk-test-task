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
	s.mux.HandleFunc("/signup", s.SignUp)
	s.mux.HandleFunc("/login", s.Login)

	s.mux.HandleFunc("/actor/create", s.AuthMiddleware(s.CreateActor))
	s.mux.HandleFunc("/actor/update", s.AuthMiddleware(s.UpdateActor))
	s.mux.HandleFunc("/actor/delete", s.AuthMiddleware(s.DeleteActor))
	s.mux.HandleFunc("/actor/load", s.AuthMiddleware(s.LoadActorList))

	s.mux.HandleFunc("/movie/create", s.AuthMiddleware(s.CreateMovie))
	s.mux.HandleFunc("/movie/update", s.AuthMiddleware(s.UpdateMovie))
	s.mux.HandleFunc("/movie/delete", s.AuthMiddleware(s.DeleteMovie))
	s.mux.HandleFunc("/movies", s.AuthMiddleware(s.GetMovieList))
	s.mux.HandleFunc("/movie/find", s.AuthMiddleware(s.FindMovieListByTitleAndActorName))

	s.log.Infoln("сервер успешно запущен на порту :9000")
	if err := http.ListenAndServe(":9000", s.mux); err != nil {
		s.log.Fatalln("не удалось начать прослушивание, ошибка:", err)
	}
}
