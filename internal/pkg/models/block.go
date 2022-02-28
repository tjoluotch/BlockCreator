package models

import "time"

type Block struct {
	//sha1 hash
	Hash         BlockHash
	CreatedAt    time.Time
	LastModified time.Time
}

type BlockHash string
