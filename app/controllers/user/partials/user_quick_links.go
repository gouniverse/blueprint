package partials

import (
	"net/http"
	"project/app/links"
	"strings"

	"github.com/gouniverse/hb"
	"github.com/samber/lo"
)

func UserQuickLinks(req *http.Request) *hb.Tag {
	errorMessage := ""

	if errorMessage != "" {
		return hb.Section().Text(errorMessage)
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
			Title:              "My Dashboard",
			Background:         "rgba(141, 210, 220, 0.5)",
			BackgroundSelected: "rgba(141, 210, 220, 1)",
			Count:              hb.I().Class("bi bi-file-earmark-text-fill").ToHTML(),
			URL:                links.NewUserLinks().Home(map[string]string{}),
			IsSelected:         lo.If(req.URL.Path == links.USER_HOME, true).Else(false),
		},
		{
			Title:              "My Account",
			Background:         "rgba(172, 220, 141, 0.5)",
			BackgroundSelected: "rgba(172, 220, 141, 1)",
			Count:              hb.I().Class("bi bi-person-circle").ToHTML(),
			URL:                links.NewUserLinks().Profile(map[string]string{}),
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
		heading := hb.Heading5().
			Class("fs-3").
			Style("text-decoration: none; color: #333;").
			StyleIf(cardData.IsSelected, `color: #fff;`).
			Child(hb.Div().
				HTML(cardData.Count).
				Class("fs-1 ms-1 me-3")).
			Child(hb.Div().
				Text(cardData.Title).
				Class("fs-4 ms-1 me-3"))

		card := hb.Div().
			Class("card").
			StyleIfElse(cardData.IsSelected, `background: `+cardData.BackgroundSelected, `background: `+cardData.Background).
			Style(`cursor: pointer;`).
			StyleIf(cardData.IsSelected, `box-shadow: 0 1rem 1rem rgba(0, 0, 0, 0.3)`).
			// Child(
			// 	hb.Div().
			// 		Class("card-header").
			// 		Child(
			// 			hb.Span().
			// 				Class("float-end fs-6").
			// 				Style("text-decoration: underline;color: #333;").
			// 				Text("see all"),
			// 		)).
			Child(
				hb.Div().
					Class("card-body").
					Child(heading),
			)

		link := hb.Hyperlink().
			Href(cardData.URL).
			Child(card).
			Style("text-decoration: none;")

		return hb.Div().
			Class("col-lg-4 col-md-6 col-sm-12 mb-3").
			Child(link)
	})

	return hb.Section().
		ID("SectionQuickLinks").
		Child(hb.Div().
			Class("container").
			Child(
				hb.Div().
					Class("row").
					Children(columns)))
}
