package tesla

type VehicleResponse struct {
	Response VehicleResponseInner `json:"response"`
}

type VehicleResponseInner struct {
	ID              int            `json:"id"`
	VehicleID       int            `json:"vehicle_id"`
	Vin             string         `json:"vin"`
	AccessType      string         `json:"access_type"`
	GranularAccess  GranularAccess `json:"granular_access"`
	Tokens          []string       `json:"tokens"`
	State           string         `json:"state"`
	InService       bool           `json:"in_service"`
	IDS             string         `json:"id_s"`
	CalendarEnabled bool           `json:"calendar_enabled"`
	APIVersion      int            `json:"api_version"`
}

type VehicleDataResponse struct {
	Response VehicleDataResponseInner `json:"response"`
}
type GranularAccess struct {
	HidePrivate bool `json:"hide_private"`
}
type ChargeState struct {
	BatteryHeaterOn           bool    `json:"battery_heater_on"`
	BatteryLevel              int     `json:"battery_level"`
	BatteryRange              float32 `json:"battery_range"`
	ChargeAmps                int     `json:"charge_amps"`
	ChargeCurrentRequest      int     `json:"charge_current_request"`
	ChargeCurrentRequestMax   int     `json:"charge_current_request_max"`
	ChargeEnableRequest       bool    `json:"charge_enable_request"`
	ChargeEnergyAdded         float32 `json:"charge_energy_added"`
	ChargeLimitSoc            int     `json:"charge_limit_soc"`
	ChargeLimitSocMax         int     `json:"charge_limit_soc_max"`
	ChargeLimitSocMin         int     `json:"charge_limit_soc_min"`
	ChargeLimitSocStd         int     `json:"charge_limit_soc_std"`
	ChargeMilesAddedIdeal     float32 `json:"charge_miles_added_ideal"`
	ChargeMilesAddedRated     float32 `json:"charge_miles_added_rated"`
	ChargePortColdWeatherMode bool    `json:"charge_port_cold_weather_mode"`
	ChargePortColor           string  `json:"charge_port_color"`
	ChargePortDoorOpen        bool    `json:"charge_port_door_open"`
	ChargePortLatch           string  `json:"charge_port_latch"`
	ChargeRate                float32 `json:"charge_rate"`
	ChargerActualCurrent      int     `json:"charger_actual_current"`
	//ChargerPhases                  any     `json:"charger_phases"`
	ChargerPilotCurrent   int     `json:"charger_pilot_current"`
	ChargerPower          int     `json:"charger_power"`
	ChargerVoltage        int     `json:"charger_voltage"`
	ChargingState         string  `json:"charging_state"`
	ConnChargeCable       string  `json:"conn_charge_cable"`
	EstBatteryRange       float32 `json:"est_battery_range"`
	FastChargerBrand      string  `json:"fast_charger_brand"`
	FastChargerPresent    bool    `json:"fast_charger_present"`
	FastChargerType       string  `json:"fast_charger_type"`
	IdealBatteryRange     float32 `json:"ideal_battery_range"`
	ManagedChargingActive bool    `json:"managed_charging_active"`
	//ManagedChargingStartTime       any     `json:"managed_charging_start_time"`
	ManagedChargingUserCanceled bool `json:"managed_charging_user_canceled"`
	MaxRangeChargeCounter       int  `json:"max_range_charge_counter"`
	MinutesToFullCharge         int  `json:"minutes_to_full_charge"`
	//NotEnoughPowerToHeat           any     `json:"not_enough_power_to_heat"`
	OffPeakChargingEnabled   bool   `json:"off_peak_charging_enabled"`
	OffPeakChargingTimes     string `json:"off_peak_charging_times"`
	OffPeakHoursEndTime      int    `json:"off_peak_hours_end_time"`
	PreconditioningEnabled   bool   `json:"preconditioning_enabled"`
	PreconditioningTimes     string `json:"preconditioning_times"`
	ScheduledChargingMode    string `json:"scheduled_charging_mode"`
	ScheduledChargingPending bool   `json:"scheduled_charging_pending"`
	//ScheduledChargingStartTime     any     `json:"scheduled_charging_start_time"`
	ScheduledDepartureTime         int     `json:"scheduled_departure_time"`
	ScheduledDepartureTimeMinutes  int     `json:"scheduled_departure_time_minutes"`
	SuperchargerSessionTripPlanner bool    `json:"supercharger_session_trip_planner"`
	TimeToFullCharge               float32 `json:"time_to_full_charge"`
	Timestamp                      int64   `json:"timestamp"`
	TripCharging                   bool    `json:"trip_charging"`
	UsableBatteryLevel             int     `json:"usable_battery_level"`
	//UserChargeEnableRequest        any     `json:"user_charge_enable_request"`
}
type ClimateState struct {
	AllowCabinOverheatProtection           bool    `json:"allow_cabin_overheat_protection"`
	AutoSeatClimateLeft                    bool    `json:"auto_seat_climate_left"`
	AutoSeatClimateRight                   bool    `json:"auto_seat_climate_right"`
	AutoSteeringWheelHeat                  bool    `json:"auto_steering_wheel_heat"`
	BatteryHeater                          bool    `json:"battery_heater"`
	BatteryHeaterNoPower                   any     `json:"battery_heater_no_power"`
	BioweaponMode                          bool    `json:"bioweapon_mode"`
	CabinOverheatProtection                string  `json:"cabin_overheat_protection"`
	CabinOverheatProtectionActivelyCooling bool    `json:"cabin_overheat_protection_actively_cooling"`
	ClimateKeeperMode                      string  `json:"climate_keeper_mode"`
	CopActivationTemperature               string  `json:"cop_activation_temperature"`
	DefrostMode                            int     `json:"defrost_mode"`
	DriverTempSetting                      float32 `json:"driver_temp_setting"`
	FanStatus                              int     `json:"fan_status"`
	HvacAutoRequest                        string  `json:"hvac_auto_request"`
	InsideTemp                             float32 `json:"inside_temp"`
	IsAutoConditioningOn                   bool    `json:"is_auto_conditioning_on"`
	IsClimateOn                            bool    `json:"is_climate_on"`
	IsFrontDefrosterOn                     bool    `json:"is_front_defroster_on"`
	IsPreconditioning                      bool    `json:"is_preconditioning"`
	IsRearDefrosterOn                      bool    `json:"is_rear_defroster_on"`
	LeftTempDirection                      int     `json:"left_temp_direction"`
	MaxAvailTemp                           float32 `json:"max_avail_temp"`
	MinAvailTemp                           float32 `json:"min_avail_temp"`
	OutsideTemp                            float32 `json:"outside_temp"`
	PassengerTempSetting                   float32 `json:"passenger_temp_setting"`
	RemoteHeaterControlEnabled             bool    `json:"remote_heater_control_enabled"`
	RightTempDirection                     int     `json:"right_temp_direction"`
	SeatHeaterLeft                         int     `json:"seat_heater_left"`
	SeatHeaterRearCenter                   int     `json:"seat_heater_rear_center"`
	SeatHeaterRearLeft                     int     `json:"seat_heater_rear_left"`
	SeatHeaterRearRight                    int     `json:"seat_heater_rear_right"`
	SeatHeaterRight                        int     `json:"seat_heater_right"`
	SideMirrorHeaters                      bool    `json:"side_mirror_heaters"`
	SteeringWheelHeatLevel                 int     `json:"steering_wheel_heat_level"`
	SteeringWheelHeater                    bool    `json:"steering_wheel_heater"`
	SupportsFanOnlyCabinOverheatProtection bool    `json:"supports_fan_only_cabin_overheat_protection"`
	Timestamp                              int64   `json:"timestamp"`
	WiperBladeHeater                       bool    `json:"wiper_blade_heater"`
}
type DriveState struct {
	ActiveRouteLatitude            float32 `json:"active_route_latitude"`
	ActiveRouteLongitude           float32 `json:"active_route_longitude"`
	ActiveRouteTrafficMinutesDelay int     `json:"active_route_traffic_minutes_delay"`
	GpsAsOf                        int     `json:"gps_as_of"`
	Heading                        int     `json:"heading"`
	Latitude                       float32 `json:"latitude"`
	Longitude                      float32 `json:"longitude"`
	NativeLatitude                 float32 `json:"native_latitude"`
	NativeLocationSupported        int     `json:"native_location_supported"`
	NativeLongitude                float32 `json:"native_longitude"`
	NativeType                     string  `json:"native_type"`
	Power                          int     `json:"power"`
	ShiftState                     any     `json:"shift_state"`
	Speed                          any     `json:"speed"`
	Timestamp                      int64   `json:"timestamp"`
}
type GuiSettings struct {
	Gui24HourTime        bool   `json:"gui_24_hour_time"`
	GuiChargeRateUnits   string `json:"gui_charge_rate_units"`
	GuiDistanceUnits     string `json:"gui_distance_units"`
	GuiRangeDisplay      string `json:"gui_range_display"`
	GuiTemperatureUnits  string `json:"gui_temperature_units"`
	GuiTirepressureUnits string `json:"gui_tirepressure_units"`
	ShowRangeUnits       bool   `json:"show_range_units"`
	Timestamp            int64  `json:"timestamp"`
}
type VehicleConfig struct {
	AuxParkLamps                string `json:"aux_park_lamps"`
	BadgeVersion                int    `json:"badge_version"`
	CanAcceptNavigationRequests bool   `json:"can_accept_navigation_requests"`
	CanActuateTrunks            bool   `json:"can_actuate_trunks"`
	CarSpecialType              string `json:"car_special_type"`
	CarType                     string `json:"car_type"`
	ChargePortType              string `json:"charge_port_type"`
	CopUserSetTempSupported     bool   `json:"cop_user_set_temp_supported"`
	DashcamClipSaveSupported    bool   `json:"dashcam_clip_save_supported"`
	DefaultChargeToMax          bool   `json:"default_charge_to_max"`
	DriverAssist                string `json:"driver_assist"`
	EceRestrictions             bool   `json:"ece_restrictions"`
	EfficiencyPackage           string `json:"efficiency_package"`
	EuVehicle                   bool   `json:"eu_vehicle"`
	ExteriorColor               string `json:"exterior_color"`
	ExteriorTrim                string `json:"exterior_trim"`
	ExteriorTrimOverride        string `json:"exterior_trim_override"`
	HasAirSuspension            bool   `json:"has_air_suspension"`
	HasLudicrousMode            bool   `json:"has_ludicrous_mode"`
	HasSeatCooling              bool   `json:"has_seat_cooling"`
	HeadlampType                string `json:"headlamp_type"`
	InteriorTrimType            string `json:"interior_trim_type"`
	KeyVersion                  int    `json:"key_version"`
	MotorizedChargePort         bool   `json:"motorized_charge_port"`
	PaintColorOverride          string `json:"paint_color_override"`
	PerformancePackage          string `json:"performance_package"`
	Plg                         bool   `json:"plg"`
	Pws                         bool   `json:"pws"`
	RearDriveUnit               string `json:"rear_drive_unit"`
	RearSeatHeaters             int    `json:"rear_seat_heaters"`
	RearSeatType                int    `json:"rear_seat_type"`
	Rhd                         bool   `json:"rhd"`
	RoofColor                   string `json:"roof_color"`
	SeatType                    any    `json:"seat_type"`
	SpoilerType                 string `json:"spoiler_type"`
	SunRoofInstalled            any    `json:"sun_roof_installed"`
	SupportsQrPairing           bool   `json:"supports_qr_pairing"`
	ThirdRowSeats               string `json:"third_row_seats"`
	Timestamp                   int64  `json:"timestamp"`
	TrimBadging                 string `json:"trim_badging"`
	UseRangeBadging             bool   `json:"use_range_badging"`
	UtcOffset                   int    `json:"utc_offset"`
	WebcamSelfieSupported       bool   `json:"webcam_selfie_supported"`
	WebcamSupported             bool   `json:"webcam_supported"`
	WheelType                   string `json:"wheel_type"`
}
type MediaInfo struct {
	A2DpSourceName       string  `json:"a2dp_source_name"`
	AudioVolume          float32 `json:"audio_volume"`
	AudioVolumeIncrement float32 `json:"audio_volume_increment"`
	AudioVolumeMax       float32 `json:"audio_volume_max"`
	MediaPlaybackStatus  string  `json:"media_playback_status"`
	NowPlayingAlbum      string  `json:"now_playing_album"`
	NowPlayingArtist     string  `json:"now_playing_artist"`
	NowPlayingDuration   int     `json:"now_playing_duration"`
	NowPlayingElapsed    int     `json:"now_playing_elapsed"`
	NowPlayingSource     string  `json:"now_playing_source"`
	NowPlayingStation    string  `json:"now_playing_station"`
	NowPlayingTitle      string  `json:"now_playing_title"`
}
type MediaState struct {
	RemoteControlEnabled bool `json:"remote_control_enabled"`
}
type SoftwareUpdate struct {
	DownloadPerc        int    `json:"download_perc"`
	ExpectedDurationSec int    `json:"expected_duration_sec"`
	InstallPerc         int    `json:"install_perc"`
	Status              string `json:"status"`
	Version             string `json:"version"`
}
type SpeedLimitMode struct {
	Active          bool    `json:"active"`
	CurrentLimitMph float32 `json:"current_limit_mph"`
	MaxLimitMph     float32 `json:"max_limit_mph"`
	MinLimitMph     float32 `json:"min_limit_mph"`
	PinCodeSet      bool    `json:"pin_code_set"`
}
type VehicleState struct {
	APIVersion                 int            `json:"api_version"`
	AutoparkStateV3            string         `json:"autopark_state_v3"`
	AutoparkStyle              string         `json:"autopark_style"`
	CalendarSupported          bool           `json:"calendar_supported"`
	CarVersion                 string         `json:"car_version"`
	CenterDisplayState         int            `json:"center_display_state"`
	DashcamClipSaveAvailable   bool           `json:"dashcam_clip_save_available"`
	DashcamState               string         `json:"dashcam_state"`
	Df                         int            `json:"df"`
	Dr                         int            `json:"dr"`
	FdWindow                   int            `json:"fd_window"`
	FeatureBitmask             string         `json:"feature_bitmask"`
	FpWindow                   int            `json:"fp_window"`
	Ft                         int            `json:"ft"`
	HomelinkDeviceCount        int            `json:"homelink_device_count"`
	HomelinkNearby             bool           `json:"homelink_nearby"`
	IsUserPresent              bool           `json:"is_user_present"`
	LastAutoparkError          string         `json:"last_autopark_error"`
	Locked                     bool           `json:"locked"`
	MediaInfo                  MediaInfo      `json:"media_info"`
	MediaState                 MediaState     `json:"media_state"`
	NotificationsSupported     bool           `json:"notifications_supported"`
	Odometer                   float32        `json:"odometer"`
	ParsedCalendarSupported    bool           `json:"parsed_calendar_supported"`
	Pf                         int            `json:"pf"`
	Pr                         int            `json:"pr"`
	RdWindow                   int            `json:"rd_window"`
	RemoteStart                bool           `json:"remote_start"`
	RemoteStartEnabled         bool           `json:"remote_start_enabled"`
	RemoteStartSupported       bool           `json:"remote_start_supported"`
	RpWindow                   int            `json:"rp_window"`
	Rt                         int            `json:"rt"`
	SantaMode                  int            `json:"santa_mode"`
	SentryMode                 bool           `json:"sentry_mode"`
	SentryModeAvailable        bool           `json:"sentry_mode_available"`
	ServiceMode                bool           `json:"service_mode"`
	ServiceModePlus            bool           `json:"service_mode_plus"`
	SmartSummonAvailable       bool           `json:"smart_summon_available"`
	SoftwareUpdate             SoftwareUpdate `json:"software_update"`
	SpeedLimitMode             SpeedLimitMode `json:"speed_limit_mode"`
	SummonStandbyModeEnabled   bool           `json:"summon_standby_mode_enabled"`
	Timestamp                  int64          `json:"timestamp"`
	TpmsHardWarningFl          bool           `json:"tpms_hard_warning_fl"`
	TpmsHardWarningFr          bool           `json:"tpms_hard_warning_fr"`
	TpmsHardWarningRl          bool           `json:"tpms_hard_warning_rl"`
	TpmsHardWarningRr          bool           `json:"tpms_hard_warning_rr"`
	TpmsLastSeenPressureTimeFl int            `json:"tpms_last_seen_pressure_time_fl"`
	TpmsLastSeenPressureTimeFr int            `json:"tpms_last_seen_pressure_time_fr"`
	TpmsLastSeenPressureTimeRl int            `json:"tpms_last_seen_pressure_time_rl"`
	TpmsLastSeenPressureTimeRr int            `json:"tpms_last_seen_pressure_time_rr"`
	TpmsPressureFl             float32        `json:"tpms_pressure_fl"`
	TpmsPressureFr             float32        `json:"tpms_pressure_fr"`
	TpmsPressureRl             float32        `json:"tpms_pressure_rl"`
	TpmsPressureRr             float32        `json:"tpms_pressure_rr"`
	TpmsRcpFrontValue          float32        `json:"tpms_rcp_front_value"`
	TpmsRcpRearValue           float32        `json:"tpms_rcp_rear_value"`
	TpmsSoftWarningFl          bool           `json:"tpms_soft_warning_fl"`
	TpmsSoftWarningFr          bool           `json:"tpms_soft_warning_fr"`
	TpmsSoftWarningRl          bool           `json:"tpms_soft_warning_rl"`
	TpmsSoftWarningRr          bool           `json:"tpms_soft_warning_rr"`
	ValetMode                  bool           `json:"valet_mode"`
	ValetPinNeeded             bool           `json:"valet_pin_needed"`
	VehicleName                string         `json:"vehicle_name"`
	VehicleSelfTestProgress    int            `json:"vehicle_self_test_progress"`
	VehicleSelfTestRequested   bool           `json:"vehicle_self_test_requested"`
	WebcamAvailable            bool           `json:"webcam_available"`
}
type VehicleDataResponseInner struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	VehicleID int    `json:"vehicle_id"`
	Vin       string `json:"vin"`
	//Color           any            `json:"color"`
	AccessType      string         `json:"access_type"`
	GranularAccess  GranularAccess `json:"granular_access"`
	Tokens          []string       `json:"tokens"`
	State           string         `json:"state"`
	InService       bool           `json:"in_service"`
	IDS             string         `json:"id_s"`
	CalendarEnabled bool           `json:"calendar_enabled"`
	APIVersion      int            `json:"api_version"`
	//BackseatToken          any            `json:"backseat_token"`
	//BackseatTokenUpdatedAt any            `json:"backseat_token_updated_at"`
	ChargeState   ChargeState   `json:"charge_state"`
	ClimateState  ClimateState  `json:"climate_state"`
	DriveState    DriveState    `json:"drive_state"`
	GuiSettings   GuiSettings   `json:"gui_settings"`
	VehicleConfig VehicleConfig `json:"vehicle_config"`
	VehicleState  VehicleState  `json:"vehicle_state"`
}

type VehicleCommandResponse struct {
	Response VehicleCommandResponseInner `json:"response"`
}

type VehicleCommandResponseInner struct {
	Result bool   `json:"result"`
	Reason string `json:"reason"`
}

type SetChargeLimitRequest struct {
	Percent int `json:"percent"`
}

type SetChargeAmpsRequest struct {
	ChargingAmps int `json:"charging_amps"`
}
