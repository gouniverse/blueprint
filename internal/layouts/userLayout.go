package layouts

import (
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/links"

	"github.com/gouniverse/cdn"
	"github.com/gouniverse/dashboard"
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
		firtsName, lastName, err := userUntokenized(*authUser)
		if err == nil {
			dashboardUser = dashboard.User{
				FirstName: firtsName,
				LastName:  lastName,
			}
		}
	}

	// Prepare script URLs
	scriptURLs := []string{} // prepend any if required
	scriptURLs = append(scriptURLs, options.ScriptURLs...)
	scriptURLs = append(scriptURLs, cdn.Htmx_1_9_9())

	// Prepare scripts
	scripts := []string{} // prepend any if required
	scripts = append(scripts, options.Scripts...)

	// Prepare styles
	styles := []string{ // prepend any if required
		`nav#Toolbar {border-bottom: 4px solid blue;}`,
	}
	styles = append(styles, options.Styles...)

	homeLink := links.NewUserLinks().Home(map[string]string{})

	dashboard := dashboard.NewDashboard(dashboard.Config{
		HTTPRequest:     r,
		Content:         options.Content.ToHTML(),
		Title:           options.Title + " | User | " + config.AppName,
		LoginURL:        links.NewAuthLinks().Login(homeLink),
		Menu:            userLayoutMainMenu(authUser),
		LogoImageURL:    "/media/user/dashboard-logo.jpg",
		LogoRawHtml:     userLogoHtml(),
		LogoRedirectURL: homeLink,
		User:            dashboardUser,
		UserMenu:        userLayoutUserMenu(authUser),
		ThemeHandlerUrl: links.NewWebsiteLinks().Theme(map[string]string{"redirect": r.URL.Path}),
		Scripts:         scripts,
		ScriptURLs:      scriptURLs,
		Styles:          styles,
		StyleURLs:       options.StyleURLs,
		FaviconURL:      FaviconURL(),
		// Theme: dashboard.THEME_MINTY,
	})

	return dashboard
}
