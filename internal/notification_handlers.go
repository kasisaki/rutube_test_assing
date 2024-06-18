package internal

import (
	"github.com/go-mail/mail"
	"time"
)

func SendBirthdayNotifications() {
	var subscriptions []Subscription
	Db.Find(&subscriptions)

	for _, subscription := range subscriptions {
		var employee Employee
		var user User

		if err := Db.First(&employee, subscription.EmployeeID).Error; err != nil {
			continue
		}

		if err := Db.First(&user, subscription.UserID).Error; err != nil {
			continue
		}

		// Проверяем, сколько дней осталось до дня рождения
		today := time.Now()
		birthday := time.Date(today.Year(), employee.Birthday.Month(), employee.Birthday.Day(), 0, 0, 0, 0, today.Location())

		// Если день рождения уже прошел в этом году, то рассматриваем день рождения в следующем году
		if birthday.Before(today) {
			birthday = birthday.AddDate(1, 0, 0)
		}

		daysUntilBirthday := int(birthday.Sub(today).Hours() / 24)
		if daysUntilBirthday == subscription.NotifyDays {
			sendEmail(user.Email, "Happy Birthday!", "Today is "+employee.Name+"'s birthday!")
		}
	}
}

func sendEmail(to string, subject string, body string) {
	m := mail.NewMessage()
	m.SetHeader("From", "your-email@example.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := mail.NewDialer("smtp.example.com", 587, "user", "password")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
