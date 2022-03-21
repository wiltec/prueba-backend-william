package DataBase

import (
	"gorm.io/gorm"
)

type IDHandler interface {
	Connect() (*gorm.DB, error)
}
