package shared

import (
	"log/slog"
	"net/http"
	"project/app/links"
	"project/config"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/shopstore"
	"github.com/spf13/cast"
)

func Header(store shopstore.StoreInterface, logger *slog.Logger, r *http.Request) hb.TagInterface {
	if store == nil {
		logger.Error("shop store is nil")
		return nil
	}

	if config.ShopStore == nil {
		logger.Error("shop store is nil")
		return nil
	}

	linkHome := hb.NewHyperlink().
		HTML("Dashboard").
		Href(links.NewAdminLinks().Home(map[string]string{})).
		Class("nav-link")

	linkShop := hb.NewHyperlink().
		HTML("Shop").
		Href(links.NewAdminLinks().Shop(map[string]string{})).
		Class("nav-link")

	linkOrders := hb.Hyperlink().
		HTML("Orders").
		Href(links.NewAdminLinks().Shop(map[string]string{
			"controller": "orders",
		})).
		Class("nav-link")

	linkDiscounts := hb.Hyperlink().
		HTML("Discounts").
		Href(links.NewAdminLinks().Shop(map[string]string{
			"controller": "discounts",
		})).
		Class("nav-link")

	linkProducts := hb.Hyperlink().
		HTML("Products ").
		Href(links.NewAdminLinks().Shop(map[string]string{
			"controller": "products",
		})).
		Class("nav-link")

	productsCount, err := config.ShopStore.ProductCount(r.Context(), shopstore.NewProductQuery())

	if err != nil {
		logger.Error(err.Error())
		productsCount = -1
	}

	ordersCount, err := config.ShopStore.OrderCount(r.Context(), shopstore.NewOrderQuery())

	if err != nil {
		logger.Error(err.Error())
		ordersCount = -1
	}

	discountsCount, err := config.ShopStore.DiscountCount(r.Context(), shopstore.NewDiscountQuery())

	if err != nil {
		logger.Error(err.Error())
		discountsCount = -1
	}

	ulNav := hb.NewUL().
		Class("nav  nav-pills justify-content-center").
		Child(hb.NewLI().
			Class("nav-item").Child(linkHome)).
		Child(hb.NewLI().
			Class("nav-item").Child(linkShop)).
		Child(hb.LI().
			Class("nav-item").
			Child(linkOrders.
				Child(hb.Span().
					Class("badge bg-secondary ms-2").
					HTML(cast.ToString(ordersCount))))).
		Child(hb.LI().
			Child(linkProducts.
				Child(hb.Span().
					Class("badge bg-secondary ms-2").
					HTML(cast.ToString(productsCount))))).
		Child(hb.LI().
			Child(linkDiscounts.
				Child(hb.Span().
					Class("badge bg-secondary ms-2").
					HTML(cast.ToString(discountsCount)))))

	divCard := hb.NewDiv().Class("card card-default mt-3 mb-3")
	divCardBody := hb.NewDiv().Class("card-body").Style("padding: 2px;")
	return divCard.AddChild(divCardBody.AddChild(ulNav))
}
