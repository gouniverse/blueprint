package shared

import "project/internal/links"

type Links struct{}

func NewLinks() *Links {
	return &Links{}
}

func (*Links) Home(params map[string]string) string {
	params["controller"] = "home"
	return links.NewAdminLinks().Shop(params)
}

func (*Links) Discounts(params map[string]string) string {
	params["controller"] = "discounts"
	return links.NewAdminLinks().Shop(params)
}

func (*Links) Orders(params map[string]string) string {
	params["controller"] = "orders"
	return links.NewAdminLinks().Shop(params)
}

func (*Links) ProductCreate(params map[string]string) string {
	params["controller"] = "product_create"
	return links.NewAdminLinks().Shop(params)
}

func (*Links) ProductDelete(params map[string]string) string {
	params["controller"] = "product_delete"
	return links.NewAdminLinks().Shop(params)
}

func (*Links) Products(params map[string]string) string {
	params["controller"] = "products"
	return links.NewAdminLinks().Shop(params)
}

func (*Links) ProductUpdate(params map[string]string) string {
	params["controller"] = "product_update"
	return links.NewAdminLinks().Shop(params)
}
