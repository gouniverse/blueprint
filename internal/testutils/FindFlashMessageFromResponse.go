package testutils

import (
	"net/http"
	"project/config"
	"project/internal/types"

	"github.com/gouniverse/utils"
)

func FlashMessageFind(messageID string) (msg *types.FlashMessage, err error) {
	msgData, err := config.CacheStore.GetJSON(messageID+"_flash_message", "")
	if err != nil {
		return msg, err
	}

	if msgData == "" {
		return msg, nil
	}

	msgDataAny := msgData.(map[string]interface{})
	dataMap := &types.FlashMessage{
		Type:    utils.ToString(msgDataAny["type"]),
		Message: utils.ToString(msgDataAny["message"]),
		Url:     utils.ToString(msgDataAny["url"]),
		Time:    utils.ToString(msgDataAny["time"]),
	}

	return dataMap, nil
}

func FlashMessageFindFromBody(body string) (msg *types.FlashMessage, err error) {
	flashMessageID := utils.StrLeftFrom(utils.StrRightFrom(body, `/flash?message_id=`), `"`)
	return FlashMessageFind(flashMessageID)
}

func FlashMessageFindFromResponse(r *http.Response) (msg *types.FlashMessage, err error) {
	location := r.Header.Get("Location")
	flashMessageID := utils.StrRightFrom(location, `/flash?message_id=`)
	return FlashMessageFind(flashMessageID)
}
