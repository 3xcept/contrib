// Code generated by ent, DO NOT EDIT.

package skipedgeexample

import (
	"github.com/3xcept/contrib/entproto/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(sql.FieldLTE(FieldID, id))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SkipEdgeExample) predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SkipEdgeExample) predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SkipEdgeExample) predicate.SkipEdgeExample {
	return predicate.SkipEdgeExample(sql.NotPredicates(p))
}
