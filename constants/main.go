package constants

type key string

// TODO: remove POSTS_PAGE_LIMIT
const POSTS_PAGE_LIMIT int = 10
const POSTS_PAGE_FIRST int = 10
const CURRENT_USER_KEY key = "CURRENT_USER"
const CURRENT_ONLINE_USER_KEY key = "CURRENT_ONLINE_USER"

var MB int64 = 1 << 20
