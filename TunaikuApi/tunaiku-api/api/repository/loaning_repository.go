package repository

import (
	"github.com/jinzhu/gorm"
	"tunaiku-api/struct"
	. "tunaiku-api/struct"
	. "tunaiku-api/api"
)

type LoaningRepository struct {
	Conn *gorm.DB
}

func (lr *LoaningRepository) GetId(id uint) (*Loaning, error) {
	var loaning Loaning

	err := lr.Conn.Find(&loaning, id).Error

	if err != nil {
		return nil, err
	}

	return &loaning, nil
}

func (lr *LoaningRepository) Create(loaning *_struct.Loaning) error {
	err := lr.Conn.Create(&loaning).Error

	if err != nil {
		return err
	}

	var ktp_ KtpDetail

	ktp_.ID = loaning.KtpDetailID
	ktp_.YearID  = loaning.KtpDetailID % 100
	ktp_.MonthID = (loaning.KtpDetailID / 100) % 100
	ktp_.DateID = (loaning.KtpDetailID / 10000) % 100
	ktp_.SubDistrictID = (loaning.KtpDetailID / 1000000) % 100
	ktp_.ProvincialID = (loaning.KtpDetailID / 100000000) % 100
	err = lr.Conn.Create(&ktp_).Error

	if err != nil {
		return err
	}

	return nil
}

func (lr *LoaningRepository) Delete(loaning *_struct.Loaning) error {
	err := lr.Conn.Delete(&loaning).Error

	if err != nil {
		return err
	}

	return nil
}

func (lr *LoaningRepository) Get() ([]*_struct.Loaning, error) {
	var loanings []*Loaning

	err := lr.Conn.Find(&loanings).Error

	if err != nil {
		return nil, err
	}

	return loanings, nil
}

func (lr *LoaningRepository) Update(loaning *_struct.Loaning) error {
	err := lr.Conn.Save(&loaning).Error

	if err != nil {
		return err
	}

	return nil
}

func NewLoaningRepository(Conn *gorm.DB) RepositoryLoaning {
	return &LoaningRepository{Conn}
}