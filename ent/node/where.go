// Code generated by ent, DO NOT EDIT.

package node

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/davidroman0O/go-tempolite/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Node {
	return predicate.Node(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Node {
	return predicate.Node(sql.FieldContainsFold(FieldID, id))
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHandlerTask applies the HasEdge predicate on the "handler_task" edge.
func HasHandlerTask() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, HandlerTaskTable, HandlerTaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHandlerTaskWith applies the HasEdge predicate on the "handler_task" edge with a given conditions (other predicates).
func HasHandlerTaskWith(preds ...predicate.HandlerTask) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := newHandlerTaskStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSagaStepTask applies the HasEdge predicate on the "saga_step_task" edge.
func HasSagaStepTask() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SagaStepTaskTable, SagaStepTaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSagaStepTaskWith applies the HasEdge predicate on the "saga_step_task" edge with a given conditions (other predicates).
func HasSagaStepTaskWith(preds ...predicate.SagaTask) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := newSagaStepTaskStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSideEffectTask applies the HasEdge predicate on the "side_effect_task" edge.
func HasSideEffectTask() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SideEffectTaskTable, SideEffectTaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSideEffectTaskWith applies the HasEdge predicate on the "side_effect_task" edge with a given conditions (other predicates).
func HasSideEffectTaskWith(preds ...predicate.SideEffectTask) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := newSideEffectTaskStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCompensationTask applies the HasEdge predicate on the "compensation_task" edge.
func HasCompensationTask() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, CompensationTaskTable, CompensationTaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompensationTaskWith applies the HasEdge predicate on the "compensation_task" edge with a given conditions (other predicates).
func HasCompensationTaskWith(preds ...predicate.CompensationTask) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := newCompensationTaskStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Node) predicate.Node {
	return predicate.Node(sql.NotPredicates(p))
}
