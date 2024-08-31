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

func (l *adminLinks) BlogPostCreate(params map[string]string) string {
	return URL(ADMIN_BLOG_POST_CREATE, params)
}

func (l *adminLinks) BlogPostDelete(params map[string]string) string {
	return URL(ADMIN_BLOG_POST_DELETE, params)
}

func (l *adminLinks) BlogPostManager(params map[string]string) string {
	return URL(ADMIN_BLOG_POST_MANAGER, params)
}

func (l *adminLinks) BlogPostUpdate(params map[string]string) string {
	return URL(ADMIN_BLOG_POST_UPDATE, params)
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
