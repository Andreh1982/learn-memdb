package worker

import "time"

// SqsEntryEntity is the interface that receive message from SQS
type SqsEntryEntity struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	EntryID   string    `json:"entryId"`
	Name      string    `json:"name"`
	IsFromSqs bool      `json:"isFromSqs"`
}
