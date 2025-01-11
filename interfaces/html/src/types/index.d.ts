interface ChargeSetting {
    enabled: boolean;
    charge_start_threshold: number;
    power_usage_decrease_threshold: number;
    power_usage_increase_threshold: number;
    update_interval: number;
    minimum_setting: MinimumChargeSetting;
}

interface MinimumChargeSetting {
    threshold: number;
    time_range_start: string;
    time_range_end: string;
    amperage: number;
}
