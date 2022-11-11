package main

import (
	"learn-memdb/internal/domain/learnmemdb"
	"learn-memdb/internal/infrastructure/api"
	"learn-memdb/internal/infrastructure/aws"
	"learn-memdb/internal/infrastructure/database"
	"learn-memdb/internal/infrastructure/environment"
	"learn-memdb/internal/infrastructure/logger"
	"learn-memdb/internal/infrastructure/logger/logwrapper"

	"go.uber.org/zap"
)

func main() {

	env := environment.GetInstance()
	zaplogger, dispose := logger.New()
	defer dispose()

	logger := logwrapper.New(&logwrapper.Zap{Logger: *zaplogger}).Version(env.APP_VERSION)
	logger.Info("Starting Learn-MemDB APP")

	logger.Info("env",
		zap.String("LOG_LEVEL", env.LOG_LEVEL),
		zap.Bool("DEFAULT_PERSISTENT", env.DEFAULT_PERSISTENT),
		zap.String("APP_PORT", env.APP_PORT),
		zap.String("ENVIRONMENT", env.ENVIRONMENT),
	)

	learnmemdbUseCases, err := setupLearnMemdb()

	if err != nil {
		logger.Error("failed to setup Learn-MemDB", zap.Error(err))
	}

	setupApi(logger, learnmemdbUseCases)
}

func setupLearnMemdb() (learnmemdb.UseCases, error) {
	dynamodb, err := setupDynamoDB()
	if err != nil {
		return nil, err
	}

	memdbInput := &learnmemdb.Input{
		Repository: dynamodb,
	}
	learnmemdbUseCases := learnmemdb.New(memdbInput)
	return learnmemdbUseCases, nil
}

func setupDynamoDB() (learnmemdb.Repository, error) {
	env := environment.GetInstance()
	if !env.DEFAULT_PERSISTENT {
		return database.NewMemoryDatabase(), nil
	}

	awsRegion := env.AWS_REGION
	awsEndpoint := env.DYNAMO_AWS_ENDPOINT
	table := env.DYNAMO_TABLE_NAME
	cfg, err := aws.EndpointResolverWithOptionsFunc(awsEndpoint, awsRegion)
	if err != nil {
		return nil, err
	}
	return database.NewDynamoDB(cfg, table), nil
}

func setupApi(logger logwrapper.LoggerWrapper, learnmemdbUseCases learnmemdb.UseCases) {
	input := api.Input{
		Logger:             logger,
		LearnMemdbUseCases: learnmemdbUseCases,
	}
	api.Start(input)
}
