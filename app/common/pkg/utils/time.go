package utils

import (
	"gorm.io/gorm"
	"time"
)

func ParseDeleteTime(deletedAt gorm.DeletedAt) string {
	if deletedAt.Valid {
		return deletedAt.Time.Format(time.DateTime)
	}
	return ""
}
