package db

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jackc/pgx/v5/pgtype"
)

// MarshalUUID allows uuid to be marshalled by graphql
func MarshalUUID(id pgtype.UUID) graphql.Marshaler {
	str := fmt.Sprintf("%x-%x-%x-%x-%x", id.Bytes[0:4], id.Bytes[4:6], id.Bytes[6:8], id.Bytes[8:10], id.Bytes[10:16])
	return graphql.MarshalString(str)
}

// UnmarshalUUID allows uuid to be unmarshalled by graphql
func UnmarshalUUID(v interface{}) (pgtype.UUID, error) {
	idAsString, ok := v.(pgtype.UUID)
	if !ok {
		return pgtype.UUID{}, errors.New("id should be a valid UUID")
	}

	return idAsString, nil
}

type UserSettingValue string

const (
	UserSettingValueOff UserSettingValue = "off"
	UserSettingValueOn  UserSettingValue = "on"
)

var AllUserSettingValue = []UserSettingValue{
	UserSettingValueOff,
	UserSettingValueOn,
}

func (e UserSettingValue) IsValid() bool {
	switch e {
	case UserSettingValueOff, UserSettingValueOn:
		return true
	}
	return false
}

func (e UserSettingValue) String() string {
	return string(e)
}

func (e *UserSettingValue) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserSettingValue(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserSettingValue", str)
	}
	return nil
}

func (e UserSettingValue) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserSettings struct {
	IsRepliesVisible UserSettingValue `json:"is_replies_visible"`
	IsMediaVisible   UserSettingValue `json:"is_media_visible"`
	IsLikesVisible   UserSettingValue `json:"is_likes_visible"`
}
