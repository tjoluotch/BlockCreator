package models

import "time"

type Block struct {
	//sha1 hash
	Hash string
	CreatedAt time.Time
	LastModified time.Time
	//base64 encoded string
	TransactionID string
}
