package models

import (
	"github.com/jinzhu/gorm"
)

// Organization : Organization
type Organization struct {
	gorm.Model `json:"-" xml:"-"`
	Name       string
	Address    string
	Groups     []Group
	Users      []User
}

// Group : Group
type Group struct {
	gorm.Model     `json:"-" xml:"-"`
	Name           string
	OrganizationID int
	ScanResults    []ScanResult
}

// User : User
type User struct {
	gorm.Model     `json:"-" xml:"-"`
	OrganizationID string
	Name           string
	Email          string
	Password       string
	Auth           []Auth
}

// Auth : Auth
type Auth struct {
	gorm.Model `json:"-" xml:"-"`
	Group      Group
	User       User
	Auth       string
}

// Task : Task
type Task struct {
	gorm.Model     `json:"-" xml:"-"`
	CveID          string
	ServerName     string
	OrganizationID string
	Users          []User
	Comments       []Comment
}

// Comment : Comment
type Comment struct {
	gorm.Model `json:"-" xml:"-"`
	Content    string
	UserID     string
}
