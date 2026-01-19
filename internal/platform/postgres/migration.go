package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrationsWithRetry(dbURL string) error {
	var m *migrate.Migrate
	var err error

	const (
		maxAttempts = 5
		retryDelay  = 2 * time.Second
	)

	for i := 1; i <= maxAttempts; i++ {
		// file://migrations points to your local folder
		m, err = migrate.New("file://migrations", dbURL)
		if err == nil {
			log.Printf("Successfully connected to database on attempt %d", i)
			break
		}

		log.Printf("Database not ready (attempt %d/%d): %v. Retrying in %v...", i, maxAttempts, err, retryDelay)
		time.Sleep(retryDelay)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to database after %d attempts: %w", maxAttempts, err)
	}

	// Execute the migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	if err == migrate.ErrNoChange {
		log.Println("Database is already up to date.")
	} else {
		log.Println("Database migrations applied successfully!")
	}

	return nil
}
