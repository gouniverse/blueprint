package layouts

import (
	"project/internal/links"
	"project/pkg/userstore"

	"github.com/gouniverse/dashboard"
)

// userLayoutUserMenu generates the user menu items for the dashboard.
//
// Parameters:
// - `authUser` (*models.User): The authenticated user.
//
// Returns:
// - `[]dashboard.MenuItem`: The user menu items.
func userLayoutUserMenu(authUser *userstore.User) []dashboard.MenuItem {
	adminDashboardMenuItem := dashboard.MenuItem{
		Title: "To Admin Dashboard",
		URL:   links.NewAdminLinks().Home(map[string]string{}),
	}

	logoutMenuItem := dashboard.MenuItem{
		Title: "Logout",
		URL:   links.NewAuthLinks().Logout(),
	}

	profileMenuItem := dashboard.MenuItem{
		Title: "Profile",
		URL:   links.NewUserLinks().Profile(map[string]string{}),
	}

	items := []dashboard.MenuItem{profileMenuItem}

	if authUser != nil && authUser.IsAdministrator() {
		items = append(items, adminDashboardMenuItem)
	}

	items = append(items, logoutMenuItem)

	return items
}
