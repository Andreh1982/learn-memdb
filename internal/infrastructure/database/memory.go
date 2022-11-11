package database

import (
	"fmt"
	"learn-memdb/internal/domain/learnmemdb"
)

func NewMemoryDatabase() learnmemdb.Repository {
	return &memoryDatabase{
		records: make(map[string]*learnmemdb.EntryEntity),
	}
}

type memoryDatabase struct {
	records map[string]*learnmemdb.EntryEntity
}

func (m *memoryDatabase) Find(key string) (*learnmemdb.EntryEntity, error) {
	record := m.records[key]

	if record == nil {
		return nil, nil
	}

	return record, nil
}

func (m *memoryDatabase) Insert(entryEntity learnmemdb.EntryEntity) (*learnmemdb.EntryEntity, error) {
	m.records[fmt.Sprint(entryEntity.EntryID)] = &entryEntity

	return &entryEntity, nil
}

func (m *memoryDatabase) Upsert(applicationEntity learnmemdb.EntryEntity) (*learnmemdb.EntryEntity, error) {
	return m.Insert(applicationEntity)
}

func (m *memoryDatabase) Delete(key string) error {
	delete(m.records, key)

	return nil
}

func (m *memoryDatabase) List() (*[]learnmemdb.EntryEntity, error) {
	var records []learnmemdb.EntryEntity

	for _, record := range m.records {
		records = append(records, *record)
	}

	return &records, nil
}
