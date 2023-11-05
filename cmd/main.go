package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"os/signal"

	"github.com/ashmortar/go-watchlist/assets"
	"github.com/ashmortar/go-watchlist/router"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"

	"github.com/spf13/viper"
)

func main() {
	// Load environment variables from .env file
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Initialize SQLite database
	// DB, err := db.Open()
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close(DB)

	// Initialize Echo
	e := echo.New()

	// middlewares
	e.Use(middleware.RequestID())
	logger := zerolog.New(os.Stdout)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("logger", logger)
			return next(c)
		}
	})
	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		c.Set("DB", DB)
	// 		return next(c)
	// 	}
	// })
	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		c.Set("config", viper)
	// 		return next(c)
	// 	}
	// })
	assets.Serve(e)
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	router.InitRoutes(e)

	// Start server
	go func() {
		port := viper.GetString("PORT")
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
