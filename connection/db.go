package connection

import (
	"BackendEsp32/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SetupConnection() *sql.DB {

	dbUser := config.GetEnv("DB_USER")
	dbPass := config.GetEnv("DB_PASS")
	dbHost := config.GetEnv("DB_HOST")
	dbPort := config.GetEnv("DB_PORT")
	dbName := config.GetEnv("DB_NAME")

	// format:
	// user:pass@tcp(host:port)/dbname?parseTime=true
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Gagal koneksi database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database tidak merespon: ", err)
	}

	log.Println("Database connected successfully")

	return db
}
