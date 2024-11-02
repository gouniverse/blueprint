package admin

import (
	"fmt"
	"net/http"
	"project/config"
	"project/internal/layouts"
	"project/internal/links"

	"github.com/goravel/framework/support/carbon"
	"github.com/gouniverse/bs"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/gouniverse/statsstore"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

// == CONTROLLER ===============================================================

type homeController struct{}

var _ router.HTMLControllerInterface = (*homeController)(nil)

// == CONSTRUCTOR ==============================================================

func NewHomeController() *homeController {
	return &homeController{}
}

// == PUBLIC METHODS ===========================================================

func (controller *homeController) Handler(w http.ResponseWriter, r *http.Request) string {
	return layouts.NewAdminLayout(r, layouts.Options{
		Title:   "Stats",
		Content: controller.view(),
		ScriptURLs: []string{
			cdn.Jquery_3_7_1(),
			// `https://cdnjs.cloudflare.com/ajax/libs/Chart.js/1.0.2/Chart.min.js`,
			`https://cdn.jsdelivr.net/npm/chart.js`,
		},
		Styles: []string{},
	}).ToHTML()
}

// == PRIVATE METHODS ==========================================================

func (c *homeController) view() *hb.Tag {
	breadcrumbs := layouts.Breadcrumbs([]layouts.Breadcrumb{
		{
			Name: "Home",
			URL:  links.NewAdminLinks().Home(map[string]string{}),
		},
		{
			Name: "Stats",
			URL:  links.NewAdminLinks().Stats(map[string]string{}),
		},
	})

	title := hb.Heading1().
		HTML("Stats. Home")

	options :=
		hb.Section().
			Class("mb-3 mt-3").
			Style("background-color: #f8f9fa;").
			Child(
				hb.UL().
					Class("list-group").
					Child(hb.LI().
						Class("list-group-item").
						Child(hb.A().
							Href(links.NewAdminLinks().Stats(map[string]string{
								"controller": "visitor-activity",
							})).
							Text("Visitor Activity")).
						Child(hb.LI().
							Class("list-group-item").
							Child(hb.A().
								Href(links.NewAdminLinks().Stats(map[string]string{
									"controller": "visitor-paths",
								})).
								Text("Visitor Paths")))))

	return hb.Div().
		Class("container").
		Child(title).
		Child(breadcrumbs).
		Child(options).
		Child(c.cardStatsSummary())
}

// == PRIVATE METHODS ==========================================================

func (c *homeController) cardStatsSummary() hb.TagInterface {
	dates, uniqueVisits, totalVisits, err := c.visitorsData()

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
				HTML("Stats Summary"))).
		Child(hb.Div().
			Class("card-body").
			Child(c.chartStatsSummary(dates, uniqueVisits, totalVisits)).
			Child(hb.BR()).
			Child(c.tableStatsSummary(dates, uniqueVisits, totalVisits)).
			Child(hb.BR())).
		Child(hb.Div().
			Class("card-footer"))
}

func (c *homeController) chartStatsSummary(dates []string, uniqueVisits []int64, totalVisits []int64) hb.TagInterface {
	labels := dates
	uniqueVisitValues := uniqueVisits
	totalVisitValues := totalVisits

	labelsJSON, err := utils.ToJSON(labels)

	if err != nil {
		return hb.Div().Class("alert alert-danger").Text(err.Error())
	}

	uniqueVisitvaluesJSON, err := utils.ToJSON(uniqueVisitValues)

	if err != nil {
		return hb.Div().Class("alert alert-danger").Text(err.Error())
	}

	totalVisitValuesJSON, err := utils.ToJSON(totalVisitValues)

	if err != nil {
		return hb.Div().Class("alert alert-danger").Text(err.Error())
	}

	script := hb.Script(`
			setTimeout(function () {
				generateVisitorsChart();
			}, 1000);
			function generateVisitorsChart() {
				var visitorData = {
					labels: ` + labelsJSON + `,
					datasets:
							[
								{
									label: "Unique Visitors",
									fillColor: "rgba(172,194,132,0.4)",
									strokeColor: "#ACC26D",
									pointColor: "#fff",
									pointStrokeColor: "#9DB86D",
									data: ` + uniqueVisitvaluesJSON + `
								},
								{
									label: "Total Visitors",
									fillColor: "rgba(91,192,222,0.4)",
									strokeColor: "#5BC0DE",
									pointColor: "#fff",
									pointStrokeColor: "#39B7CD",
									data: ` + totalVisitValuesJSON + `
								}
							]
				};

				var visitorContext = document.getElementById('StatsSummary').getContext('2d');
				
				new Chart(visitorContext, {
					type: 'bar',
					data: visitorData
				});
			}
		`)

	canvas := hb.Canvas().ID("StatsSummary").Style("width:100%;height:300px;")
	return hb.Wrap().
		Child(canvas).
		Child(script)
}

func (c *homeController) tableStatsSummary(dates []string, uniqueVisits []int64, totalVisits []int64) hb.TagInterface {
	avgUniqueVisits := float64(lo.Sum(uniqueVisits)) / float64(len(dates))
	avgTotalVisits := float64(lo.Sum(totalVisits)) / float64(len(dates))

	cardAvgUniqueVisits := hb.Heading3().Text(fmt.Sprintf("Average Unique Visitors: %.2f", avgUniqueVisits))
	cardAvgTotalVisits := hb.Heading3().Text(fmt.Sprintf("Average Total Visitors: %.2f", avgTotalVisits))

	table := hb.Table().
		Class("table table-striped table-bordered").
		Children([]hb.TagInterface{
			hb.Thead().Children([]hb.TagInterface{
				hb.TR().Children([]hb.TagInterface{
					hb.TH().Text("Date"),
					hb.TH().Text("Unique Visitors"),
					hb.TH().Text("Total Visitors"),
				}),
			}),
			hb.Tbody().Children(lo.Map(dates, func(date string, index int) hb.TagInterface {
				return hb.TR().Children([]hb.TagInterface{
					hb.TD().Text(date),
					hb.TD().Text(cast.ToString(uniqueVisits[index])),
					hb.TD().Text(cast.ToString(totalVisits[index])),
				})
			})),
			hb.Tfoot().Children([]hb.TagInterface{
				hb.TR().Children([]hb.TagInterface{
					hb.TH().Text("Total"),
					hb.TH().Text(cast.ToString(lo.Sum(uniqueVisits))),
					hb.TH().Text(cast.ToString(lo.Sum(totalVisits))),
				}),
			}),
		})

	return hb.Wrap().
		Child(bs.Row().
			Class("g-4").
			Child(bs.Column(6).
				Child(cardAvgUniqueVisits)).
			Child(bs.Column(6).
				Child(cardAvgTotalVisits))).
		Child(table)
}

func (c *homeController) datesInRange(timeStart, timeEnd carbon.Carbon) []string {
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

func (c *homeController) visitorsData() (dates []string, uniqueVisits []int64, totalVisits []int64, err error) {
	datesInRange := c.datesInRange(carbon.Now().SubDays(31), carbon.Now())

	for _, date := range datesInRange {
		uniqueVisitorCount, err := config.StatsStore.VisitorCount(statsstore.VisitorQueryOptions{
			CreatedAtGte: date + " 00:00:00",
			CreatedAtLte: date + " 23:59:59",
			Distinct:     statsstore.COLUMN_IP_ADDRESS,
		})

		if err != nil {
			return nil, nil, nil, err
		}

		totalVisitorCount, err := config.StatsStore.VisitorCount(statsstore.VisitorQueryOptions{
			CreatedAtGte: date + " 00:00:00",
			CreatedAtLte: date + " 23:59:59",
		})

		if err != nil {
			return nil, nil, nil, err
		}

		dates = append(dates, date)
		uniqueVisits = append(uniqueVisits, uniqueVisitorCount)
		totalVisits = append(totalVisits, totalVisitorCount)
	}

	return dates, uniqueVisits, totalVisits, nil
}
