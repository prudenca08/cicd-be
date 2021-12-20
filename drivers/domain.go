package drivers

import (
	"finalproject/features/admins"
	adminDB "finalproject/features/admins/data"
	"finalproject/features/docses"
	docsesDB "finalproject/features/docses/data"
	"finalproject/features/doctor"
	doctorDB "finalproject/features/doctor/data"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) admins.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}
func NewDoctorRepository(conn *gorm.DB) doctor.Repository {
	return doctorDB.NewMysqlDoctorRepository(conn)
}
func NewDocsesRepository(conn *gorm.DB) docses.Repository {
	return docsesDB.NewMysqlDocsesRepository(conn)
}
