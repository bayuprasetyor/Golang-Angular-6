package repository

import (
	"github.com/jinzhu/gorm"
	"tunaiku-api/api"
	"tunaiku-api/struct"
	. "tunaiku-api/struct"
)

type MonthRepository struct {
	Conn *gorm.DB
}

func (mr *MonthRepository) Create(month *_struct.Month) error {
	err := mr.Conn.Create(&month).Error

	if err != nil {
		return err
	}

	return nil
}

func (mr *MonthRepository) Delete(month *_struct.Month) error {
	err := mr.Conn.Delete(&month).Error

	if err != nil {
		return err
	}

	return nil
}

func (mr *MonthRepository) Get() ([]*_struct.Month, error) {
	var months []*Month

	err := mr.Conn.Find(&months).Error

	if err != nil {
		return nil, err
	}

	return months, nil
}

func (mr *MonthRepository) GetId(id uint) (*_struct.Month, error) {
	var month Month

	err := mr.Conn.Find(&month, id).Error

	if err != nil {
		return nil, err
	}

	return &month, nil
}

func (mr *MonthRepository) Update(month *_struct.Month) error {
	err := mr.Conn.Save(&month).Error

	if err != nil {
		return err
	}

	return nil
}

func NewMonthRepository(Conn *gorm.DB) api.RepositoryMonth{
	return &MonthRepository{Conn}
}