package main

import (
	"net/http"
	"os"
	"path/filepath"
	"retro/config"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/pkg/browser"
	"go.uber.org/zap"
)

var logger *zap.Logger

func serve(content string, handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", content)
		handle.ServeHTTP(w, r)
	})
}

func main() {
	// set up logger
	logger, _ = zap.NewDevelopment()

	config.Initialize(logger)

	router := mux.NewRouter()

	// set up api endpoints
	router.HandleFunc("/api/game", getGameInformation)
	router.HandleFunc("/api/game/order", setGameOrder)
	router.HandleFunc("/api/metadata", getMetadata)
	router.HandleFunc("/api/metadata/mode", setGameMode)
	router.HandleFunc("/api/metadata/refresh", setRefresh)
	router.HandleFunc("/api/layout/list", getLayoutList)
	router.HandleFunc("/api/layout/add", addLayout)
	router.HandleFunc("/api/layout/remove", removeLayout)
	router.HandleFunc("/api/layout/update", updateLayout)
	router.HandleFunc("/api/layout/active", getActiveLayout)
	router.HandleFunc("/api/layout/active/set", setActiveLayout)
	router.HandleFunc("/api/profile/list", getProfiles)
	router.HandleFunc("/api/profile/add", addProfile)
	router.HandleFunc("/api/profile/remove", removeLayout)
	router.HandleFunc("/api/profile/active", getActiveProfile)
	router.HandleFunc("/api/profile/active/set", setActiveProfile)

	// calculate our working directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logger.Fatal("Could not calculate directory",
			zap.Error(err),
		)
	}
	// static directory to serve frontend
	fs := http.FileServer(http.Dir(dir + "/ui"))
	//http.Handle("/", fs)

	router.PathPrefix("/css").Handler(serve(
		"text/css; charset=utf-8",
		http.FileServer(http.Dir(dir+"/ui")),
	))
	router.PathPrefix("/js").Handler(serve(
		"text/javascript; charset=utf-8",
		http.FileServer(http.Dir(dir+"/ui")),
	))
	router.PathPrefix("/").Handler(fs)

	// wait a second then open frontend
	go func() {
		time.Sleep(1 * time.Second)
		logger.Info("Loading UI (http://localhost:3000/)")
		browser.OpenURL("http://localhost:3000/")
	}()

	// start the server itself
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         ":3000",
		Handler: handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(router),
	}
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatal("Closed http server",
			zap.Error(err),
		)
	}
}
