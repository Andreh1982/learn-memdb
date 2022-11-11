package learnmemdb

import "time"

type EntryEntity struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	EntryID   string    `json:"entry_id"`
	Name      string    `json:"name"`
}
