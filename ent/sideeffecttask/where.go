// Code generated by ent, DO NOT EDIT.

package sideeffecttask

import (
	"entgo.io/ent/dialect/sql"
	"github.com/davidroman0O/go-tempolite/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SideEffectTask {
	return predicate.SideEffectTask(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SideEffectTask {
	return predicate.SideEffectTask(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SideEffectTask {
	return predicate.SideEffectTask(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SideEffectTask {
	return predicate.SideEffectTask(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SideEffectTask {
	return predicate.SideEffectTask(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SideEffectTask {
	return predicate.SideEffectTask(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SideEffectTask {
	return predicate.SideEffectTask(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SideEffectTask {
	return predicate.SideEffectTask(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SideEffectTask {
	return predicate.SideEffectTask(sql.FieldLTE(FieldID, id))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SideEffectTask) predicate.SideEffectTask {
	return predicate.SideEffectTask(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SideEffectTask) predicate.SideEffectTask {
	return predicate.SideEffectTask(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SideEffectTask) predicate.SideEffectTask {
	return predicate.SideEffectTask(sql.NotPredicates(p))
}
