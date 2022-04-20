package client

import (
	blockresolver "BlockCreator/internal/pkg/blockresolver/configs/grpc"
	"BlockCreator/internal/pkg/blockticker"
	"context"
	"log"
	"time"
)

//func pBB()
func ProcessBlock() *blockresolver.Block {
	block := <-blockticker.BlockChan
	return &blockresolver.Block{
		Hash:         []byte(block.Hash),
		CreatedAt:    block.CreatedAt.String(),
		LastModified: block.LastModified.String(),
	}
}

// PushBlock receives a block object through the channel from the other go routine and sends it
func PushBlock(client blockresolver.BlockResolverClient, block *blockresolver.Block) error {
	log.Println("Received block from channel:", string(block.Hash))
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	newBlockId, err := client.SendBlock(ctx, block)
	if err != nil {
		log.Println("%v.SendBlock(_) = _, %v\n", client, err)
		return err
	}
	log.Printf("Got Block ID: [%s]\n", newBlockId.Id)
	return nil
}

// TODO on project extension look to add a bidirectional stream
//// RunSendBlocks receives a block from the chan and sends it
//func RunSendBlocks(client blockresolver.BlockResolverClient) {
//	//var blockIdInterface interface{}
//	blockProcessor = make(chan blockresolver.Block, 1)
//	block := <-blockticker.BlockChan
//	newBlock := blockresolver.Block{Hash: []byte(block.Hash), CreatedAt: block.CreatedAt.String(),
//		LastModified: block.LastModified.String()}
//	log.Println("Received block from channel:", string(newBlock.Hash))
//	blockProcessor <- newBlock
//
//	//ctx, cancel := context.WithTimeout(context.TODO(), 100*time.Second)
//	//defer cancel()
//
//	stream, err := client.SendBlocks(context.TODO())
//	if err != nil {
//		log.Fatalf("%v.SendBlocks(_) = _, %v\n", client, err)
//	}
//	waitChan := make(chan struct{})
//
//	go func() {
//		for {
//			// received input from blockprocessor
//			<-blockProcessor
//			if err := stream.Send(&newBlock); err != nil {
//				log.Fatalf("Failed to send a block: %v\n", err)
//			}
//
//			newBlockId, err := stream.Recv()
//			if err == io.EOF {
//				close(waitChan)
//				return
//			}
//			if err != nil {
//				log.Fatalf("Failed to receive an ID : %v\n", err)
//			}
//
//			log.Printf("Got Block ID: [%s]\n", newBlockId.Id)
//			close(waitChan)
//			return
//		}
//	}()
//
//	//if err := stream.Send(&newBlock); err != nil {
//	//	log.Fatalf("Failed to send a block: %v\n", err)
//	//}
//	<-waitChan
//	//stream.CloseSend()
//}
