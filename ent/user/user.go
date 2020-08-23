// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldUserID holds the string denoting the userid field in the database.
	FieldUserID = "user_id"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldUserID,
	FieldPassword,
}

var (
	// UserIDValidator is a validator for the "UserID" field. It is called by the builders before save.
	UserIDValidator func(string) error
	// PasswordValidator is a validator for the "Password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
