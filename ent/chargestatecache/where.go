// Code generated by ent, DO NOT EDIT.

package chargestatecache

import (
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldID, id))
}

// Vin applies equality check predicate on the "vin" field. It's identical to VinEQ.
func Vin(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldVin, v))
}

// BatteryLevel applies equality check predicate on the "battery_level" field. It's identical to BatteryLevelEQ.
func BatteryLevel(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldBatteryLevel, v))
}

// BatteryRange applies equality check predicate on the "battery_range" field. It's identical to BatteryRangeEQ.
func BatteryRange(v float32) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldBatteryRange, v))
}

// ChargeAmps applies equality check predicate on the "charge_amps" field. It's identical to ChargeAmpsEQ.
func ChargeAmps(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargeAmps, v))
}

// ChargeCurrentRequest applies equality check predicate on the "charge_current_request" field. It's identical to ChargeCurrentRequestEQ.
func ChargeCurrentRequest(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargeCurrentRequest, v))
}

// ChargeCurrentRequestMax applies equality check predicate on the "charge_current_request_max" field. It's identical to ChargeCurrentRequestMaxEQ.
func ChargeCurrentRequestMax(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargeCurrentRequestMax, v))
}

// ChargeEnableRequest applies equality check predicate on the "charge_enable_request" field. It's identical to ChargeEnableRequestEQ.
func ChargeEnableRequest(v bool) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargeEnableRequest, v))
}

// ChargeLimitSoc applies equality check predicate on the "charge_limit_soc" field. It's identical to ChargeLimitSocEQ.
func ChargeLimitSoc(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargeLimitSoc, v))
}

// ChargePortDoorOpen applies equality check predicate on the "charge_port_door_open" field. It's identical to ChargePortDoorOpenEQ.
func ChargePortDoorOpen(v bool) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargePortDoorOpen, v))
}

// ChargePortLatch applies equality check predicate on the "charge_port_latch" field. It's identical to ChargePortLatchEQ.
func ChargePortLatch(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargePortLatch, v))
}

// ChargerActualCurrent applies equality check predicate on the "charger_actual_current" field. It's identical to ChargerActualCurrentEQ.
func ChargerActualCurrent(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargerActualCurrent, v))
}

// ChargerVoltage applies equality check predicate on the "charger_voltage" field. It's identical to ChargerVoltageEQ.
func ChargerVoltage(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargerVoltage, v))
}

// ChargingState applies equality check predicate on the "charging_state" field. It's identical to ChargingStateEQ.
func ChargingState(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargingState, v))
}

// MinutesToFullCharge applies equality check predicate on the "minutes_to_full_charge" field. It's identical to MinutesToFullChargeEQ.
func MinutesToFullCharge(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldMinutesToFullCharge, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldTimestamp, v))
}

// UsableBatteryLevel applies equality check predicate on the "usable_battery_level" field. It's identical to UsableBatteryLevelEQ.
func UsableBatteryLevel(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldUsableBatteryLevel, v))
}

// VinEQ applies the EQ predicate on the "vin" field.
func VinEQ(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldVin, v))
}

// VinNEQ applies the NEQ predicate on the "vin" field.
func VinNEQ(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldVin, v))
}

// VinIn applies the In predicate on the "vin" field.
func VinIn(vs ...string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldVin, vs...))
}

// VinNotIn applies the NotIn predicate on the "vin" field.
func VinNotIn(vs ...string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldVin, vs...))
}

// VinGT applies the GT predicate on the "vin" field.
func VinGT(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldVin, v))
}

// VinGTE applies the GTE predicate on the "vin" field.
func VinGTE(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldVin, v))
}

// VinLT applies the LT predicate on the "vin" field.
func VinLT(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldVin, v))
}

// VinLTE applies the LTE predicate on the "vin" field.
func VinLTE(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldVin, v))
}

// VinContains applies the Contains predicate on the "vin" field.
func VinContains(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldContains(FieldVin, v))
}

// VinHasPrefix applies the HasPrefix predicate on the "vin" field.
func VinHasPrefix(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldHasPrefix(FieldVin, v))
}

// VinHasSuffix applies the HasSuffix predicate on the "vin" field.
func VinHasSuffix(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldHasSuffix(FieldVin, v))
}

// VinEqualFold applies the EqualFold predicate on the "vin" field.
func VinEqualFold(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEqualFold(FieldVin, v))
}

// VinContainsFold applies the ContainsFold predicate on the "vin" field.
func VinContainsFold(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldContainsFold(FieldVin, v))
}

// BatteryLevelEQ applies the EQ predicate on the "battery_level" field.
func BatteryLevelEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldBatteryLevel, v))
}

// BatteryLevelNEQ applies the NEQ predicate on the "battery_level" field.
func BatteryLevelNEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldBatteryLevel, v))
}

// BatteryLevelIn applies the In predicate on the "battery_level" field.
func BatteryLevelIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldBatteryLevel, vs...))
}

// BatteryLevelNotIn applies the NotIn predicate on the "battery_level" field.
func BatteryLevelNotIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldBatteryLevel, vs...))
}

// BatteryLevelGT applies the GT predicate on the "battery_level" field.
func BatteryLevelGT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldBatteryLevel, v))
}

// BatteryLevelGTE applies the GTE predicate on the "battery_level" field.
func BatteryLevelGTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldBatteryLevel, v))
}

// BatteryLevelLT applies the LT predicate on the "battery_level" field.
func BatteryLevelLT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldBatteryLevel, v))
}

// BatteryLevelLTE applies the LTE predicate on the "battery_level" field.
func BatteryLevelLTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldBatteryLevel, v))
}

// BatteryRangeEQ applies the EQ predicate on the "battery_range" field.
func BatteryRangeEQ(v float32) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldBatteryRange, v))
}

// BatteryRangeNEQ applies the NEQ predicate on the "battery_range" field.
func BatteryRangeNEQ(v float32) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldBatteryRange, v))
}

// BatteryRangeIn applies the In predicate on the "battery_range" field.
func BatteryRangeIn(vs ...float32) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldBatteryRange, vs...))
}

// BatteryRangeNotIn applies the NotIn predicate on the "battery_range" field.
func BatteryRangeNotIn(vs ...float32) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldBatteryRange, vs...))
}

// BatteryRangeGT applies the GT predicate on the "battery_range" field.
func BatteryRangeGT(v float32) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldBatteryRange, v))
}

// BatteryRangeGTE applies the GTE predicate on the "battery_range" field.
func BatteryRangeGTE(v float32) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldBatteryRange, v))
}

// BatteryRangeLT applies the LT predicate on the "battery_range" field.
func BatteryRangeLT(v float32) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldBatteryRange, v))
}

// BatteryRangeLTE applies the LTE predicate on the "battery_range" field.
func BatteryRangeLTE(v float32) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldBatteryRange, v))
}

// ChargeAmpsEQ applies the EQ predicate on the "charge_amps" field.
func ChargeAmpsEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargeAmps, v))
}

// ChargeAmpsNEQ applies the NEQ predicate on the "charge_amps" field.
func ChargeAmpsNEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldChargeAmps, v))
}

// ChargeAmpsIn applies the In predicate on the "charge_amps" field.
func ChargeAmpsIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldChargeAmps, vs...))
}

// ChargeAmpsNotIn applies the NotIn predicate on the "charge_amps" field.
func ChargeAmpsNotIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldChargeAmps, vs...))
}

// ChargeAmpsGT applies the GT predicate on the "charge_amps" field.
func ChargeAmpsGT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldChargeAmps, v))
}

// ChargeAmpsGTE applies the GTE predicate on the "charge_amps" field.
func ChargeAmpsGTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldChargeAmps, v))
}

// ChargeAmpsLT applies the LT predicate on the "charge_amps" field.
func ChargeAmpsLT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldChargeAmps, v))
}

// ChargeAmpsLTE applies the LTE predicate on the "charge_amps" field.
func ChargeAmpsLTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldChargeAmps, v))
}

// ChargeCurrentRequestEQ applies the EQ predicate on the "charge_current_request" field.
func ChargeCurrentRequestEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargeCurrentRequest, v))
}

// ChargeCurrentRequestNEQ applies the NEQ predicate on the "charge_current_request" field.
func ChargeCurrentRequestNEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldChargeCurrentRequest, v))
}

// ChargeCurrentRequestIn applies the In predicate on the "charge_current_request" field.
func ChargeCurrentRequestIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldChargeCurrentRequest, vs...))
}

// ChargeCurrentRequestNotIn applies the NotIn predicate on the "charge_current_request" field.
func ChargeCurrentRequestNotIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldChargeCurrentRequest, vs...))
}

// ChargeCurrentRequestGT applies the GT predicate on the "charge_current_request" field.
func ChargeCurrentRequestGT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldChargeCurrentRequest, v))
}

// ChargeCurrentRequestGTE applies the GTE predicate on the "charge_current_request" field.
func ChargeCurrentRequestGTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldChargeCurrentRequest, v))
}

// ChargeCurrentRequestLT applies the LT predicate on the "charge_current_request" field.
func ChargeCurrentRequestLT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldChargeCurrentRequest, v))
}

// ChargeCurrentRequestLTE applies the LTE predicate on the "charge_current_request" field.
func ChargeCurrentRequestLTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldChargeCurrentRequest, v))
}

// ChargeCurrentRequestMaxEQ applies the EQ predicate on the "charge_current_request_max" field.
func ChargeCurrentRequestMaxEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargeCurrentRequestMax, v))
}

// ChargeCurrentRequestMaxNEQ applies the NEQ predicate on the "charge_current_request_max" field.
func ChargeCurrentRequestMaxNEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldChargeCurrentRequestMax, v))
}

// ChargeCurrentRequestMaxIn applies the In predicate on the "charge_current_request_max" field.
func ChargeCurrentRequestMaxIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldChargeCurrentRequestMax, vs...))
}

// ChargeCurrentRequestMaxNotIn applies the NotIn predicate on the "charge_current_request_max" field.
func ChargeCurrentRequestMaxNotIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldChargeCurrentRequestMax, vs...))
}

// ChargeCurrentRequestMaxGT applies the GT predicate on the "charge_current_request_max" field.
func ChargeCurrentRequestMaxGT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldChargeCurrentRequestMax, v))
}

// ChargeCurrentRequestMaxGTE applies the GTE predicate on the "charge_current_request_max" field.
func ChargeCurrentRequestMaxGTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldChargeCurrentRequestMax, v))
}

// ChargeCurrentRequestMaxLT applies the LT predicate on the "charge_current_request_max" field.
func ChargeCurrentRequestMaxLT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldChargeCurrentRequestMax, v))
}

// ChargeCurrentRequestMaxLTE applies the LTE predicate on the "charge_current_request_max" field.
func ChargeCurrentRequestMaxLTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldChargeCurrentRequestMax, v))
}

// ChargeEnableRequestEQ applies the EQ predicate on the "charge_enable_request" field.
func ChargeEnableRequestEQ(v bool) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargeEnableRequest, v))
}

// ChargeEnableRequestNEQ applies the NEQ predicate on the "charge_enable_request" field.
func ChargeEnableRequestNEQ(v bool) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldChargeEnableRequest, v))
}

// ChargeLimitSocEQ applies the EQ predicate on the "charge_limit_soc" field.
func ChargeLimitSocEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargeLimitSoc, v))
}

// ChargeLimitSocNEQ applies the NEQ predicate on the "charge_limit_soc" field.
func ChargeLimitSocNEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldChargeLimitSoc, v))
}

// ChargeLimitSocIn applies the In predicate on the "charge_limit_soc" field.
func ChargeLimitSocIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldChargeLimitSoc, vs...))
}

// ChargeLimitSocNotIn applies the NotIn predicate on the "charge_limit_soc" field.
func ChargeLimitSocNotIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldChargeLimitSoc, vs...))
}

// ChargeLimitSocGT applies the GT predicate on the "charge_limit_soc" field.
func ChargeLimitSocGT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldChargeLimitSoc, v))
}

// ChargeLimitSocGTE applies the GTE predicate on the "charge_limit_soc" field.
func ChargeLimitSocGTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldChargeLimitSoc, v))
}

// ChargeLimitSocLT applies the LT predicate on the "charge_limit_soc" field.
func ChargeLimitSocLT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldChargeLimitSoc, v))
}

// ChargeLimitSocLTE applies the LTE predicate on the "charge_limit_soc" field.
func ChargeLimitSocLTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldChargeLimitSoc, v))
}

// ChargePortDoorOpenEQ applies the EQ predicate on the "charge_port_door_open" field.
func ChargePortDoorOpenEQ(v bool) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargePortDoorOpen, v))
}

// ChargePortDoorOpenNEQ applies the NEQ predicate on the "charge_port_door_open" field.
func ChargePortDoorOpenNEQ(v bool) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldChargePortDoorOpen, v))
}

// ChargePortLatchEQ applies the EQ predicate on the "charge_port_latch" field.
func ChargePortLatchEQ(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargePortLatch, v))
}

// ChargePortLatchNEQ applies the NEQ predicate on the "charge_port_latch" field.
func ChargePortLatchNEQ(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldChargePortLatch, v))
}

// ChargePortLatchIn applies the In predicate on the "charge_port_latch" field.
func ChargePortLatchIn(vs ...string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldChargePortLatch, vs...))
}

// ChargePortLatchNotIn applies the NotIn predicate on the "charge_port_latch" field.
func ChargePortLatchNotIn(vs ...string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldChargePortLatch, vs...))
}

// ChargePortLatchGT applies the GT predicate on the "charge_port_latch" field.
func ChargePortLatchGT(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldChargePortLatch, v))
}

// ChargePortLatchGTE applies the GTE predicate on the "charge_port_latch" field.
func ChargePortLatchGTE(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldChargePortLatch, v))
}

// ChargePortLatchLT applies the LT predicate on the "charge_port_latch" field.
func ChargePortLatchLT(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldChargePortLatch, v))
}

// ChargePortLatchLTE applies the LTE predicate on the "charge_port_latch" field.
func ChargePortLatchLTE(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldChargePortLatch, v))
}

// ChargePortLatchContains applies the Contains predicate on the "charge_port_latch" field.
func ChargePortLatchContains(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldContains(FieldChargePortLatch, v))
}

// ChargePortLatchHasPrefix applies the HasPrefix predicate on the "charge_port_latch" field.
func ChargePortLatchHasPrefix(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldHasPrefix(FieldChargePortLatch, v))
}

// ChargePortLatchHasSuffix applies the HasSuffix predicate on the "charge_port_latch" field.
func ChargePortLatchHasSuffix(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldHasSuffix(FieldChargePortLatch, v))
}

// ChargePortLatchEqualFold applies the EqualFold predicate on the "charge_port_latch" field.
func ChargePortLatchEqualFold(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEqualFold(FieldChargePortLatch, v))
}

// ChargePortLatchContainsFold applies the ContainsFold predicate on the "charge_port_latch" field.
func ChargePortLatchContainsFold(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldContainsFold(FieldChargePortLatch, v))
}

// ChargerActualCurrentEQ applies the EQ predicate on the "charger_actual_current" field.
func ChargerActualCurrentEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargerActualCurrent, v))
}

// ChargerActualCurrentNEQ applies the NEQ predicate on the "charger_actual_current" field.
func ChargerActualCurrentNEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldChargerActualCurrent, v))
}

// ChargerActualCurrentIn applies the In predicate on the "charger_actual_current" field.
func ChargerActualCurrentIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldChargerActualCurrent, vs...))
}

// ChargerActualCurrentNotIn applies the NotIn predicate on the "charger_actual_current" field.
func ChargerActualCurrentNotIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldChargerActualCurrent, vs...))
}

// ChargerActualCurrentGT applies the GT predicate on the "charger_actual_current" field.
func ChargerActualCurrentGT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldChargerActualCurrent, v))
}

// ChargerActualCurrentGTE applies the GTE predicate on the "charger_actual_current" field.
func ChargerActualCurrentGTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldChargerActualCurrent, v))
}

// ChargerActualCurrentLT applies the LT predicate on the "charger_actual_current" field.
func ChargerActualCurrentLT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldChargerActualCurrent, v))
}

// ChargerActualCurrentLTE applies the LTE predicate on the "charger_actual_current" field.
func ChargerActualCurrentLTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldChargerActualCurrent, v))
}

// ChargerVoltageEQ applies the EQ predicate on the "charger_voltage" field.
func ChargerVoltageEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargerVoltage, v))
}

// ChargerVoltageNEQ applies the NEQ predicate on the "charger_voltage" field.
func ChargerVoltageNEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldChargerVoltage, v))
}

// ChargerVoltageIn applies the In predicate on the "charger_voltage" field.
func ChargerVoltageIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldChargerVoltage, vs...))
}

// ChargerVoltageNotIn applies the NotIn predicate on the "charger_voltage" field.
func ChargerVoltageNotIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldChargerVoltage, vs...))
}

// ChargerVoltageGT applies the GT predicate on the "charger_voltage" field.
func ChargerVoltageGT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldChargerVoltage, v))
}

// ChargerVoltageGTE applies the GTE predicate on the "charger_voltage" field.
func ChargerVoltageGTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldChargerVoltage, v))
}

// ChargerVoltageLT applies the LT predicate on the "charger_voltage" field.
func ChargerVoltageLT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldChargerVoltage, v))
}

// ChargerVoltageLTE applies the LTE predicate on the "charger_voltage" field.
func ChargerVoltageLTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldChargerVoltage, v))
}

// ChargingStateEQ applies the EQ predicate on the "charging_state" field.
func ChargingStateEQ(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldChargingState, v))
}

// ChargingStateNEQ applies the NEQ predicate on the "charging_state" field.
func ChargingStateNEQ(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldChargingState, v))
}

// ChargingStateIn applies the In predicate on the "charging_state" field.
func ChargingStateIn(vs ...string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldChargingState, vs...))
}

// ChargingStateNotIn applies the NotIn predicate on the "charging_state" field.
func ChargingStateNotIn(vs ...string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldChargingState, vs...))
}

// ChargingStateGT applies the GT predicate on the "charging_state" field.
func ChargingStateGT(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldChargingState, v))
}

// ChargingStateGTE applies the GTE predicate on the "charging_state" field.
func ChargingStateGTE(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldChargingState, v))
}

// ChargingStateLT applies the LT predicate on the "charging_state" field.
func ChargingStateLT(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldChargingState, v))
}

// ChargingStateLTE applies the LTE predicate on the "charging_state" field.
func ChargingStateLTE(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldChargingState, v))
}

// ChargingStateContains applies the Contains predicate on the "charging_state" field.
func ChargingStateContains(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldContains(FieldChargingState, v))
}

// ChargingStateHasPrefix applies the HasPrefix predicate on the "charging_state" field.
func ChargingStateHasPrefix(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldHasPrefix(FieldChargingState, v))
}

// ChargingStateHasSuffix applies the HasSuffix predicate on the "charging_state" field.
func ChargingStateHasSuffix(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldHasSuffix(FieldChargingState, v))
}

// ChargingStateEqualFold applies the EqualFold predicate on the "charging_state" field.
func ChargingStateEqualFold(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEqualFold(FieldChargingState, v))
}

// ChargingStateContainsFold applies the ContainsFold predicate on the "charging_state" field.
func ChargingStateContainsFold(v string) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldContainsFold(FieldChargingState, v))
}

// MinutesToFullChargeEQ applies the EQ predicate on the "minutes_to_full_charge" field.
func MinutesToFullChargeEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldMinutesToFullCharge, v))
}

// MinutesToFullChargeNEQ applies the NEQ predicate on the "minutes_to_full_charge" field.
func MinutesToFullChargeNEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldMinutesToFullCharge, v))
}

// MinutesToFullChargeIn applies the In predicate on the "minutes_to_full_charge" field.
func MinutesToFullChargeIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldMinutesToFullCharge, vs...))
}

// MinutesToFullChargeNotIn applies the NotIn predicate on the "minutes_to_full_charge" field.
func MinutesToFullChargeNotIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldMinutesToFullCharge, vs...))
}

// MinutesToFullChargeGT applies the GT predicate on the "minutes_to_full_charge" field.
func MinutesToFullChargeGT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldMinutesToFullCharge, v))
}

// MinutesToFullChargeGTE applies the GTE predicate on the "minutes_to_full_charge" field.
func MinutesToFullChargeGTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldMinutesToFullCharge, v))
}

// MinutesToFullChargeLT applies the LT predicate on the "minutes_to_full_charge" field.
func MinutesToFullChargeLT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldMinutesToFullCharge, v))
}

// MinutesToFullChargeLTE applies the LTE predicate on the "minutes_to_full_charge" field.
func MinutesToFullChargeLTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldMinutesToFullCharge, v))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldTimestamp, v))
}

// UsableBatteryLevelEQ applies the EQ predicate on the "usable_battery_level" field.
func UsableBatteryLevelEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldEQ(FieldUsableBatteryLevel, v))
}

// UsableBatteryLevelNEQ applies the NEQ predicate on the "usable_battery_level" field.
func UsableBatteryLevelNEQ(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNEQ(FieldUsableBatteryLevel, v))
}

// UsableBatteryLevelIn applies the In predicate on the "usable_battery_level" field.
func UsableBatteryLevelIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldIn(FieldUsableBatteryLevel, vs...))
}

// UsableBatteryLevelNotIn applies the NotIn predicate on the "usable_battery_level" field.
func UsableBatteryLevelNotIn(vs ...int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldNotIn(FieldUsableBatteryLevel, vs...))
}

// UsableBatteryLevelGT applies the GT predicate on the "usable_battery_level" field.
func UsableBatteryLevelGT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGT(FieldUsableBatteryLevel, v))
}

// UsableBatteryLevelGTE applies the GTE predicate on the "usable_battery_level" field.
func UsableBatteryLevelGTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldGTE(FieldUsableBatteryLevel, v))
}

// UsableBatteryLevelLT applies the LT predicate on the "usable_battery_level" field.
func UsableBatteryLevelLT(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLT(FieldUsableBatteryLevel, v))
}

// UsableBatteryLevelLTE applies the LTE predicate on the "usable_battery_level" field.
func UsableBatteryLevelLTE(v int) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.FieldLTE(FieldUsableBatteryLevel, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ChargeStateCache) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ChargeStateCache) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ChargeStateCache) predicate.ChargeStateCache {
	return predicate.ChargeStateCache(sql.NotPredicates(p))
}
