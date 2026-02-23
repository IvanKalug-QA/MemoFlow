package main

import (
	"bytes"
	"encoding/json"
	"memoflow/internal/auth"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginSuccess(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "user@mail.ru",
		Password: "20022002",
	})

	resp, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected %d got %d", http.StatusOK, resp.StatusCode)
	}
}
