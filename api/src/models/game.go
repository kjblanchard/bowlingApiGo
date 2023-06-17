package models

type Game struct {
	UserId int
	Score int
}

type Score struct {
	Score int `json:"score"`
}