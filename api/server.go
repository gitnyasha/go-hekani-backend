package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/gitnyasha/go-hekani-backend/db/sqlc"
	"github.com/gitnyasha/go-hekani-backend/token"
	"github.com/gitnyasha/go-hekani-backend/util"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenKey)

	if err != nil {
		return nil, fmt.Errorf("connot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	server.setupRouter()
	return server, err
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/questions", server.createQuestion)
	router.GET("/questions/:id", server.getQuestion)
	router.GET("/questions", server.listQuestion)
	router.GET("/users", server.listUser)
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.GET("/users/:id", server.getUser)

	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
