package partials

import (
	"net/http"
	"project/internal/links"
	"strings"

	"github.com/gouniverse/hb"
	"github.com/samber/lo"
)

func UserQuickLinks(req *http.Request) *hb.Tag {
	errorMessage := ""

	if errorMessage != "" {
		return hb.NewSection().Text(errorMessage)
	}

	cardsData := []struct {
		Title              string
		Background         string
		BackgroundSelected string
		Count              string
		URL                string
		IsSelected         bool
	}{
		{
			Title:              "Dashboard",
			Background:         "#87cefa",
			BackgroundSelected: "cornflowerblue",
			Count:              hb.NewI().Class("bi bi-file-earmark-text-fill").ToHTML(),
			URL:                links.NewUserLinks().Home(),
			IsSelected:         lo.If(req.URL.Path == links.USER_HOME, true).Else(false),
		},
		{
			Title:              "My Account",
			Background:         "#c2ff99",
			BackgroundSelected: "forestgreen",
			Count:              hb.NewI().Class("bi bi-person-circle").ToHTML(),
			URL:                links.NewUserLinks().Profile(),
			IsSelected:         strings.Contains(req.URL.Path, "/profile"),
		},
	}

	columns := lo.Map(cardsData, func(cardData struct {
		Title              string
		Background         string
		BackgroundSelected string
		Count              string
		URL                string
		IsSelected         bool
	}, index int) hb.TagInterface {
		heading := hb.NewHeading5().
			Class("fs-3").
			Style("text-decoration: none; color: #333;").
			StyleIf(cardData.IsSelected, `color: #fff;`).
			Child(hb.NewDiv().
				HTML(cardData.Count).
				Class("fs-1 ms-1 me-3")).
			Child(hb.NewDiv().
				Text(cardData.Title).
				Class("fs-4 ms-1 me-3"))

		card := hb.NewDiv().
			Class("card").
			StyleIfElse(cardData.IsSelected, `background: `+cardData.BackgroundSelected, `background: `+cardData.Background).
			Style(`cursor: pointer;`).
			StyleIf(cardData.IsSelected, `box-shadow: 0 1rem 1rem rgba(0, 0, 0, 0.3)`).
			// Child(
			// 	hb.NewDiv().
			// 		Class("card-header").
			// 		Child(
			// 			hb.NewSpan().
			// 				Class("float-end fs-6").
			// 				Style("text-decoration: underline;color: #333;").
			// 				Text("see all"),
			// 		)).
			Child(
				hb.NewDiv().
					Class("card-body").
					Child(heading),
			)

		link := hb.NewHyperlink().
			Href(cardData.URL).
			Child(card).
			Style("text-decoration: none;")

		return hb.NewDiv().
			Class("col-lg-4 col-md-6 col-sm-12 mb-3").
			Child(link)
	})

	return hb.NewSection().
		ID("SectionQuickLinks").
		Child(hb.NewDiv().
			Class("container").
			Child(
				hb.NewDiv().
					Class("row").
					Children(columns)))
}
