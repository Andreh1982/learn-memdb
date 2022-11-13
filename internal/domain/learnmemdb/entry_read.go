package learnmemdb

import (
	"fmt"
	"learn-memdb/internal/domain/appcontext"

	"go.uber.org/zap"
)

type Reader interface {
	Read(ctx appcontext.Context, entryID string) (learnmemdbEntity *EntryEntity, err error)
}

func (l *learnmemdb) Read(ctx appcontext.Context, entryID string) (learnmemdbEntity *EntryEntity, err error) {

	logger := ctx.Logger()
	logger.Info("Reading entry", zap.String("entryID", fmt.Sprint(entryID)), zap.String("where", "read"))

	if entryID == "" {
		return nil, DomainErrorFactory(BadRequest, "entryID is required")
	}

	result, err := l.repository.Find(entryID)
	if err != nil {
		logger.Error("error reading entry", zap.Error(err), zap.String("where", "read"))
		return nil, err
	}

	return result, nil
}
