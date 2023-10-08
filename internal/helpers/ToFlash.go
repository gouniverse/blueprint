package helpers

import (
	"fmt"
	"net/http"
	"project/config"
	"project/internal/links"

	"github.com/gouniverse/uid"
)

// GuestFlash redirects the user to a flash page
func ToFlash(w http.ResponseWriter, r *http.Request, messageType string, message string, url string, seconds int) string {
	id := uid.HumanUid()
	config.Cms.CacheStore.Set(id+"_flash_message", message, int64(seconds)+10)
	config.Cms.CacheStore.Set(id+"_flash_type", messageType, int64(seconds)+10)
	config.Cms.CacheStore.Set(id+"_flash_url", url, int64(seconds)+10)
	config.Cms.CacheStore.Set(id+"_flash_time", fmt.Sprint(seconds), int64(seconds)+10)
	http.Redirect(w, r, links.NewWebsiteLinks().Flash(map[string]string{
		"message_id": id,
	}), http.StatusSeeOther)
	return ""
}
