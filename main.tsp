// https://typespec.io/playground

import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;

enum ErrorCode {
  FailedPrecondition,
  ValidationError,
  NotFound,
  InternalServerError,
}

@error
model Error {
  code: ErrorCode;
  message: string;
}

@service({
  title: "Charging Management Service",
})
@server("http://localhost:8088", "API server endpoint")
@route("/vehicle")
namespace Vehicle {
  model ChargeState {
    @minValue(0)
    @maxValue(100)
    battery_level: integer;

    battery_range: float;
    charge_amps: integer;
    charge_current_request: integer;
    charge_current_request_max: integer;
    charge_enable_request: boolean;

    @minValue(0)
    @maxValue(100)
    charge_limit_soc: integer;

    charge_port_door_open: boolean;

    @example("Engaged")
    charge_port_latch: string;

    charger_actual_current: integer;
    charger_voltage: integer;

    @example("Charging")
    charging_state: string;

    minutes_to_full_charge: integer;
    time_to_full_charge: float;
    timestamp: int64;

    @minValue(0)
    @maxValue(100)
    usable_battery_level: integer;
  }
  model VehicleData {
    @minLength(1)
    vin: string;

    state: "online" | "offline";
    charge_state: ChargeState;
  }
  op getVehicleData(): {
    @statusCode statusCode: 200;
    @body body: VehicleData;
  } | {
    @statusCode statusCode: 400;
    @body error: Error;
  } | {
    @statusCode statusCode: 500;
    @body error: Error;
  };

  @route("/charge")
  namespace Charge {
    @route("/setting")
    namespace Setting {
      model Setting {
        enabled: boolean;
      }

      op getVehicleChargeSetting(): {
        @statusCode statusCode: 200;
        @body body: Setting;
      };
    }
  }
}
