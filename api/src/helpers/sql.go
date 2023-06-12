package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/kjblanchard/BowlingWebApp/models"
)

func Connect() (*sql.DB, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "db:3306",
		DBName:               "bowling",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db, nil
}

func GetUserByName(username string) (*models.User, error) {
	db, _ := Connect()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return nil, fmt.Errorf("Failed on query %q: %v", username, err)
	}
	defer rows.Close()
	// // Loop through users returned in the call, and assign them to the user
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password_hash, &user.Email); err != nil {
			return nil, fmt.Errorf("Failed scanning %q: %v", username, err)
		}
		fmt.Printf("User found: Username: %s\nEmail: %s", user.Username, user.Email)
		return &user, nil
	}
	return nil, fmt.Errorf("Failed to find user %q", username)
}