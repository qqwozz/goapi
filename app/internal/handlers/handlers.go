package handlers

import (
	"encoding/json"
	"goapi/internal/database"
	"goapi/internal/models"
	"net/http"
	"strconv"
	"strings"
)

type Handlers struct {
	store *database.TaskStore
}

func NewHandler(store *database.TaskStore) *Handlers {
	return &Handlers{
		store: store,
	}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, map[string]string{"error": message})
}

func extractTaskID(r *http.Request) (int, error) {
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/tasks/"), "/")
	idString := pathParts[0]
	return strconv.Atoi(idString)
}

func handleNotFoundOrInternalError(w http.ResponseWriter, err error) {
	if strings.Contains(err.Error(), "record not found") {
		respondWithError(w, http.StatusNotFound, err.Error())
	} else {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *Handlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.store.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Ошибка получения задач")
		return
	}
	respondWithJSON(w, http.StatusOK, tasks)
}

func (h *Handlers) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id, err := extractTaskID(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "некорректный id задачи")
		return
	}

	task, err := h.store.GetById(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	var input models.CreateTaskInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "некорректно отправлены данные")
		return
	}

	if strings.TrimSpace(input.Title) == "" {
		respondWithError(w, http.StatusBadRequest, "заголовок задачи должен присутствовать")
		return
	}

	task, err := h.store.Create(input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, task)
}

func (h *Handlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := extractTaskID(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "некорректный id задачи")
		return
	}

	var input models.UpdateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "некорректные данные")
		return
	}

	if input.Title != nil && strings.TrimSpace(*input.Title) == "" {
		respondWithError(w, http.StatusBadRequest, "заголовок обязателен")
		return
	}

	task, err := h.store.Update(id, input)
	if err != nil {
		handleNotFoundOrInternalError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

func (h *Handlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := extractTaskID(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "некорректный id задачи")
		return
	}

	if err := h.store.Delete(id); err != nil {
		handleNotFoundOrInternalError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "success"})
}