package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Lupusdog/server-recruit-challenge-2024/model"
	"github.com/Lupusdog/server-recruit-challenge-2024/service"
	"github.com/gorilla/mux"
)

type albumController struct {
	service service.AlbumService
}

func NewAlbumController(s service.AlbumService) *albumController {
	return &albumController{service: s}
}

// GET /albums のハンドラー
func (c *albumController) GetAlbumListHandler(w http.ResponseWriter, r *http.Request) {
	albums, err := c.service.GetAlbumListService(r.Context())
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(albums)
}

// GET /albums/{id} のハンドラー
func (c *albumController) GetAlbumDetailHandler(w http.ResponseWriter, r *http.Request) {
	albumID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return 
	}

	album, err := c.service.GetAlbumService(r.Context(), model.AlbumID(albumID))
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return 
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(album)
}

// POST /albums のハンドラー
func (c *albumController) PostAlbumHandler(w http.ResponseWriter, r *http.Request) {
	var album *model.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		err = fmt.Errorf("invalid body param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return 
	}

	if err := c.service.PostAlbumService(r.Context(), album); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(album)
}

// DELETE /albums/{id} のハンドラー
func (c *albumController) DeleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
	albumID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	if err := c.service.DeleteAlbumService(r.Context(), model.AlbumID(albumID)); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.WriteHeader(204)
}