package internal

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"uniqueIndex" json:"username"`
	Password string
	Email    string `gorm:"uniqueIndex" json:"email"`
}

type Employee struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
}

type Subscription struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint `json:"user_id"`
	EmployeeID uint `json:"employee_id"`
	NotifyDays int  // Дней до дня рождения для оповещения
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&User{}, &Employee{}, &Subscription{})
}
