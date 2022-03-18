package gorm20

import (
	"fmt"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestConnect(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println(db)
}
