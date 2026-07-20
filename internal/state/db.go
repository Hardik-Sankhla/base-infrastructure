package state

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"log/slog"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// Init initializes the SQLite database
func Init(dbPath string) error {
	slog.Info("Initializing state database", "path", dbPath)
	
	// Ensure directory exists
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create database directory: %w", err)
	}

	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Create tables if not exists
	return createSchema()
}

func createSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS device_info (
		uuid TEXT PRIMARY KEY,
		hostname TEXT,
		os TEXT,
		arch TEXT,
		kernel_version TEXT,
		last_boot TEXT
	);

	CREATE TABLE IF NOT EXISTS plugins (
		name TEXT PRIMARY KEY,
		version TEXT,
		installed_at DATETIME,
		status TEXT,
		health_score INTEGER
	);

	CREATE TABLE IF NOT EXISTS health_log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		plugin TEXT,
		score INTEGER,
		diagnostics TEXT
	);

	CREATE TABLE IF NOT EXISTS capabilities (
		id TEXT PRIMARY KEY,
		provider TEXT,
		version TEXT,
		state TEXT,
		confidence INTEGER
	);

	CREATE TABLE IF NOT EXISTS execution_runs (
		id TEXT PRIMARY KEY,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		plan_version TEXT,
		status TEXT
	);
	`
	_, err := DB.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}
	
	slog.Debug("State database schema initialized")
	return nil
}
