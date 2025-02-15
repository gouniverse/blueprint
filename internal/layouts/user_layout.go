package layouts

import (
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/links"

	"github.com/gouniverse/cdn"
	"github.com/gouniverse/cmsstore"
	"github.com/gouniverse/dashboard"
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
		firstName, lastName, err := getUserData(r, authUser)
		if err == nil {
			dashboardUser = dashboard.User{
				FirstName: firstName,
				LastName:  lastName,
			}
		}
	}

	// googleTagScriptURL := "https://www.googletagmanager.com/gtag/js?id=G-247NHE839P"
	// googleTagScript := `window.dataLayer = window.dataLayer || []; function gtag(){dataLayer.push(arguments);} gtag('js', new Date()); gtag('config', 'G-247NHE839P');`
	// googleAdsScriptURL := "https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-8821108004642146"
	// statcounterScript := `<script type="text/javascript">var sc_project=12939246;var sc_invisible=1;var sc_security="2c1cdc75";</script><script	src="https://www.statcounter.com/counter/counter.js" async></script><noscript><img	src="https://c.statcounter.com/12939246/0/2c1cdc75/1/" alt=""></noscript>`

	// Prepare script URLs
	scriptURLs := []string{} // prepend any if required
	scriptURLs = append(scriptURLs, options.ScriptURLs...)
	scriptURLs = append(scriptURLs, cdn.Htmx_2_0_0())

	// Prepare scripts
	scripts := []string{} // prepend any if required
	scripts = append(scripts, options.Scripts...)

	// Prepare styles
	styles := []string{ // prepend any if required
		`nav#Toolbar {border-bottom: 8px solid blue;}`,
	}
	styles = append(styles, options.Styles...)

	homeLink := links.NewUserLinks().Home(map[string]string{})

	titlePostfix := ` | ` + lo.Ternary(authUser == nil, `Guest`, `User`) + ` | ` + config.AppName

	_, isPage := r.Context().Value("page").(cmsstore.PageInterface)

	if isPage {
		titlePostfix = "" // no postfix for CMS pages
	}

	dashboard := dashboard.NewDashboard(dashboard.Config{
		HTTPRequest: r,
		Content:     options.Content.ToHTML(),
		Title:       options.Title + titlePostfix,
		LoginURL:    links.NewAuthLinks().Login(homeLink),
		MenuItems:   userLayoutMainMenuItems(authUser),
		// LogoImageURL:              "/media/user/dashboard-logo.jpg",
		NavbarBackgroundColorMode: "primary",
		LogoRawHtml:               userLogoHtml(),
		LogoRedirectURL:           homeLink,
		User:                      dashboardUser,
		UserMenu:                  userLayoutUserMenuItems(authUser),
		ThemeHandlerUrl:           links.NewWebsiteLinks().Theme(map[string]string{"redirect": r.URL.Path}),
		Scripts:                   scripts,
		ScriptURLs:                scriptURLs,
		Styles:                    styles,
		StyleURLs:                 options.StyleURLs,
		FaviconURL:                FaviconURL(),
		// Theme: dashboard.THEME_MINTY,
	})

	return dashboard
}
