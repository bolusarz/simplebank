package gapi

import (
	"fmt"
	db "github.com/bolusarz/simplebank/db/sqlc"
	"github.com/bolusarz/simplebank/pb"
	"github.com/bolusarz/simplebank/token"
	"github.com/bolusarz/simplebank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
