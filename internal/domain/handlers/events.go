package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/orensimple/otus_hw1_8/internal/domain/models"
	"github.com/orensimple/otus_hw1_8/internal/domain/services"
	"github.com/orensimple/otus_hw1_8/internal/logger"
	"github.com/orensimple/otus_hw1_8/internal/maindb"
	"github.com/orensimple/otus_hw1_8/internal/memory"
	"github.com/spf13/viper"
)

type Handler struct {
	Handlers           *http.Handler
	MainEventStorage   *memory.MemEventStorage
	MainEventService   *services.EventService
	MaindbEventStorage *maindb.PgEventStorage
}

func (h *Handler) CreateEvent(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	id, tSt, tEn, err := validate(req)
	if err != nil {
		resp.WriteHeader(400)
		return
	}

	event, err := h.MainEventService.CreateEvent(req.Context(), id, req.Form.Get("owner"), req.Form.Get("title"), req.Form.Get("text"), tSt, tEn)
	if err == nil {
		data["result"] = event
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func (h *Handler) UpdateEvent(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	id, tSt, tEn, err := validate(req)
	if err != nil {
		resp.WriteHeader(400)
		return
	}

	event, err := h.MainEventService.UpdateEvent(req.Context(), id, req.Form.Get("owner"), req.Form.Get("title"), req.Form.Get("text"), tSt, tEn)
	if err == nil {
		data["result"] = event
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func (h *Handler) DeleteEvent(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(400)
		return
	}
	id, _ := strconv.ParseInt(req.Form.Get("id"), 10, 64)
	if id == 0 {
		resp.WriteHeader(400)
		return
	}
	err = h.MainEventService.DeleteEvent(req.Context(), id)
	if err == nil {
		data["result"] = "success"
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func (h *Handler) GetEventsByDay(resp http.ResponseWriter, req *http.Request) {
	data := make(map[string][]*models.Event)
	result, _ := h.MainEventService.GetEventsByTime(req.Context(), "day")
	data["result"] = result
	response, _ := json.Marshal(data)

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.WriteHeader(200)
	resp.Write(response)
}

func (h *Handler) GetEventsByWeek(resp http.ResponseWriter, req *http.Request) {
	data := make(map[string][]*models.Event)
	result, _ := h.MainEventService.GetEventsByTime(req.Context(), "week")
	data["result"] = result
	response, _ := json.Marshal(data)

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.WriteHeader(200)
	resp.Write(response)
}

func (h *Handler) GetEventsByMonth(resp http.ResponseWriter, req *http.Request) {
	data := make(map[string][]*models.Event)
	result, _ := h.MainEventService.GetEventsByTime(req.Context(), "month")
	data["result"] = result
	response, _ := json.Marshal(data)

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.WriteHeader(200)
	resp.Write(response)
}

func (h *Handler) AddForTest() {
	h.MainEventStorage = memory.NewMemEventStorage()

	h.MainEventService = &services.EventService{
		EventStorage: h.MainEventStorage,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	tSt, _ := time.Parse("2006-01-02 15:04", "2019-11-25 20:00")
	tEn, _ := time.Parse("2006-01-02 15:04", "2019-11-25 20:59")
	h.MainEventService.CreateEvent(ctx, 1, `a`, `b`, `c`, tSt, tEn)
	h.MainEventService.CreateEvent(ctx, 2, `a`, `b`, `c`, tSt, tEn)
}

func (h *Handler) InitDB() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/events", viper.GetString("postgres.user"), viper.GetString("postgres.passwd"), viper.GetString("postgres.ip"), viper.GetString("postgres.port"))

	var err error
	h.MaindbEventStorage, err = maindb.NewPgEventStorage(dsn)
	if err != nil {
		logger.ContextLogger.Infof("Problem connect to db", dsn, err.Error())
	}

	h.MainEventService = &services.EventService{
		EventStorage: h.MaindbEventStorage,
	}
}

func validate(req *http.Request) (int64, time.Time, time.Time, error) {
	err := req.ParseForm()
	if err != nil {
		logger.ContextLogger.Infof("form", "uri", err)
		return 0, time.Now(), time.Now(), err
	}

	id, _ := strconv.ParseInt(req.Form.Get("id"), 10, 64)
	if id == 0 {
		logger.ContextLogger.Infof("id", "uri", err)
		return 0, time.Now(), time.Now(), err
	}

	tSt, err := time.Parse("2006-01-02", req.Form.Get("tSt"))
	if err != nil {
		logger.ContextLogger.Infof("st", "uri", err)
		return 0, time.Now(), time.Now(), err
	}
	tEn, err := time.Parse("2006-01-02", req.Form.Get("tEn"))
	if err != nil {
		logger.ContextLogger.Infof("st", "uri", err)
		return 0, time.Now(), time.Now(), err
	}

	return id, tSt, tEn, err
}
