package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/quandev306/webcrawl/config"
)

func NewMongoDatabase(env *Env) config.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connectionURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", env.DBUser, env.DBPass, env.DBHost, env.DBPort)
	client, err := config.NewClient(ctx, connectionURI)
	if err != nil {
		log.Fatalf("mongo: create client error: %v", err)
	}
	if err := client.Connect(ctx); err != nil {
		log.Fatalf("mongo: connect error: %v", err)
	}
	if err := client.Ping(ctx); err != nil {
		log.Fatalf("mongo: ping error: %v", err)
	}
	return client
}

func CloseMongoDBConnection(client config.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		log.Printf("mongo: disconnect error: %v", err)
	}
}
