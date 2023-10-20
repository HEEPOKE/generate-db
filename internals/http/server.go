package http

import (
	"log"
	"net/http"

	"github.com/HEEPOKE/generate-db/internals/app/handlers"
	"github.com/HEEPOKE/generate-db/internals/app/services"
	"github.com/HEEPOKE/generate-db/internals/core/interfaces"
	_ "github.com/HEEPOKE/generate-db/pkg/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	echo            *echo.Echo
	generateHandler *handlers.GenerateHandler
	insertHandler   *handlers.InsertHandler
}

func NewServer(generateRepository interfaces.GenerateRepository, insertRepository interfaces.InsertRepository, utilitiesRepository interfaces.UtilitiesRepository) *Server {
	e := echo.New()

	loggerConfig := middleware.LoggerConfig{
		Format:           "URI::${uri}, METHOD::${method},  STATUS::${status}, HEADER::${header}, QUERY::${query}, ERROR::${error}",
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

	insertService := services.NewInsertService(insertRepository)
	insertHandler := handlers.NewInsertHandler(*insertService)

	utilitiesService := services.NewUtilitiesService()

	return &Server{
		echo:            e,
		generateHandler: generateHandler,
		insertHandler:   insertHandler,
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
	s.echo.GET("/", func(c echo.Context) error {
		response := map[string]string{
			"message": "HEEPOKE",
		}
		return c.JSON(http.StatusOK, response)
	})

	apis := s.echo.Group("/apis")

	apis.GET("/docs/*", echoSwagger.WrapHandler)

	generate := apis.Group("/generate")
	generate.GET("/get-details", s.generateHandler.GetListGenerateAll)
	generate.POST("/mockup-data", s.generateHandler.MockupData)

	// insert := apis.Group("/insert")
	// insert.POST("/", s.insertHandler.)
}
