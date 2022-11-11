package learnmemdb

import (
	"learn-memdb/internal/domain/appcontext"

	"go.uber.org/zap"
)

type Lister interface {
	ListAll(ctx appcontext.Context) (*[]EntryEntity, error)
}

func (l *learnmemdb) ListAll(ctx appcontext.Context) (*[]EntryEntity, error) {
	logger := ctx.Logger()
	logger.Info("Listing entries", zap.String("where", "listall"))

	result, err := l.repository.List()

	return result, err
}
