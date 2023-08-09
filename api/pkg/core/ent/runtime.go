// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/isutare412/bloated/api/pkg/core/ent/bannedip"
	"github.com/isutare412/bloated/api/pkg/core/ent/schema"
	"github.com/isutare412/bloated/api/pkg/core/ent/todo"
	"github.com/isutare412/bloated/api/pkg/core/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bannedipMixin := schema.BannedIP{}.Mixin()
	bannedipMixinFields0 := bannedipMixin[0].Fields()
	_ = bannedipMixinFields0
	bannedipFields := schema.BannedIP{}.Fields()
	_ = bannedipFields
	// bannedipDescCreateTime is the schema descriptor for create_time field.
	bannedipDescCreateTime := bannedipMixinFields0[0].Descriptor()
	// bannedip.DefaultCreateTime holds the default value on creation for the create_time field.
	bannedip.DefaultCreateTime = bannedipDescCreateTime.Default.(func() time.Time)
	// bannedipDescUpdateTime is the schema descriptor for update_time field.
	bannedipDescUpdateTime := bannedipMixinFields0[1].Descriptor()
	// bannedip.DefaultUpdateTime holds the default value on creation for the update_time field.
	bannedip.DefaultUpdateTime = bannedipDescUpdateTime.Default.(func() time.Time)
	// bannedip.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	bannedip.UpdateDefaultUpdateTime = bannedipDescUpdateTime.UpdateDefault.(func() time.Time)
	// bannedipDescIP is the schema descriptor for ip field.
	bannedipDescIP := bannedipFields[0].Descriptor()
	// bannedip.IPValidator is a validator for the "ip" field. It is called by the builders before save.
	bannedip.IPValidator = func() func(string) error {
		validators := bannedipDescIP.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(ip string) error {
			for _, fn := range fns {
				if err := fn(ip); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// bannedipDescCountry is the schema descriptor for country field.
	bannedipDescCountry := bannedipFields[1].Descriptor()
	// bannedip.CountryValidator is a validator for the "country" field. It is called by the builders before save.
	bannedip.CountryValidator = func() func(string) error {
		validators := bannedipDescCountry.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(country string) error {
			for _, fn := range fns {
				if err := fn(country); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	todoMixin := schema.Todo{}.Mixin()
	todoMixinFields0 := todoMixin[0].Fields()
	_ = todoMixinFields0
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescCreateTime is the schema descriptor for create_time field.
	todoDescCreateTime := todoMixinFields0[0].Descriptor()
	// todo.DefaultCreateTime holds the default value on creation for the create_time field.
	todo.DefaultCreateTime = todoDescCreateTime.Default.(func() time.Time)
	// todoDescUpdateTime is the schema descriptor for update_time field.
	todoDescUpdateTime := todoMixinFields0[1].Descriptor()
	// todo.DefaultUpdateTime holds the default value on creation for the update_time field.
	todo.DefaultUpdateTime = todoDescUpdateTime.Default.(func() time.Time)
	// todo.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	todo.UpdateDefaultUpdateTime = todoDescUpdateTime.UpdateDefault.(func() time.Time)
	// todoDescTask is the schema descriptor for task field.
	todoDescTask := todoFields[1].Descriptor()
	// todo.TaskValidator is a validator for the "task" field. It is called by the builders before save.
	todo.TaskValidator = func() func(string) error {
		validators := todoDescTask.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(task string) error {
			for _, fn := range fns {
				if err := fn(task); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescUserName is the schema descriptor for user_name field.
	userDescUserName := userFields[2].Descriptor()
	// user.UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	user.UserNameValidator = func() func(string) error {
		validators := userDescUserName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(user_name string) error {
			for _, fn := range fns {
				if err := fn(user_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescGivenName is the schema descriptor for given_name field.
	userDescGivenName := userFields[3].Descriptor()
	// user.GivenNameValidator is a validator for the "given_name" field. It is called by the builders before save.
	user.GivenNameValidator = func() func(string) error {
		validators := userDescGivenName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(given_name string) error {
			for _, fn := range fns {
				if err := fn(given_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescFamilyName is the schema descriptor for family_name field.
	userDescFamilyName := userFields[4].Descriptor()
	// user.FamilyNameValidator is a validator for the "family_name" field. It is called by the builders before save.
	user.FamilyNameValidator = func() func(string) error {
		validators := userDescFamilyName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(family_name string) error {
			for _, fn := range fns {
				if err := fn(family_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPhotoURL is the schema descriptor for photo_url field.
	userDescPhotoURL := userFields[5].Descriptor()
	// user.PhotoURLValidator is a validator for the "photo_url" field. It is called by the builders before save.
	user.PhotoURLValidator = userDescPhotoURL.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
