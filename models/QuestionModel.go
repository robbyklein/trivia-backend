package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Question string
	Answer   string
}
