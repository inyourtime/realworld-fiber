package sql

import "gorm.io/gorm"

type articleRepo struct {
	db *gorm.DB
}
