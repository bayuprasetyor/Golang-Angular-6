package _struct

import "github.com/jinzhu/gorm"

type Date struct {
	gorm.Model
	DateValue string
	Gender string

}
