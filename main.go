package main

import (
	"context"
	"log"
	"net"

	"github.com/go-systems-lab/go-backend-masterclass/api"
	db "github.com/go-systems-lab/go-backend-masterclass/db/sqlc"
	"github.com/go-systems-lab/go-backend-masterclass/gapi"
	"github.com/go-systems-lab/go-backend-masterclass/pb"
	"github.com/go-systems-lab/go-backend-masterclass/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	defer conn.Close()

	store := db.NewStore(conn)
	runGrpcServer(config, store)
	// runGinServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server: ", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
