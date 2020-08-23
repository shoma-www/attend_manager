// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/shoma-www/attend_manager/ent/schema"
	"github.com/shoma-www/attend_manager/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserID is the schema descriptor for UserID field.
	userDescUserID := userFields[1].Descriptor()
	// user.UserIDValidator is a validator for the "UserID" field. It is called by the builders before save.
	user.UserIDValidator = func() func(string) error {
		validators := userDescUserID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_UserID string) error {
			for _, fn := range fns {
				if err := fn(_UserID); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPassword is the schema descriptor for Password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "Password" field. It is called by the builders before save.
	user.PasswordValidator = func() func(string) error {
		validators := userDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Password string) error {
			for _, fn := range fns {
				if err := fn(_Password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
