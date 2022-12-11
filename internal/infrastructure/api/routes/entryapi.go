package routes

import (
	"fmt"
	"net/http"

	"learn-memdb/internal/domain/appcontext"
	"learn-memdb/internal/domain/learnmemdb"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func MakeEntriesRoute(r *gin.Engine, managerUseCases learnmemdb.UseCases) {
	grp := r.Group("/entries")

	grp.POST("/", func(c *gin.Context) {
		createEntry(c, managerUseCases)
	})

	grp.DELETE("/:entry_id", func(c *gin.Context) {
		deleteEntry(c, managerUseCases)
	})

	grp.GET("/:entry_id", func(c *gin.Context) {
		readEntry(c, managerUseCases)
	})

	grp.GET("/all", func(c *gin.Context) {
		listEntries(c, managerUseCases)
	})
}

func createEntry(c *gin.Context, managerUseCases learnmemdb.UseCases) {
	context := getContext(c)
	var learnmemdbEntity learnmemdb.EntryEntity
	err := c.ShouldBind(&learnmemdbEntity)
	if err != nil {
		context.Logger().Error("error binding json", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	result, err := managerUseCases.Create(context, learnmemdbEntity)
	respond(c, result, err)
}

func deleteEntry(c *gin.Context, managerUseCases learnmemdb.UseCases) {
	context := getContext(c)
	entryID := c.Param("entry_id")
	_ = managerUseCases.Delete(context, entryID)
	c.JSON(http.StatusOK, gin.H{"body": "deleted " + entryID})
}

func readEntry(c *gin.Context, managerUseCases learnmemdb.UseCases) {
	context := getContext(c)
	entryID := c.Param("entry_id")
	result, err := managerUseCases.Read(context, entryID)
	respond(c, result, err)
}

func listEntries(c *gin.Context, managerUseCases learnmemdb.UseCases) {
	context := getContext(c)
	result, err := managerUseCases.ListAll(context)
	respond(c, result, err)
}

func getContext(c *gin.Context) appcontext.Context {
	return c.Value(string(appcontext.AppContextKey)).(appcontext.Context)
}

func respond(c *gin.Context, result interface{}, err error) {
	if err != nil {
		if re, ok := err.(*learnmemdb.DomainError); ok {
			fmt.Printf("re: %v\n", re)
			c.JSON(re.StatusCode, gin.H{"error": re.Err.Error(), "retryable": re.Retryable, "message": re.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "retryable": true})
		}
		return
	}
	c.JSON(http.StatusOK, result)
}
