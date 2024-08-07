// Code generated by ent, DO NOT EDIT.

package chargecommandhistory

import (
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldLTE(FieldID, id))
}

// Vin applies equality check predicate on the "vin" field. It's identical to VinEQ.
func Vin(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldEQ(FieldVin, v))
}

// Operation applies equality check predicate on the "operation" field. It's identical to OperationEQ.
func Operation(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldEQ(FieldOperation, v))
}

// Amps applies equality check predicate on the "amps" field. It's identical to AmpsEQ.
func Amps(v int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldEQ(FieldAmps, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldEQ(FieldTimestamp, v))
}

// VinEQ applies the EQ predicate on the "vin" field.
func VinEQ(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldEQ(FieldVin, v))
}

// VinNEQ applies the NEQ predicate on the "vin" field.
func VinNEQ(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldNEQ(FieldVin, v))
}

// VinIn applies the In predicate on the "vin" field.
func VinIn(vs ...string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldIn(FieldVin, vs...))
}

// VinNotIn applies the NotIn predicate on the "vin" field.
func VinNotIn(vs ...string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldNotIn(FieldVin, vs...))
}

// VinGT applies the GT predicate on the "vin" field.
func VinGT(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldGT(FieldVin, v))
}

// VinGTE applies the GTE predicate on the "vin" field.
func VinGTE(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldGTE(FieldVin, v))
}

// VinLT applies the LT predicate on the "vin" field.
func VinLT(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldLT(FieldVin, v))
}

// VinLTE applies the LTE predicate on the "vin" field.
func VinLTE(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldLTE(FieldVin, v))
}

// VinContains applies the Contains predicate on the "vin" field.
func VinContains(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldContains(FieldVin, v))
}

// VinHasPrefix applies the HasPrefix predicate on the "vin" field.
func VinHasPrefix(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldHasPrefix(FieldVin, v))
}

// VinHasSuffix applies the HasSuffix predicate on the "vin" field.
func VinHasSuffix(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldHasSuffix(FieldVin, v))
}

// VinEqualFold applies the EqualFold predicate on the "vin" field.
func VinEqualFold(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldEqualFold(FieldVin, v))
}

// VinContainsFold applies the ContainsFold predicate on the "vin" field.
func VinContainsFold(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldContainsFold(FieldVin, v))
}

// OperationEQ applies the EQ predicate on the "operation" field.
func OperationEQ(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldEQ(FieldOperation, v))
}

// OperationNEQ applies the NEQ predicate on the "operation" field.
func OperationNEQ(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldNEQ(FieldOperation, v))
}

// OperationIn applies the In predicate on the "operation" field.
func OperationIn(vs ...string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldIn(FieldOperation, vs...))
}

// OperationNotIn applies the NotIn predicate on the "operation" field.
func OperationNotIn(vs ...string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldNotIn(FieldOperation, vs...))
}

// OperationGT applies the GT predicate on the "operation" field.
func OperationGT(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldGT(FieldOperation, v))
}

// OperationGTE applies the GTE predicate on the "operation" field.
func OperationGTE(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldGTE(FieldOperation, v))
}

// OperationLT applies the LT predicate on the "operation" field.
func OperationLT(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldLT(FieldOperation, v))
}

// OperationLTE applies the LTE predicate on the "operation" field.
func OperationLTE(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldLTE(FieldOperation, v))
}

// OperationContains applies the Contains predicate on the "operation" field.
func OperationContains(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldContains(FieldOperation, v))
}

// OperationHasPrefix applies the HasPrefix predicate on the "operation" field.
func OperationHasPrefix(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldHasPrefix(FieldOperation, v))
}

// OperationHasSuffix applies the HasSuffix predicate on the "operation" field.
func OperationHasSuffix(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldHasSuffix(FieldOperation, v))
}

// OperationEqualFold applies the EqualFold predicate on the "operation" field.
func OperationEqualFold(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldEqualFold(FieldOperation, v))
}

// OperationContainsFold applies the ContainsFold predicate on the "operation" field.
func OperationContainsFold(v string) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldContainsFold(FieldOperation, v))
}

// AmpsEQ applies the EQ predicate on the "amps" field.
func AmpsEQ(v int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldEQ(FieldAmps, v))
}

// AmpsNEQ applies the NEQ predicate on the "amps" field.
func AmpsNEQ(v int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldNEQ(FieldAmps, v))
}

// AmpsIn applies the In predicate on the "amps" field.
func AmpsIn(vs ...int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldIn(FieldAmps, vs...))
}

// AmpsNotIn applies the NotIn predicate on the "amps" field.
func AmpsNotIn(vs ...int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldNotIn(FieldAmps, vs...))
}

// AmpsGT applies the GT predicate on the "amps" field.
func AmpsGT(v int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldGT(FieldAmps, v))
}

// AmpsGTE applies the GTE predicate on the "amps" field.
func AmpsGTE(v int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldGTE(FieldAmps, v))
}

// AmpsLT applies the LT predicate on the "amps" field.
func AmpsLT(v int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldLT(FieldAmps, v))
}

// AmpsLTE applies the LTE predicate on the "amps" field.
func AmpsLTE(v int) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldLTE(FieldAmps, v))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.FieldLTE(FieldTimestamp, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ChargeCommandHistory) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ChargeCommandHistory) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ChargeCommandHistory) predicate.ChargeCommandHistory {
	return predicate.ChargeCommandHistory(sql.NotPredicates(p))
}
