package api

import (
	"net/http"

	"github.com/pulse227/server-recruit-challenge-sample/api/middleware"
	"github.com/pulse227/server-recruit-challenge-sample/controller"
	"github.com/pulse227/server-recruit-challenge-sample/infra/mysqldb"
	"github.com/pulse227/server-recruit-challenge-sample/service"
)

func NewRouter(
	dbUser, dbPass, dbHost, dbName string,
) (http.Handler, error) {
	dbClient, err := mysqldb.Initialize(dbUser, dbPass, dbHost, dbName)
	if err != nil {
		return nil, err
	}
	// 接続確認
	if err := dbClient.Ping(); err != nil {
		return nil, err
	}

	singerRepo := mysqldb.NewSingerRepository(dbClient)
	singerService := service.NewSingerService(singerRepo)
	singerController := controller.NewSingerController(singerService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /singers", singerController.GetSingerListHandler)
	mux.HandleFunc("GET /singers/{id}", singerController.GetSingerDetailHandler)
	mux.HandleFunc("POST /singers", singerController.PostSingerHandler)
	mux.HandleFunc("DELETE /singers/{id}", singerController.DeleteSingerHandler)

	wrappedMux := middleware.LoggingMiddleware(mux)

	return wrappedMux, nil
}
