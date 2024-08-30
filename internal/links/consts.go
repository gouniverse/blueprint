package links

const CATCHALL = "/*"

// == AUTH LINKS =============================================================

const AUTH_AUTH = "/auth/auth"
const AUTH_LOGIN = "/auth/login"
const AUTH_LOGOUT = "/auth/logout"
const AUTH_REGISTER = "/auth/register"

// == ADMIN LINKS ============================================================

const ADMIN_HOME = "/admin"
const ADMIN_BLOG = "/admin/blog"
const ADMIN_CMS = "/admin/cms"
const ADMIN_MEDIA = "/admin/media"
const ADMIN_USERS = "/admin/users"

// == USER LINKS =============================================================

const USER_HOME = "/user"
const USER_PROFILE = USER_HOME + "/profile"
const USER_PROFILE_UPDATE = USER_HOME + "/profile/update"

// == WEBSITE LINKS ==========================================================

const HOME = "/"
const BLOG = "/blog"
const BLOG_POST = "/blog/post"
const BLOG_POST_WITH_REGEX = BLOG_POST + "/{id:[0-9]+}"
const BLOG_POST_WITH_REGEX2 = BLOG_POST + "/{id:[0-9]+}/{title}"
const CONTACT = "/contact"
const FLASH = "/flash"
const MEDIA = "/media" + CATCHALL

const THEME = "/theme"
const THUMB = "/th/{extension:[a-z]+}/{size:[0-9x]+}/{quality:[0-9]+}/*"
const WIDGET = "/widget"
