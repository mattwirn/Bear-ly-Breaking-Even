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

type Home_Uts struct {
	gorm.Model
	Username    string
	ExpenseName string
	Amount      uint
}

type Travel struct {
	gorm.Model
	Username    string
	ExpenseName string
	Amount      uint
}
type Food struct {
	gorm.Model
	Username    string
	ExpenseName string
	Amount      uint
}
type Entertainment struct {
	gorm.Model
	Username    string
	ExpenseName string
	Amount      uint
}
type Health struct {
	gorm.Model
	Username    string
	ExpenseName string
	Amount      uint
}
