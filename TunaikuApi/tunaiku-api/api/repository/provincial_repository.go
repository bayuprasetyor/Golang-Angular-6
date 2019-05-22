package repository

import (
	"github.com/jinzhu/gorm"
	"tunaiku-api/api"
	"tunaiku-api/struct"
	. "tunaiku-api/struct"
)

type ProvincialRepository struct {
	Conn *gorm.DB
}

func (pr *ProvincialRepository) Create(provincial *_struct.Provincial) error {
	err := pr.Conn.Create(&provincial).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *ProvincialRepository) Update(provincial *_struct.Provincial) error {
	err := pr.Conn.Save(&provincial).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *ProvincialRepository) Delete(provincial *_struct.Provincial) error {
	err := pr.Conn.Delete(&provincial).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *ProvincialRepository) Get() ([]*_struct.Provincial, error) {
	var provincials []*Provincial

	err := pr.Conn.Find(&provincials).Error

	if err != nil {
		return nil, err
	}

	return provincials, nil
}

func (pr *ProvincialRepository) GetId(id uint) (*_struct.Provincial, error) {
	var provincial Provincial

	err := pr.Conn.Find(&provincial, id).Error

	if err != nil {
		return nil, err
	}

	return &provincial, nil
}

func NewProvincialRepository(Conn *gorm.DB) api.RepositoryProvincial{
	return &ProvincialRepository{Conn}
}