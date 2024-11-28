package delivery

import (
	"server/config"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	host string
}

func (s *Server) setupController(){

}

func (s *Server) Run(){
	s.setupController()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("Server error", err.Error())
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("config : ", err.Error())
	}
	
	db, err := config.NewDbConnection(cfg)
	if err != nil {
		log.Fatal("db connection : ", err.Error())
	}

	db.Conn()
	// Temporary before connecting on repository

	engine := gin.Default()
	return &Server{
		engine: engine,
		host: ":8080",
	}
}