// Code generated by ent, DO NOT EDIT.

package todo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/isutare412/bloated/api/pkg/core/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldUpdateTime, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldUserID, v))
}

// Task applies equality check predicate on the "task" field. It's identical to TaskEQ.
func Task(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldTask, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldUpdateTime, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContainsFold(FieldUserID, v))
}

// TaskEQ applies the EQ predicate on the "task" field.
func TaskEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldTask, v))
}

// TaskNEQ applies the NEQ predicate on the "task" field.
func TaskNEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldTask, v))
}

// TaskIn applies the In predicate on the "task" field.
func TaskIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldTask, vs...))
}

// TaskNotIn applies the NotIn predicate on the "task" field.
func TaskNotIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldTask, vs...))
}

// TaskGT applies the GT predicate on the "task" field.
func TaskGT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldTask, v))
}

// TaskGTE applies the GTE predicate on the "task" field.
func TaskGTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldTask, v))
}

// TaskLT applies the LT predicate on the "task" field.
func TaskLT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldTask, v))
}

// TaskLTE applies the LTE predicate on the "task" field.
func TaskLTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldTask, v))
}

// TaskContains applies the Contains predicate on the "task" field.
func TaskContains(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContains(FieldTask, v))
}

// TaskHasPrefix applies the HasPrefix predicate on the "task" field.
func TaskHasPrefix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasPrefix(FieldTask, v))
}

// TaskHasSuffix applies the HasSuffix predicate on the "task" field.
func TaskHasSuffix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasSuffix(FieldTask, v))
}

// TaskEqualFold applies the EqualFold predicate on the "task" field.
func TaskEqualFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEqualFold(FieldTask, v))
}

// TaskContainsFold applies the ContainsFold predicate on the "task" field.
func TaskContainsFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContainsFold(FieldTask, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		p(s.Not())
	})
}