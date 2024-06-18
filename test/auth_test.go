package test

import (
	"bytes"
	"encoding/json"
	"github.com/kasisaki/rutube_test_assing/internal"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestLogin(t *testing.T) {
	internal.Db.Create(&internal.User{Username: "testuser", Password: "password"})

	creds := internal.Credentials{Username: "testuser", Password: "password"}
	body, _ := json.Marshal(creds)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(internal.Login)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)

	tokenString, ok := response["token"]
	if !ok {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}

	claims := &internal.Claims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return internal.JwtKey, nil
	})

	if err != nil || claims.Username != "testuser" {
		t.Errorf("handler returned unexpected username: got %v want %v", claims.Username, "testuser")
	}
}

func TestAuthenticate(t *testing.T) {
	req, err := http.NewRequest("GET", "/protected", nil)
	if err != nil {
		t.Fatal(err)
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &internal.Claims{
		Username: "testuser",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(internal.JwtKey)

	req.Header.Set("Authorization", tokenString)

	rr := httptest.NewRecorder()
	handler := internal.Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
