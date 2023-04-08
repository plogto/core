package convertor

import (
	"strings"

	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
)

func ModelPrimaryColorToDB(primaryColor model.PrimaryColor) db.PrimaryColor {
	color := primaryColor.String()

	return db.PrimaryColor(strings.ToLower(color))
}

func ModelBackgroundColorToDB(backgroundColor model.BackgroundColor) db.BackgroundColor {
	color := backgroundColor.String()

	return db.BackgroundColor(strings.ToLower(color))
}
