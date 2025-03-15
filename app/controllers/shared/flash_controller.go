package shared

import (
	"net/http"
	"project/app/layouts"
	"project/app/links"
	"project/config"
	"project/internal/helpers"

	"strings"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/icons"
	"github.com/gouniverse/router"
	"github.com/gouniverse/utils"
)

// == CONTROLLER ==============================================================

type flashController struct{}

var _ router.HTMLControllerInterface = (*flashController)(nil)

// == CONSTRUCTOR =============================================================

func NewFlashController() *flashController {
	return &flashController{}
}

// == PUBLIC METHODS ==========================================================

func (controller flashController) Handler(w http.ResponseWriter, r *http.Request) string {
	authUser := helpers.GetAuthUser(r)

	if authUser != nil && authUser.IsRegistrationCompleted() {
		return layouts.NewUserLayout(r, layouts.Options{
			Title:      "System Message",
			Content:    controller.pageHTML(r),
			ScriptURLs: []string{},
			Styles:     []string{`.Center > div{padding:0px !important;margin:0px !important;}`},
		}).ToHTML()
	}

	if config.CmsStoreUsed {
		return layouts.NewWebsiteLayout(layouts.Options{
			Request:    r,
			Title:      "System Message",
			Content:    controller.pageHTML(r),
			ScriptURLs: []string{},
			Styles:     []string{`.Center > div{padding:0px !important;margin:0px !important;}`},
		}).ToHTML()
	} else {
		return layouts.NewUserLayout(r, layouts.Options{
			Title:      "System Message",
			Content:    controller.pageHTML(r),
			ScriptURLs: []string{},
		}).ToHTML()
	}
}

func (c flashController) pageHTML(r *http.Request) hb.TagInterface {
	messageID := utils.Req(r, "message_id", "")
	msgData, err := config.CacheStore.GetJSON(messageID+"_flash_message", "")

	msgType := "error"
	message := "The message is no longer available"
	url := links.NewWebsiteLinks().Home()
	time := "5"

	if err != nil {
		message = "The message is no longer available"
	}

	if msgData == "" {
		message = "The message is no longer available"
	}

	if msgData != "" {
		msgDataAny := msgData.(map[string]interface{})
		msgType = utils.ToString(msgDataAny["type"])
		message = utils.ToString(msgDataAny["message"])
		url = utils.ToString(msgDataAny["url"])
		time = utils.ToString(msgDataAny["time"])
	}

	alert := hb.Div()
	alertIcon := ""
	if msgType == "error" {
		alert.Class("alert alert-danger")
		alertIcon = icons.BootstrapExclamationOctagon
	} else if msgType == "success" {
		alert.Class("alert alert-success")
		alertIcon = icons.BootstrapCheckCircle
	} else if msgType == "warning" {
		alert.Class("alert alert-warning")
		alertIcon = icons.BootstrapExclamationTriangle
	} else {
		alert.Class("alert alert-info")
		alertIcon = icons.BootstrapInfoCircle
	}

	css := ""
	css += "div.alert-success{color:green;}"
	css += "div.alert-danger{color:red;}"
	css += "div.alert-info{color:blue;}"
	css += "div.alert-warning{color:warning;}"

	icon := strings.ReplaceAll(alertIcon, "height=\"16\"", "height=\"24\"")
	icon = strings.ReplaceAll(icon, "width=\"16\"", "width=\"24\"")
	alert.Child(hb.Span().
		Child(hb.Span().
			HTML(icon).
			Style("position:absolute;top:-16px;")).
		Style("position:relative; margin:0px 20px 0px 0px; display:inline-table;width:24px;"))
	alert.Child(hb.Span().HTML(message))

	container := hb.Div().
		Class("container").
		Style("padding:0px 0px 20px 0px;text-align:left;").
		Child(hb.Style(css)).
		Child(alert)

	if url != "" {
		link := hb.Hyperlink().Href(url).HTML("Click here to continue")
		divLink := hb.Div()
		divLink.AddChild(link).Style("padding:20px 0px 20px 0px;")
		container.AddChild(divLink)
	}

	if url != "" && time != "" {
		script := hb.Script("setTimeout(()=>{location.href=\"" + url + "\"}, " + time + "*1000)")
		container.AddChild(script)
	}

	return hb.Section().
		Children([]hb.TagInterface{
			container,
		}).
		Style("padding: 80px 0px 40px 0px;")
}
