package repository

import (
	"github.com/jinzhu/gorm"
	"tunaiku-api/api"
	"tunaiku-api/struct"
	. "tunaiku-api/struct"

)

type DateRepository struct {
	Conn *gorm.DB
}

func (dr *DateRepository) Create(date *_struct.Date) error {
	err := dr.Conn.Create(&date).Error

	if err != nil {
		return err
	}

	return nil
}

func (dr *DateRepository) Update(date *_struct.Date) error {
	err := dr.Conn.Save(&date).Error

	if err != nil {
		return err
	}

	return nil
}

func (dr *DateRepository) Delete(date *_struct.Date) error {
	err := dr.Conn.Delete(&date).Error

	if err != nil {
		return err
	}

	return nil
}

func (dr *DateRepository) Get() ([]*_struct.Date, error) {
	var dates []*Date

	err := dr.Conn.Find(&dates).Error

	if err != nil {
		return nil, err
	}

	return dates, nil
}

func (dr *DateRepository) GetId(id uint) (*_struct.Date, error) {
	var date Date

	err := dr.Conn.Find(&date, id).Error

	if err != nil {
		return nil, err
	}

	return &date, nil
}

func NewDateRepository(Conn *gorm.DB) api.RepositoryDate{
	return &DateRepository{Conn}
}