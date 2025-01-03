// Code generated by ent, DO NOT EDIT.

package oastypes

import (
	"time"

	"github.com/3xcept/contrib/entoas/internal/oastypes/predicate"
	"github.com/3xcept/contrib/entoas/internal/oastypes/schema"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldID, id))
}

// Int applies equality check predicate on the "int" field. It's identical to IntEQ.
func Int(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldInt, v))
}

// Int8 applies equality check predicate on the "int8" field. It's identical to Int8EQ.
func Int8(v int8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldInt8, v))
}

// Int16 applies equality check predicate on the "int16" field. It's identical to Int16EQ.
func Int16(v int16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldInt16, v))
}

// Int32 applies equality check predicate on the "int32" field. It's identical to Int32EQ.
func Int32(v int32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldInt32, v))
}

// Int64 applies equality check predicate on the "int64" field. It's identical to Int64EQ.
func Int64(v int64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldInt64, v))
}

// Uint applies equality check predicate on the "uint" field. It's identical to UintEQ.
func Uint(v uint) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldUint, v))
}

// Uint8 applies equality check predicate on the "uint8" field. It's identical to Uint8EQ.
func Uint8(v uint8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldUint8, v))
}

// Uint16 applies equality check predicate on the "uint16" field. It's identical to Uint16EQ.
func Uint16(v uint16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldUint16, v))
}

// Uint32 applies equality check predicate on the "uint32" field. It's identical to Uint32EQ.
func Uint32(v uint32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldUint32, v))
}

// Uint64 applies equality check predicate on the "uint64" field. It's identical to Uint64EQ.
func Uint64(v uint64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldUint64, v))
}

// Float32 applies equality check predicate on the "float32" field. It's identical to Float32EQ.
func Float32(v float32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldFloat32, v))
}

// Float64 applies equality check predicate on the "float64" field. It's identical to Float64EQ.
func Float64(v float64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldFloat64, v))
}

// StringField applies equality check predicate on the "string_field" field. It's identical to StringFieldEQ.
func StringField(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldStringField, v))
}

// Bool applies equality check predicate on the "bool" field. It's identical to BoolEQ.
func Bool(v bool) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldBool, v))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldUUID, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v time.Time) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldTime, v))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldText, v))
}

// Bytes applies equality check predicate on the "bytes" field. It's identical to BytesEQ.
func Bytes(v []byte) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldBytes, v))
}

// Other applies equality check predicate on the "other" field. It's identical to OtherEQ.
func Other(v *schema.Link) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldOther, v))
}

// Optional applies equality check predicate on the "optional" field. It's identical to OptionalEQ.
func Optional(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldOptional, v))
}

// Nillable applies equality check predicate on the "nillable" field. It's identical to NillableEQ.
func Nillable(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldNillable, v))
}

// OptionalAndNillable applies equality check predicate on the "optional_and_nillable" field. It's identical to OptionalAndNillableEQ.
func OptionalAndNillable(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldOptionalAndNillable, v))
}

// IntEQ applies the EQ predicate on the "int" field.
func IntEQ(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldInt, v))
}

// IntNEQ applies the NEQ predicate on the "int" field.
func IntNEQ(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldInt, v))
}

// IntIn applies the In predicate on the "int" field.
func IntIn(vs ...int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldInt, vs...))
}

// IntNotIn applies the NotIn predicate on the "int" field.
func IntNotIn(vs ...int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldInt, vs...))
}

// IntGT applies the GT predicate on the "int" field.
func IntGT(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldInt, v))
}

// IntGTE applies the GTE predicate on the "int" field.
func IntGTE(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldInt, v))
}

// IntLT applies the LT predicate on the "int" field.
func IntLT(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldInt, v))
}

// IntLTE applies the LTE predicate on the "int" field.
func IntLTE(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldInt, v))
}

// Int8EQ applies the EQ predicate on the "int8" field.
func Int8EQ(v int8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldInt8, v))
}

// Int8NEQ applies the NEQ predicate on the "int8" field.
func Int8NEQ(v int8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldInt8, v))
}

// Int8In applies the In predicate on the "int8" field.
func Int8In(vs ...int8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldInt8, vs...))
}

// Int8NotIn applies the NotIn predicate on the "int8" field.
func Int8NotIn(vs ...int8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldInt8, vs...))
}

// Int8GT applies the GT predicate on the "int8" field.
func Int8GT(v int8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldInt8, v))
}

// Int8GTE applies the GTE predicate on the "int8" field.
func Int8GTE(v int8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldInt8, v))
}

// Int8LT applies the LT predicate on the "int8" field.
func Int8LT(v int8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldInt8, v))
}

// Int8LTE applies the LTE predicate on the "int8" field.
func Int8LTE(v int8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldInt8, v))
}

// Int16EQ applies the EQ predicate on the "int16" field.
func Int16EQ(v int16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldInt16, v))
}

// Int16NEQ applies the NEQ predicate on the "int16" field.
func Int16NEQ(v int16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldInt16, v))
}

// Int16In applies the In predicate on the "int16" field.
func Int16In(vs ...int16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldInt16, vs...))
}

// Int16NotIn applies the NotIn predicate on the "int16" field.
func Int16NotIn(vs ...int16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldInt16, vs...))
}

// Int16GT applies the GT predicate on the "int16" field.
func Int16GT(v int16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldInt16, v))
}

// Int16GTE applies the GTE predicate on the "int16" field.
func Int16GTE(v int16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldInt16, v))
}

// Int16LT applies the LT predicate on the "int16" field.
func Int16LT(v int16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldInt16, v))
}

// Int16LTE applies the LTE predicate on the "int16" field.
func Int16LTE(v int16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldInt16, v))
}

// Int32EQ applies the EQ predicate on the "int32" field.
func Int32EQ(v int32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldInt32, v))
}

// Int32NEQ applies the NEQ predicate on the "int32" field.
func Int32NEQ(v int32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldInt32, v))
}

// Int32In applies the In predicate on the "int32" field.
func Int32In(vs ...int32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldInt32, vs...))
}

// Int32NotIn applies the NotIn predicate on the "int32" field.
func Int32NotIn(vs ...int32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldInt32, vs...))
}

// Int32GT applies the GT predicate on the "int32" field.
func Int32GT(v int32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldInt32, v))
}

// Int32GTE applies the GTE predicate on the "int32" field.
func Int32GTE(v int32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldInt32, v))
}

// Int32LT applies the LT predicate on the "int32" field.
func Int32LT(v int32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldInt32, v))
}

// Int32LTE applies the LTE predicate on the "int32" field.
func Int32LTE(v int32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldInt32, v))
}

// Int64EQ applies the EQ predicate on the "int64" field.
func Int64EQ(v int64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldInt64, v))
}

// Int64NEQ applies the NEQ predicate on the "int64" field.
func Int64NEQ(v int64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldInt64, v))
}

// Int64In applies the In predicate on the "int64" field.
func Int64In(vs ...int64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldInt64, vs...))
}

// Int64NotIn applies the NotIn predicate on the "int64" field.
func Int64NotIn(vs ...int64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldInt64, vs...))
}

// Int64GT applies the GT predicate on the "int64" field.
func Int64GT(v int64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldInt64, v))
}

// Int64GTE applies the GTE predicate on the "int64" field.
func Int64GTE(v int64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldInt64, v))
}

// Int64LT applies the LT predicate on the "int64" field.
func Int64LT(v int64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldInt64, v))
}

// Int64LTE applies the LTE predicate on the "int64" field.
func Int64LTE(v int64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldInt64, v))
}

// UintEQ applies the EQ predicate on the "uint" field.
func UintEQ(v uint) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldUint, v))
}

// UintNEQ applies the NEQ predicate on the "uint" field.
func UintNEQ(v uint) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldUint, v))
}

// UintIn applies the In predicate on the "uint" field.
func UintIn(vs ...uint) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldUint, vs...))
}

// UintNotIn applies the NotIn predicate on the "uint" field.
func UintNotIn(vs ...uint) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldUint, vs...))
}

// UintGT applies the GT predicate on the "uint" field.
func UintGT(v uint) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldUint, v))
}

// UintGTE applies the GTE predicate on the "uint" field.
func UintGTE(v uint) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldUint, v))
}

// UintLT applies the LT predicate on the "uint" field.
func UintLT(v uint) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldUint, v))
}

// UintLTE applies the LTE predicate on the "uint" field.
func UintLTE(v uint) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldUint, v))
}

// Uint8EQ applies the EQ predicate on the "uint8" field.
func Uint8EQ(v uint8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldUint8, v))
}

// Uint8NEQ applies the NEQ predicate on the "uint8" field.
func Uint8NEQ(v uint8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldUint8, v))
}

// Uint8In applies the In predicate on the "uint8" field.
func Uint8In(vs ...uint8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldUint8, vs...))
}

// Uint8NotIn applies the NotIn predicate on the "uint8" field.
func Uint8NotIn(vs ...uint8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldUint8, vs...))
}

// Uint8GT applies the GT predicate on the "uint8" field.
func Uint8GT(v uint8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldUint8, v))
}

// Uint8GTE applies the GTE predicate on the "uint8" field.
func Uint8GTE(v uint8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldUint8, v))
}

// Uint8LT applies the LT predicate on the "uint8" field.
func Uint8LT(v uint8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldUint8, v))
}

// Uint8LTE applies the LTE predicate on the "uint8" field.
func Uint8LTE(v uint8) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldUint8, v))
}

// Uint16EQ applies the EQ predicate on the "uint16" field.
func Uint16EQ(v uint16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldUint16, v))
}

// Uint16NEQ applies the NEQ predicate on the "uint16" field.
func Uint16NEQ(v uint16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldUint16, v))
}

// Uint16In applies the In predicate on the "uint16" field.
func Uint16In(vs ...uint16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldUint16, vs...))
}

// Uint16NotIn applies the NotIn predicate on the "uint16" field.
func Uint16NotIn(vs ...uint16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldUint16, vs...))
}

// Uint16GT applies the GT predicate on the "uint16" field.
func Uint16GT(v uint16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldUint16, v))
}

// Uint16GTE applies the GTE predicate on the "uint16" field.
func Uint16GTE(v uint16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldUint16, v))
}

// Uint16LT applies the LT predicate on the "uint16" field.
func Uint16LT(v uint16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldUint16, v))
}

// Uint16LTE applies the LTE predicate on the "uint16" field.
func Uint16LTE(v uint16) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldUint16, v))
}

// Uint32EQ applies the EQ predicate on the "uint32" field.
func Uint32EQ(v uint32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldUint32, v))
}

// Uint32NEQ applies the NEQ predicate on the "uint32" field.
func Uint32NEQ(v uint32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldUint32, v))
}

// Uint32In applies the In predicate on the "uint32" field.
func Uint32In(vs ...uint32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldUint32, vs...))
}

// Uint32NotIn applies the NotIn predicate on the "uint32" field.
func Uint32NotIn(vs ...uint32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldUint32, vs...))
}

// Uint32GT applies the GT predicate on the "uint32" field.
func Uint32GT(v uint32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldUint32, v))
}

// Uint32GTE applies the GTE predicate on the "uint32" field.
func Uint32GTE(v uint32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldUint32, v))
}

// Uint32LT applies the LT predicate on the "uint32" field.
func Uint32LT(v uint32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldUint32, v))
}

// Uint32LTE applies the LTE predicate on the "uint32" field.
func Uint32LTE(v uint32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldUint32, v))
}

// Uint64EQ applies the EQ predicate on the "uint64" field.
func Uint64EQ(v uint64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldUint64, v))
}

// Uint64NEQ applies the NEQ predicate on the "uint64" field.
func Uint64NEQ(v uint64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldUint64, v))
}

// Uint64In applies the In predicate on the "uint64" field.
func Uint64In(vs ...uint64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldUint64, vs...))
}

// Uint64NotIn applies the NotIn predicate on the "uint64" field.
func Uint64NotIn(vs ...uint64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldUint64, vs...))
}

// Uint64GT applies the GT predicate on the "uint64" field.
func Uint64GT(v uint64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldUint64, v))
}

// Uint64GTE applies the GTE predicate on the "uint64" field.
func Uint64GTE(v uint64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldUint64, v))
}

// Uint64LT applies the LT predicate on the "uint64" field.
func Uint64LT(v uint64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldUint64, v))
}

// Uint64LTE applies the LTE predicate on the "uint64" field.
func Uint64LTE(v uint64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldUint64, v))
}

// Float32EQ applies the EQ predicate on the "float32" field.
func Float32EQ(v float32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldFloat32, v))
}

// Float32NEQ applies the NEQ predicate on the "float32" field.
func Float32NEQ(v float32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldFloat32, v))
}

// Float32In applies the In predicate on the "float32" field.
func Float32In(vs ...float32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldFloat32, vs...))
}

// Float32NotIn applies the NotIn predicate on the "float32" field.
func Float32NotIn(vs ...float32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldFloat32, vs...))
}

// Float32GT applies the GT predicate on the "float32" field.
func Float32GT(v float32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldFloat32, v))
}

// Float32GTE applies the GTE predicate on the "float32" field.
func Float32GTE(v float32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldFloat32, v))
}

// Float32LT applies the LT predicate on the "float32" field.
func Float32LT(v float32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldFloat32, v))
}

// Float32LTE applies the LTE predicate on the "float32" field.
func Float32LTE(v float32) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldFloat32, v))
}

// Float64EQ applies the EQ predicate on the "float64" field.
func Float64EQ(v float64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldFloat64, v))
}

// Float64NEQ applies the NEQ predicate on the "float64" field.
func Float64NEQ(v float64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldFloat64, v))
}

// Float64In applies the In predicate on the "float64" field.
func Float64In(vs ...float64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldFloat64, vs...))
}

// Float64NotIn applies the NotIn predicate on the "float64" field.
func Float64NotIn(vs ...float64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldFloat64, vs...))
}

// Float64GT applies the GT predicate on the "float64" field.
func Float64GT(v float64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldFloat64, v))
}

// Float64GTE applies the GTE predicate on the "float64" field.
func Float64GTE(v float64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldFloat64, v))
}

// Float64LT applies the LT predicate on the "float64" field.
func Float64LT(v float64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldFloat64, v))
}

// Float64LTE applies the LTE predicate on the "float64" field.
func Float64LTE(v float64) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldFloat64, v))
}

// StringFieldEQ applies the EQ predicate on the "string_field" field.
func StringFieldEQ(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldStringField, v))
}

// StringFieldNEQ applies the NEQ predicate on the "string_field" field.
func StringFieldNEQ(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldStringField, v))
}

// StringFieldIn applies the In predicate on the "string_field" field.
func StringFieldIn(vs ...string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldStringField, vs...))
}

// StringFieldNotIn applies the NotIn predicate on the "string_field" field.
func StringFieldNotIn(vs ...string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldStringField, vs...))
}

// StringFieldGT applies the GT predicate on the "string_field" field.
func StringFieldGT(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldStringField, v))
}

// StringFieldGTE applies the GTE predicate on the "string_field" field.
func StringFieldGTE(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldStringField, v))
}

// StringFieldLT applies the LT predicate on the "string_field" field.
func StringFieldLT(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldStringField, v))
}

// StringFieldLTE applies the LTE predicate on the "string_field" field.
func StringFieldLTE(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldStringField, v))
}

// StringFieldContains applies the Contains predicate on the "string_field" field.
func StringFieldContains(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldContains(FieldStringField, v))
}

// StringFieldHasPrefix applies the HasPrefix predicate on the "string_field" field.
func StringFieldHasPrefix(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldHasPrefix(FieldStringField, v))
}

// StringFieldHasSuffix applies the HasSuffix predicate on the "string_field" field.
func StringFieldHasSuffix(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldHasSuffix(FieldStringField, v))
}

// StringFieldEqualFold applies the EqualFold predicate on the "string_field" field.
func StringFieldEqualFold(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEqualFold(FieldStringField, v))
}

// StringFieldContainsFold applies the ContainsFold predicate on the "string_field" field.
func StringFieldContainsFold(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldContainsFold(FieldStringField, v))
}

// BoolEQ applies the EQ predicate on the "bool" field.
func BoolEQ(v bool) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldBool, v))
}

// BoolNEQ applies the NEQ predicate on the "bool" field.
func BoolNEQ(v bool) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldBool, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldUUID, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v time.Time) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v time.Time) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...time.Time) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...time.Time) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v time.Time) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v time.Time) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v time.Time) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v time.Time) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldTime, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldContainsFold(FieldText, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldState, vs...))
}

// BytesEQ applies the EQ predicate on the "bytes" field.
func BytesEQ(v []byte) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldBytes, v))
}

// BytesNEQ applies the NEQ predicate on the "bytes" field.
func BytesNEQ(v []byte) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldBytes, v))
}

// BytesIn applies the In predicate on the "bytes" field.
func BytesIn(vs ...[]byte) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldBytes, vs...))
}

// BytesNotIn applies the NotIn predicate on the "bytes" field.
func BytesNotIn(vs ...[]byte) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldBytes, vs...))
}

// BytesGT applies the GT predicate on the "bytes" field.
func BytesGT(v []byte) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldBytes, v))
}

// BytesGTE applies the GTE predicate on the "bytes" field.
func BytesGTE(v []byte) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldBytes, v))
}

// BytesLT applies the LT predicate on the "bytes" field.
func BytesLT(v []byte) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldBytes, v))
}

// BytesLTE applies the LTE predicate on the "bytes" field.
func BytesLTE(v []byte) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldBytes, v))
}

// OtherEQ applies the EQ predicate on the "other" field.
func OtherEQ(v *schema.Link) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldOther, v))
}

// OtherNEQ applies the NEQ predicate on the "other" field.
func OtherNEQ(v *schema.Link) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldOther, v))
}

// OtherIn applies the In predicate on the "other" field.
func OtherIn(vs ...*schema.Link) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldOther, vs...))
}

// OtherNotIn applies the NotIn predicate on the "other" field.
func OtherNotIn(vs ...*schema.Link) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldOther, vs...))
}

// OtherGT applies the GT predicate on the "other" field.
func OtherGT(v *schema.Link) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldOther, v))
}

// OtherGTE applies the GTE predicate on the "other" field.
func OtherGTE(v *schema.Link) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldOther, v))
}

// OtherLT applies the LT predicate on the "other" field.
func OtherLT(v *schema.Link) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldOther, v))
}

// OtherLTE applies the LTE predicate on the "other" field.
func OtherLTE(v *schema.Link) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldOther, v))
}

// OptionalEQ applies the EQ predicate on the "optional" field.
func OptionalEQ(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldOptional, v))
}

// OptionalNEQ applies the NEQ predicate on the "optional" field.
func OptionalNEQ(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldOptional, v))
}

// OptionalIn applies the In predicate on the "optional" field.
func OptionalIn(vs ...int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldOptional, vs...))
}

// OptionalNotIn applies the NotIn predicate on the "optional" field.
func OptionalNotIn(vs ...int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldOptional, vs...))
}

// OptionalGT applies the GT predicate on the "optional" field.
func OptionalGT(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldOptional, v))
}

// OptionalGTE applies the GTE predicate on the "optional" field.
func OptionalGTE(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldOptional, v))
}

// OptionalLT applies the LT predicate on the "optional" field.
func OptionalLT(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldOptional, v))
}

// OptionalLTE applies the LTE predicate on the "optional" field.
func OptionalLTE(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldOptional, v))
}

// OptionalIsNil applies the IsNil predicate on the "optional" field.
func OptionalIsNil() predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIsNull(FieldOptional))
}

// OptionalNotNil applies the NotNil predicate on the "optional" field.
func OptionalNotNil() predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotNull(FieldOptional))
}

// NillableEQ applies the EQ predicate on the "nillable" field.
func NillableEQ(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldNillable, v))
}

// NillableNEQ applies the NEQ predicate on the "nillable" field.
func NillableNEQ(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldNillable, v))
}

// NillableIn applies the In predicate on the "nillable" field.
func NillableIn(vs ...int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldNillable, vs...))
}

// NillableNotIn applies the NotIn predicate on the "nillable" field.
func NillableNotIn(vs ...int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldNillable, vs...))
}

// NillableGT applies the GT predicate on the "nillable" field.
func NillableGT(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldNillable, v))
}

// NillableGTE applies the GTE predicate on the "nillable" field.
func NillableGTE(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldNillable, v))
}

// NillableLT applies the LT predicate on the "nillable" field.
func NillableLT(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldNillable, v))
}

// NillableLTE applies the LTE predicate on the "nillable" field.
func NillableLTE(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldNillable, v))
}

// OptionalAndNillableEQ applies the EQ predicate on the "optional_and_nillable" field.
func OptionalAndNillableEQ(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldEQ(FieldOptionalAndNillable, v))
}

// OptionalAndNillableNEQ applies the NEQ predicate on the "optional_and_nillable" field.
func OptionalAndNillableNEQ(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNEQ(FieldOptionalAndNillable, v))
}

// OptionalAndNillableIn applies the In predicate on the "optional_and_nillable" field.
func OptionalAndNillableIn(vs ...int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIn(FieldOptionalAndNillable, vs...))
}

// OptionalAndNillableNotIn applies the NotIn predicate on the "optional_and_nillable" field.
func OptionalAndNillableNotIn(vs ...int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotIn(FieldOptionalAndNillable, vs...))
}

// OptionalAndNillableGT applies the GT predicate on the "optional_and_nillable" field.
func OptionalAndNillableGT(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGT(FieldOptionalAndNillable, v))
}

// OptionalAndNillableGTE applies the GTE predicate on the "optional_and_nillable" field.
func OptionalAndNillableGTE(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldGTE(FieldOptionalAndNillable, v))
}

// OptionalAndNillableLT applies the LT predicate on the "optional_and_nillable" field.
func OptionalAndNillableLT(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLT(FieldOptionalAndNillable, v))
}

// OptionalAndNillableLTE applies the LTE predicate on the "optional_and_nillable" field.
func OptionalAndNillableLTE(v int) predicate.OASTypes {
	return predicate.OASTypes(sql.FieldLTE(FieldOptionalAndNillable, v))
}

// OptionalAndNillableIsNil applies the IsNil predicate on the "optional_and_nillable" field.
func OptionalAndNillableIsNil() predicate.OASTypes {
	return predicate.OASTypes(sql.FieldIsNull(FieldOptionalAndNillable))
}

// OptionalAndNillableNotNil applies the NotNil predicate on the "optional_and_nillable" field.
func OptionalAndNillableNotNil() predicate.OASTypes {
	return predicate.OASTypes(sql.FieldNotNull(FieldOptionalAndNillable))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OASTypes) predicate.OASTypes {
	return predicate.OASTypes(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OASTypes) predicate.OASTypes {
	return predicate.OASTypes(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OASTypes) predicate.OASTypes {
	return predicate.OASTypes(sql.NotPredicates(p))
}
