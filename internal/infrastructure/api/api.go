package api

import (
	"learn-memdb/internal/domain/learnmemdb"
	"learn-memdb/internal/infrastructure/api/middlewares"
	"learn-memdb/internal/infrastructure/api/routes"
	"learn-memdb/internal/infrastructure/environment"
	"learn-memdb/internal/infrastructure/logger/logwrapper"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Input struct {
	Logger             logwrapper.LoggerWrapper
	LearnMemdbUseCases learnmemdb.UseCases
}

func Start(input Input) {
	r := gin.New()

	logger := input.Logger
	logger.Info("Starting Learn-MemDB API")

	applicationPort := resolvePort()
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.ContextMiddleware())
	r.Use(middlewares.TraceMiddleware())
	r.Use(middlewares.Logger(logger))
	if !environment.GetInstance().IsDevelopment() {
		r.Use(middlewares.Recovery(true))
	}
	// r.Use(middlewares.Metrics(metricService))

	routes.MakeHealthRoute(r)
	routes.MakeMetricRoute(r)
	routes.MakeEntriesRoute(r, input.LearnMemdbUseCases)

	if err := r.Run(applicationPort); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}

func resolvePort() string {
	const CHAR string = ":"
	env := environment.GetInstance()
	port := env.APP_PORT
	fisrtChar := port[:1]
	if fisrtChar != CHAR {
		port = CHAR + port
	}
	return port
}
