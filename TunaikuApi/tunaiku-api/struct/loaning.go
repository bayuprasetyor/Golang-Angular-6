package _struct

import (
	"github.com/jinzhu/gorm"
)

type Loaning struct {
	gorm.Model

	Name string
	Date string
	KtpDetailID uint
	BirthDate string
	Gender string
	Amount int
	Period int
}
