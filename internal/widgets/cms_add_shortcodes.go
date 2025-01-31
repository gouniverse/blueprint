package widgets

import (
	"project/config"

	"github.com/gouniverse/cms"
)

func CmsAddShortcodes() {
	shortcodes := []cms.ShortcodeInterface{}

	list := WidgetRegistry()

	for _, widget := range list {
		shortcodes = append(shortcodes, widget)
	}

	config.Cms.ShortcodesAdd(shortcodes)
}
