package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-redis/redis"
	"github.com/inconshreveable/log15"
	cveModels "github.com/kotakanbe/go-cve-dictionary/models"
	//	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/vattle/sqlboiler/boil"
)

var redisClient *redis.Client

func init() {
}

// OpenDBs : OpenDBs
func OpenDBs() (err error) {
	if err = prepareViper(); err != nil {
		log15.Warn("Failed to read viper config file", "method", "main.main", "err", err)
		return err
	}
	var postgresURL = fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.dbname"),
		viper.GetString("postgres.sslmode"),
		viper.GetString("postgres.pass"),
	)

	if err = connectPostgres(postgresURL); err != nil {
		log15.Error("Failed to connect postgres", "method", "db.OpenDBs", "err", err)
		return err
	}
	if err = connectRedis(viper.GetString("redis.url")); err != nil {
		log15.Error("Failed to connect redis", "method", "db.OpenDBs", "err", err)
		return err
	}
	return nil
}

// GetCveByID : GetCveById
func GetCveByID(cveID string) (cveDetail *cveModels.CveDetail, err error) {
	var result *redis.StringStringMapCmd
	if result = redisClient.HGetAll(cveID); result.Err() != nil {
		log15.Error("Failed to HGetALL", "method", "db.GetCveByID", "err", result.Err())
		err = result.Err()
		return
	}
	if len(result.Val()) == 0 {
		log15.Info("HGetAll no result", "method", "db.GetCveByID")
		return nil, nil
	}

	jvn := cveModels.Jvn{}
	if j, ok := result.Val()["Jvn"]; ok {
		if err = json.Unmarshal([]byte(j), &jvn); err != nil {
			log15.Error("Failed to Unmarshal json", "method", "db.GetCveByID", "err", err)
			return
		}
	}

	nvd := cveModels.Nvd{}
	if j, ok := result.Val()["Nvd"]; ok {
		if err = json.Unmarshal([]byte(j), &nvd); err != nil {
			log15.Error("Failed to Unmarshal json", "method", "db.GetCveByID", "err", err)
			return
		}
	}

	cveDetail = &cveModels.CveDetail{
		CveID: cveID,
		Nvd:   nvd,
		Jvn:   jvn,
	}
	return
}

func connectRedis(url string) error {
	var err error
	var option *redis.Options
	if option, err = redis.ParseURL(url); err != nil {
		log15.Error("Failed to open redis", "url", url, "err", err)
		return err
	}
	if redisClient = redis.NewClient(option); redisClient.Ping().Err() != nil {
		return redisClient.Ping().Err()
	}
	log15.Info("connected to redis ", "url", url)
	return nil
}

func connectPostgres(url string) error {
	postgres, err := sql.Open("postgres", url)
	if err != nil {
		log15.Error("Failed to open postgres", "url", url, "err", err)
		return err
	}

	boil.SetDB(postgres)
	log15.Info("connected to postgres", "url", url)
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

	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}

	err = viper.ReadInConfig()
	return err
}
