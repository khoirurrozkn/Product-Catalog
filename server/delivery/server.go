package delivery

import (
	"log"
	"server/config"
	"server/delivery/controller"
	"server/middleware"
	"server/repository"
	"server/usecase"
	"server/utils/common"

	"github.com/gin-gonic/gin"
)

type Server struct {
	authMiddlware middleware.AuthMiddleware
	userUc usecase.UserUsecase
	productUc usecase.ProductUsecase
	favoriteUc usecase.FavoriteUsecase

	engine *gin.Engine
	routerGroup *gin.RouterGroup
	host string
}

func (s *Server) setupController(){
	controller.NewUserController(s.userUc, s.routerGroup, s.authMiddlware).Route()
	controller.NewProductController(s.productUc, s.routerGroup, s.authMiddlware).Route()
	controller.NewFavoriteController(s.favoriteUc, s.routerGroup, s.authMiddlware).Route()

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

	jwt_token := common.NewJwtToken(cfg.TokenConfig)
	authMiddleware := middleware.NewAuthMiddleware(jwt_token)

	userRefreshTokenRepo := repository.NewUserRefreshTokenRepository(db.Conn())
	userRefreshTokenUc := usecase.NewUserRefreshTokenUsecase(userRefreshTokenRepo, jwt_token)

	userRepo := repository.NewUserRepository(db.Conn())
	userUc := usecase.NewUserUsecase(userRepo, userRefreshTokenUc, jwt_token)

	productRepo := repository.NewProductRepository(db.Conn())
	productUc := usecase.NewProductUsecase(productRepo)

	favoriteRepo := repository.NewFavoriteRepository(db.Conn())
	favoriteUc := usecase.NewFavoriteUsecase(favoriteRepo)

	engine := gin.Default()
	engine.Use(middleware.NewCorsMiddleware())
	routerGroup := engine.RouterGroup.Group("/api")

	return &Server{
		authMiddlware: authMiddleware,
		userUc: userUc,
		productUc: productUc,
		favoriteUc: favoriteUc,

		engine: engine,
		routerGroup: routerGroup,
		host: ":8080",
	}
}