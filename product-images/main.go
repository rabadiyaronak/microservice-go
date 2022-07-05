package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/rabadiyaronak/microservice-go/product-images/files"
	"github.com/rabadiyaronak/microservice-go/product-images/handlers"

	"github.com/gorilla/mux"
	hclog "github.com/hashicorp/go-hclog"
)

func main() {
	var logLevel = getLogLevel()
	var basePathForResourceServer = getBasePath()

	logger := hclog.New(&hclog.LoggerOptions{
		Name:  "Product-images",
		Level: hclog.LevelFromString(logLevel),
	})

	//Create server logger from default logger
	serverLogger := logger.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

	//create local image store
	//max file size 5 MB
	store, err := files.NewLocal(basePathForResourceServer, 1024*5*1000)
	if err != nil {
		logger.Error("unable to create storage ", "error", err)
		os.Exit(-1)
	}

	//create file handler
	fh := handlers.NewFiles(store, logger)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	//handle CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	//GZip handler Middleware
	mw := handlers.NewGZipHandler(logger)

	// filename regex: {filename:[a-zA-Z]+\\.[a-z]{3}}
	// problem with FileServer is that it is dumb
	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.UploadREST)
	ph.HandleFunc("/", fh.UploadMultipart)

	// get files
	gh := sm.Methods(http.MethodGet).Subrouter()
	gh.Handle(
		"/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
		http.StripPrefix("/images/", http.FileServer(http.Dir(basePathForResourceServer))),
	)
	gh.Use(mw.GZipMiddleware)

	s := http.Server{
		Addr:         ":9091",
		Handler:      ch(sm),
		ErrorLog:     serverLogger,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Info("Starting server", "bind_address", ":9091")

		err := s.ListenAndServe()
		if err != nil {
			logger.Error("Unable to start server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	logger.Info("Shutting down server with", "signal", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}

func getLogLevel() string {
	var logLevel, isAvailable = os.LookupEnv("LOG_LEVEL")

	if !isAvailable {
		logLevel = "INFO"
	}

	return logLevel
}

func getBasePath() string {
	var basePath, isAvailable = os.LookupEnv("BASE_PATH")

	if !isAvailable {
		basePath = "./imageStore"
	}

	return basePath
}
