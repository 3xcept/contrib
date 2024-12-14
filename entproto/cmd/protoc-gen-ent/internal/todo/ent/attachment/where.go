// Code generated by ent, DO NOT EDIT.

package attachment

import (
	"github.com/3xcept/contrib/entproto/cmd/protoc-gen-ent/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContainsFold(FieldID, id))
}

// Contents applies equality check predicate on the "contents" field. It's identical to ContentsEQ.
func Contents(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldContents, v))
}

// ContentsEQ applies the EQ predicate on the "contents" field.
func ContentsEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldContents, v))
}

// ContentsNEQ applies the NEQ predicate on the "contents" field.
func ContentsNEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldContents, v))
}

// ContentsIn applies the In predicate on the "contents" field.
func ContentsIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldContents, vs...))
}

// ContentsNotIn applies the NotIn predicate on the "contents" field.
func ContentsNotIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldContents, vs...))
}

// ContentsGT applies the GT predicate on the "contents" field.
func ContentsGT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldContents, v))
}

// ContentsGTE applies the GTE predicate on the "contents" field.
func ContentsGTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldContents, v))
}

// ContentsLT applies the LT predicate on the "contents" field.
func ContentsLT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldContents, v))
}

// ContentsLTE applies the LTE predicate on the "contents" field.
func ContentsLTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldContents, v))
}

// ContentsContains applies the Contains predicate on the "contents" field.
func ContentsContains(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContains(FieldContents, v))
}

// ContentsHasPrefix applies the HasPrefix predicate on the "contents" field.
func ContentsHasPrefix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasPrefix(FieldContents, v))
}

// ContentsHasSuffix applies the HasSuffix predicate on the "contents" field.
func ContentsHasSuffix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasSuffix(FieldContents, v))
}

// ContentsEqualFold applies the EqualFold predicate on the "contents" field.
func ContentsEqualFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEqualFold(FieldContents, v))
}

// ContentsContainsFold applies the ContainsFold predicate on the "contents" field.
func ContentsContainsFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContainsFold(FieldContents, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(sql.NotPredicates(p))
}
