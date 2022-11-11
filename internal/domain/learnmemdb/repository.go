package learnmemdb

type RepositoryReader interface {
	Find(entryID string) (*EntryEntity, error)
	List() (*[]EntryEntity, error)
}

type RepositoryWriter interface {
	Insert(entryEntity EntryEntity) (*EntryEntity, error)
	Delete(entryID string) error
	Upsert(entryEntity EntryEntity) (*EntryEntity, error)
}

type Repository interface {
	RepositoryReader
	RepositoryWriter
}
