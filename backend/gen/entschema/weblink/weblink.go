// Code generated by entc, DO NOT EDIT.

package weblink

import (
	"time"
)

const (
	// Label holds the string label denoting the weblink type in the database.
	Label = "web_link"
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
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldFileType holds the string denoting the file_type field in the database.
	FieldFileType = "file_type"
	// FieldSequence holds the string denoting the sequence field in the database.
	FieldSequence = "sequence"
	// FieldForeverDelete holds the string denoting the forever_delete field in the database.
	FieldForeverDelete = "forever_delete"
	// FieldUserUID holds the string denoting the user_uid field in the database.
	FieldUserUID = "user_uid"
	// FieldWorkspaceUID holds the string denoting the workspace_uid field in the database.
	FieldWorkspaceUID = "workspace_uid"
	// FieldFolderUID holds the string denoting the folder_uid field in the database.
	FieldFolderUID = "folder_uid"
	// Table holds the table name of the weblink in the database.
	Table = "web_link"
)

// Columns holds all SQL columns for weblink fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeactivatedAt,
	FieldTitle,
	FieldImage,
	FieldLink,
	FieldDescription,
	FieldFileType,
	FieldSequence,
	FieldForeverDelete,
	FieldUserUID,
	FieldWorkspaceUID,
	FieldFolderUID,
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
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultFileType holds the default value on creation for the "file_type" field.
	DefaultFileType string
	// DefaultSequence holds the default value on creation for the "sequence" field.
	DefaultSequence int
	// DefaultForeverDelete holds the default value on creation for the "forever_delete" field.
	DefaultForeverDelete bool
)
