package layouts

import (
	"net/http"
	"project/app/links"
	"project/config"
	"project/internal/helpers"

	"github.com/gouniverse/cdn"
	"github.com/gouniverse/dashboard"
	"github.com/samber/lo"
)

func NewAdminLayout(r *http.Request, options Options) *dashboard.Dashboard {
	return adminLayout(r, options)
}

// layout generates a dashboard based on the provided request and layout options.
//
// Parameters:
// - r: a pointer to an http.Request object representing the incoming HTTP request.
// - opts: a layoutOptions struct containing the layout options for the dashboard.
//
// Returns:
// - a pointer to a dashboard.Dashboard object representing the generated dashboard.
func adminLayout(r *http.Request, options Options) *dashboard.Dashboard {
	authUser := helpers.GetAuthUser(r)

	dashboardUser := dashboard.User{}
	if authUser != nil {
		firstName, lastName, err := getUserData(r, authUser)
		if err == nil {
			dashboardUser = dashboard.User{
				FirstName: firstName,
				LastName:  lastName,
			}
		}
	}

	// Prepare script URLs
	scriptURLs := []string{} // prepend any if required
	scriptURLs = append(scriptURLs, options.ScriptURLs...)
	scriptURLs = append(scriptURLs, cdn.Htmx_2_0_0())

	// Prepare scripts
	scripts := []string{} // prepend any if required
	scripts = append(scripts, options.Scripts...)

	// Prepare styles
	styles := []string{ // prepend any if required
		`nav#Toolbar {border-bottom: 8px solid red;}`,
	}
	styles = append(styles, options.Styles...)

	homeLink := links.NewAdminLinks().Home(map[string]string{})

	path := lo.IfF(r != nil, func() string {
		return r.URL.Path
	}).ElseF(func() string {
		return ""
	})
	themeLink := links.NewWebsiteLinks().Theme(map[string]string{"redirect": path})

	dashboard := dashboard.NewDashboard(dashboard.Config{
		HTTPRequest:     r,
		Content:         options.Content.ToHTML(),
		Title:           options.Title + " | Admin | " + config.AppName,
		LoginURL:        links.NewAuthLinks().Login(homeLink),
		MenuItems:       adminLayoutMainMenu(authUser),
		LogoImageURL:    "/media/user/dashboard-logo.jpg",
		LogoRawHtml:     adminLogoHtml(),
		LogoRedirectURL: homeLink,
		User:            dashboardUser,
		UserMenu:        adminLayoutUserMenu(authUser),
		ThemeHandlerUrl: themeLink,
		Scripts:         scripts,
		ScriptURLs:      scriptURLs,
		Styles:          styles,
		StyleURLs:       options.StyleURLs,
		FaviconURL:      FaviconURL(),
		// Theme: dashboard.THEME_MINTY,
	})

	return dashboard
}
