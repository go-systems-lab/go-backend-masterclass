package gapi

import (
	"fmt"

	db "github.com/go-systems-lab/go-backend-masterclass/db/sqlc"
	"github.com/go-systems-lab/go-backend-masterclass/pb"
	"github.com/go-systems-lab/go-backend-masterclass/token"
	"github.com/go-systems-lab/go-backend-masterclass/util"
	"github.com/go-systems-lab/go-backend-masterclass/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServiceServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
