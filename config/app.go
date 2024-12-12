package config

import (
	"github.com/TrinityKnights/Backend/internal/builder"
	handlerEvent "github.com/TrinityKnights/Backend/internal/delivery/http/handler/event"
	handlerUser "github.com/TrinityKnights/Backend/internal/delivery/http/handler/user"
	handlerVenue "github.com/TrinityKnights/Backend/internal/delivery/http/handler/venue"
	"github.com/TrinityKnights/Backend/internal/delivery/http/middleware"
	"github.com/TrinityKnights/Backend/internal/delivery/http/route"
	repositoryEvent "github.com/TrinityKnights/Backend/internal/repository/event"
	repositoryUser "github.com/TrinityKnights/Backend/internal/repository/user"
	repositoryVenue "github.com/TrinityKnights/Backend/internal/repository/venue"
	serviceEvent "github.com/TrinityKnights/Backend/internal/service/event"
	serviceUser "github.com/TrinityKnights/Backend/internal/service/user"
	serviceVenue "github.com/TrinityKnights/Backend/internal/service/venue"
	"github.com/TrinityKnights/Backend/pkg/cache"
	"github.com/TrinityKnights/Backend/pkg/jwt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	Cache    *cache.ImplCache
	App      *echo.Echo
	Log      *logrus.Logger
	Validate *validator.Validate
	JWT      *jwt.JWTConfig
	Midtrans MidtransConfig
}

func Bootstrap(config *BootstrapConfig) error {
	// Initialize JWT service
	jwtService := jwt.NewJWTService(config.JWT)

	// Initialize repository
	userRepository := repositoryUser.NewUserRepository(config.DB, config.Log)
	venueRepository := repositoryVenue.NewVenueRepository(config.DB, config.Log)
	eventRepository := repositoryEvent.NewEventRepository(config.DB, config.Log)

	// Initialize service
	userService := serviceUser.NewUserServiceImpl(config.DB, config.Log, config.Validate, userRepository, jwtService)
	venueService := serviceVenue.NewVenueServiceImpl(config.DB, config.Cache, config.Log, config.Validate, venueRepository)
	eventService := serviceEvent.NewEventServiceImpl(config.DB, config.Cache, config.Log, config.Validate, eventRepository)

	// Initialize handler
	userHandler := handlerUser.NewUserHandler(config.Log, userService)
	venueHandler := handlerVenue.NewVenueHandler(config.Log, venueService)
	eventHandler := handlerEvent.NewEventHandler(config.Log, eventService)

	// Initialize graphql

	// Initialize middleware
	authMiddleware := middleware.AuthMiddleware(jwtService)

	// Initialize route
	routeConfig := route.Config{
		App:          config.App,
		UserHandler:  userHandler,
		VenueHandler: venueHandler.(*handlerVenue.VenueHandlerImpl),
		EventHandler: eventHandler.(*handlerEvent.EventHandlerImpl),
	}

	// Build routes
	b := builder.Config{
		App:            config.App,
		UserHandler:    userHandler,
		VenueHandler:   venueHandler.(*handlerVenue.VenueHandlerImpl),
		EventHandler:   eventHandler.(*handlerEvent.EventHandlerImpl),
		AuthMiddleware: authMiddleware,
		Routes:         &routeConfig,
	}
	b.BuildRoutes()

	config.Log.Infof("Application is ready")
	return nil
}
