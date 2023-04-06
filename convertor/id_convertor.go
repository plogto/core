package convertor

import (
	"github.com/google/uuid"
)

func StringsToUUIDs(ids []string) []uuid.UUID {
	var result []uuid.UUID

	for _, id := range ids {
		result = append(result, uuid.MustParse(id))
	}

	return result
}
