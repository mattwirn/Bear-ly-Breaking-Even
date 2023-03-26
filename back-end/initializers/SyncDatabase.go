package initializers

import "github.com/mattwirn/Bear-ly-Breaking-Even/back-end/models"

func SyncDatabases() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Income{})
	DB.AutoMigrate(&models.Cat1{})
}
