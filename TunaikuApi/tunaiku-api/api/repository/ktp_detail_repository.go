package repository

import (
	"github.com/jinzhu/gorm"
	"tunaiku-api/api"
	. "tunaiku-api/struct"
)

type KtpDetailRepository struct {
	Conn *gorm.DB
}

func (kr *KtpDetailRepository) Create(ktp uint) error {
	var ktp_ KtpDetail

	ktp_.ID = ktp
	ktp_.ProvincialID = 01
	err := kr.Conn.Save(&ktp_).Error

	if err != nil {
		return err
	}

	return nil
}

func NewKtpDetailRepository(Conn *gorm.DB) api.RepositoryKtpDetail {
	return &KtpDetailRepository{Conn}
}
