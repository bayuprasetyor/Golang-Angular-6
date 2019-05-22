package _struct

import "github.com/jinzhu/gorm"

type KtpDetail struct {
	gorm.Model

	ProvincialID uint
	DateID uint
	MonthID uint
	SubDistrictID uint
	YearID uint

}

