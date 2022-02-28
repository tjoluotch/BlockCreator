package main

import (
	"BlockCreator/internal/pkg/blockticker"
	"fmt"
)

func main() {
	fmt.Println("Block creator start up...")

	// TODO refactor to error handle & gracefully exit through channels
	go blockticker.BlockTicker()

	// This print statement will be executed before
	// the first block is created and prints in the console
	fmt.Println("Initiate block creation...")

	// here we use an empty select{} in order to keep
	// our main function alive indefinitely as it would
	// complete before our BlockTicker has a chance
	// to execute if we didn't.
	// TODO possibly refactor to use wait group
	select {}
}
