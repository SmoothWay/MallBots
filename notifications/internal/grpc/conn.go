package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Dial(ctx context.Context, endpoint string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("failed to close grpc connection: %v", err)
		}
		go func() {
			<-ctx.Done()
			if err := conn.Close(); err != nil {
				log.Printf("failed to close grpc connection: %v", err)
			}
		}()
	}()

	return conn, nil
}
