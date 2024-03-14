package rest

import (
	"encoding/json"
	"fmt"
	"strconv"
	"vk-film-library/internal/entity/actor"

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

	actorID, err := s.Usecase.Actors.CreateActor(ts, createActorParam)
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

	if err := s.Usecase.Actors.UpdateActor(ts, updateActorParam); err != nil {
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

	if err := s.Usecase.Actors.DeleteActor(ts, actorID); err != nil {
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
