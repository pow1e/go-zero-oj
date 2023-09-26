package utils

import (
	"errors"
	"gorm.io/gorm"
)

func GormFirstError(err error) error {
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
