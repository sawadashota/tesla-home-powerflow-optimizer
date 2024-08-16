# Tesla Home PowerFlow Optimizer

## Requirements

* Tesla car
* [Tesla API Client](https://developer.tesla.com/)
  * OAuth Grant Type: `Authorization Code and Machine-to-Machine`
  * Scope:
    * Vehicle Information: `vehicle_device_data`
    * Vehicle Charging Management: `vehicle_charging_cmds`
  * [Command-authentication private key](https://github.com/teslamotors/vehicle-command?tab=readme-ov-file#generating-a-command-authentication-private-key)
    * You should host public key on your website like this: https://github.com/sawadashota/tesla-home-powerflow-optimizer-web
* Something to get your home power usage data
  * AiSEG2

## Installation

```shell
go install github.com/sawadashota/tesla-home-powerflow-optimizer
```

## Setup

### Environment Variables

```shell
cp .envrc.sample .envrc
vi .envrc
direnv allow
```

### Sign in with Tesla

```shell
tesla-home-powerflow-optimizer authentication
```

### Check API Connections

> [!IMPORTANT]  
> Check your Tesla car is plugged and ready to charge.

```shell
tesla-home-powerflow-optimizer test start-charging
```

and check collector (e.g. AiSEG2) connection

```shell
tesla-home-powerflow-optimizer test collector
```

## Run

```shell
tesla-home-powerflow-optimizer serve
```

You can access the setting page at http://localhost:8088
