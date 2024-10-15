package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/service"
)

type singerController struct {
	service service.SingerService
}

func NewSingerController(s service.SingerService) *singerController {
	return &singerController{service: s}
}

// GET /singers のハンドラー
func (c *singerController) GetSingerListHandler(w http.ResponseWriter, r *http.Request) {
	singers, err := c.service.GetSingerListService(r.Context())
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(singers)
}

// GET /singers/{id} のハンドラー
func (c *singerController) GetSingerDetailHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	singerID, err := strconv.Atoi(idString)
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	singer, err := c.service.GetSingerService(r.Context(), model.SingerID(singerID))
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(singer)
}

// POST /singers のハンドラー
func (c *singerController) PostSingerHandler(w http.ResponseWriter, r *http.Request) {
	var singer *model.Singer
	if err := json.NewDecoder(r.Body).Decode(&singer); err != nil {
		err = fmt.Errorf("invalid body param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	if err := c.service.PostSingerService(r.Context(), singer); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(singer)
}

// DELETE /singers/{id} のハンドラー
func (c *singerController) DeleteSingerHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	singerID, err := strconv.Atoi(idString)
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	if err := c.service.DeleteSingerService(r.Context(), model.SingerID(singerID)); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.WriteHeader(204)
}
