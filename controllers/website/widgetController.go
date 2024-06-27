package website

import (
	"net/http"
	"project/internal/widgets"

	"github.com/gouniverse/utils"
)

// == CONTROLLER ===============================================================

type widgetController struct{}

// == CONSTRUCTOR ==============================================================

func NewWidgetController() *widgetController {
	return &widgetController{}
}

// == PUBLIC METHODS ==========================================================

func (controller *widgetController) Handler(w http.ResponseWriter, r *http.Request) string {
	alias := utils.Req(r, "alias", "")

	if alias == "" {
		return "Widget type not specified"
	}

	widgetList := []widgets.Widget{
		// widgets.NewAuthenticatedWidget(),
		// widgets.NewPrintWidget(),
		// widgets.NewVisibleWidget(),
		// widgets.NewUnauthenticatedWidget(),
	}

	for _, widget := range widgetList {
		if widget.Alias() == alias {
			return widget.Render(r, "", nil)
		}
	}

	return alias
}
