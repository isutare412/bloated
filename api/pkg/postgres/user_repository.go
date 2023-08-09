package postgres

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/core/ent/user"
	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

type UserRepository struct {
	conn *Connection
}

func NewUserRepository(conn *Connection) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	found, err := r.conn.txClient(ctx).User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, pkgerror.Known{
			Code:   pkgerror.CodeNotFound,
			Origin: err,
			Simple: fmt.Errorf("todo with id(%d) does not exist", id),
		}
	case err != nil:
		return nil, err
	}

	return found, nil
}

func (r *UserRepository) Upsert(ctx context.Context, usr *ent.User) (*ent.User, error) {
	userID, err := r.conn.txClient(ctx).User.
		Create().
		SetEmail(usr.Email).
		SetNillableUserName(pointerIfNotZero(usr.UserName)).
		SetNillableGivenName(pointerIfNotZero(usr.GivenName)).
		SetNillableFamilyName(pointerIfNotZero(usr.FamilyName)).
		SetNillablePhotoURL(pointerIfNotZero(usr.PhotoURL)).
		SetOrigin(usr.Origin).
		OnConflict(
			sql.ConflictColumns(user.FieldEmail),
			sql.ResolveWithNewValues(),
			sql.ResolveWith(func(us *sql.UpdateSet) {
				us.SetIgnore(user.FieldID)

				if usr.UserName == "" {
					us.SetNull(user.FieldUserName)
				}
				if usr.GivenName == "" {
					us.SetNull(user.FieldGivenName)
				}
				if usr.FamilyName == "" {
					us.SetNull(user.FieldFamilyName)
				}
				if usr.PhotoURL == "" {
					us.SetNull(user.FieldPhotoURL)
				}
			}),
		).
		ID(ctx)
	if err != nil {
		return nil, err
	}

	found, err := r.conn.txClient(ctx).User.
		Query().
		Where(user.ID(userID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return found, nil
}
