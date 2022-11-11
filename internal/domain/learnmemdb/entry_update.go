package learnmemdb

import "learn-memdb/internal/domain/appcontext"

type Updater interface {
	Update(ctx appcontext.Context, entryEntity EntryEntity) (*string, error)
}

func (l *learnmemdb) Update(ctx appcontext.Context, entryEntity EntryEntity) (*string, error) {
	return nil, nil
}
