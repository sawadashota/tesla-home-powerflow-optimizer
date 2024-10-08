// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/chargecommandhistory"
)

// ChargeCommandHistory is the model entity for the ChargeCommandHistory schema.
type ChargeCommandHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Vin holds the value of the "vin" field.
	Vin string `json:"vin,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation string `json:"operation,omitempty"`
	// Amps holds the value of the "amps" field.
	Amps int `json:"amps,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp    time.Time `json:"timestamp,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChargeCommandHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case chargecommandhistory.FieldID, chargecommandhistory.FieldAmps:
			values[i] = new(sql.NullInt64)
		case chargecommandhistory.FieldVin, chargecommandhistory.FieldOperation:
			values[i] = new(sql.NullString)
		case chargecommandhistory.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChargeCommandHistory fields.
func (cch *ChargeCommandHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chargecommandhistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cch.ID = int(value.Int64)
		case chargecommandhistory.FieldVin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vin", values[i])
			} else if value.Valid {
				cch.Vin = value.String
			}
		case chargecommandhistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				cch.Operation = value.String
			}
		case chargecommandhistory.FieldAmps:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amps", values[i])
			} else if value.Valid {
				cch.Amps = int(value.Int64)
			}
		case chargecommandhistory.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				cch.Timestamp = value.Time
			}
		default:
			cch.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ChargeCommandHistory.
// This includes values selected through modifiers, order, etc.
func (cch *ChargeCommandHistory) Value(name string) (ent.Value, error) {
	return cch.selectValues.Get(name)
}

// Update returns a builder for updating this ChargeCommandHistory.
// Note that you need to call ChargeCommandHistory.Unwrap() before calling this method if this ChargeCommandHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (cch *ChargeCommandHistory) Update() *ChargeCommandHistoryUpdateOne {
	return NewChargeCommandHistoryClient(cch.config).UpdateOne(cch)
}

// Unwrap unwraps the ChargeCommandHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cch *ChargeCommandHistory) Unwrap() *ChargeCommandHistory {
	_tx, ok := cch.config.driver.(*txDriver)
	if !ok {
		panic("ent: ChargeCommandHistory is not a transactional entity")
	}
	cch.config.driver = _tx.drv
	return cch
}

// String implements the fmt.Stringer.
func (cch *ChargeCommandHistory) String() string {
	var builder strings.Builder
	builder.WriteString("ChargeCommandHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cch.ID))
	builder.WriteString("vin=")
	builder.WriteString(cch.Vin)
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(cch.Operation)
	builder.WriteString(", ")
	builder.WriteString("amps=")
	builder.WriteString(fmt.Sprintf("%v", cch.Amps))
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(cch.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ChargeCommandHistories is a parsable slice of ChargeCommandHistory.
type ChargeCommandHistories []*ChargeCommandHistory
