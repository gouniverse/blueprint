package layouts

import (
	"project/internal/links"

	"github.com/gouniverse/dashboard"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/userstore"
)

// userLayoutMainMenu generates the main menu items for the user dashboard.
//
// Parameters:
// - `user` (*models.User): The authenticated user.
//
// Returns:
// - `[]dashboard.MenuItem`: The main menu items.
func userLayoutMainMenuItems(user userstore.UserInterface) []dashboard.MenuItem {
	websiteHomeLink := links.NewWebsiteLinks().Home()
	dashboardLink := links.NewUserLinks().Home(map[string]string{})
	loginLink := links.NewAuthLinks().Login(dashboardLink)
	logoutLink := links.NewAuthLinks().Logout()

	homeMenuItem := dashboard.MenuItem{
		Icon:  hb.I().Class("bi bi-house").Style("margin-right:10px;").ToHTML(),
		Title: "Home",
		URL:   websiteHomeLink,
	}

	profileMenuItem := dashboard.MenuItem{
		Icon:  hb.I().Class("bi bi-person").Style("margin-right:10px;").ToHTML(),
		Title: "My Account",
		URL:   links.NewUserLinks().Profile(map[string]string{}),
	}

	loginMenuItem := dashboard.MenuItem{
		Icon:  hb.I().Class("bi bi-door-open").Style("margin-right:10px;").ToHTML(),
		Title: "Login",
		URL:   loginLink,
	}

	// shopMenuItem := dashboard.MenuItem{
	// 	Icon:  hb.I().Class("bi bi-shop").Style("margin-right:10px;").ToHTML(),
	// 	Title: "Your Shop",
	// 	URL:   links.NewUserLinks().Shop(map[string]string{}),
	// }

	// inviteFriendMenuItem := dashboard.MenuItem{
	// 	Icon:  hb.I().Class("bi bi-people").Style("margin-right:10px;").ToHTML(),
	// 	Title: "Invite a Friend",
	// 	URL:   links.NewUserLinks().InviteFriend(),
	// }

	websiteMenuItem := dashboard.MenuItem{
		Icon:   hb.I().Class("bi bi-globe").Style("margin-right:10px;").ToHTML(),
		Title:  "To Website",
		URL:    websiteHomeLink,
		Target: "_blank",
	}

	logoutMenuItem := dashboard.MenuItem{
		Icon:  hb.I().Class("bi bi-arrow-right").Style("margin-right:10px;").ToHTML(),
		Title: "Logout",
		URL:   logoutLink,
	}

	dashboardMenuItem := dashboard.MenuItem{
		Icon:  hb.I().Class("bi bi-speedometer").Style("margin-right:10px;").ToHTML(),
		Title: "Dashboard",
		URL:   dashboardLink,
	}

	menuItems := []dashboard.MenuItem{}

	if user != nil {
		menuItems = append(menuItems, dashboardMenuItem)
		// menuItems = append(menuItems, shopMenuItem)
		menuItems = append(menuItems, profileMenuItem)
		// menuItems = append(menuItems, inviteFriendMenuItem)
		menuItems = append(menuItems, websiteMenuItem)
		menuItems = append(menuItems, logoutMenuItem)
	} else {
		menuItems = append(menuItems, homeMenuItem)
		menuItems = append(menuItems, loginMenuItem)
	}

	return menuItems
}
