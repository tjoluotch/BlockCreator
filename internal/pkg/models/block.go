package models

import "time"

type Block struct {
	Hash string
	createdAt time.Time
	lastModified time.Time
}
