// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/isutare412/bloated/api/pkg/core/ent/predicate"
	"github.com/isutare412/bloated/api/pkg/core/ent/todo"
	"github.com/isutare412/bloated/api/pkg/core/ent/user"
	"github.com/isutare412/bloated/api/pkg/core/enum"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// ClearEmail clears the value of the "email" field.
func (uu *UserUpdate) ClearEmail() *UserUpdate {
	uu.mutation.ClearEmail()
	return uu
}

// SetUserName sets the "user_name" field.
func (uu *UserUpdate) SetUserName(s string) *UserUpdate {
	uu.mutation.SetUserName(s)
	return uu
}

// SetNillableUserName sets the "user_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUserName(s *string) *UserUpdate {
	if s != nil {
		uu.SetUserName(*s)
	}
	return uu
}

// ClearUserName clears the value of the "user_name" field.
func (uu *UserUpdate) ClearUserName() *UserUpdate {
	uu.mutation.ClearUserName()
	return uu
}

// SetGivenName sets the "given_name" field.
func (uu *UserUpdate) SetGivenName(s string) *UserUpdate {
	uu.mutation.SetGivenName(s)
	return uu
}

// SetNillableGivenName sets the "given_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableGivenName(s *string) *UserUpdate {
	if s != nil {
		uu.SetGivenName(*s)
	}
	return uu
}

// ClearGivenName clears the value of the "given_name" field.
func (uu *UserUpdate) ClearGivenName() *UserUpdate {
	uu.mutation.ClearGivenName()
	return uu
}

// SetFamilyName sets the "family_name" field.
func (uu *UserUpdate) SetFamilyName(s string) *UserUpdate {
	uu.mutation.SetFamilyName(s)
	return uu
}

// SetNillableFamilyName sets the "family_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFamilyName(s *string) *UserUpdate {
	if s != nil {
		uu.SetFamilyName(*s)
	}
	return uu
}

// ClearFamilyName clears the value of the "family_name" field.
func (uu *UserUpdate) ClearFamilyName() *UserUpdate {
	uu.mutation.ClearFamilyName()
	return uu
}

// SetPhotoURL sets the "photo_url" field.
func (uu *UserUpdate) SetPhotoURL(s string) *UserUpdate {
	uu.mutation.SetPhotoURL(s)
	return uu
}

// SetNillablePhotoURL sets the "photo_url" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhotoURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhotoURL(*s)
	}
	return uu
}

// ClearPhotoURL clears the value of the "photo_url" field.
func (uu *UserUpdate) ClearPhotoURL() *UserUpdate {
	uu.mutation.ClearPhotoURL()
	return uu
}

// SetOrigin sets the "origin" field.
func (uu *UserUpdate) SetOrigin(e enum.Issuer) *UserUpdate {
	uu.mutation.SetOrigin(e)
	return uu
}

// AddTodoIDs adds the "todos" edge to the Todo entity by IDs.
func (uu *UserUpdate) AddTodoIDs(ids ...int) *UserUpdate {
	uu.mutation.AddTodoIDs(ids...)
	return uu
}

// AddTodos adds the "todos" edges to the Todo entity.
func (uu *UserUpdate) AddTodos(t ...*Todo) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTodoIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearTodos clears all "todos" edges to the Todo entity.
func (uu *UserUpdate) ClearTodos() *UserUpdate {
	uu.mutation.ClearTodos()
	return uu
}

// RemoveTodoIDs removes the "todos" edge to Todo entities by IDs.
func (uu *UserUpdate) RemoveTodoIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveTodoIDs(ids...)
	return uu
}

// RemoveTodos removes "todos" edges to Todo entities.
func (uu *UserUpdate) RemoveTodos(t ...*Todo) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTodoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.UserName(); ok {
		if err := user.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`ent: validator failed for field "User.user_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.GivenName(); ok {
		if err := user.GivenNameValidator(v); err != nil {
			return &ValidationError{Name: "given_name", err: fmt.Errorf(`ent: validator failed for field "User.given_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.FamilyName(); ok {
		if err := user.FamilyNameValidator(v); err != nil {
			return &ValidationError{Name: "family_name", err: fmt.Errorf(`ent: validator failed for field "User.family_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.PhotoURL(); ok {
		if err := user.PhotoURLValidator(v); err != nil {
			return &ValidationError{Name: "photo_url", err: fmt.Errorf(`ent: validator failed for field "User.photo_url": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Origin(); ok {
		if err := user.OriginValidator(v); err != nil {
			return &ValidationError{Name: "origin", err: fmt.Errorf(`ent: validator failed for field "User.origin": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uu.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uu.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if uu.mutation.UserNameCleared() {
		_spec.ClearField(user.FieldUserName, field.TypeString)
	}
	if value, ok := uu.mutation.GivenName(); ok {
		_spec.SetField(user.FieldGivenName, field.TypeString, value)
	}
	if uu.mutation.GivenNameCleared() {
		_spec.ClearField(user.FieldGivenName, field.TypeString)
	}
	if value, ok := uu.mutation.FamilyName(); ok {
		_spec.SetField(user.FieldFamilyName, field.TypeString, value)
	}
	if uu.mutation.FamilyNameCleared() {
		_spec.ClearField(user.FieldFamilyName, field.TypeString)
	}
	if value, ok := uu.mutation.PhotoURL(); ok {
		_spec.SetField(user.FieldPhotoURL, field.TypeString, value)
	}
	if uu.mutation.PhotoURLCleared() {
		_spec.ClearField(user.FieldPhotoURL, field.TypeString)
	}
	if value, ok := uu.mutation.Origin(); ok {
		_spec.SetField(user.FieldOrigin, field.TypeEnum, value)
	}
	if uu.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TodosTable,
			Columns: []string{user.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTodosIDs(); len(nodes) > 0 && !uu.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TodosTable,
			Columns: []string{user.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TodosTable,
			Columns: []string{user.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// ClearEmail clears the value of the "email" field.
func (uuo *UserUpdateOne) ClearEmail() *UserUpdateOne {
	uuo.mutation.ClearEmail()
	return uuo
}

// SetUserName sets the "user_name" field.
func (uuo *UserUpdateOne) SetUserName(s string) *UserUpdateOne {
	uuo.mutation.SetUserName(s)
	return uuo
}

// SetNillableUserName sets the "user_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUserName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUserName(*s)
	}
	return uuo
}

// ClearUserName clears the value of the "user_name" field.
func (uuo *UserUpdateOne) ClearUserName() *UserUpdateOne {
	uuo.mutation.ClearUserName()
	return uuo
}

// SetGivenName sets the "given_name" field.
func (uuo *UserUpdateOne) SetGivenName(s string) *UserUpdateOne {
	uuo.mutation.SetGivenName(s)
	return uuo
}

// SetNillableGivenName sets the "given_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGivenName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetGivenName(*s)
	}
	return uuo
}

// ClearGivenName clears the value of the "given_name" field.
func (uuo *UserUpdateOne) ClearGivenName() *UserUpdateOne {
	uuo.mutation.ClearGivenName()
	return uuo
}

// SetFamilyName sets the "family_name" field.
func (uuo *UserUpdateOne) SetFamilyName(s string) *UserUpdateOne {
	uuo.mutation.SetFamilyName(s)
	return uuo
}

// SetNillableFamilyName sets the "family_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFamilyName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetFamilyName(*s)
	}
	return uuo
}

// ClearFamilyName clears the value of the "family_name" field.
func (uuo *UserUpdateOne) ClearFamilyName() *UserUpdateOne {
	uuo.mutation.ClearFamilyName()
	return uuo
}

// SetPhotoURL sets the "photo_url" field.
func (uuo *UserUpdateOne) SetPhotoURL(s string) *UserUpdateOne {
	uuo.mutation.SetPhotoURL(s)
	return uuo
}

// SetNillablePhotoURL sets the "photo_url" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhotoURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhotoURL(*s)
	}
	return uuo
}

// ClearPhotoURL clears the value of the "photo_url" field.
func (uuo *UserUpdateOne) ClearPhotoURL() *UserUpdateOne {
	uuo.mutation.ClearPhotoURL()
	return uuo
}

// SetOrigin sets the "origin" field.
func (uuo *UserUpdateOne) SetOrigin(e enum.Issuer) *UserUpdateOne {
	uuo.mutation.SetOrigin(e)
	return uuo
}

// AddTodoIDs adds the "todos" edge to the Todo entity by IDs.
func (uuo *UserUpdateOne) AddTodoIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddTodoIDs(ids...)
	return uuo
}

// AddTodos adds the "todos" edges to the Todo entity.
func (uuo *UserUpdateOne) AddTodos(t ...*Todo) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTodoIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearTodos clears all "todos" edges to the Todo entity.
func (uuo *UserUpdateOne) ClearTodos() *UserUpdateOne {
	uuo.mutation.ClearTodos()
	return uuo
}

// RemoveTodoIDs removes the "todos" edge to Todo entities by IDs.
func (uuo *UserUpdateOne) RemoveTodoIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveTodoIDs(ids...)
	return uuo
}

// RemoveTodos removes "todos" edges to Todo entities.
func (uuo *UserUpdateOne) RemoveTodos(t ...*Todo) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTodoIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.UserName(); ok {
		if err := user.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`ent: validator failed for field "User.user_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.GivenName(); ok {
		if err := user.GivenNameValidator(v); err != nil {
			return &ValidationError{Name: "given_name", err: fmt.Errorf(`ent: validator failed for field "User.given_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.FamilyName(); ok {
		if err := user.FamilyNameValidator(v); err != nil {
			return &ValidationError{Name: "family_name", err: fmt.Errorf(`ent: validator failed for field "User.family_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.PhotoURL(); ok {
		if err := user.PhotoURLValidator(v); err != nil {
			return &ValidationError{Name: "photo_url", err: fmt.Errorf(`ent: validator failed for field "User.photo_url": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Origin(); ok {
		if err := user.OriginValidator(v); err != nil {
			return &ValidationError{Name: "origin", err: fmt.Errorf(`ent: validator failed for field "User.origin": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uuo.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uuo.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if uuo.mutation.UserNameCleared() {
		_spec.ClearField(user.FieldUserName, field.TypeString)
	}
	if value, ok := uuo.mutation.GivenName(); ok {
		_spec.SetField(user.FieldGivenName, field.TypeString, value)
	}
	if uuo.mutation.GivenNameCleared() {
		_spec.ClearField(user.FieldGivenName, field.TypeString)
	}
	if value, ok := uuo.mutation.FamilyName(); ok {
		_spec.SetField(user.FieldFamilyName, field.TypeString, value)
	}
	if uuo.mutation.FamilyNameCleared() {
		_spec.ClearField(user.FieldFamilyName, field.TypeString)
	}
	if value, ok := uuo.mutation.PhotoURL(); ok {
		_spec.SetField(user.FieldPhotoURL, field.TypeString, value)
	}
	if uuo.mutation.PhotoURLCleared() {
		_spec.ClearField(user.FieldPhotoURL, field.TypeString)
	}
	if value, ok := uuo.mutation.Origin(); ok {
		_spec.SetField(user.FieldOrigin, field.TypeEnum, value)
	}
	if uuo.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TodosTable,
			Columns: []string{user.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTodosIDs(); len(nodes) > 0 && !uuo.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TodosTable,
			Columns: []string{user.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TodosTable,
			Columns: []string{user.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
