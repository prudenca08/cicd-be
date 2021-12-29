package data

import (
	"finalproject/features/doctor"
	"finalproject/features/doctor/bussiness"

	"gorm.io/gorm"
)

type MysqlDoctorRepository struct {
	Conn *gorm.DB
}

func NewMysqlDoctorRepository(conn *gorm.DB) doctor.Repository {
	return &MysqlDoctorRepository{
		Conn: conn,
	}
}

func (rep *MysqlDoctorRepository) Register(domain *doctor.Domain) (doctor.Domain, error) {

	doc := fromDomain(*domain)

	result := rep.Conn.Create(&doc)

	if result.Error != nil {
		return doctor.Domain{}, result.Error
	}

	return toDomain(doc), nil
}

func (rep *MysqlDoctorRepository) Login(username, password string) (doctor.Domain, error) {
	var doc Doctor
	err := rep.Conn.First(&doc, "username = ?", username).Error

	if err != nil {
		return doctor.Domain{}, bussiness.ErrEmailorPass
	}

	return toDomain(doc), nil
}
func (rep *MysqlDoctorRepository) Update(docID int, domain *doctor.Domain) (doctor.Domain, error) {

	profileUpdate := fromDomain(*domain)

	profileUpdate.ID = docID

	result := rep.Conn.Where("id = ?", docID).Updates(&profileUpdate)

	if result.Error != nil {
		return doctor.Domain{}, bussiness.ErrNotFound
	}

	return toDomainUpdate(profileUpdate), nil
}
func (rep *MysqlDoctorRepository) DoctorByID(id int) (doctor.Domain, error) {

	var doc Doctor

	result := rep.Conn.Where("id = ?", id).First(&doc)

	if result.Error != nil {
		return doctor.Domain{}, result.Error
	}

	return toDomain(doc), nil
}
func (rep *MysqlDoctorRepository) Delete(docID int, id int) (string, error) {
	rec := Doctor{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", bussiness.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", bussiness.ErrNotFound
	}

	return "Product has been delete", nil

}
func (rep *MysqlDoctorRepository) AllDoctor() ([]doctor.Domain, error) {

	var pat []Doctor

	result := rep.Conn.Find(&pat)

	if result.Error != nil {
		return []doctor.Domain{}, result.Error
	}

	return toDomainList(pat), nil

}
