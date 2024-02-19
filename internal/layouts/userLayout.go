package layouts

import (
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/links"
	"project/pkg/userstore"

	"github.com/gouniverse/cdn"
	"github.com/gouniverse/dashboard"
	"github.com/gouniverse/hb"
	"github.com/samber/lo"
)

func NewUserLayout(r *http.Request, options Options) *dashboard.Dashboard {
	return userLayout(r, options)
}

// layout generates a dashboard based on the provided request and layout options.
//
// Parameters:
// - r: a pointer to an http.Request object representing the incoming HTTP request.
// - opts: a layoutOptions struct containing the layout options for the dashboard.
//
// Returns:
// - a pointer to a dashboard.Dashboard object representing the generated dashboard.
func userLayout(r *http.Request, options Options) *dashboard.Dashboard {
	authUser := helpers.GetAuthUser(r)

	dashboardUser := dashboard.User{}
	if authUser != nil {
		firstName := lo.If(authUser.FirstName() == "", authUser.Email()).Else(authUser.FirstName())
		dashboardUser = dashboard.User{
			FirstName: firstName,
			LastName:  authUser.LastName(),
		}
	}

	// Prepare script URLs
	scriptURLs := []string{}
	scriptURLs = append(scriptURLs, options.ScriptURLs...)
	scriptURLs = append(scriptURLs, cdn.Htmx_1_9_9())

	// Prepare scripts
	scripts := []string{}
	scripts = append(scripts, options.Scripts...)

	homeLink := links.NewUserLinks().Home()

	dashboard := dashboard.NewDashboard(dashboard.Config{
		HTTPRequest:     r,
		Content:         options.Content.ToHTML(),
		Title:           options.Title + " | User | " + config.AppName,
		LoginURL:        links.NewAuthLinks().Login(homeLink),
		Menu:            userMenu(authUser),
		LogoImageURL:    "/media/user/dashboard-logo.jpg",
		LogoRedirectURL: links.NewUserLinks().Home(),
		User:            dashboardUser,
		UserMenu:        userDashboardUserMenu(authUser),
		ThemeHandlerUrl: links.NewWebsiteLinks().Theme(map[string]string{"redirect": r.URL.Path}),
		Scripts:         scripts,
		ScriptURLs:      scriptURLs,
		Styles:          options.Styles,
		StyleURLs:       options.StyleURLs,
		FaviconURL:      links.URL("favicon.svg", map[string]string{}),
		// Theme: dashboard.THEME_MINTY,
	})

	return dashboard
}

func userMenu(user *userstore.User) []dashboard.MenuItem {
	websiteHomeLink := links.NewWebsiteLinks().Home()
	dashboardLink := links.NewUserLinks().Home()
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

	// shopMenuItem := dashboard.MenuItem{
	// 	Icon:  hb.NewI().Class("bi bi-shop").Style("margin-right:10px;").ToHTML(),
	// 	Title: "Your Shop",
	// 	URL:   links.NewUserLinks().Shop(map[string]string{}),
	// }

	// pendingAssessmentsMenuItem := dashboard.MenuItem{
	// 	Icon:  hb.NewI().Class("bi bi-clock").Style("margin-right:10px;").ToHTML(),
	// 	Title: "Ordered Exams",
	// 	URL:   links.NewUserLinks().ExamsOrdered(),
	// }

	// failedAssessmentsMenuItem := dashboard.MenuItem{
	// 	Icon:  hb.NewI().Class("bi bi-patch-exclamation").Style("margin-right:10px;").ToHTML(),
	// 	Title: "Failed Exams",
	// 	URL:   links.NewUserLinks().ExamsFailed(),
	// }

	// passedAssessmentsMenuItem := dashboard.MenuItem{
	// 	Icon:  hb.NewI().Class("bi bi-patch-check").Style("margin-right:10px;").ToHTML(),
	// 	Title: "Passed Exams",
	// 	URL:   links.NewUserLinks().ExamsPassed(),
	// }

	// inviteFriendMenuItem := dashboard.MenuItem{
	// 	Icon:  hb.NewI().Class("bi bi-people").Style("margin-right:10px;").ToHTML(),
	// 	Title: "Invite a Friend",
	// 	URL:   links.NewUserLinks().InviteFriend(),
	// }

	websiteMenuItem := dashboard.MenuItem{
		Icon:   hb.NewI().Class("bi bi-globe").Style("margin-right:10px;").ToHTML(),
		Title:  "To Website",
		URL:    websiteHomeLink,
		Target: "_blank",
	}

	if user != nil {
		homeMenuItem = dashboard.MenuItem{
			Icon:  hb.NewI().Class("bi bi-speedometer").Style("margin-right:10px;").ToHTML(),
			Title: "Dashboard",
			URL:   dashboardLink,
		}

		loginMenuItem = dashboard.MenuItem{
			Icon:  hb.NewI().Class("bi bi-arrow-right").Style("margin-right:10px;").ToHTML(),
			Title: "Logout",
			URL:   logoutLink,
		}
	}

	menuItems := []dashboard.MenuItem{
		homeMenuItem,
	}

	if user != nil {
		// menuItems = append(menuItems, shopMenuItem)
		// menuItems = append(menuItems, pendingAssessmentsMenuItem)
		// menuItems = append(menuItems, passedAssessmentsMenuItem)
		// menuItems = append(menuItems, failedAssessmentsMenuItem)
		// menuItems = append(menuItems, inviteFriendMenuItem)
		menuItems = append(menuItems, websiteMenuItem)
	}

	menuItems = append(menuItems, loginMenuItem)

	return menuItems
}

// userDashboardUserMenu generates the user menu items for the dashboard.
//
// Parameters:
// - `authUser` (*models.User): The authenticated user.
//
// Returns:
// - `[]dashboard.MenuItem`: The user menu items.
func userDashboardUserMenu(authUser *userstore.User) []dashboard.MenuItem {
	items := []dashboard.MenuItem{
		{
			Title: "Profile",
			// URL:   links.NewUserLinks().Profile(),
		},
	}

	if authUser != nil && authUser.IsAdministrator() {
		items = append(items, dashboard.MenuItem{
			Title: "To Admin Panel",
			URL:   links.NewAdminLinks().Home(),
		})
	}

	items = append(items, dashboard.MenuItem{
		Title: "Logout",
		URL:   links.NewAuthLinks().Logout(),
	})

	return items
}
