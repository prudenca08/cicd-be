package drivers

import (
	"finalproject/features/admins"
	adminDB "finalproject/features/admins/data"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) admins.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}
