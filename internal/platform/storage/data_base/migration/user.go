package migration

import (
	"github.com/alexperezortuno/go-auth/internal/platform/storage/data_base/model"
	"gorm.io/gorm"
)

func UserMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})

	if err != nil {
		return
	}
}
