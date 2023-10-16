package links

type adminLinks struct{}

func NewAdminLinks() *adminLinks {
	return &adminLinks{}
}

func (l *adminLinks) Home() string {
	return URL(ADMIN_HOME, nil)
}

func (l *adminLinks) Blog() string {
	return URL(ADMIN_BLOG, nil)
}

func (l *adminLinks) Cms() string {
	return URL(ADMIN_CMS, nil)
}

func (l *adminLinks) FileManager() string {
	return URL(ADMIN_MEDIA, nil)
}

func (l *adminLinks) FileManagerWithParams(params map[string]string) string {
	return URL(ADMIN_MEDIA, params)
}

func (l *adminLinks) Users() string {
	return URL(ADMIN_USERS, nil)
}

func (l *adminLinks) UsersWithParams(params map[string]string) string {
	return URL(ADMIN_USERS, params)
}
