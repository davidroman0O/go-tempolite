// Code generated by ent, DO NOT EDIT.

package compensationtask

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/davidroman0O/go-tempolite/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CompensationTask {
	return predicate.CompensationTask(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CompensationTask {
	return predicate.CompensationTask(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CompensationTask {
	return predicate.CompensationTask(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CompensationTask {
	return predicate.CompensationTask(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CompensationTask {
	return predicate.CompensationTask(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CompensationTask {
	return predicate.CompensationTask(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CompensationTask {
	return predicate.CompensationTask(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CompensationTask {
	return predicate.CompensationTask(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CompensationTask {
	return predicate.CompensationTask(sql.FieldLTE(FieldID, id))
}

// HasNode applies the HasEdge predicate on the "node" edge.
func HasNode() predicate.CompensationTask {
	return predicate.CompensationTask(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, NodeTable, NodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodeWith applies the HasEdge predicate on the "node" edge with a given conditions (other predicates).
func HasNodeWith(preds ...predicate.Node) predicate.CompensationTask {
	return predicate.CompensationTask(func(s *sql.Selector) {
		step := newNodeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CompensationTask) predicate.CompensationTask {
	return predicate.CompensationTask(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CompensationTask) predicate.CompensationTask {
	return predicate.CompensationTask(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CompensationTask) predicate.CompensationTask {
	return predicate.CompensationTask(sql.NotPredicates(p))
}