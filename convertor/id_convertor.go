package convertor

import (
	"github.com/google/uuid"
)

func StringsToUUIDs(ids []string) []uuid.UUID {
	var result []uuid.UUID

	for _, id := range ids {
		ID, _ := uuid.Parse(id)
		result = append(result, ID)
	}

	return result
}
