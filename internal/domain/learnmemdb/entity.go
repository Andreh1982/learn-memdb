package learnmemdb

import "time"

type EntryEntity struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	EntryID   string    `json:"entryId"`
	Name      string    `json:"name"`
	IsFromSqs bool      `json:"isFromSqs"`
}
