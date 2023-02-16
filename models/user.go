package models

type user struct {
	ID           uint
	Username     string
	Email        string
	PasswordHash string
	TodoItems    []todoItem
}
