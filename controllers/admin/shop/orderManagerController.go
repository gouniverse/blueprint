package admin

import (
	"net/http"

	"github.com/gouniverse/router"
)

type orderManagerController struct{}

func NewOrderManagerController() *orderManagerController {
	return &orderManagerController{}
}

var _ router.HTMLControllerInterface = (*orderManagerController)(nil)

func (orderManagerController *orderManagerController) Handler(w http.ResponseWriter, r *http.Request) string {
	return "Order Manager"
}
