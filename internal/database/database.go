package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type DB struct {
	*sql.DB
}

var dbInstance *DB

var ConnectedChan = make(chan bool)

func GetDB() *DB {
	ConnectionMonitor()

	return dbInstance
}

func loadDBConnectionString(key string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf(".env file not found - Could not load environment variables")
	}

	raw := os.Getenv(key)
	if raw == "" {
		return "", fmt.Errorf("%s environment variable not set", key)
	}

	dataSourceName := os.ExpandEnv(raw)

	return dataSourceName, nil
}

func openConnection() error {
	dsn, err := loadDBConnectionString("DATABASE_URL")
	if err != nil {
		return fmt.Errorf("failed to load connection string: %w", err)
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("error connecting to database - check connection url or server availability: %w", err)
	}

	fmt.Println("Connected to the database!!!")

	dbInstance = &DB{db}

	return nil
}

func CloseConnection() error {
	if dbInstance == nil {
		return nil
	}

	err := dbInstance.Close()
	if err != nil {
		return fmt.Errorf("error closing database connection: %w", err)
	}

	return nil
}

func ConnectionMonitor() {
	if dbInstance == nil || dbInstance.Ping() != nil {
		go func() {
			for {
				fmt.Println("\033[33mAttempting to connect to the database...\033[0m")

				err := openConnection()
				if err != nil {
					time.Sleep(10 * time.Second)
					fmt.Println("\033[31mFailed to connect: \033[0m", err)
					continue
				}

				break
			}

			ConnectedChan <- true

			defer fmt.Println("\033[32mConnection established!\033[0m")
		}()
	}
}
