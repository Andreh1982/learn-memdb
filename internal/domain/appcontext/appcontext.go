package appcontext

import (
	"context"

	"learn-memdb/internal/infrastructure/logger/logwrapper"

	"github.com/gin-gonic/gin"
)

type ContextKey string

const (
	AppContextKey               ContextKey = "appContext"
	ginContextKey               ContextKey = "ginContext"
	defaultBackgroundContextKey ContextKey = "ctx"
)

type Context interface {
	Done()
	SetLogger(logger logwrapper.LoggerWrapper)
	Logger() logwrapper.LoggerWrapper
	Context() context.Context
}

func New(ctx context.Context, ginContext *gin.Context) Context {
	return &appContext{
		defaultBackgroundContext: ctx,
		ginContext:               ginContext,
	}
}

func NewBackground() Context {
	ctx := context.Background()

	return &appContext{
		defaultBackgroundContext: ctx,
		ginContext:               nil,
	}
}

func GetAppContext(c *gin.Context) Context {
	return c.MustGet(string(AppContextKey)).(Context)
}

type appContext struct {
	logger                   logwrapper.LoggerWrapper
	defaultBackgroundContext context.Context
	ginContext               *gin.Context
}

func (appContext *appContext) SetLogger(logger logwrapper.LoggerWrapper) {
	appContext.logger = logger
}

func (appContext *appContext) Logger() logwrapper.LoggerWrapper {
	return appContext.logger
}

func (appContext *appContext) Context() context.Context {
	return appContext.defaultBackgroundContext
}

func (appContext *appContext) Done() {
	appContext.ginContext = nil
	appContext.defaultBackgroundContext = nil
	appContext.logger = nil
}
