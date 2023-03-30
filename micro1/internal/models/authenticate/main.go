package authenticate

import (
	"time"
)

type TokenAes struct {
	Id      string `gorm:"primaryKey"`
	Expired time.Time
	UserId  string
}
