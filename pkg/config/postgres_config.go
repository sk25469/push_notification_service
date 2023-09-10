package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"github.com/sk25469/push_noti_service/pkg/utils"
)

var conn *sql.DB

func InitPostgres() {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		utils.DB_HOST, utils.DB_PORT, utils.DB_USER, utils.DB_PASSWORD, utils.DB_NAME)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	fmt.Println("Connected to PostgreSQL!")
}

func GetPostgresConnection() *sql.DB {
	return conn
}
