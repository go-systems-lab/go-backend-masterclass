package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/go-systems-lab/go-backend-masterclass/db/sqlc"
	"github.com/go-systems-lab/go-backend-masterclass/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store      db.Store
	router     *gin.Engine
	config     util.Config
	tokenMaker *util.TokenMaker
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := util.NewTokenMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{store: store, config: config, tokenMaker: tokenMaker}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
