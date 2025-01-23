package delivery

import (
	"log"
	"server/config"
	"server/delivery/controller"
	"server/middleware"
	"server/repository"
	"server/usecase"

	"github.com/gin-gonic/gin"
)

type Server struct {
	userUc usecase.UserUsecase
	productUc usecase.ProductUsecase

	engine *gin.Engine
	routerGroup *gin.RouterGroup
	host string
}

func (s *Server) setupController(){
	controller.NewUserController(s.userUc, s.routerGroup).Route()
	controller.NewProductController(s.productUc, s.routerGroup).Route()

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

	userRepo := repository.NewUserRepository(db.Conn())
	userUc := usecase.NewUserUsecase(userRepo)

	productRepo := repository.NewProductRepository(db.Conn())
	productUc := usecase.NewProductUsecase(productRepo)

	engine := gin.Default()
	engine.Use(middleware.NewCorsMiddleware())
	routerGroup := engine.RouterGroup.Group("/api")

	return &Server{
		userUc: userUc,
		productUc: productUc,

		engine: engine,
		routerGroup: routerGroup,
		host: ":8080",
	}
}