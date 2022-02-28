package models

import "time"

type Block struct {
	//sha1 hash
	Hash         BlockHash
	CreatedAt    time.Time
	LastModified time.Time
	// TODO Id in later version of the project will be updated following reception from python micro service through rpc
	//  thus changing the above LastModified Block variable
	//Id string
}

type BlockHash string
