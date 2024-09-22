package links

type adminLinks struct{}

func NewAdminLinks() *adminLinks {
	return &adminLinks{}
}

func (l *adminLinks) Home(params map[string]string) string {
	return URL(ADMIN_HOME, params)
}

func (l *adminLinks) Blog(params map[string]string) string {
	return URL(ADMIN_BLOG, params)
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

func (l *adminLinks) FileManager(params map[string]string) string {
	return URL(ADMIN_MEDIA, params)
}

func (l *adminLinks) Users(params map[string]string) string {
	return URL(ADMIN_USERS, params)
}

func (l *adminLinks) UsersUserCreate(params map[string]string) string {
	return URL(ADMIN_USERS_USER_CREATE, params)
}

func (l *adminLinks) UsersUserDelete(params map[string]string) string {
	return URL(ADMIN_USERS_USER_DELETE, params)
}

func (l *adminLinks) UsersUserImpersonate(params map[string]string) string {
	return URL(ADMIN_USERS_USER_IMPERSONATE, params)
}

func (l *adminLinks) UsersUserManager(params map[string]string) string {
	return URL(ADMIN_USERS_USER_MANAGER, params)
}

func (l *adminLinks) UsersUserUpdate(params map[string]string) string {
	return URL(ADMIN_USERS_USER_UPDATE, params)
}
