package admin

import (
	"net/http"
	"project/internal/layouts"
	"project/internal/links"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/icons"
	"github.com/gouniverse/router"
	"github.com/samber/lo"
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
		Title:      "Home",
		Content:    controller.view(),
		ScriptURLs: []string{},
		Styles:     []string{},
	}).ToHTML()
}

// == PRIVATE METHODS ==========================================================

func (c *homeController) view() *hb.Tag {
	header := hb.Heading1().
		HTML("Admin Home").
		Style("margin-bottom:30px;margin-top:30px;")

	sectionTiles := hb.Section().Children([]hb.TagInterface{
		bs.Row().Class("g-4").Children(c.tiles()),
	})

	return hb.Wrap().Child(header).Child(sectionTiles)
}

func (c *homeController) quickSearch(_ *http.Request) []*hb.Tag {
	heading := hb.Heading2().HTML("Quick Search")

	return []*hb.Tag{
		hb.BR(),
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
			"title": "Blog Manager",
			"icon":  "bi-newspaper",
			"link":  links.NewAdminLinks().Blog(map[string]string{}),
		},
		{
			"title": "User Manager",
			"icon":  "bi-people",
			"link":  links.NewAdminLinks().Users(map[string]string{}),
		},
		// {
		// 	"title": "FAQ Manager",
		// 	"icon":  "bi-question-circle",
		// 	"link":  links.NewAdminLinks().Faq(map[string]string{}),
		// },
		{
			"title": "File Manager",
			"icon":  "bi-box",
			"link":  links.NewAdminLinks().FileManager(map[string]string{}),
		},
		// {
		// 	"title":  "CDN Manager",
		// 	"icon":   "bi-folder-symlink",
		// 	"link":   "https://gitlab.com/repo/cdn",
		// 	"target": "_blank",
		// },
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
						hb.Div().
							Child(icons.Icon(tile["icon"], 36, 36, "red")).
							Style("text-align:center;padding:10px;"),
						hb.Heading5().
							HTML(tile["title"]).
							Style("text-align:center;padding:10px;"),
					}),
			})

		link := hb.Hyperlink().
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
