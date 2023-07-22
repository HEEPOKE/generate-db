package http

import (
	"log"
	"net/http"

	"github.com/HEEPOKE/generate-db/internals/app/handlers"
	"github.com/HEEPOKE/generate-db/internals/app/services"
	"github.com/HEEPOKE/generate-db/internals/core/interfaces"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo            *echo.Echo
	generateHandler *handlers.GenerateHandler
}

func NewServer(generateRepository interfaces.GenerateRepository) *Server {
	e := echo.New()

	loggerConfig := middleware.LoggerConfig{
		Format:           "URI::${uri}\n, METHOD::${method},  STATUS::${status}, HEADER::${header}\n, QUERY::${query}\n, ERROR::${error}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(middleware.LoggerWithConfig(loggerConfig))
	e.Use(middleware.Recover())

	generateService := services.NewGenerateService(generateRepository)
	generateHandler := handlers.NewGenerateHandler(*generateService)

	return &Server{
		echo:            e,
		generateHandler: generateHandler,
	}
}

func (s *Server) RouteInit(address string) {
	s.routeConfig()

	err := s.echo.Start(address)
	if err != nil {
		log.Fatalf("Failed To Start The Server: %v", err)
	}
}

func (s *Server) routeConfig() {
	apis := s.echo.Group("/apis")

	generate := apis.Group("/generate")

	generate.GET("/mockup-data", s.generateHandler.GetListGenerateAll)
}
