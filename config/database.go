package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDatabase() error {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL n'est pas définie dans les variables d'environnement")
	}

	var err error
	DB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture de la connexion: %w", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("erreur lors du ping de la base de données: %w", err)
	}

	if err = createTables(); err != nil {
		return fmt.Errorf("erreur lors de la création des tables: %w", err)
	}

	log.Println("✓ Base de données initialisée")
	return nil
}

func createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		phone VARCHAR(50),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	`

	_, err := DB.Exec(query)
	return err
}

func CloseDatabase() {
	if DB != nil {
		DB.Close()
		log.Println("✓ Connexion à la base de données fermée")
	}
}
