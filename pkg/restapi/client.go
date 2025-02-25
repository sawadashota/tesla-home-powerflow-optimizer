// Package restapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package restapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// Defines values for ErrorCode.
const (
	FailedPrecondition  ErrorCode = "FailedPrecondition"
	InternalServerError ErrorCode = "InternalServerError"
	NotFound            ErrorCode = "NotFound"
	ValidationError     ErrorCode = "ValidationError"
)

// Defines values for VehicleDataState.
const (
	Offline VehicleDataState = "offline"
	Online  VehicleDataState = "online"
)

// ChargeSettingMinimumSetting defines model for Charge.Setting.MinimumSetting.
type ChargeSettingMinimumSetting struct {
	Amperage       uint8  `json:"amperage"`
	Threshold      uint8  `json:"threshold"`
	TimeRangeEnd   string `json:"time_range_end"`
	TimeRangeStart string `json:"time_range_start"`
}

// ChargeSettingSetting defines model for Charge.Setting.Setting.
type ChargeSettingSetting struct {
	ChargeStartThreshold int  `json:"charge_start_threshold"`
	Enabled              bool `json:"enabled"`

	// MinimumSetting setting to charge the vehicle until the battery reaches the threshold
	MinimumSetting              ChargeSettingMinimumSetting `json:"minimum_setting"`
	PowerUsageDecreaseThreshold int                         `json:"power_usage_decrease_threshold"`
	PowerUsageIncreaseThreshold int                         `json:"power_usage_increase_threshold"`

	// UpdateInterval minutes of update interval
	UpdateInterval int `json:"update_interval"`
}

// ChargeState defines model for ChargeState.
type ChargeState struct {
	BatteryLevel            int     `json:"battery_level"`
	BatteryRange            float32 `json:"battery_range"`
	ChargeAmps              int     `json:"charge_amps"`
	ChargeCurrentRequest    int     `json:"charge_current_request"`
	ChargeCurrentRequestMax int     `json:"charge_current_request_max"`
	ChargeEnableRequest     bool    `json:"charge_enable_request"`
	ChargeLimitSoc          int     `json:"charge_limit_soc"`
	ChargePortDoorOpen      bool    `json:"charge_port_door_open"`
	ChargePortLatch         string  `json:"charge_port_latch"`
	ChargerActualCurrent    int     `json:"charger_actual_current"`
	ChargerVoltage          int     `json:"charger_voltage"`
	ChargingState           string  `json:"charging_state"`
	MinutesToFullCharge     int     `json:"minutes_to_full_charge"`
	Timestamp               int64   `json:"timestamp"`
	UsableBatteryLevel      int     `json:"usable_battery_level"`
}

// Error defines model for Error.
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

// ErrorCode defines model for ErrorCode.
type ErrorCode string

// VehicleData defines model for VehicleData.
type VehicleData struct {
	ChargeState ChargeState      `json:"charge_state"`
	State       VehicleDataState `json:"state"`
	Vin         string           `json:"vin"`
}

// VehicleDataState defines model for VehicleData.State.
type VehicleDataState string

// SettingSaveVehicleChargeSettingJSONRequestBody defines body for SettingSaveVehicleChargeSetting for application/json ContentType.
type SettingSaveVehicleChargeSettingJSONRequestBody = ChargeSettingSetting

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /vehicle)
	GetVehicleData(w http.ResponseWriter, r *http.Request)

	// (GET /vehicle/charge/setting)
	SettingGetVehicleChargeSetting(w http.ResponseWriter, r *http.Request)

	// (PUT /vehicle/charge/setting)
	SettingSaveVehicleChargeSetting(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (GET /vehicle)
func (_ Unimplemented) GetVehicleData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /vehicle/charge/setting)
func (_ Unimplemented) SettingGetVehicleChargeSetting(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (PUT /vehicle/charge/setting)
func (_ Unimplemented) SettingSaveVehicleChargeSetting(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetVehicleData operation middleware
func (siw *ServerInterfaceWrapper) GetVehicleData(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetVehicleData(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// SettingGetVehicleChargeSetting operation middleware
func (siw *ServerInterfaceWrapper) SettingGetVehicleChargeSetting(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SettingGetVehicleChargeSetting(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// SettingSaveVehicleChargeSetting operation middleware
func (siw *ServerInterfaceWrapper) SettingSaveVehicleChargeSetting(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SettingSaveVehicleChargeSetting(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/vehicle", wrapper.GetVehicleData)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/vehicle/charge/setting", wrapper.SettingGetVehicleChargeSetting)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/vehicle/charge/setting", wrapper.SettingSaveVehicleChargeSetting)
	})

	return r
}

type GetVehicleDataRequestObject struct {
}

type GetVehicleDataResponseObject interface {
	VisitGetVehicleDataResponse(w http.ResponseWriter) error
}

type GetVehicleData200JSONResponse VehicleData

func (response GetVehicleData200JSONResponse) VisitGetVehicleDataResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetVehicleData400JSONResponse Error

func (response GetVehicleData400JSONResponse) VisitGetVehicleDataResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetVehicleData500JSONResponse Error

func (response GetVehicleData500JSONResponse) VisitGetVehicleDataResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type SettingGetVehicleChargeSettingRequestObject struct {
}

type SettingGetVehicleChargeSettingResponseObject interface {
	VisitSettingGetVehicleChargeSettingResponse(w http.ResponseWriter) error
}

type SettingGetVehicleChargeSetting200JSONResponse ChargeSettingSetting

func (response SettingGetVehicleChargeSetting200JSONResponse) VisitSettingGetVehicleChargeSettingResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type SettingGetVehicleChargeSetting500JSONResponse Error

func (response SettingGetVehicleChargeSetting500JSONResponse) VisitSettingGetVehicleChargeSettingResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type SettingSaveVehicleChargeSettingRequestObject struct {
	Body *SettingSaveVehicleChargeSettingJSONRequestBody
}

type SettingSaveVehicleChargeSettingResponseObject interface {
	VisitSettingSaveVehicleChargeSettingResponse(w http.ResponseWriter) error
}

type SettingSaveVehicleChargeSetting201Response struct {
}

func (response SettingSaveVehicleChargeSetting201Response) VisitSettingSaveVehicleChargeSettingResponse(w http.ResponseWriter) error {
	w.WriteHeader(201)
	return nil
}

type SettingSaveVehicleChargeSetting400JSONResponse Error

func (response SettingSaveVehicleChargeSetting400JSONResponse) VisitSettingSaveVehicleChargeSettingResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type SettingSaveVehicleChargeSetting500JSONResponse Error

func (response SettingSaveVehicleChargeSetting500JSONResponse) VisitSettingSaveVehicleChargeSettingResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /vehicle)
	GetVehicleData(ctx context.Context, request GetVehicleDataRequestObject) (GetVehicleDataResponseObject, error)

	// (GET /vehicle/charge/setting)
	SettingGetVehicleChargeSetting(ctx context.Context, request SettingGetVehicleChargeSettingRequestObject) (SettingGetVehicleChargeSettingResponseObject, error)

	// (PUT /vehicle/charge/setting)
	SettingSaveVehicleChargeSetting(ctx context.Context, request SettingSaveVehicleChargeSettingRequestObject) (SettingSaveVehicleChargeSettingResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetVehicleData operation middleware
func (sh *strictHandler) GetVehicleData(w http.ResponseWriter, r *http.Request) {
	var request GetVehicleDataRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetVehicleData(ctx, request.(GetVehicleDataRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetVehicleData")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetVehicleDataResponseObject); ok {
		if err := validResponse.VisitGetVehicleDataResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// SettingGetVehicleChargeSetting operation middleware
func (sh *strictHandler) SettingGetVehicleChargeSetting(w http.ResponseWriter, r *http.Request) {
	var request SettingGetVehicleChargeSettingRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.SettingGetVehicleChargeSetting(ctx, request.(SettingGetVehicleChargeSettingRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "SettingGetVehicleChargeSetting")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(SettingGetVehicleChargeSettingResponseObject); ok {
		if err := validResponse.VisitSettingGetVehicleChargeSettingResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// SettingSaveVehicleChargeSetting operation middleware
func (sh *strictHandler) SettingSaveVehicleChargeSetting(w http.ResponseWriter, r *http.Request) {
	var request SettingSaveVehicleChargeSettingRequestObject

	var body SettingSaveVehicleChargeSettingJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.SettingSaveVehicleChargeSetting(ctx, request.(SettingSaveVehicleChargeSettingRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "SettingSaveVehicleChargeSetting")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(SettingSaveVehicleChargeSettingResponseObject); ok {
		if err := validResponse.VisitSettingSaveVehicleChargeSettingResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xX32/bNhD+Vwiuj0KsZm0R6G3L2iHAuhXI0Jc0E87SWWJBkRp5dBME/t8HkrIlWZK9",
	"oFv3sqfI5PF+fN/H4+WJF7pptUJFlmdP3BY1NhA+r2swFV7cIpFQ1cV7oUTjmu6nN2iNbtGQwGAOTYsG",
	"KvTfJW7ASeLZyzcJ32jTAPGMO6Hoiie8gQfvKe420S3PXiecHlvkGReKsELDdwmn2qCttSxHXtOTTtN0",
	"4DWd9SoazA2oCnNUY9c8fZOlKU84PkDTSn/w8jKutECERvGM//HpU/l0ucvinxf8EMOS8diMQ1gCQ+Mg",
	"e5dfFWSXcIN/OmGw5NndAKqZ6JOak56v+4Nnvf6MBfn0j7hfJL0IdjFIPiLrwEDgY8oBKlhLDKbd5lpr",
	"iaD8Znc4t31ckPK3Dc/unvgLgxue8e9WvXJXnWxXpzW7u094ibYwoiWhPchdAEaaxVIY1ci2WItCInOK",
	"hAwr60DLIzMIRY02rPXl7hLe6i9ocmehwrzEwiBYPAJkJNApIEMPQs16OAOpa0sgf5jQbEFGyQ2rbYRy",
	"hJbpDYu27GCbDL1PnR9pbU9esiSAs+WcRWxazlQWy8q9JSCcyrWjMZe4RTkl5XTX2B8Ol2igW+WadbTo",
	"wICmtYP9gYvOoHDGoKLcQ4qWnmObN/Bw0j5SM+N6cL06UykaQbnVxXOR6M632lBeam1y3aI6GSqYSqCi",
	"9mZ923urKqiCkCYdNJ40ORTkQO5xOFW7ybdaEozIOTYSqvJijeLo87ju9uYS6W5NTjrfOCnzGGw+hu+y",
	"lqBp/fbhjRKK3rzis1fWBra+QpdHN3Ps6Vi0Y4ku6vGk+JaUNiOrJaXMyWKR8Cm7EyoXORoSsgD2XA95",
	"a4w2M4+dLgPtpx6gcPTaG3rloLVjOS683cFzb7+Y03WXASqvhjv+DoTE8oPBQqtShC6f8I8gRQn+R6wj",
	"4b9qeqddePVvfDdVIG/RbNFEg/sZ0X+MT+BPQHDq1aezgAz78S7h/d3rStBKCuWL15tN+JrLZitU9wD+",
	"gqqimmcvz41D/sg+XDLOeAqvPyzURgemBI1aAnsPCipsUBHzqInCO9yisfFRTS/Si9Qn6aUNreAZ/z4s",
	"hWGuDoCtuonCf1cYmpjHM5B0U/KM/4w0RNyXYlutbMT7Mk2j/hR1LRDaVooinF99tlr1o/s5PoZhQt3j",
	"EeH3Gll3o1kNlllXFIgllhe+xFf/YCJRewsp2CBPVmgnS6Y0MadKNJZAlWH22qdYOvSzm1BbL3pmHxXB",
	"Q0j19bdINd4ihvv9XXKgehUltxqMsLPMd7NpL4DuwnTH/kUlLMz4zxbFf4J0wlu3jOYtbHERzlDHj7p8",
	"/CZI9j2JjMPdhM+X0zF9EW7m5Q9M4Rdm0GpnCgwGa0TF/ARN3sYy8NtO0v939u/cWf8qhTUb/r0cW//w",
	"4WZfGKqy1SJMJM5InvGaqM1WK6kLkLW2lF2lV1d8d7/7KwAA///TwpZxVREAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
