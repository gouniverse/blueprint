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
const ADMIN_BLOG_POST_MANAGER = "/admin/blog/post-manager"
const ADMIN_BLOG_POST_CREATE = "/admin/blog/post-create"
const ADMIN_BLOG_POST_DELETE = "/admin/blog/post-delete"
const ADMIN_BLOG_POST_UPDATE = "/admin/blog/post-update"
const ADMIN_CMS = "/admin/cms"
const ADMIN_FILE_MANAGER = "/admin/file-manager"
const ADMIN_MEDIA = "/admin/media"
const ADMIN_SHOP = ADMIN_HOME + "/shop"
const ADMIN_STATS = ADMIN_HOME + "/stats"
const ADMIN_TASKS = ADMIN_HOME + "/tasks"
const ADMIN_USERS = ADMIN_HOME + "/users"
const ADMIN_USERS_USER_CREATE = ADMIN_USERS + "/user-create"
const ADMIN_USERS_USER_DELETE = ADMIN_USERS + "/user-delete"
const ADMIN_USERS_USER_IMPERSONATE = ADMIN_USERS + "/user-impersonate"
const ADMIN_USERS_USER_MANAGER = ADMIN_USERS + "/user-manager"
const ADMIN_USERS_USER_UPDATE = ADMIN_USERS + "/user-update"

// == USER LINKS =============================================================

const USER_HOME = "/user"
const USER_ORDERS = USER_HOME + "/orders"
const USER_ORDER_CREATE = USER_ORDERS + "/create"
const USER_ORDER_CREATE_PAYMENT_BEGIN = USER_ORDER_CREATE + "/payment-begin"
const USER_ORDER_DELETE = USER_ORDERS + "/delete"
const USER_ORDER_LIST = USER_ORDERS + "/list"
const USER_PROFILE = USER_HOME + "/profile"
const USER_PROFILE_UPDATE = USER_HOME + "/profile/update"

// == WEBSITE LINKS ==========================================================

const HOME = "/"
const BLOG = "/blog"
const BLOG_POST = "/blog/post"
const BLOG_POST_WITH_REGEX = BLOG_POST + "/{id:[0-9]+}"
const BLOG_POST_WITH_REGEX2 = BLOG_POST + "/{id:[0-9]+}/{title}"
const CONTACT = "/contact"
const FILES = "/files" + CATCHALL
const FLASH = "/flash"
const MEDIA = "/media" + CATCHALL
const PAYMENT_CANCELED = "/payment/canceled"
const PAYMENT_SUCCESS = "/payment/success"
const RESOURCES = "/resources"
const THEME = "/theme"
const THUMB = "/th/{extension:[a-z]+}/{size:[0-9x]+}/{quality:[0-9]+}/*"
const WIDGET = "/widget"
