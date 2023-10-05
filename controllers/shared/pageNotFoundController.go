package shared

import "net/http"

type pageNotFoundController struct{}

func NewPageNotFoundControllerController() *pageNotFoundController {
	return &pageNotFoundController{}
}

func (controller *pageNotFoundController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	return "Sorry, page not found."
}
