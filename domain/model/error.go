package model

type ErrorCode string

const (
	ErrCodeValidationError ErrorCode = "ValidationError"

	ErrCodeInternalServer  ErrorCode = "InternalServer"
	ErrCodeUnauthorized    ErrorCode = "Unauthorized"
	ErrCodeForbidden       ErrorCode = "Forbidden"
	ErrCodeNotFound        ErrorCode = "NotFound"
	ErrCodeTooManyRequests ErrorCode = "TooManyRequests"

	ErrCodeDeviceOffline ErrorCode = "DeviceOffline"

	ErrCodeMetricNotFound ErrorCode = "MetricNotFound"
)
