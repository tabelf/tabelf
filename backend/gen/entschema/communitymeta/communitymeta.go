// Code generated by entc, DO NOT EDIT.

package communitymeta

import (
	"time"
)

const (
	// Label holds the string label denoting the communitymeta type in the database.
	Label = "community_meta"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeactivatedAt holds the string denoting the deactivated_at field in the database.
	FieldDeactivatedAt = "deactivated_at"
	// FieldCommunityUID holds the string denoting the community_uid field in the database.
	FieldCommunityUID = "community_uid"
	// FieldUserUID holds the string denoting the user_uid field in the database.
	FieldUserUID = "user_uid"
	// FieldHasPraise holds the string denoting the has_praise field in the database.
	FieldHasPraise = "has_praise"
	// FieldHasStar holds the string denoting the has_star field in the database.
	FieldHasStar = "has_star"
	// FieldHasView holds the string denoting the has_view field in the database.
	FieldHasView = "has_view"
	// FieldHasUsed holds the string denoting the has_used field in the database.
	FieldHasUsed = "has_used"
	// Table holds the table name of the communitymeta in the database.
	Table = "community_meta"
)

// Columns holds all SQL columns for communitymeta fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeactivatedAt,
	FieldCommunityUID,
	FieldUserUID,
	FieldHasPraise,
	FieldHasStar,
	FieldHasView,
	FieldHasUsed,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUID holds the default value on creation for the "uid" field.
	DefaultUID func() string
	// UIDValidator is a validator for the "uid" field. It is called by the builders before save.
	UIDValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultHasPraise holds the default value on creation for the "has_praise" field.
	DefaultHasPraise bool
	// DefaultHasStar holds the default value on creation for the "has_star" field.
	DefaultHasStar bool
	// DefaultHasView holds the default value on creation for the "has_view" field.
	DefaultHasView bool
	// DefaultHasUsed holds the default value on creation for the "has_used" field.
	DefaultHasUsed bool
)
