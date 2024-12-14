package bootstrap

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

	// Start server in a goroutine
	server := &http.Server{
		Addr:         ":" + cfg.AppPort,
		ReadTimeout:  time.Duration(cfg.AppReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.AppWriteTimeout) * time.Second,
	}

	go func() {
		if err := e.StartServer(server); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Graceful shutdown logic
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-quit
	log.Println("Shutdown signal received. Shutting down server...")

	// Set a timeout for the shutdown process
	timeoutFunc := time.AfterFunc(15*time.Second, func() {
		log.Printf("timeout %d seconds has been elapsed, force exit", 15*time.Second)
		os.Exit(0)
	})

	defer timeoutFunc.Stop()

	// Attempt a graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}
	log.Println("Server is shutting down completed")

	// Close the database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting SQL DB: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
	log.Println("Database connection closing completed")

	log.Println("Server exited gracefully")
}
