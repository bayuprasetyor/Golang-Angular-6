package _struct

import "github.com/jinzhu/gorm"

type Month struct {
	gorm.Model

	NameMonth string
}