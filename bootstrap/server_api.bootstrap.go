package bootstrap

import (
	"log"
	"net/http"
	"sync"
	"time"

	"service-api/config"
	"service-api/middlewares"
	"service-api/models"
	"service-api/seeders"
	util "service-api/utils"

	"service-api/routes"

	"github.com/joho/godotenv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
initialize server api
@param wg *sync.WaitGroup
@author Roby Parlan
*/
func ServerApi(wg *sync.WaitGroup) {
	defer wg.Done()

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())

	godotenv.Load()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	db, err := config.NewDB(cfg.DB)
	if err != nil {
		log.Fatalf("Error initializing database: %s", err)
	}

	// Auto migrate tables
	if err := db.AutoMigrate(&models.Brands{}, &models.Products{}, &models.Users{}); err != nil {
		log.Fatal("Failed to migrate tables:", err)
	}

	// Seeder
	seeders.Seed(db)

	// Inisialisasi logger
	logger, err := util.NewLogger(cfg.LogsDir) // Path ke file log
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Middleware logging
	e.Use(middlewares.LoggingMiddleware(logger))

	e.Validator = &util.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = util.HTTPErrorHandler

	e.Use(middleware.RequestID())

	// config cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		MaxAge:       86400, // 1 day
	}))

	/*
		Secure middleware provides protection against
		cross-site scripting (XSS) attack, content type sniffing, clickjacking,
		insecure connection and other code injection attacks.
		Set By Default Echo
	*/
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            3600,
		HSTSExcludeSubdomains: true,
	}))

	// config rate limit
	// configRateLimiter := util.GetRateLimiter()
	// e.Use(middleware.RateLimiterWithConfig(configRateLimiter))

	// Register routes
	routes.RegisterRoutes(e, db, cfg)

	// Start server
	server := &http.Server{
		Addr:         ":" + cfg.AppPort,
		ReadTimeout:  time.Duration(cfg.AppReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.AppWriteTimeout) * time.Second,
	}
	e.Logger.Fatal(e.StartServer(server))
}
