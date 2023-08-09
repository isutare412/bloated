// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/isutare412/bloated/api/pkg/core/ent/predicate"
	"github.com/isutare412/bloated/api/pkg/core/ent/tokenhistory"
	"github.com/isutare412/bloated/api/pkg/core/enum"
)

// TokenHistoryUpdate is the builder for updating TokenHistory entities.
type TokenHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *TokenHistoryMutation
}

// Where appends a list predicates to the TokenHistoryUpdate builder.
func (thu *TokenHistoryUpdate) Where(ps ...predicate.TokenHistory) *TokenHistoryUpdate {
	thu.mutation.Where(ps...)
	return thu
}

// SetUpdateTime sets the "update_time" field.
func (thu *TokenHistoryUpdate) SetUpdateTime(t time.Time) *TokenHistoryUpdate {
	thu.mutation.SetUpdateTime(t)
	return thu
}

// SetEmail sets the "email" field.
func (thu *TokenHistoryUpdate) SetEmail(s string) *TokenHistoryUpdate {
	thu.mutation.SetEmail(s)
	return thu
}

// SetUserName sets the "user_name" field.
func (thu *TokenHistoryUpdate) SetUserName(s string) *TokenHistoryUpdate {
	thu.mutation.SetUserName(s)
	return thu
}

// SetIssuedFrom sets the "issued_from" field.
func (thu *TokenHistoryUpdate) SetIssuedFrom(e enum.Issuer) *TokenHistoryUpdate {
	thu.mutation.SetIssuedFrom(e)
	return thu
}

// Mutation returns the TokenHistoryMutation object of the builder.
func (thu *TokenHistoryUpdate) Mutation() *TokenHistoryMutation {
	return thu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (thu *TokenHistoryUpdate) Save(ctx context.Context) (int, error) {
	thu.defaults()
	return withHooks(ctx, thu.sqlSave, thu.mutation, thu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (thu *TokenHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := thu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (thu *TokenHistoryUpdate) Exec(ctx context.Context) error {
	_, err := thu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thu *TokenHistoryUpdate) ExecX(ctx context.Context) {
	if err := thu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (thu *TokenHistoryUpdate) defaults() {
	if _, ok := thu.mutation.UpdateTime(); !ok {
		v := tokenhistory.UpdateDefaultUpdateTime()
		thu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (thu *TokenHistoryUpdate) check() error {
	if v, ok := thu.mutation.Email(); ok {
		if err := tokenhistory.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "TokenHistory.email": %w`, err)}
		}
	}
	if v, ok := thu.mutation.UserName(); ok {
		if err := tokenhistory.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`ent: validator failed for field "TokenHistory.user_name": %w`, err)}
		}
	}
	if v, ok := thu.mutation.IssuedFrom(); ok {
		if err := tokenhistory.IssuedFromValidator(v); err != nil {
			return &ValidationError{Name: "issued_from", err: fmt.Errorf(`ent: validator failed for field "TokenHistory.issued_from": %w`, err)}
		}
	}
	return nil
}

func (thu *TokenHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := thu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tokenhistory.Table, tokenhistory.Columns, sqlgraph.NewFieldSpec(tokenhistory.FieldID, field.TypeInt))
	if ps := thu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := thu.mutation.UpdateTime(); ok {
		_spec.SetField(tokenhistory.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := thu.mutation.Email(); ok {
		_spec.SetField(tokenhistory.FieldEmail, field.TypeString, value)
	}
	if value, ok := thu.mutation.UserName(); ok {
		_spec.SetField(tokenhistory.FieldUserName, field.TypeString, value)
	}
	if value, ok := thu.mutation.IssuedFrom(); ok {
		_spec.SetField(tokenhistory.FieldIssuedFrom, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, thu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tokenhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	thu.mutation.done = true
	return n, nil
}

// TokenHistoryUpdateOne is the builder for updating a single TokenHistory entity.
type TokenHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TokenHistoryMutation
}

// SetUpdateTime sets the "update_time" field.
func (thuo *TokenHistoryUpdateOne) SetUpdateTime(t time.Time) *TokenHistoryUpdateOne {
	thuo.mutation.SetUpdateTime(t)
	return thuo
}

// SetEmail sets the "email" field.
func (thuo *TokenHistoryUpdateOne) SetEmail(s string) *TokenHistoryUpdateOne {
	thuo.mutation.SetEmail(s)
	return thuo
}

// SetUserName sets the "user_name" field.
func (thuo *TokenHistoryUpdateOne) SetUserName(s string) *TokenHistoryUpdateOne {
	thuo.mutation.SetUserName(s)
	return thuo
}

// SetIssuedFrom sets the "issued_from" field.
func (thuo *TokenHistoryUpdateOne) SetIssuedFrom(e enum.Issuer) *TokenHistoryUpdateOne {
	thuo.mutation.SetIssuedFrom(e)
	return thuo
}

// Mutation returns the TokenHistoryMutation object of the builder.
func (thuo *TokenHistoryUpdateOne) Mutation() *TokenHistoryMutation {
	return thuo.mutation
}

// Where appends a list predicates to the TokenHistoryUpdate builder.
func (thuo *TokenHistoryUpdateOne) Where(ps ...predicate.TokenHistory) *TokenHistoryUpdateOne {
	thuo.mutation.Where(ps...)
	return thuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (thuo *TokenHistoryUpdateOne) Select(field string, fields ...string) *TokenHistoryUpdateOne {
	thuo.fields = append([]string{field}, fields...)
	return thuo
}

// Save executes the query and returns the updated TokenHistory entity.
func (thuo *TokenHistoryUpdateOne) Save(ctx context.Context) (*TokenHistory, error) {
	thuo.defaults()
	return withHooks(ctx, thuo.sqlSave, thuo.mutation, thuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (thuo *TokenHistoryUpdateOne) SaveX(ctx context.Context) *TokenHistory {
	node, err := thuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (thuo *TokenHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := thuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thuo *TokenHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := thuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (thuo *TokenHistoryUpdateOne) defaults() {
	if _, ok := thuo.mutation.UpdateTime(); !ok {
		v := tokenhistory.UpdateDefaultUpdateTime()
		thuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (thuo *TokenHistoryUpdateOne) check() error {
	if v, ok := thuo.mutation.Email(); ok {
		if err := tokenhistory.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "TokenHistory.email": %w`, err)}
		}
	}
	if v, ok := thuo.mutation.UserName(); ok {
		if err := tokenhistory.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`ent: validator failed for field "TokenHistory.user_name": %w`, err)}
		}
	}
	if v, ok := thuo.mutation.IssuedFrom(); ok {
		if err := tokenhistory.IssuedFromValidator(v); err != nil {
			return &ValidationError{Name: "issued_from", err: fmt.Errorf(`ent: validator failed for field "TokenHistory.issued_from": %w`, err)}
		}
	}
	return nil
}

func (thuo *TokenHistoryUpdateOne) sqlSave(ctx context.Context) (_node *TokenHistory, err error) {
	if err := thuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tokenhistory.Table, tokenhistory.Columns, sqlgraph.NewFieldSpec(tokenhistory.FieldID, field.TypeInt))
	id, ok := thuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TokenHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := thuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tokenhistory.FieldID)
		for _, f := range fields {
			if !tokenhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tokenhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := thuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := thuo.mutation.UpdateTime(); ok {
		_spec.SetField(tokenhistory.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := thuo.mutation.Email(); ok {
		_spec.SetField(tokenhistory.FieldEmail, field.TypeString, value)
	}
	if value, ok := thuo.mutation.UserName(); ok {
		_spec.SetField(tokenhistory.FieldUserName, field.TypeString, value)
	}
	if value, ok := thuo.mutation.IssuedFrom(); ok {
		_spec.SetField(tokenhistory.FieldIssuedFrom, field.TypeEnum, value)
	}
	_node = &TokenHistory{config: thuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, thuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tokenhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	thuo.mutation.done = true
	return _node, nil
}