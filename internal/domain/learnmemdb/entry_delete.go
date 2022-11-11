package learnmemdb

import "learn-memdb/internal/domain/appcontext"

type Deleter interface {
	Delete(ctx appcontext.Context, entryID string) (*string, error)
}

func (l *learnmemdb) Delete(ctx appcontext.Context, entryID string) (*string, error) {
	return nil, nil
}
