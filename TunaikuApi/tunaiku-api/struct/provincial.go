package _struct

import "github.com/jinzhu/gorm"

type Provincial struct {
	gorm.Model

	NameProvincial string
}
