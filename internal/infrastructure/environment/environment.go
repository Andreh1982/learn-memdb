package environment

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

// Single is the singleton instance of the environment
type Single struct {
	ENVIRONMENT         string // nolint: golint
	APP_VERSION         string // nolint: golint
	APP_PORT            string // nolint: golint
	APP_URL             string // nolint: golint
	AWS_ENDPOINT        string // nolint: golint
	AWS_REGION          string // nolint: golint
	AWS_PROFILE         string // nolint: golint
	LOG_LEVEL           string // nolint: golint
	DYNAMO_AWS_ENDPOINT string
	DYNAMO_TABLE_NAME   string

	DEFAULT_PERSISTENT bool
}

func init() {
	if os.Getenv("ENVIRONMENT") == "development" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Println("Error loading .env.local file")
		}
	}
	env := GetInstance()
	env.Setup()
}

func (e *Single) Setup() {
	e.ENVIRONMENT = os.Getenv("ENVIRONMENT")
	e.APP_VERSION = os.Getenv("APP_VERSION")
	e.APP_PORT = getenv("APPLICATION_PORT", "9001")
	e.APP_URL = getenv("APPLICATION_URL", "http://localhost")

	e.AWS_REGION = getenv("AWS_REGION", "us-east-1")
	e.AWS_ENDPOINT = getenv("AWS_ENDPOINT", "")
	e.AWS_PROFILE = getenv("AWS_PROFILE", "")

	e.LOG_LEVEL = getenv("LOG_LEVEL", "debug")

	e.DEFAULT_PERSISTENT = getenvBool("DEFAULT_PERSISTENT", "false")

	e.DYNAMO_AWS_ENDPOINT = getenv("DYNAMO_AWS_ENDPOINT", e.AWS_ENDPOINT)
	e.DYNAMO_TABLE_NAME = getenv("DYNAMO_TABLE_NAME", "learn-memdb")
}

func (e *Single) IsDevelopment() bool {
	return e.ENVIRONMENT == "development"
}

func getenvBool(key, fallback string) bool {
	value := getenv(key, fallback)
	valueBool, _ := strconv.ParseBool(value)
	return valueBool
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &Single{}
			singleInstance.Setup()
		} else {
			fmt.Println("Single instance already created.")
		}
	}
	return singleInstance
}
