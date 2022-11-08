package constants

type key string

const POSTS_PAGE_LIMIT int = 10
const TRENDS_PAGE_LIMIT int = 8
const USERS_PAGE_LIMIT int = 20
const TAGS_PAGE_LIMIT int = 20
const CURRENT_USER_KEY key = "CURRENT_USER"
const CURRENT_ONLINE_USER_KEY key = "CURRENT_ONLINE_USER"

var MB int64 = 1 << 20

const MENTION_PATTERN = "@(\\w|_)+"
const KEY_PATTERN = "(\\$\\$\\$___[0123456789abcdefg-]+___\\$\\$\\$)"

// loader keys
const USER_LOADER_KEY = "USER_LOADER_KEY"
const POST_LOADER_KEY = "POST_LOADER_KEY"
