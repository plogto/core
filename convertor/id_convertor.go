package convertor

import (
	"encoding/hex"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func StringsToUUIDs(ids []string) []pgtype.UUID {
	var result []pgtype.UUID

	for _, id := range ids {
		result = append(result, StringToUUID(id))
	}

	return result
}

func StringToUUID(s string) pgtype.UUID {
	data, err := parseUUID(s)
	if err != nil {
		panic(err)
	}
	return pgtype.UUID{
		Bytes: data,
		Valid: true,
	}
}

func parseUUID(src string) (dst [16]byte, err error) {
	switch len(src) {
	case 36:
		src = src[0:8] + src[9:13] + src[14:18] + src[19:23] + src[24:]
	case 32:
		// dashes already stripped, assume valid
	default:
		// assume invalid.
		fmt.Printf("cannot parse UUID %v", src)
		return dst, nil
	}

	buf, err := hex.DecodeString(src)
	if err != nil {
		return dst, err
	}

	copy(dst[:], buf)
	return dst, err
}

func UUIDToString(t pgtype.UUID) string {
	if !t.Valid {
		return ""
	}
	src := t.Bytes
	return fmt.Sprintf("%x-%x-%x-%x-%x", src[0:4], src[4:6], src[6:8], src[8:10], src[10:16])
}
