package links

type authLinks struct {
}

func NewAuthLinks() *authLinks {
	return &authLinks{}
}

func (l *authLinks) Auth() string {
	return URL(AUTH_AUTH, nil)
}

func (l *authLinks) AuthKnightLogin(backUrl string) string {
	params := map[string]string{
		"back_url": backUrl,
		"next_url": l.Auth(),
	}
	return "https://authknight.com/app/login" + query(params)
}

func (l *authLinks) Login(backUrl string) string {
	params := map[string]string{}

	if backUrl != "" {
		params["back_url"] = backUrl
	}

	return URL(AUTH_LOGIN, params)
}

func (l *authLinks) Logout() string {
	return URL(AUTH_LOGOUT, nil)
}

func (l *authLinks) Register(params map[string]string) string {
	return URL(AUTH_REGISTER, params)
}
