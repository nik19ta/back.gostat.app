package server

import (
	"context"
	middlewareCors "gostat/pkg/middleware/cors"
	"gostat/services/auth"
	authHttp "gostat/services/auth/delivery/http"
	authPostgres "gostat/services/auth/repository/postgres"
	authUseCase "gostat/services/auth/usecase"
	"gostat/services/links"
	linksHttp "gostat/services/links/delivery/http"
	linksPostgres "gostat/services/links/repository/postgres"
	linksUseCase "gostat/services/links/usecase"

	"gostat/services/stat"
	statHttp "gostat/services/stat/delivery/http"
	statPostgres "gostat/services/stat/repository/postgres"
	statUseCase "gostat/services/stat/usecase"
	"log"
	"net/http"

	"os"
	"os/signal"
	"time"

	database "gostat/pkg/database"

	gin "github.com/gin-gonic/gin"

	docs "gostat/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	httpServer *http.Server

	statsUC stat.UseCase
	authUC  auth.UseCase
	linksUC links.UseCase
}

func NewApp() *App {
	db := database.InitDB()

	statsRepo := statPostgres.NewUserRepository(db)
	authRepo := authPostgres.NewUserRepository(db)
	linksRepo := linksPostgres.NewLinksRepository(db)

	return &App{
		statsUC: statUseCase.NewStatUseCase(statsRepo),
		authUC:  authUseCase.NewAuthUseCase(authRepo),
		linksUC: linksUseCase.NewLinkUseCase(linksRepo),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		middlewareCors.CORSMiddleware(),
	)

	statHttp.RegisterHTTPEndpoints(router, a.statsUC)
	authHttp.RegisterHTTPEndpoints(router, a.authUC)
	linksHttp.RegisterHTTPEndpoints(router, a.linksUC)

	// * Swagger
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
