package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/vadskev/go_final_project/internal/config"
	"github.com/vadskev/go_final_project/internal/lib/logger"
	"github.com/vadskev/go_final_project/internal/transport/handlers/task"
	mwLogger "github.com/vadskev/go_final_project/internal/transport/middleware/logger"
	"go.uber.org/zap"
)

const (
	taskPath     = "/api/task"
	tasksPath    = "/api/tasks"
	nextDatePath = "/api/nextdate"
	taskDonePath = "/api/task/done"
	singPath     = "/api/signin"
)
const (
	ReadTimeout        = 4 * time.Second
	WriteTimeout       = 4 * time.Second
	IdleTimeout        = 60 * time.Second
	shutDownCtxTimeout = 1 * time.Second
)

type App struct {
	configProvider *configProvider
	httpServer     *http.Server
}

func NewApp() (*App, error) {
	app := &App{}

	err := app.loadDeps()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) loadDeps() error {
	err := a.loadConfig()
	if err != nil {
		return err
	}
	log.Println("Config loaded")

	err = a.loadConfigProvider()
	if err != nil {
		return err
	}
	log.Println("Config provider loaded")

	err = a.loadLogger()
	if err != nil {
		return err
	}
	logger.Info("Logger loaded")

	return nil
}

func (a *App) loadConfig() error {
	err := config.Load()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) loadConfigProvider() error {
	a.configProvider = newConfigProvider()
	return nil
}

func (a *App) loadLogger() error {
	err := logger.Init(a.configProvider.LogConfig().Level())
	if err != nil {
		return err
	}
	return nil
}

func (a *App) RunServer(ctx context.Context) error {
	router := chi.NewRouter()

	router.Use(mwLogger.New())

	router.Route(tasksPath, func(r chi.Router) {
		r.Get("/", task.NewTask())
	})

	a.httpServer = &http.Server{
		Addr:         a.configProvider.HTTPConfig().Address(),
		Handler:      router,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
		IdleTimeout:  IdleTimeout,
	}

	a.FileServer(router)

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Info("failed to start server")
			os.Exit(1)
		}
	}()

	logger.Info("HTTP server is running on ", zap.String("address", a.httpServer.Addr))

	// wait for gracefully shutdown
	<-ctx.Done()

	logger.Info("Shutting down server gracefully")

	shutDownCtx, cancel := context.WithTimeout(context.Background(), shutDownCtxTimeout)
	defer cancel()

	if err := a.httpServer.Shutdown(shutDownCtx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}
	<-shutDownCtx.Done()

	return nil
}

func (a *App) FileServer(router *chi.Mux) {
	root := "./web"
	fs := http.FileServer(http.Dir(root))

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}
