package model

import "time"

// Account Account Database Struct
type Account struct {
	ID                uint64
	UserName          string
	PasswordEncrypted string
	PasswordKey       string
	CreateTime        time.Time
}
