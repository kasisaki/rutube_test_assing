package test

import (
	"github.com/kasisaki/rutube_test_assing/internal"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Настройка перед запуском тестов
	setup()

	// Запуск тестов
	code := m.Run()

	teardown()

	// Выход с кодом завершения тестов
	os.Exit(code)
}

func setup() {
	// Здесь может быть код для настройки тестовой среды
	// Например, создание тестовой базы данных
	internal.Db.Create(&internal.User{Username: "testuser", Password: "password"})
}

// teardown чистит директорию после завершения тестов
func teardown() {
	err := os.Remove("test.db")
	if err != nil {
		log.Println("Ошибка удаления тестовой БД")
		return
	}
}
