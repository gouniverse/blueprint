package admin

import (
	"net/http"
	"project/internal/links"

	"github.com/gouniverse/router"
	"github.com/gouniverse/utils"

	shopDiscounts "project/app/controllers/admin/shop/discounts"
	shopProducts "project/app/controllers/admin/shop/products"
)

func ShopRoutes() []router.RouteInterface {
	handler := func(w http.ResponseWriter, r *http.Request) string {
		controller := utils.Req(r, "controller", "")

		if controller == "discounts" {
			return shopDiscounts.NewDiscountController().AnyIndex(w, r)
		}

		if controller == "product_create" {
			return shopProducts.NewProductCreateController().Handler(w, r)
		}

		if controller == "product_delete" {
			return shopProducts.NewProductDeleteController().Handler(w, r)
		}

		if controller == "products" {
			return shopProducts.NewProductManagerController().Handler(w, r)
		}

		if controller == "product_update" {
			return shopProducts.NewProductUpdateController().Handler(w, r)
		}

		if controller == "orders" {
			return NewOrderManagerController().Handler(w, r)
		}

		return NewHomeController().Handler(w, r)
	}

	return []router.RouteInterface{
		&router.Route{
			Name:        "Admin > Shop > Orders",
			Path:        links.ADMIN_SHOP,
			HTMLHandler: handler,
		},
		&router.Route{
			Name:        "Admin > Shop > Catchall",
			Path:        links.ADMIN_USERS + links.CATCHALL,
			HTMLHandler: handler,
		},
	}
}
