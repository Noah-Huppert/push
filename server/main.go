package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/Noah-Huppert/push/server/config"

	"github.com/Noah-Huppert/golog"
	"github.com/golang-migrate/migrate/v4"
	migratePsql "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// {{{1 Setup logger
	logger := golog.NewStdLogger("push-server")

	// {{{1 Load configuration
	cfg, err := config.NewConfigFromEnv()
	if err != nil {
		logger.Fatalf("error loading configuration: %s", err.Error())
	}

	// {{{1 Make handle ctrl+c
	ctx, ctxCancel := context.WithCancel(context.Background())

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	go func() {
		<-interruptChan
		ctxCancel()
	}()

	// numDoneItems a count of the number of items which should send to the
	// doneChan before the program should exit. Incremented throughout
	// the main function
	numDoneItems := 0

	// doneChan receives any value when a part of the program is done
	doneChan := make(chan int)

	// {{{1 Connect to DB
	// {{{2 Connect
	stdDb, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s "+
		"sslmode=disable", cfg.DBName, cfg.DBUser))
	if err != nil {
		logger.Fatalf("error connecting to database: %s", err.Error())
	}

	db := sqlx.NewDb(stdDb, "postgres")

	// {{{2 Test connection
	if err = db.Ping(); err != nil {
		logger.Fatalf("error testing database connection: %s",
			err.Error())
	}

	logger.Info("connected to database")

	// {{{2 Run migrations if requested
	var runMigrations bool

	flag.BoolVar(&runMigrations, "migrate", false, "Indicates program "+
		"should run DB migrations and exit (Without starting API "+
		"server)")

	flag.Parse()

	if runMigrations {
		logger.Info("running db migrations")

		// {{{3 Setup migration Postgres driver
		psqlDriver, err := migratePsql.WithInstance(stdDb,
			&migratePsql.Config{})
		if err != nil {
			logger.Fatalf("error creating migrations Postgres "+
				"driver: %s", err.Error())
		}

		// {{{3 Setup migrate client
		migrateClient, err := migrate.NewWithDatabaseInstance(
			"file://./migrations", "postgres", psqlDriver)

		if err != nil {
			logger.Fatalf("error creating migration client: %s",
				err.Error())
		}

		// {{{3 Run migrations
		if err = migrateClient.Up(); err != nil {
			logger.Fatalf("error running db migrations: %s",
				err.Error())
		}

		logger.Info("ran db migrations successfully, exiting...")

		os.Exit(0)
	}

	// {{{1 Setup HTTP server
	numDoneItems++

	// {{{2 Setup handlers
	handler := mux.NewRouter()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler: handler,
	}

	// {{{2 Start server
	go func() {
		if err = server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			logger.Fatalf("error serving API: %s", err.Error())
		}
		logger.Info("stopped API server")
	}()

	// {{{2 Handle context close
	go func() {
		<-ctx.Done()
		logger.Info("stopping API server")
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("error shutting down API server: %s",
				err.Error())
		}
		doneChan <- 0
	}()

	logger.Infof("started API server on port %d", cfg.HTTPPort)

	// {{{1 Wait for context to close
	recDoneItems := 0
	for recDoneItems < numDoneItems {
		<-doneChan
		recDoneItems++
	}

	logger.Info("program done")
}
