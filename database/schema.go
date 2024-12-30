package database

import (
	"database/sql"
	"log"
)

// CreateSchemaAndTables sets up the database schema and tables
func CreateSchemaAndTables(db *sql.DB) {
	// Define the CREATE TABLE SQL statement
	createTableSQL := `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

	// Execute the SQL statement
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create orders table: %v", err)
	}

	log.Println("Table 'users' created successfully")

}
