package config

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gouniverse/taskstore"
)

func init() {
	if TaskStoreUsed {
		addDatabaseInit(TaskStoreInitialize)
		addDatabaseMigration(TaskStoreAutoMigrate)
	}
}

func TaskStoreInitialize(db *sql.DB) error {
	if !TaskStoreUsed {
		return nil
	}

	taskStoreInstance, err := taskstore.NewStore(taskstore.NewStoreOptions{
		DB:             db,
		TaskTableName:  "snv_tasks_task",
		QueueTableName: "snv_tasks_queue",
	})

	if err != nil {
		return errors.Join(errors.New("taskstore.NewStore"), err)
	}

	if taskStoreInstance == nil {
		return errors.Join(errors.New("taskStoreInstance is nil"))
	}

	TaskStore = taskStoreInstance

	return nil
}

func TaskStoreAutoMigrate(_ context.Context) error {
	if !TaskStoreUsed {
		return nil
	}

	if TaskStore == nil {
		return errors.New("taskstore.AutoMigrate: TaskStore is nil")
	}

	err := TaskStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("taskstore.AutoMigrate"), err)
	}

	return nil
}
