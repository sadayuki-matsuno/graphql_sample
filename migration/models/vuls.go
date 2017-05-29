package models

import (
	"time"

	"github.com/future-architect/vuls/models"
	"github.com/jinzhu/gorm"
)

// ScanResult : copy from models.ScanResult
type ScanResult struct {
	gorm.Model `json:"-" xml:"-"`
	ScannedAt  time.Time

	Lang       string
	ServerName string // TOML Section key
	Family     string
	Release    string
	Errors     []string        `gorm:"type:varchar(100)[]"`
	Optional   [][]interface{} `gorm:"-"`
}

// Container : copy from models.Container
type Container struct {
	gorm.Model `json:"-" xml:"-"`
	models.Container
}

// Platform : copy from models.Platform
type Platform struct {
	gorm.Model `json:"-" xml:"-"`
	models.Platform
}

// VulnInfo : copy from models.VulnInfo
type VulnInfo struct {
	gorm.Model `json:"-" xml:"-"`
	CveID      string
	CpeNames   []string `gorm:"type:varchar(100)[]"`
}

// Confidence : copy from models.Confidence
type Confidence struct {
	gorm.Model `json:"-" xml:"-"`
	models.Confidence
}

// DistroAdvisory : copy from models.DistroAdvisory
type DistroAdvisory struct {
	gorm.Model `json:"-" xml:"-"`
	models.DistroAdvisory
}

// PackageInfo : copy from models.PackageInfo
type PackageInfo struct {
	gorm.Model `json:"-" xml:"-"`
	models.PackageInfo
}

// Changelog : copy from models.Changelog
type Changelog struct {
	gorm.Model `json:"-" xml:"-"`
	models.Changelog
}
