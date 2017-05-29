package db

import (
	"database/sql"

	"../models"
	"github.com/inconshreveable/log15"
	_ "github.com/lib/pq"
	"github.com/vattle/sqlboiler/boil"
)

func init() {
	db, err := sql.Open("postgres", `host=0.0.0.0  user=postgres dbname=saas sslmode=disable password=vuls`)
	if err != nil {
		log15.Error("Failed to open database", "err", err)

	}
	log15.Info("connectd to db")
	boil.SetDB(db)
}

// SelectSome : select
func SelectSome() {
	cveDetail, err := models.CveDetailsG().One()
	if err != nil {
		log15.Error("Failed to Read cveDetail", "err", err)
	}
	log15.Info("Fetched cveDetail", "CVEID", cveDetail.CveID)

	cpe, err := models.CpesG().One()
	if err != nil {
		log15.Error("Failed to Read cpe", "err", err)
	}
	log15.Info("Fetched cpe", "cpe", cpe)

}
