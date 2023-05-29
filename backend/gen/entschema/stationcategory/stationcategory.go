// Code generated by entc, DO NOT EDIT.

package stationcategory

import (
	"time"
)

const (
	// Label holds the string label denoting the stationcategory type in the database.
	Label = "station_category"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSequence holds the string denoting the sequence field in the database.
	FieldSequence = "sequence"
	// Table holds the table name of the stationcategory in the database.
	Table = "station_category"
)

// Columns holds all SQL columns for stationcategory fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeactivatedAt,
	FieldName,
	FieldStatus,
	FieldSequence,
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus bool
)
