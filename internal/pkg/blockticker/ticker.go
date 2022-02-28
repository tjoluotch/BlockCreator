package blockticker

import (
	"BlockCreator/internal/pkg/blockutils"
	"BlockCreator/internal/pkg/models"
	"crypto/sha1"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// genID returns a generated UUID as a string, if there was an
// error an empty string is returned followed by the error
func genID() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return id.String(), err
}

// genHash creates a sha1 hash from a blockId parameter and returns this as a BlockHash type
func genHash(blockId string) models.BlockHash {
	hash := sha1.New()
	hash.Write([]byte(blockId))
	blockHashBytes := hash.Sum(nil)
	fmt.Printf("Block hash: %x\n", string(blockHashBytes))
	return models.BlockHash(fmt.Sprintf("%x", blockHashBytes))
}

func createBlock(createdAt time.Time) (*models.Block, error) {
	blockId, err := genID()
	if err != nil {
		return nil, fmt.Errorf("failed to create block, block ID error: %v", err)
	}

	return &models.Block{
		Hash:         genHash(blockId),
		CreatedAt:    createdAt,
		LastModified: createdAt,
	}, nil
}

func BlockTicker() error {
	var block *models.Block
	ticker := time.NewTicker(blockutils.BLOCK_INTERVALS * time.Second)
	for t := range ticker.C {
		if b, err := createBlock(t); err != nil {
			return err
		} else {
			block = b
		}
		fmt.Printf("created new block: %+v\n\n", *block)

	}
	return nil
}
