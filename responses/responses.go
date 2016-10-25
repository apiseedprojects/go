package responses

import (
	"encoding/json"
	"net/http"
)

type M map[string]interface{}

type responseEnvelope struct {
	HTTPCode     int         `json:"http_code"`
	HTTPMessage  string      `json:"http_message"`
	ErrorMessage string      `json:"error_message"`
	Response     interface{} `json:"response"`
}

func Generic(w http.ResponseWriter, code int, errorMessage string, response interface{}) {
	d1 := &responseEnvelope{
		HTTPCode:     code,
		HTTPMessage:  http.StatusText(code),
		ErrorMessage: errorMessage,
		Response:     response,
	}
	b1, err := json.Marshal(d1)
	if err != nil {
		d2 := &responseEnvelope{
			HTTPCode:     http.StatusInternalServerError,
			HTTPMessage:  http.StatusText(http.StatusInternalServerError),
			ErrorMessage: http.StatusText(code) + " : " + err.Error(),
			Response:     M{},
		}
		b2, err := json.Marshal(d2)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(b2)
		return
	}
	w.WriteHeader(code)
	w.Write(b1)
}

func GenericSimplified(w http.ResponseWriter, code int, response interface{}) {
	Generic(
		w,
		code,
		http.StatusText(code),
		response,
	)
}

func Continue(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusContinue, response)
}

func SwitchingProtocols(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusSwitchingProtocols, response)
}

func OK(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusOK, response)
}

func Created(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusCreated, response)
}

func Accepted(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusAccepted, response)
}

func NonAuthoritativeInfo(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusNonAuthoritativeInfo, response)
}

func NoContent(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusNoContent, response)
}

func ResetContent(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusResetContent, response)
}

func PartialContent(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusPartialContent, response)
}

func MultipleChoices(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusMultipleChoices, response)
}

func MovedPermanently(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusMovedPermanently, response)
}

func Found(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusFound, response)
}

func SeeOther(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusSeeOther, response)
}

func NotModified(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusNotModified, response)
}

func UseProxy(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusUseProxy, response)
}

func TemporaryRedirect(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusTemporaryRedirect, response)
}

func BadRequest(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusBadRequest, response)
}

func Unauthorized(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusUnauthorized, response)
}

func PaymentRequired(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusPaymentRequired, response)
}

func Forbidden(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusForbidden, response)
}

func NotFound(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusNotFound, response)
}

func MethodNotAllowed(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusMethodNotAllowed, response)
}

func NotAcceptable(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusNotAcceptable, response)
}

func ProxyAuthRequired(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusProxyAuthRequired, response)
}

func RequestTimeout(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusRequestTimeout, response)
}

func Conflict(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusConflict, response)
}

func Gone(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusGone, response)
}

func LengthRequired(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusLengthRequired, response)
}

func PreconditionFailed(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusPreconditionFailed, response)
}

func RequestEntityTooLarge(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusRequestEntityTooLarge, response)
}

func RequestURITooLong(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusRequestURITooLong, response)
}

func UnsupportedMediaType(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusUnsupportedMediaType, response)
}

func RequestedRangeNotSatisfiable(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusRequestedRangeNotSatisfiable, response)
}

func ExpectationFailed(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusExpectationFailed, response)
}

func Teapot(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusTeapot, response)
}

func PreconditionRequired(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusPreconditionRequired, response)
}

func TooManyRequests(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusTooManyRequests, response)
}

func RequestHeaderFieldsTooLarge(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusRequestHeaderFieldsTooLarge, response)
}

func UnavailableForLegalReasons(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusUnavailableForLegalReasons, response)
}

func InternalServerError(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusInternalServerError, response)
}

func NotImplemented(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusNotImplemented, response)
}

func BadGateway(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusBadGateway, response)
}

func ServiceUnavailable(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusServiceUnavailable, response)
}

func GatewayTimeout(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusGatewayTimeout, response)
}

func HTTPVersionNotSupported(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusHTTPVersionNotSupported, response)
}

func NetworkAuthenticationRequired(w http.ResponseWriter, response interface{}) {
	GenericSimplified(w, http.StatusNetworkAuthenticationRequired, response)
}
