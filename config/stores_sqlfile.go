package config

import (
	"database/sql"
	"errors"

	"github.com/gouniverse/filesystem"
)

func init() {
	if SqlFileStoreUsed {
		addDatabaseInit(SqlFileStoreInitialize)
		addDatabaseMigration(SqlFileStoreAutoMigrate)
	}
}

func SqlFileStoreInitialize(db *sql.DB) error {
	if !SqlFileStoreUsed {
		return nil
	}

	sqlFileStorageInstance, err := filesystem.NewStorage(filesystem.Disk{
		DiskName:  filesystem.DRIVER_SQL,
		Driver:    filesystem.DRIVER_SQL,
		Url:       "/files",
		DB:        db,
		TableName: "snv_files_file",
	})

	if err != nil {
		return errors.Join(errors.New("filesystem.NewStorage"), err)
	}

	if sqlFileStorageInstance == nil {
		return errors.Join(errors.New("sqlFileStorageInstance is nil"))
	}

	SqlFileStorage = sqlFileStorageInstance

	return nil
}

func SqlFileStoreAutoMigrate() error {
	if !SqlFileStoreUsed {
		return nil
	}

	// !!! No need. Migrated during initialize
	// err := SqlFileStorage.AutoMigrate()

	// if err != nil {
	// 	return errors.Join(errors.New("filesystem.AutoMigrate"), err)
	// }

	return nil
}
