package widgets

import (
	"project/config"

	"github.com/gouniverse/cmsstore"
)

// CmsAddShortcodes adds the shortcodes to the CMS store
//
// Business Logic:
//   - Check if the CMS store is used
//   - Check if the CMS store is nil
//   - Add the shortcodes to the CMS store
//   - Loaded in the main.go file
//
// Parameters:
//   - None
//
// Returns:
//   - None
func CmsAddShortcodes() {
	if !config.CmsStoreUsed {
		return
	}

	if config.CmsStore == nil {
		return
	}

	shortcodes := []cmsstore.ShortcodeInterface{}

	list := WidgetRegistry()

	for _, widget := range list {
		shortcodes = append(shortcodes, widget)
	}

	config.CmsStore.AddShortcodes(shortcodes)
}
