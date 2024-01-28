package initializers

import "github.com/fymorGod/jwt-go/models"

func SyncDatabases() {
	DB.AutoMigrate(&models.User{})
}
