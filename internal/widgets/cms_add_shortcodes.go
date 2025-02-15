package widgets

import (
	"project/config"

	"github.com/gouniverse/cmsstore"
)

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
