package repository

import (
	"github.com/jinzhu/gorm"
	"tunaiku-api/api"
	"tunaiku-api/struct"
	. "tunaiku-api/struct"
)

type SubDistrictRepository struct {
	Conn *gorm.DB
}

func (pr *SubDistrictRepository) Create(subdistrict *_struct.SubDistrict) error {
	err := pr.Conn.Create(&subdistrict).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *SubDistrictRepository) Update(subdistrict *_struct.SubDistrict) error {
	err := pr.Conn.Save(&subdistrict).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *SubDistrictRepository) Delete(subdistrict *_struct.SubDistrict) error {
	err := pr.Conn.Delete(&subdistrict).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *SubDistrictRepository) Get() ([]*_struct.SubDistrict, error) {
	var subdistricts []*SubDistrict

	err := pr.Conn.Find(&subdistricts).Error

	if err != nil {
		return nil, err
	}

	return subdistricts, nil
}

func (pr *SubDistrictRepository) GetId(id uint) (*_struct.SubDistrict, error) {
	var subdistrict SubDistrict

	err := pr.Conn.Find(&subdistrict, id).Error

	if err != nil {
		return nil, err
	}

	return &subdistrict, nil
}

func NewSubDistrictRepository(Conn *gorm.DB) api.RepositorySubDistrict{
	return &SubDistrictRepository{Conn}
}
