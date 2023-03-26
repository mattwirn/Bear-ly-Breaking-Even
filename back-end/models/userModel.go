package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Password     string
	SessionToken string `gorm:"unique"`
}

type Income struct {
	gorm.Model
	Username string `gorm:"unique"`
	Amount   uint
}

// Expense Categories

type Cat1 struct {
	gorm.Model
	Username    string
	ExpenseName string
	Amount      uint
}
