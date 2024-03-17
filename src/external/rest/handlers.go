package rest

import (
	"encoding/json"
	"fmt"
	"strconv"
	"vk-film-library/internal/entity/actor"
	"vk-film-library/internal/entity/movie"

	"net/http"
)

func (s *Server) CreateActor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	var createActorParam actor.CreateActorParam

	if err := json.NewDecoder(r.Body).Decode(&createActorParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ts := s.SessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось открыть транзакцию, ошибка:", err)
		return
	}
	defer ts.Rollback()

	actorID, err := s.Usecase.Actor.CreateActor(ts, createActorParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ts.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось закрыть транзакцию, ошибка:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "актер успешно добавлен, id актера = %d", actorID)
}

func (s *Server) UpdateActor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	var updateActorParam actor.UpdateActorParam

	if err := json.NewDecoder(r.Body).Decode(&updateActorParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ts := s.SessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось открыть транзакцию, ошибка:", err)
		return
	}
	defer ts.Rollback()

	if err := s.Usecase.Actor.UpdateActor(ts, updateActorParam); err != nil {
		return
	}

	if err := ts.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось закрыть транзакцию, ошибка:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "данные актера успешно изменены")
}

func (s *Server) DeleteActor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	actorID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		s.log.Errorln("не удалось получить id актера, ошибка:", err)
		return
	}

	ts := s.SessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось открыть транзакцию, ошибка:", err)
		return
	}
	defer ts.Rollback()

	if err := s.Usecase.Actor.DeleteActor(ts, actorID); err != nil {
		return
	}

	if err := ts.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось закрыть транзакцию, ошибка:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "актер успешно удален")
}

func (s *Server) CreateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	var createMovieParam movie.CreateMovieParam

	if err := json.NewDecoder(r.Body).Decode(&createMovieParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ts := s.SessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось открыть транзакцию, ошибка:", err)
		return
	}
	defer ts.Rollback()

	movieID, err := s.Usecase.Movie.CreateMovie(ts, createMovieParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ts.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось закрыть транзакцию, ошибка:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "фильм успешно добавлен, id фильма = %d", movieID)
}

func (s *Server) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	var updateMovieParam movie.UpdateMovieParam

	if err := json.NewDecoder(r.Body).Decode(&updateMovieParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ts := s.SessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось открыть транзакцию, ошибка:", err)
		return
	}
	defer ts.Rollback()

	if err := s.Usecase.Movie.UpdateMovie(ts, updateMovieParam); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ts.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось закрыть транзакцию, ошибка:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "данные фильма успешно изменены")
}

func (s *Server) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	movieID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		s.log.Errorln("не удалось получить id фильма, ошибка:", err)
		return
	}

	ts := s.SessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось открыть транзакцию, ошибка:", err)
		return
	}
	defer ts.Rollback()

	if err := s.Usecase.Movie.DeleteMovie(ts, movieID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ts.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось закрыть транзакцию, ошибка:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "фильм успешно удален")
}

func (s *Server) GetMovieList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	sortString := r.URL.Query().Get("sort_by")

	ts := s.SessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось открыть транзакцию, ошибка:", err)
		return
	}
	defer ts.Rollback()

	movieList, err := s.Usecase.Movie.GetMovieList(ts, sortString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ts.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось закрыть транзакцию, ошибка:", err)
		return
	}

	if err := json.NewEncoder(w).Encode(movieList); err != nil {
		http.Error(w, "не удалось отправить данные с сервера", http.StatusInternalServerError)
		return
	}
}
