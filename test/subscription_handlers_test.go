package test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/kasisaki/rutube_test_assing/internal"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAddEmployee(t *testing.T) {
	birthday, _ := time.Parse(time.TimeOnly, "1991-07-07")
	employee := internal.Employee{Name: "John Doe", Email: "john@example.com", Birthday: birthday}
	body, _ := json.Marshal(employee)

	req, err := http.NewRequest("POST", "/employees", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(internal.AddEmployee)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response internal.Employee
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)

	if response.Name != employee.Name || response.Email != employee.Email {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}

func TestSubscribe(t *testing.T) {
	birthday, _ := time.Parse(time.TimeOnly, "1991-11-05")

	internal.Db.Create(&internal.User{Username: "testuser", Password: "password"})
	internal.Db.Create(&internal.Employee{Name: "John Doe", Email: "john@example.com", Birthday: birthday})

	subscription := internal.Subscription{EmployeeID: 1, NotifyDays: 7}
	body, _ := json.Marshal(subscription)

	req, err := http.NewRequest("POST", "/subscribe", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	req = req.WithContext(context.WithValue(req.Context(), "username", "testuser"))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(internal.Subscribe)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var response internal.Subscription
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)

	if response.EmployeeID != subscription.EmployeeID || response.NotifyDays != subscription.NotifyDays {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}

func TestUnsubscribe(t *testing.T) {
	birthday, _ := time.Parse(time.TimeOnly, "2000-01-01")
	internal.Db.Create(&internal.User{Username: "testuser", Password: "password"})
	internal.Db.Create(&internal.Employee{Name: "John Doe", Email: "john@example.com", Birthday: birthday})
	internal.Db.Create(&internal.Subscription{UserID: 1, EmployeeID: 1, NotifyDays: 7})

	req, err := http.NewRequest("DELETE", "/unsubscribe/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = req.WithContext(context.WithValue(req.Context(), "username", "testuser"))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(internal.Unsubscribe)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)

	if response["message"] != "Unsubscribed" {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}
