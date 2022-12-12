package worker

import (
	"learn-memdb/internal/domain/appcontext"
	"learn-memdb/internal/domain/learnmemdb"
	"learn-memdb/internal/infrastructure/logger/logwrapper"
)

// Input is the input to the worker
type Input struct {
	Logger             logwrapper.LoggerWrapper
	LearnMemdbUseCases learnmemdb.UseCases
}

// Start starts the worker
func Start(input Input) {
	appctx := appcontext.NewBackground()
	appctx.SetLogger(input.Logger)
	go createPolling(appctx, input, new(consumerWhoCreates))
	go createPolling(appctx, input, new(consumerWhoDeletes))
}
