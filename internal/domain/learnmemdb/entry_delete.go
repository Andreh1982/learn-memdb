package learnmemdb

import (
	"fmt"
	"learn-memdb/internal/domain/appcontext"

	"go.uber.org/zap"
)

type Deleter interface {
	Delete(ctx appcontext.Context, entryID string) error
}

func (l *learnmemdb) Delete(ctx appcontext.Context, entryID string) error {

	logger := ctx.Logger()
	logger.Info("Deleting entry", zap.String("entryID", fmt.Sprint(entryID)), zap.String("where", "delete"))

	if entryID == "" {
		return DomainErrorFactory(BadRequest, "entryID is required")
	}

	err := l.repository.Delete(entryID)
	if err != nil {
		logger.Error("error reading entry", zap.Error(err), zap.String("where", "read"))
		return err
	}

	logger.Info("Deleted entry", zap.String("entryID", fmt.Sprint(entryID)), zap.String("where", "delete"))

	return err
}
