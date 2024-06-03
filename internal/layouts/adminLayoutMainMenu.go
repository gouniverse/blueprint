package layouts

import (
	"project/internal/links"
	"project/pkg/userstore"

	"github.com/gouniverse/dashboard"
	"github.com/gouniverse/hb"
)

func adminLayoutMainMenu(user *userstore.User) []dashboard.MenuItem {
	websiteHomeLink := links.NewWebsiteLinks().Home()
	dashboardLink := links.NewAdminLinks().Home()
	loginLink := links.NewAuthLinks().Login(dashboardLink)
	logoutLink := links.NewAuthLinks().Logout()

	homeMenuItem := dashboard.MenuItem{
		Icon:  hb.NewI().Class("bi bi-house").Style("margin-right:10px;").ToHTML(),
		Title: "Home",
		URL:   websiteHomeLink,
	}

	loginMenuItem := dashboard.MenuItem{
		Icon:  hb.NewI().Class("bi bi-arrow-right").Style("margin-right:10px;").ToHTML(),
		Title: "Login",
		URL:   loginLink,
	}

	websiteMenuItem := dashboard.MenuItem{
		Icon:   hb.NewI().Class("bi bi-globe").Style("margin-right:10px;").ToHTML(),
		Title:  "To Website",
		URL:    websiteHomeLink,
		Target: "_blank",
	}

	logoutMenuItem := dashboard.MenuItem{
		Icon:  hb.NewI().Class("bi bi-arrow-right").Style("margin-right:10px;").ToHTML(),
		Title: "Logout",
		URL:   logoutLink,
	}

	dashboardMenuItem := dashboard.MenuItem{
		Icon:  hb.NewI().Class("bi bi-speedometer").Style("margin-right:10px;").ToHTML(),
		Title: "Dashboard",
		URL:   dashboardLink,
	}

	menuItems := []dashboard.MenuItem{}

	if user != nil {
		menuItems = append(menuItems, dashboardMenuItem)
		menuItems = append(menuItems, websiteMenuItem)
		menuItems = append(menuItems, logoutMenuItem)
	} else {
		menuItems = append(menuItems, homeMenuItem)
		menuItems = append(menuItems, loginMenuItem)
	}

	return menuItems
}
