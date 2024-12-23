// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/3xcept/contrib/entproto/internal/entprototest/ent/messagewithints"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// MessageWithInts is the model entity for the MessageWithInts schema.
type MessageWithInts struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Int32s holds the value of the "int32s" field.
	Int32s []int32 `json:"int32s,omitempty"`
	// Int64s holds the value of the "int64s" field.
	Int64s []int64 `json:"int64s,omitempty"`
	// Uint32s holds the value of the "uint32s" field.
	Uint32s []uint32 `json:"uint32s,omitempty"`
	// Uint64s holds the value of the "uint64s" field.
	Uint64s      []uint64 `json:"uint64s,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MessageWithInts) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case messagewithints.FieldInt32s, messagewithints.FieldInt64s, messagewithints.FieldUint32s, messagewithints.FieldUint64s:
			values[i] = new([]byte)
		case messagewithints.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MessageWithInts fields.
func (mwi *MessageWithInts) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case messagewithints.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mwi.ID = int(value.Int64)
		case messagewithints.FieldInt32s:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field int32s", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &mwi.Int32s); err != nil {
					return fmt.Errorf("unmarshal field int32s: %w", err)
				}
			}
		case messagewithints.FieldInt64s:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field int64s", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &mwi.Int64s); err != nil {
					return fmt.Errorf("unmarshal field int64s: %w", err)
				}
			}
		case messagewithints.FieldUint32s:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field uint32s", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &mwi.Uint32s); err != nil {
					return fmt.Errorf("unmarshal field uint32s: %w", err)
				}
			}
		case messagewithints.FieldUint64s:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field uint64s", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &mwi.Uint64s); err != nil {
					return fmt.Errorf("unmarshal field uint64s: %w", err)
				}
			}
		default:
			mwi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MessageWithInts.
// This includes values selected through modifiers, order, etc.
func (mwi *MessageWithInts) Value(name string) (ent.Value, error) {
	return mwi.selectValues.Get(name)
}

// Update returns a builder for updating this MessageWithInts.
// Note that you need to call MessageWithInts.Unwrap() before calling this method if this MessageWithInts
// was returned from a transaction, and the transaction was committed or rolled back.
func (mwi *MessageWithInts) Update() *MessageWithIntsUpdateOne {
	return NewMessageWithIntsClient(mwi.config).UpdateOne(mwi)
}

// Unwrap unwraps the MessageWithInts entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mwi *MessageWithInts) Unwrap() *MessageWithInts {
	_tx, ok := mwi.config.driver.(*txDriver)
	if !ok {
		panic("ent: MessageWithInts is not a transactional entity")
	}
	mwi.config.driver = _tx.drv
	return mwi
}

// String implements the fmt.Stringer.
func (mwi *MessageWithInts) String() string {
	var builder strings.Builder
	builder.WriteString("MessageWithInts(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mwi.ID))
	builder.WriteString("int32s=")
	builder.WriteString(fmt.Sprintf("%v", mwi.Int32s))
	builder.WriteString(", ")
	builder.WriteString("int64s=")
	builder.WriteString(fmt.Sprintf("%v", mwi.Int64s))
	builder.WriteString(", ")
	builder.WriteString("uint32s=")
	builder.WriteString(fmt.Sprintf("%v", mwi.Uint32s))
	builder.WriteString(", ")
	builder.WriteString("uint64s=")
	builder.WriteString(fmt.Sprintf("%v", mwi.Uint64s))
	builder.WriteByte(')')
	return builder.String()
}

// MessageWithIntsSlice is a parsable slice of MessageWithInts.
type MessageWithIntsSlice []*MessageWithInts
