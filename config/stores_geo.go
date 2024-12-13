package config

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gouniverse/geostore"
)

func init() {
	if GeoStoreUsed {
		addDatabaseInit(GeoStoreInitialize)
		addDatabaseMigration(GeoStoreAutoMigrate)
	}
}

func GeoStoreInitialize(db *sql.DB) error {
	if !GeoStoreUsed {
		return nil
	}

	geoStoreInstance, err := geostore.NewStore(geostore.NewStoreOptions{
		DB:                db,
		CountryTableName:  "snv_geo_country",
		StateTableName:    "snv_geo_state",
		TimezoneTableName: "snv_geo_timezone",
	})

	if err != nil {
		return errors.Join(errors.New("geostore.NewStore"), err)
	}

	if geoStoreInstance == nil {
		return errors.Join(errors.New("geoStoreInstance is nil"))
	}

	GeoStore = *geoStoreInstance

	return nil
}

func GeoStoreAutoMigrate(_ context.Context) error {
	if !GeoStoreUsed {
		return nil
	}

	err := GeoStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("geostore.AutoMigrate"), err)
	}

	return nil
}
