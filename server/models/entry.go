package models

import (
	"time"
)

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	Entries_type string  `json:"entries_type"`
	Owner string `json:"owner"`
	Destination_account int64 `json: "destination_account"`
}

type EntryResponse struct {
	
	AccountID int64 `json:"account_id"`
	// can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	Entries_type string  `json:"entries_type"`
}
