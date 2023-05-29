// Code generated by entc, DO NOT EDIT.

package teamfolder

import (
	"time"
)

const (
	// Label holds the string label denoting the teamfolder type in the database.
	Label = "team_folder"
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
	// FieldTeamUID holds the string denoting the team_uid field in the database.
	FieldTeamUID = "team_uid"
	// FieldFileName holds the string denoting the file_name field in the database.
	FieldFileName = "file_name"
	// Table holds the table name of the teamfolder in the database.
	Table = "team_folder"
)

// Columns holds all SQL columns for teamfolder fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeactivatedAt,
	FieldTeamUID,
	FieldFileName,
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
)