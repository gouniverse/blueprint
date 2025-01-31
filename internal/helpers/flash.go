package helpers

import (
	"net/http"
	"project/config"
	"project/internal/links"
	"strings"

	"github.com/gouniverse/uid"
)

const FLASH_ERROR = "error"
const FLASH_SUCCESS = "success"
const FLASH_INFO = "info"
const FLASH_WARNING = "warning"

func IsFlashRoute(r *http.Request) bool {
	return strings.Contains(r.URL.Path, "/flash")
}

// ToFlashURL return a flash message URL
func ToFlashURL(messageType string, message string, url string, seconds int) string {
	id := uid.HumanUid()
	config.CacheStore.SetJSON(id+"_flash_message", map[string]any{
		"type":    messageType,
		"message": message,
		"url":     url,
		"time":    seconds,
	}, int64(seconds)+10)

	return links.NewWebsiteLinks().Flash(map[string]string{
		"message_id": id,
	})
}

// ToFlash redirects the user to a flash page
func ToFlash(w http.ResponseWriter, r *http.Request, messageType string, message string, url string, seconds int) string {
	flashUrl := ToFlashURL(messageType, message, url, seconds)

	http.Redirect(w, r, flashUrl, http.StatusSeeOther)
	return `<a href="` + flashUrl + `">See Other</a>`
}

func ToFlashError(w http.ResponseWriter, r *http.Request, message string, url string, seconds int) string {
	return ToFlash(w, r, FLASH_ERROR, message, url, seconds)
}

func ToFlashInfo(w http.ResponseWriter, r *http.Request, message string, url string, seconds int) string {
	return ToFlash(w, r, FLASH_INFO, message, url, seconds)
}

func ToFlashSuccess(w http.ResponseWriter, r *http.Request, message string, url string, seconds int) string {
	return ToFlash(w, r, FLASH_SUCCESS, message, url, seconds)
}

func ToFlashWarning(w http.ResponseWriter, r *http.Request, message string, url string, seconds int) string {
	return ToFlash(w, r, FLASH_WARNING, message, url, seconds)
}

func ToFlashErrorURL(message string, url string, seconds int) string {
	return ToFlashURL(FLASH_ERROR, message, url, seconds)
}

func ToFlashInfoURL(message string, url string, seconds int) string {
	return ToFlashURL(FLASH_INFO, message, url, seconds)
}

func ToFlashSuccessURL(message string, url string, seconds int) string {
	return ToFlashURL(FLASH_SUCCESS, message, url, seconds)
}

func ToFlashWarningURL(message string, url string, seconds int) string {
	return ToFlashURL(FLASH_WARNING, message, url, seconds)
}
