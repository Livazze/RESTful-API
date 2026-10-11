package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	// PostgreSQL bağlantı dizesi (varsayılan port 5432)
	connStr := "host=db.aauxfopnxthybtigqqxe.supabase.co port=5432 user=postgres password=a312004.a312004 dbname=postgres sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Bağlantı havuzu ayarları
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Bağlantı doğrulama
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// PostgreSQL'de AUTO_INCREMENT yerine SERIAL kullanılır
	departments := `
	CREATE TABLE IF NOT EXISTS departments (
		id SERIAL PRIMARY KEY,
		name VARCHAR(250) NOT NULL
	);`
	if _, err := db.Exec(departments); err != nil {
		log.Fatalf("department table could not be created: %v", err)
	}

	students := `
	CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		name VARCHAR(250),
		department_id INT,
		grade_average DECIMAL(3,2),
		FOREIGN KEY (department_id) REFERENCES departments(id)
	);`

	if _, err = db.Exec(students); err != nil {
		log.Fatalf("student table could not be created: %v", err)
	}

	log.Println("Connection successful")
	return db
}
