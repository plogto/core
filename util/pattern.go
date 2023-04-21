package util

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/convertor"
)

func PrepareKeyPattern(value pgtype.UUID) string {
	return "$$$___" + convertor.UUIDToString(value) + "___$$$"
}
