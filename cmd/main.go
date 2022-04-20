package main

import (
	"BlockCreator/internal/pkg/blockresolver/client"
	blockresolver "BlockCreator/internal/pkg/blockresolver/configs/grpc"
	"BlockCreator/internal/pkg/blockticker"
	"BlockCreator/internal/pkg/blockutils"
	"flag"
	"fmt"
	"log"
)

var (
	grpcServerHostname = flag.String("grpcserverhost", "127.0.0.1", "The hostname of the grpc server that connects to the client stub within this application")
)

func main() {
	flag.Parse()
	// if set server host name to flag value. Either default of some hostname that has been passed to the flag
	if *grpcServerHostname != "" {
		blockutils.SERVER_HOSTNAME = *grpcServerHostname
	}
	fmt.Println("Block creator start up...")

	// TODO refactor to error handle & gracefully exit through channels
	go blockticker.BlockTicker()

	// This print statement will be executed before
	// the first block is created and prints in the console
	fmt.Println("Initiate block creation...")

	// grpc client initated to start streaming through block objects
	conn, err := client.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("grpc connection created")
	log.Printf("----------[%s]-----------\n", conn.GetState().String())

	grpcClient := blockresolver.NewBlockResolverClient(conn)

	// this should be a separate go routine if chan closes
	// underneath block until chan closes see sendnblocks impl
	pushChannel := make(chan struct{})
	go func() {
		for {
			// start the process of sending a block once signal has been received from empty struct channel
			//<-blockticker.SignalChan
			if len(blockticker.BlockChan) < 0 {
				continue
			}
			if err = client.PushBlock(grpcClient, client.ProcessBlock()); err != nil {
				//	close channel
				close(pushChannel)
				log.Printf("Failed to send a block successfully: %v\n", err)
				return
			}
		}
	}()
	<-pushChannel
	log.Println("test")
	// here we use an empty select{} in order to keep
	// our main function alive indefinitely as it would
	// complete before our BlockTicker has a chance
	// to execute if we didn't.
	// TODO refactor to use wait group
	select {}
}
