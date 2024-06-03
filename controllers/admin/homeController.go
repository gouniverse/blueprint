package admin

import (
	"net/http"
	"project/internal/layouts"
	"project/internal/links"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/icons"
	"github.com/samber/lo"
)

type homeController struct{}

func NewHomeController() *homeController {
	return &homeController{}
}

func (controller *homeController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	return layouts.NewAdminLayout(r, layouts.Options{
		Title:      "Home",
		Content:    controller.view(),
		ScriptURLs: []string{},
		Styles: []string{
			`nav#Toolbar {border-bottom: 4px solid red;}`,
		},
	}).ToHTML()
}

func (c *homeController) view() *hb.Tag {
	header := hb.NewHeading1().
		HTML("Admin Home").
		Style("margin-bottom:30px;margin-top:30px;")

	sectionTiles := hb.NewSection().Children([]hb.TagInterface{
		bs.Row().Class("g-4").Children(c.tiles()),
	})

	return hb.NewWrap().Child(header).Child(sectionTiles)
}

func (c *homeController) quickSearch(r *http.Request) []*hb.Tag {
	heading := hb.NewHeading2().HTML("Quick Search")

	return []*hb.Tag{
		hb.NewBR(),
		heading,
	}
}

func (*homeController) tiles() []hb.TagInterface {
	tiles := []map[string]string{
		{
			"title": "Website Manager",
			"icon":  "bi-globe",
			"link":  links.NewAdminLinks().Cms(),
		},
		{
			"title": "User Manager",
			"icon":  "bi-people",
			"link":  links.NewAdminLinks().Users(),
		},
		{
			"title": "Media Manager",
			"icon":  "bi-box",
			"link":  links.NewAdminLinks().FileManager(),
		},
	}

	cards := lo.Map(tiles, func(tile map[string]string, index int) hb.TagInterface {
		target := lo.ValueOr(tile, "target", "")
		card := bs.Card().
			Class("bg-transparent border round-10 shadow-lg h-100 pt-4").
			OnMouseOver(`
			this.style.setProperty('background-color', 'beige', 'important');
			this.style.setProperty('scale', 1.1);
			this.style.setProperty('border', '4px solid moccasin', 'important');
			`).
			OnMouseOut(`
			this.style.setProperty('background-color', 'transparent', 'important');
			this.style.setProperty('scale', 1);
			this.style.setProperty('border', '0px solid moccasin', 'important');
			`).
			Style("margin:0px 0px 20px 0px;").
			Children([]hb.TagInterface{
				bs.CardBody().
					Class("d-flex flex-column justify-content-evenly").
					Children([]hb.TagInterface{
						hb.NewDiv().
							Child(icons.Icon(tile["icon"], 36, 36, "red")).
							Style("text-align:center;padding:10px;"),
						hb.NewHeading5().
							HTML(tile["title"]).
							Style("text-align:center;padding:10px;"),
					}),
			})

		link := hb.NewHyperlink().
			Href(tile["link"]).
			AttrIf(target != "", "target", target).
			Child(card)

		column := bs.Column(3).
			Class("col-sm-6 col-md-4 col-lg-3").
			Child(link)

		return column
	})
	return cards
}
