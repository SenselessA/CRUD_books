package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/SenselessA/CRUD_books"
	"github.com/SenselessA/CRUD_books/pkg/handler"
	"github.com/SenselessA/CRUD_books/pkg/hash"
	"github.com/SenselessA/CRUD_books/pkg/repository"
	"github.com/SenselessA/CRUD_books/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title CRUD_books API
// @version 1.0
// @description API Server for CRUD Books Application

// @host localhost:8080
// @BasePath /books

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbName"),
		SSLMode:  viper.GetString("db.sslMode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	hasher := hash.NewSHA1Hasher("salt")

	bookRepository := repository.NewBooks(db)
	bookService := service.NewBooks(bookRepository)

	userRepository := repository.NewUsers(db)
	userService := service.NewUsers(userRepository, hasher, []byte("secret need out in config"), viper.GetViper().GetDuration("TokenTTL"))

	handlers := handler.NewHandler(bookService, userService)

	srv := new(CRUD_books.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running http server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error shutting down http server: %s", err)
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error closing db: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
