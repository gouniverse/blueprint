package admin

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

type layoutOptions struct {
	Title      string
	Content    string
	ScriptURLs []string
	Scripts    []string
	StyleURLs  []string
	Styles     []string
}

func layout(r *http.Request, opts layoutOptions) *dashboard.Dashboard {
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
	scriptURLs := []string{} // prepend any if required
	scriptURLs = append(scriptURLs, opts.ScriptURLs...)
	scriptURLs = append(scriptURLs, cdn.BootstrapJs_5_3_2())
	scriptURLs = append(scriptURLs, cdn.Htmx_1_9_6())

	// Prepare scripts
	scripts := []string{} // prepend any if required
	scripts = append(scripts, opts.Scripts...)

	homeLink := links.NewWebsiteLinks().Home()
	logoutLink := links.NewAuthLinks().Logout()

	dashboard := dashboard.NewDashboard(dashboard.Config{
		HTTPRequest: r,
		Content:     opts.Content,
		Title:       opts.Title + " | Admin | " + config.AppName,
		LoginURL:    links.NewAuthLinks().Login(homeLink),
		Menu: []dashboard.MenuItem{
			{
				Icon:  hb.NewI().Class("bi bi-house").Style("margin-right:10px;").ToHTML(),
				Title: "Home",
				URL:   homeLink,
			},
			{
				Icon:  hb.NewI().Class("bi bi-arrow-right").Style("margin-right:10px;").ToHTML(),
				Title: "Logout",
				URL:   logoutLink,
			},
		},
		// Theme: dashboard.THEME_MINTY,
		User:            dashboardUser,
		UserMenu:        userDashboardUserMenu(authUser),
		ThemeHandlerUrl: links.NewWebsiteLinks().Theme(map[string]string{"redirect": r.URL.Path}),
		Scripts:         scripts,
		ScriptURLs:      scriptURLs,
		Styles:          opts.Styles,
		StyleURLs:       opts.StyleURLs,
		FaviconURL:      links.URL("favicon.svg", map[string]string{}),
	})

	return dashboard
}

func userDashboardUserMenu(authUser *userstore.User) []dashboard.MenuItem {
	items := []dashboard.MenuItem{
		// {
		// 	Title: "Profile",
		// 	URL:   links.NewUserLinks().Profile(),
		// },
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
