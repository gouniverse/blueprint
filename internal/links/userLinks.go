package links

type userLinks struct{}

func NewUserLinks() *userLinks {
	return &userLinks{}
}

func (l *userLinks) Home() string {
	return URL(USER_HOME, nil)
}

func (l *userLinks) Profile() string {
	return URL(USER_PROFILE, nil)
}

func (l *userLinks) ProfileSave() string {
	return URL(USER_PROFILE_UPDATE, nil)
}
