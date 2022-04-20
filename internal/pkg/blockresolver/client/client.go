package client

import (
	"BlockCreator/internal/pkg/blockutils"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func getServer() string {
	return fmt.Sprintf("%s:%d", blockutils.SERVER_HOSTNAME, blockutils.SERVER_ADDR)
}

func Start() (*grpc.ClientConn, error) {
	// creating gprc client stub
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(getServer(), opts...)
	if err != nil {
		log.Println("failed to open grpc client connection: error", err)
		return nil, err
	}
	return conn, nil
}
