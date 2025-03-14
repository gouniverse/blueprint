package admin

import (
	"net/http"
	"project/app/controllers/admin/shop/shared"
	"project/config"
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

// == CONSTRUCTOR ==============================================================

func NewHomeController() *homeController {
	return &homeController{}
}

// == INTERFACES ===============================================================

var _ router.HTMLControllerInterface = (*homeController)(nil)

// == PUBLIC METHODS ===========================================================

func (controller *homeController) Handler(w http.ResponseWriter, r *http.Request) string {
	return layouts.NewAdminLayout(r, layouts.Options{
		Title:      "Shop",
		Content:    controller.view(r),
		ScriptURLs: []string{},
		Styles:     []string{},
	}).ToHTML()
}

// == PRIVATE METHODS ==========================================================

func (controller *homeController) view(r *http.Request) *hb.Tag {
	breadcrumbs := layouts.Breadcrumbs([]layouts.Breadcrumb{
		{
			Name: "Home",
			URL:  links.NewAdminLinks().Home(map[string]string{}),
		},
		{
			Name: "Shop",
			URL:  shared.NewLinks().Home(map[string]string{}),
		},
	})

	header := hb.Heading1().
		HTML("Shop").
		Style("margin-bottom:30px;margin-top:30px;")

	sectionTiles := hb.Section().
		Children([]hb.TagInterface{
			bs.Row().
				Class("g-4").
				Children(controller.tiles()),
		})

	return hb.Wrap().
		Child(breadcrumbs).
		Child(hb.HR()).
		Child(shared.Header(config.ShopStore, &config.Logger, r)).
		Child(hb.HR()).
		Child(header).
		Child(sectionTiles)
}

func (*homeController) tiles() []hb.TagInterface {
	tiles := []map[string]string{
		{
			"title": "Order Manager",
			"icon":  "bi-cart",
			"link":  shared.NewLinks().Orders(map[string]string{}),
		},
		{
			"title": "Product Manager",
			"icon":  "bi-box",
			"link":  shared.NewLinks().Products(map[string]string{}),
		},
		{
			"title": "Discount Manager",
			"icon":  "bi-percent",
			"link":  shared.NewLinks().Discounts(map[string]string{}),
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
