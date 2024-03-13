package rest

import (
	"encoding/json"
	"grpc-test/internal/entity/actor"
	"net/http"
)

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	var act actor.Actor

	if err := json.NewDecoder(r.Body).Decode(&act); err != nil {
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

	if err := s.Usecase.Actors.CreateUser(ts, act); err != nil {
		return
	}

	if err := ts.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Errorln("не удалось закрыть транзакцию, ошибка:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
