package admin

import (
	"net/http"
	"project/config"
	"project/internal/layouts"
	"project/internal/links"

	"github.com/goravel/framework/support/carbon"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/gouniverse/statsstore"
	"github.com/samber/lo"
)

// == CONTROLLER ===============================================================

type visitorActivityController struct{}

var _ router.HTMLControllerInterface = (*visitorActivityController)(nil)

// == CONSTRUCTOR ==============================================================

func NewVisitorActivityController() *visitorActivityController {
	return &visitorActivityController{}
}

// == PUBLIC METHODS ===========================================================

func (controller *visitorActivityController) Handler(w http.ResponseWriter, r *http.Request) string {
	return layouts.NewAdminLayout(r, layouts.Options{
		Title:      "Visitor Activity | Stats",
		Content:    controller.view(),
		ScriptURLs: []string{},
		Styles:     []string{},
	}).ToHTML()
}

// == PRIVATE METHODS ==========================================================

func (c *visitorActivityController) view() *hb.Tag {
	breadcrumbs := layouts.Breadcrumbs([]layouts.Breadcrumb{
		{
			Name: "Home",
			URL:  links.NewAdminLinks().Home(map[string]string{}),
		},
		{
			Name: "Stats",
			URL:  links.NewAdminLinks().Stats(map[string]string{}),
		},
		{
			Name: "Visitor Activity",
			URL: links.NewAdminLinks().Stats(map[string]string{
				"controller": "visitor-activity",
			}),
		},
	})

	title := hb.Heading1().
		HTML("Stats. Visitor Activity")

	return hb.Div().
		Class("container").
		Child(title).
		Child(breadcrumbs).
		Child(c.cardVisitorActivity())
}

// == PRIVATE METHODS ==========================================================

func (c *visitorActivityController) cardVisitorActivity() hb.TagInterface {
	visitors, err := c.dataVisitors()

	if err != nil {
		return hb.Div().Class("alert alert-danger").Text(err.Error())
	}

	return hb.Div().
		Class("card").
		Child(hb.Div().
			Class("card-header").
			Child(hb.Heading2().
				Class("card-title").
				Style("margin-bottom: 0px;").
				HTML("Visitor Activity"))).
		Child(hb.Div().
			Class("card-body").
			Child(c.tableVisitors(visitors)).
			Child(hb.BR())).
		Child(hb.Div().
			Class("card-footer"))
}

func (c *visitorActivityController) tableVisitors(visitors []statsstore.VisitorInterface) hb.TagInterface {
	// Page Views:
	// 1
	// Latest Page View:
	// 27 Oct 2024 00:28:02
	// Resolution:
	// 414x736
	// System:
	// Chrome for Android/Android
	// Vivo Y5s

	// Total Sessions:
	// 1
	// Location:
	// [China] China
	// ISP / IP Address:
	// China Mobile (39.173.105.141)
	// Referring URL:
	// (No referring link)
	// Visit Page:
	//  https://lesichkov.co.uk/

	table := hb.Table().
		Class("table table-striped table-bordered").
		Children([]hb.TagInterface{
			hb.Thead().Children([]hb.TagInterface{
				hb.TR().Children([]hb.TagInterface{
					// hb.TH().Text("Date"),
					// hb.TH().Text("Unique Visitors"),
					// hb.TH().Text("Total Visitors"),
				}),
			}),
			hb.Tbody().Children(lo.Map(visitors, func(visitor statsstore.VisitorInterface, index int) hb.TagInterface {
				date := visitor.CreatedAtCarbon().ToDateString()
				ip := visitor.IpAddress()
				country := visitor.Country()
				device := visitor.UserDevice()
				deviceType := visitor.UserDeviceType()
				browser := visitor.UserBrowser()
				browserVersion := visitor.UserBrowserVersion()
				userAgent := visitor.UserAgent()
				userLanguage := visitor.UserAcceptLanguage()
				encoding := visitor.UserAcceptEncoding()
				os := visitor.UserOs()
				referrer := visitor.UserReferrer()
				path := visitor.Path()
				return hb.TR().Children([]hb.TagInterface{
					hb.TD().Child(hb.NewSpan().Text("Path: ")).Text(path),
					hb.TD().Child(hb.NewSpan().Text("Date: ")).Text(date),
					hb.TD().Child(hb.NewSpan().Text("IP: ")).Text(ip),
					hb.TD().Child(hb.NewSpan().Text("Country: ")).Text(country),
					hb.TD().Child(hb.NewSpan().Text("Device: ")).Text(device),
					hb.TD().Child(hb.NewSpan().Text("Device Type: ")).Text(deviceType),
					hb.TD().Child(hb.NewSpan().Text("Browser: ")).Text(browser),
					hb.TD().Child(hb.NewSpan().Text("Browser Version: ")).Text(browserVersion),
					hb.TD().Child(hb.NewSpan().Text("OS: ")).Text(os),
					hb.TD().Child(hb.NewSpan().Text("Referrer: ")).Text(referrer),
					hb.TD().Child(hb.NewSpan().Text("User Agent: ")).Text(userAgent),
					hb.TD().Child(hb.NewSpan().Text("User Language: ")).Text(userLanguage),
					hb.TD().Child(hb.NewSpan().Text("User Encoding: ")).Text(encoding),
				})
			})),
			hb.Tfoot().Children([]hb.TagInterface{
				// hb.TR().Children([]hb.TagInterface{
				// 	hb.TH().Text("Total"),
				// 	hb.TH().Text(cast.ToString(lo.Sum(uniqueVisits))),
				// 	hb.TH().Text(cast.ToString(lo.Sum(totalVisits))),
				// }),
			}),
		})

	return hb.Wrap().
		Child(table)
}

func (c *visitorActivityController) datesInRange(timeStart, timeEnd carbon.Carbon) []string {
	rangeDates := []string{}

	if timeStart.Lte(timeEnd) {
		rangeDates = append(rangeDates, timeStart.ToDateString())
		for timeStart.Lt(timeEnd) {
			timeStart = timeStart.AddDays(1) // += 86400 // add 24 hours
			rangeDates = append(rangeDates, timeStart.ToDateString())
		}
	}

	return rangeDates
}

func (c *visitorActivityController) dataVisitors() (dates []statsstore.VisitorInterface, err error) {
	startDate := carbon.Now().SubDays(31).ToDateString()
	endDate := carbon.Now().ToDateString()

	visitors, err := config.StatsStore.VisitorList(statsstore.VisitorQueryOptions{
		CreatedAtGte: startDate + " 00:00:00",
		CreatedAtLte: endDate + " 23:59:59",
	})

	return visitors, err
}
