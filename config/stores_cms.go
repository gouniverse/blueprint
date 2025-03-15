package config

import (
	"context"
	"database/sql"
	"errors"
	"project/pkg/webtheme"

	"github.com/gouniverse/cms"
	"github.com/gouniverse/cmsstore"
	"github.com/gouniverse/ui"
)

func init() {
	if CmsUsed {
		addDatabaseInit(CmsInitialize)
		addDatabaseMigration(CmsAutoMigrate)
	}

	if CmsStoreUsed {
		addDatabaseInit(CmsStoreInitialize)
		addDatabaseMigration(CmsStoreAutoMigrate)
	}
}

func CmsInitialize(db *sql.DB) error {
	if !CmsUsed {
		return nil
	}

	cmsInstance, err := cms.NewCms(cms.Config{
		Database:               Database,
		Prefix:                 "cms_",
		TemplatesEnable:        true,
		PagesEnable:            true,
		MenusEnable:            true,
		BlocksEnable:           true,
		BlockEditorDefinitions: webtheme.BlockEditorDefinitions(),
		BlockEditorRenderer: func(blocks []ui.BlockInterface) string {
			return webtheme.New(blocks).ToHtml()
		},
		//CacheAutomigrate:    true,
		//CacheEnable:         true,
		EntitiesAutomigrate: true,
		//LogsEnable:          true,
		//LogsAutomigrate:     true,
		// SettingsEnable: true,
		//SettingsAutomigrate: true,
		//SessionAutomigrate:  true,
		//SessionEnable:       true,
		Shortcodes: []cms.ShortcodeInterface{},
		//TasksEnable:         true,
		//TasksAutomigrate:    true,
		TranslationsEnable:         true,
		TranslationLanguageDefault: TranslationLanguageDefault,
		TranslationLanguages:       TranslationLanguageList,
		// CustomEntityList: customEntityList(),
	})

	if err != nil {
		return errors.Join(errors.New("cms.NewCms"), err)
	}

	if cmsInstance == nil {
		panic("cmsInstance is nil")
	}

	Cms = *cmsInstance

	return nil
}

func CmsAutoMigrate(_ context.Context) error {
	if !CmsStoreUsed {
		return nil
	}

	// !!! No need. Migrated during initialize
	// err := Cms.AutoMigrate()

	// if err != nil {
	// 	return errors.Join(errors.New("cms.AutoMigrate"), err)
	// }

	return nil
}

func CmsStoreInitialize(db *sql.DB) error {
	if !CmsStoreUsed {
		return nil
	}

	cmsStoreInstance, err := cmsstore.NewStore(cmsstore.NewStoreOptions{
		DB: db,

		BlockTableName:    "snv_cms_block",
		PageTableName:     "snv_cms_page",
		TemplateTableName: "snv_cms_template",
		SiteTableName:     "snv_cms_site",

		MenusEnabled:      true,
		MenuItemTableName: "snv_cms_menu_item",
		MenuTableName:     "snv_cms_menu",

		TranslationsEnabled:        true,
		TranslationTableName:       "snv_cms_translation",
		TranslationLanguageDefault: TranslationLanguageDefault,
		TranslationLanguages:       TranslationLanguageList,

		VersioningEnabled:   true,
		VersioningTableName: "snv_cms_version",
	})

	if err != nil {
		return errors.Join(errors.New("cmsstore.NewStore"), err)
	}

	if cmsStoreInstance == nil {
		return errors.New("cmsstore.NewStore: cmsStoreInstance is nil")
	}

	CmsStore = cmsStoreInstance

	return nil
}

func CmsStoreAutoMigrate(ctx context.Context) error {
	if !CmsStoreUsed {
		return nil
	}

	if CmsStore == nil {
		return errors.New("cmsstore.AutoMigrate: CmsStore is nil")
	}

	err := CmsStore.AutoMigrate(ctx)

	if err != nil {
		return errors.Join(errors.New("cmsstore.AutoMigrate"), err)
	}

	return nil
}
