package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) ToJson() ([]byte, error) {
	// encode the struct to json
	return json.Marshal(u)
}
