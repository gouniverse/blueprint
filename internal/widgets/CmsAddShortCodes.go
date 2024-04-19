package widgets

import (
	"project/config"

	"github.com/gouniverse/cms"
)

func CmsAddShortcodes() {
	shortcodes := []cms.ShortcodeInterface{
		NewAuthenticatedWidget(),
		NewUnauthenticatedWidget(),
		NewVisibleWidget(),
	}

	config.Cms.ShortcodesAdd(shortcodes)
}
