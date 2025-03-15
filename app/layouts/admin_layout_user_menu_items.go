package layouts

import (
	"project/app/links"

	"github.com/gouniverse/dashboard"
	"github.com/gouniverse/userstore"
)

// userDashboardUserMenu generates the user menu items for the dashboard.
//
// Parameters:
// - `authUser` (*models.User): The authenticated user.
//
// Returns:
// - `[]dashboard.MenuItem`: The user menu items.
func adminLayoutUserMenu(authUser userstore.UserInterface) []dashboard.MenuItem {
	userDashboardMenuItem := dashboard.MenuItem{
		Title: "To User Panel",
		URL:   links.NewUserLinks().Home(map[string]string{}),
	}

	logoutMenuItem := dashboard.MenuItem{
		Title: "Logout",
		URL:   links.NewAuthLinks().Logout(),
	}

	items := []dashboard.MenuItem{}

	if authUser != nil && authUser.IsAdministrator() {
		items = append(items, userDashboardMenuItem)
	}

	items = append(items, logoutMenuItem)

	return items
}
