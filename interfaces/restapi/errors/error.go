package errors

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/morikuni/failure/v2"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/logx"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/pkg/restapi"
)

func RequestErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func ResponseErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	msg := "failed to process request"
	code := failure.CodeOf(err)
	stack := failure.CallStackOf(err)
	slog.Error(
		msg,
		logx.ErrorAttr(err),
		slog.String("error_code", fmt.Sprint(code)),
		slog.String("stack", fmt.Sprintf("%+v", stack)),
	)
	http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
}

func WriteError(w http.ResponseWriter, err error, code restapi.ErrorCode) {
	w.WriteHeader(identifyStatusCode(code))
	encoder := json.NewEncoder(w)
	err = encoder.Encode(convertError(err, code))
	if err != nil {
		slog.Error(err.Error())
	}
}

func identifyStatusCode(code restapi.ErrorCode) int {
	switch code {
	case restapi.FailedPrecondition, restapi.ValidationError:
		return http.StatusBadRequest
	case restapi.NotFound:
		return http.StatusNotFound
	case restapi.InternalServerError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func convertError(err error, code restapi.ErrorCode) *restapi.Error {
	message := failure.MessageOf(err)
	return &restapi.Error{
		Code:    code,
		Message: message.String(),
	}
}
