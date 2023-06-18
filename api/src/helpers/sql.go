package helpers

import (
	"database/sql"
	"fmt"
	"log"
	_ "os"

	"github.com/go-sql-driver/mysql"
	"github.com/kjblanchard/BowlingWebApp/models"
)

var db *sql.DB

func QueryDb(query string, args ...any) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed on query: %v", err)
	}
	fmt.Println("Connection complete")
	return rows, nil
}

func Connect() error {
	// Capture connection properties.
	cfg := mysql.Config{
		// User:                 os.Getenv("DBUSER"),
		// Passwd:               os.Getenv("DBPASS"),
		User:   "root",
		Passwd: "example",
		Net:    "tcp",
		// Addr:                 "db:3306",
		Addr:                 "localhost:3306",
		DBName:               "bowling",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	databaseHandle, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	db = databaseHandle

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return nil
}
func CloseConnection() error {
	err := db.Close()
	if err != nil {
		return fmt.Errorf("what in the world can't even close %v", err)
	}
	return nil
}

func GetUserByName(username string) (*models.User, error) {
	rows, err := db.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return nil, fmt.Errorf("failed on query %q: %v", username, err)
	}
	defer rows.Close()
	rows.Next()
	var user models.User
	if err := rows.Scan(&user.ID, &user.Username, &user.Password_hash, &user.Email); err != nil {
		return nil, fmt.Errorf("failed mapping returned user %q: %v", username, err)
	}
	return &user, nil
}

func AddScore(userId int, score int) (*models.Game, error) {
	if _, err := db.Exec("insert into games (userId, score) values (?, ?)", userId, score); err != nil {
		return nil, fmt.Errorf("could not insert game, %v", err)
	}
	game := models.Game{
		UserId: userId,
		Score:  score,
	}
	return &game, nil
}
