package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "liva:312004@tcp(127.0.0.1:3306)/apidb")
	if err != nil {
		panic(err)
	}

	// Bağlantı havuzu ayarları
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// ping ile bağlantı doğrulama
	if err := db.Ping(); err != nil {
		panic(err)
	}

	departments := `
	CREATE TABLE IF NOT EXISTS departments(
		id int AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(250) NOT NULL
	);`
	if _, err := db.Exec(departments); err != nil {
		log.Fatalf("department table could not be created: %v", err)
	}

	students := `
	CREATE table IF NOT EXISTS students (
		id int AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(250),
		department_id int,
		grade_average decimal(3,2),
		FOREIGN KEY (department_id) REFERENCES departments(id)
	);`

	_, err = db.Exec(students)
	if err != nil {
		log.Fatalf("student table could not be created: %v", err)
	}

	log.Println("connecrion succesful")
	return db
}
