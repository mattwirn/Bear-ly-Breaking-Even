package initializers

import "github.com/mattwirn/Bear-ly-Breaking-Even/back-end/models"

func SyncDatabases() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Income{})
	DB.AutoMigrate(&models.Home_Uts{})
	DB.AutoMigrate(&models.Travel{})
	DB.AutoMigrate(&models.Food{})
	DB.AutoMigrate(&models.Entertainment{})
	DB.AutoMigrate(&models.Health{})
}
