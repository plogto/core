package constants

import "database/sql"

var GENERATE_CREDITS_AMOUNT sql.NullFloat64 = sql.NullFloat64{100, true}
