package admin

import (
	"net/http"
	"project/config"
	"project/internal/layouts"
	"project/internal/links"
	"strings"

	"github.com/goravel/framework/support/carbon"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/gouniverse/statsstore"
	"github.com/samber/lo"
)

// == CONTROLLER ===============================================================

type visitorPathsController struct{}

var _ router.HTMLControllerInterface = (*visitorPathsController)(nil)

// == CONSTRUCTOR ==============================================================

func NewVisitorPathsController() *visitorPathsController {
	return &visitorPathsController{}
}

// == PUBLIC METHODS ===========================================================

func (controller *visitorPathsController) Handler(w http.ResponseWriter, r *http.Request) string {
	return layouts.NewAdminLayout(r, layouts.Options{
		Title:      "Visitor Paths | Stats",
		Content:    controller.view(),
		ScriptURLs: []string{},
		Styles:     []string{},
	}).ToHTML()
}

// == PRIVATE METHODS ==========================================================

func (c *visitorPathsController) view() *hb.Tag {
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
			Name: "Visitor Paths",
			URL: links.NewAdminLinks().Stats(map[string]string{
				"controller": "visitor-paths",
			}),
		},
	})

	title := hb.Heading1().
		HTML("Stats. Visitor Paths")

	return hb.Div().
		Class("container").
		Child(title).
		Child(breadcrumbs).
		Child(c.cardVisitorPaths())
}

// == PRIVATE METHODS ==========================================================

func (c *visitorPathsController) cardVisitorPaths() hb.TagInterface {
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
				HTML("Visitor Paths"))).
		Child(hb.Div().
			Class("card-body").
			Child(c.tableVisitors(visitors)).
			Child(hb.BR())).
		Child(hb.Div().
			Class("card-footer"))
}

func (c *visitorPathsController) tableVisitors(visitors []vistorActivity) hb.TagInterface {
	cards := hb.Wrap().
		Children(lo.Map(visitors, func(visitor vistorActivity, index int) hb.TagInterface {
			ip := visitor.VisitorIP
			country := visitor.VisitorCountry
			device := visitor.VisitorDevice
			deviceType := visitor.VisitorDeviceType
			browser := visitor.VisitorBrowser
			browserVersion := visitor.VisitorBrowserVersion
			// userAgent := visitor.UserAgent()
			// userLanguage := visitor.UserAcceptLanguage()
			// encoding := visitor.UserAcceptEncoding()
			os := visitor.VisitorOS
			osVersion := visitor.VisitorOSVersion
			referrer := visitor.VisitorPaths[0].Referrer
			if referrer == "" {
				referrer = "(No referring link)"
			}

			countryFlagSrc := "https://flagicons.lipis.dev/flags/4x3/" + strings.ToLower(country) + ".svg"
			countryFlagSrc = strings.ReplaceAll(countryFlagSrc, "un", "xx")
			countryFlag := hb.Img(countryFlagSrc).
				Style("width: 20px;").
				Title("Country code: " + country)

			tableHeader := hb.Table().
				Style("width: 100%;").
				Style("font-size: 14px;").
				Child(hb.TR().
					Child(hb.TD().
						Style("width: 140px;").
						Child(countryFlag).
						Text(" ").
						Text(ip)).
					Child(hb.TD().
						Style("width: 200px;").
						Text(os).
						Text(" ").
						Text(osVersion).
						Text(", ").
						Text(deviceType).
						Text(" ").
						Text(device),
					).
					Child(hb.TD().
						Text(browser).
						Text(" ").
						Text(browserVersion),
					))

			tableBody := hb.Table().
				Style("width: 100%;").
				Child(hb.TR().
					Child(hb.TD().
						Style("width: 100px;").
						Text("")).
					Child(hb.TD().
						Style("width: 100px;").
						Text("")).
					Child(hb.TD().
						Child(hb.Div().Style("color: limegreen;").Text("Referrer: ").Text(referrer)))).
				Children(lo.Map(visitor.VisitorPaths, func(path visitorPath, index int) hb.TagInterface {
					date := path.Date.ToDateString()
					time := path.Date.ToTimeString()
					link := links.NewWebsiteLinks().Home() + path.Path
					link = strings.ReplaceAll(link, "[GET]", "")

					hyperlink := hb.Hyperlink().
						Target("_blank").
						Href(link).
						Text(path.Path)

					return hb.TR().
						Child(hb.TD().
							Style("width: 100px;").
							Text(date)).
						Child(hb.TD().
							Style("width: 100px;").
							Text(time)).
						Child(hb.TD().
							Child(hb.Div().Style("color: crimson;").Child(hyperlink)))
				}))

			return hb.Div().
				Class("card mb-3").
				Child(hb.Div().
					Class("card-header").
					Child(tableHeader)).
				Child(hb.Div().
					Class("card-body").
					Child(tableBody))
		}))

	return cards
}

func (c *visitorPathsController) dataVisitors() ([]vistorActivity, error) {
	startDate := carbon.Now().SubDays(31).ToDateString()
	endDate := carbon.Now().ToDateString()

	visits, err := config.StatsStore.VisitorList(statsstore.VisitorQueryOptions{
		CreatedAtGte: startDate + " 00:00:00",
		CreatedAtLte: endDate + " 23:59:59",
	})

	if err != nil {
		return nil, err
	}

	vistorActivitys := []vistorActivity{}

	for _, visit := range visits {
		fingerprint := visit.Fingerprint()

		activity := vistorActivity{
			VisitorCountry:        visit.Country(),
			VisitorDevice:         visit.UserDevice(),
			VisitorDeviceType:     visit.UserDeviceType(),
			VisitorFingerprint:    visit.Fingerprint(),
			VisitorIP:             visit.IpAddress(),
			VisitorOS:             visit.UserOs(),
			VisitorOSVersion:      visit.UserOsVersion(),
			VisitorBrowser:        visit.UserBrowser(),
			VisitorBrowserVersion: visit.UserBrowserVersion(),
			VisitorPaths: []visitorPath{
				{
					Date:     visit.CreatedAtCarbon(),
					Path:     visit.Path(),
					Referrer: visit.UserReferrer(),
				},
			},
		}

		_, index, isFound := lo.FindIndexOf(vistorActivitys, func(v vistorActivity) bool {
			return v.VisitorFingerprint == fingerprint
		})

		if isFound {
			vistorActivitys[index].VisitorPaths = append(vistorActivitys[index].VisitorPaths, visitorPath{
				Date:     visit.CreatedAtCarbon(),
				Path:     visit.Path(),
				Referrer: visit.UserReferrer(),
			})
		} else {
			vistorActivitys = append(vistorActivitys, activity)
		}
	}

	return vistorActivitys, err
}

type vistorActivity struct {
	VisitorCountry        string
	VisitorDevice         string
	VisitorDeviceType     string
	VisitorFingerprint    string
	VisitorIP             string
	VisitorOS             string
	VisitorOSVersion      string
	VisitorBrowser        string
	VisitorBrowserVersion string
	VisitorPaths          []visitorPath
}

type visitorPath struct {
	Date     carbon.Carbon
	Path     string
	Referrer string
}
