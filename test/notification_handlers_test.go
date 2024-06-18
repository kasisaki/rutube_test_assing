package test

import (
	"github.com/kasisaki/rutube_test_assing/internal"
	"testing"
	"time"
)

func TestSendBirthdayNotifications(t *testing.T) {
	internal.Db.Create(&internal.User{Username: "testuser", Password: "password", Email: "user@example.com"})
	internal.Db.Create(&internal.Employee{Name: "John Doe", Email: "john@example.com", Birthday: time.Now().AddDate(0, 0, 7)})
	internal.Db.Create(&internal.Subscription{UserID: 1, EmployeeID: 1, NotifyDays: 7})

	internal.SendBirthdayNotifications()

	// Тестируем отправку почты
	// Здесь мы могли бы проверить лог, мок или состояние системы,
	// чтобы убедиться, что почта отправлена.
	// Это требует более сложной инфраструктуры для тестирования отправки почты.
}
