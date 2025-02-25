// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/chargesetting"
)

// ChargeSetting is the model entity for the ChargeSetting schema.
type ChargeSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Enabled holds the value of the "enabled" field.
	Enabled bool `json:"enabled,omitempty"`
	// ChargeStartThreshold holds the value of the "charge_start_threshold" field.
	ChargeStartThreshold int `json:"charge_start_threshold,omitempty"`
	// PowerUsageIncreaseThreshold holds the value of the "power_usage_increase_threshold" field.
	PowerUsageIncreaseThreshold int `json:"power_usage_increase_threshold,omitempty"`
	// PowerUsageDecreaseThreshold holds the value of the "power_usage_decrease_threshold" field.
	PowerUsageDecreaseThreshold int `json:"power_usage_decrease_threshold,omitempty"`
	// UpdateInterval holds the value of the "update_interval" field.
	UpdateInterval int `json:"update_interval,omitempty"`
	// MinChargeThreshold holds the value of the "min_charge_threshold" field.
	MinChargeThreshold int `json:"min_charge_threshold,omitempty"`
	// MinChargeTimeRangeStart holds the value of the "min_charge_time_range_start" field.
	MinChargeTimeRangeStart string `json:"min_charge_time_range_start,omitempty"`
	// MinChargeTimeRangeEnd holds the value of the "min_charge_time_range_end" field.
	MinChargeTimeRangeEnd string `json:"min_charge_time_range_end,omitempty"`
	// MinChargeAmperage holds the value of the "min_charge_amperage" field.
	MinChargeAmperage int `json:"min_charge_amperage,omitempty"`
	selectValues      sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChargeSetting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case chargesetting.FieldEnabled:
			values[i] = new(sql.NullBool)
		case chargesetting.FieldID, chargesetting.FieldChargeStartThreshold, chargesetting.FieldPowerUsageIncreaseThreshold, chargesetting.FieldPowerUsageDecreaseThreshold, chargesetting.FieldUpdateInterval, chargesetting.FieldMinChargeThreshold, chargesetting.FieldMinChargeAmperage:
			values[i] = new(sql.NullInt64)
		case chargesetting.FieldMinChargeTimeRangeStart, chargesetting.FieldMinChargeTimeRangeEnd:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChargeSetting fields.
func (cs *ChargeSetting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chargesetting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cs.ID = int(value.Int64)
		case chargesetting.FieldEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enabled", values[i])
			} else if value.Valid {
				cs.Enabled = value.Bool
			}
		case chargesetting.FieldChargeStartThreshold:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field charge_start_threshold", values[i])
			} else if value.Valid {
				cs.ChargeStartThreshold = int(value.Int64)
			}
		case chargesetting.FieldPowerUsageIncreaseThreshold:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field power_usage_increase_threshold", values[i])
			} else if value.Valid {
				cs.PowerUsageIncreaseThreshold = int(value.Int64)
			}
		case chargesetting.FieldPowerUsageDecreaseThreshold:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field power_usage_decrease_threshold", values[i])
			} else if value.Valid {
				cs.PowerUsageDecreaseThreshold = int(value.Int64)
			}
		case chargesetting.FieldUpdateInterval:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_interval", values[i])
			} else if value.Valid {
				cs.UpdateInterval = int(value.Int64)
			}
		case chargesetting.FieldMinChargeThreshold:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field min_charge_threshold", values[i])
			} else if value.Valid {
				cs.MinChargeThreshold = int(value.Int64)
			}
		case chargesetting.FieldMinChargeTimeRangeStart:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field min_charge_time_range_start", values[i])
			} else if value.Valid {
				cs.MinChargeTimeRangeStart = value.String
			}
		case chargesetting.FieldMinChargeTimeRangeEnd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field min_charge_time_range_end", values[i])
			} else if value.Valid {
				cs.MinChargeTimeRangeEnd = value.String
			}
		case chargesetting.FieldMinChargeAmperage:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field min_charge_amperage", values[i])
			} else if value.Valid {
				cs.MinChargeAmperage = int(value.Int64)
			}
		default:
			cs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ChargeSetting.
// This includes values selected through modifiers, order, etc.
func (cs *ChargeSetting) Value(name string) (ent.Value, error) {
	return cs.selectValues.Get(name)
}

// Update returns a builder for updating this ChargeSetting.
// Note that you need to call ChargeSetting.Unwrap() before calling this method if this ChargeSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (cs *ChargeSetting) Update() *ChargeSettingUpdateOne {
	return NewChargeSettingClient(cs.config).UpdateOne(cs)
}

// Unwrap unwraps the ChargeSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cs *ChargeSetting) Unwrap() *ChargeSetting {
	_tx, ok := cs.config.driver.(*txDriver)
	if !ok {
		panic("ent: ChargeSetting is not a transactional entity")
	}
	cs.config.driver = _tx.drv
	return cs
}

// String implements the fmt.Stringer.
func (cs *ChargeSetting) String() string {
	var builder strings.Builder
	builder.WriteString("ChargeSetting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cs.ID))
	builder.WriteString("enabled=")
	builder.WriteString(fmt.Sprintf("%v", cs.Enabled))
	builder.WriteString(", ")
	builder.WriteString("charge_start_threshold=")
	builder.WriteString(fmt.Sprintf("%v", cs.ChargeStartThreshold))
	builder.WriteString(", ")
	builder.WriteString("power_usage_increase_threshold=")
	builder.WriteString(fmt.Sprintf("%v", cs.PowerUsageIncreaseThreshold))
	builder.WriteString(", ")
	builder.WriteString("power_usage_decrease_threshold=")
	builder.WriteString(fmt.Sprintf("%v", cs.PowerUsageDecreaseThreshold))
	builder.WriteString(", ")
	builder.WriteString("update_interval=")
	builder.WriteString(fmt.Sprintf("%v", cs.UpdateInterval))
	builder.WriteString(", ")
	builder.WriteString("min_charge_threshold=")
	builder.WriteString(fmt.Sprintf("%v", cs.MinChargeThreshold))
	builder.WriteString(", ")
	builder.WriteString("min_charge_time_range_start=")
	builder.WriteString(cs.MinChargeTimeRangeStart)
	builder.WriteString(", ")
	builder.WriteString("min_charge_time_range_end=")
	builder.WriteString(cs.MinChargeTimeRangeEnd)
	builder.WriteString(", ")
	builder.WriteString("min_charge_amperage=")
	builder.WriteString(fmt.Sprintf("%v", cs.MinChargeAmperage))
	builder.WriteByte(')')
	return builder.String()
}

// ChargeSettings is a parsable slice of ChargeSetting.
type ChargeSettings []*ChargeSetting
