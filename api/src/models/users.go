package models

type User struct {
	ID            int8
	Username      string
	Password_hash string
	Email         string
}