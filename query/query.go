package query

import (
	"context"
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type QueryServerImpl struct {
	UnimplementedQueryServer
}

// GetLogs implements QueryServer.
func (*QueryServerImpl) GetLogs(context.Context, *LogsRequest) (*LogsResponse, error) {
	panic("unimplemented")
}

func StartQueryServer() error {
	grpcServer := grpc.NewServer()
	queryServer := QueryServerImpl{}
	RegisterQueryServer(grpcServer, &queryServer)

	qsPort := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", qsPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("query service listenting on port %d", qsPort)

	go func() {
		err = grpcServer.Serve(lis)
		// if err != nil {
		// 	return err
		// }
		panic(err)
	}() // TODO fix the error handling here!

	return nil
}
