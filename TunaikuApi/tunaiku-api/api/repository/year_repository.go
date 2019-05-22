package repository

import (
	"github.com/jinzhu/gorm"
	"tunaiku-api/api"
	"tunaiku-api/struct"
	. "tunaiku-api/struct"
)

type YearRepository struct {
	Conn *gorm.DB
}

func (pr *YearRepository) Create(year *_struct.Year) error {
	err := pr.Conn.Create(&year).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *YearRepository) Update(year *_struct.Year) error {
	err := pr.Conn.Save(&year).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *YearRepository) Delete(year *_struct.Year) error {
	err := pr.Conn.Delete(&year).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *YearRepository) Get() ([]*_struct.Year, error) {
	var years []*Year

	err := pr.Conn.Find(&years).Error

	if err != nil {
		return nil, err
	}

	return years, nil
}

func (pr *YearRepository) GetId(id uint) (*_struct.Year, error) {
	var year Year

	err := pr.Conn.Find(&year, id).Error

	if err != nil {
		return nil, err
	}

	return &year, nil
}

func NewYearRepository(Conn *gorm.DB) api.RepositoryYear{
	return &YearRepository{Conn}
}