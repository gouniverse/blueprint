package layouts

import (
	"project/internal/links"
	"project/pkg/userstore"

	"github.com/gouniverse/dashboard"
)

// userDashboardUserMenu generates the user menu items for the dashboard.
//
// Parameters:
// - `authUser` (*models.User): The authenticated user.
//
// Returns:
// - `[]dashboard.MenuItem`: The user menu items.
func adminLayoutUserMenu(authUser *userstore.User) []dashboard.MenuItem {
	userDashboardMenuItem := dashboard.MenuItem{
		Title: "To User Panel",
		URL:   links.NewUserLinks().Home(),
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
