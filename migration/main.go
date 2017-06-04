package main

import (
	"fmt"
	"os"
	"path/filepath"

	"./models"
	//	"github.com/future-architect/vuls/models"
	"github.com/inconshreveable/log15"
	"github.com/jinzhu/gorm"
	//	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var db *gorm.DB

func main() {

	var err error
	if err = prepareViper(); err != nil {
		log15.Warn("Failed to read viper config file", "method", "main.main", "err", err)
	}
	var dsn = fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.dbname"),
		viper.GetString("postgres.sslmode"),
		viper.GetString("postgres.pass"),
	)
	log15.Info("check dsn", "method", "main.main", "dsn", dsn)
	if db, err = gorm.Open("postgres", dsn); err != nil {
		log15.Warn("Faied to open db", "method", "main.main", "dsn", dsn, "err", err)
	}
	log15.Info("Succeed to connect databse", "method", "main.main", "dsn", dsn)

	MigrateVulsDB()
	MigrateUserDB()

}

// MigrateVulsDB migrates Database
func MigrateVulsDB() error {
	log15.Info("Migrate database", "method", "main.MigrateVulsDB")
	if err := db.AutoMigrate(
		&models.ScanResult{},
		&models.Container{},
		&models.Platform{},
		&models.VulnInfo{},
		&models.Confidence{},
		&models.DistroAdvisory{},
		&models.PackageInfo{},
		&models.Changelog{},
	).Error; err != nil {
		log15.Error("Failed to migrate vuls db.", "err", err)
		return err
	}

	log15.Info("Create Index on vuls db", "method", "main.MigrateVulsDB")
	errMsg := "Failed to create index. err: %s"
	if err := db.Model(&models.ScanResult{}).
		AddIndex("idx_scanresult_id", "id").Error; err != nil {
		return fmt.Errorf(errMsg, err)
	}
	// TODO: add more key on container
	if err := db.Model(&models.Container{}).
		AddIndex("idx_container_id", "id").Error; err != nil {
		return fmt.Errorf(errMsg, err)
	}
	return nil
}

// MigrateUserDB migrates Database
func MigrateUserDB() error {
	log15.Info("Migrate database", "method", "main.MigrateUserDB")
	if err := db.AutoMigrate(
		&models.Organization{},
		&models.Group{},
		&models.User{},
		&models.Auth{},
		&models.Task{},
		&models.Comment{},
	).Error; err != nil {
		log15.Error("Failed to migrate user db.", "err", err)
		return err
	}

	//	log15.Info("Create Index on user db", "method", "main.MigrateUserDB")
	//	errMsg := "Failed to create index. err: %s"
	//	if err := db.Model(&models.ScanResult{}).
	//		AddIndex("idx_scanresult_id", "id").Error; err != nil {
	//		return fmt.Errorf(errMsg, err)
	//	}
	//	// TODO: add more key on container
	//	if err := db.Model(&models.Container{}).
	//		AddIndex("idx_container_id", "id").Error; err != nil {
	//		return fmt.Errorf(errMsg, err)
	//	}

	return nil
}

func prepareViper() error {
	log15.Info("call prepareViper", "method", "main.prepareViper")
	viper.SetConfigName("sqlboiler")

	configHome := os.Getenv("XDG_CONFIG_HOME")
	homePath := os.Getenv("HOME")
	wd, err := os.Getwd()
	if err != nil {
		wd = "./"
	}

	configPaths := []string{wd}
	if len(configHome) > 0 {
		configPaths = append(configPaths, filepath.Join(configHome, "sqlboiler"))
	} else {
		configPaths = append(configPaths, filepath.Join(homePath, ".config/sqlboiler"))
	}
	configPaths = append(configPaths, filepath.Join("../"))

	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}

	err = viper.ReadInConfig()
	return err
}
