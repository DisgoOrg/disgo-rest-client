package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/DisgoOrg/log"
)

// rest errors
var (
	ErrBadGateway   = errors.New("bad gateway could not reach destination")
	ErrUnauthorized = errors.New("not authorized for this endpoint")
	ErrBadRequest   = errors.New("bad request please check your request")
	ErrRatelimited  = errors.New("too many requests")
)

// NewRestClient constructs a new RestClient with the given http.Client, log.Logger & useragent
func NewRestClient(httpClient *http.Client, logger log.Logger, userAgent string, customHeader http.Header) RestClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &RestClientImpl{userAgent: userAgent, httpClient: httpClient, logger: logger, customHeader: customHeader}
}

type RestClient interface {
	Close()
	UserAgent() string
	HttpClient() *http.Client
	Logger() log.Logger
	Do(route *CompiledAPIRoute, rqBody interface{}, rsBody interface{}) RestError
	DoWithHeaders(route *CompiledAPIRoute, rqBody interface{}, rsBody interface{}, customHeader http.Header) RestError
}

type RestClientImpl struct {
	userAgent    string
	httpClient   *http.Client
	logger       log.Logger
	customHeader http.Header
}

func (r *RestClientImpl) Close() {
	r.httpClient.CloseIdleConnections()
}

func (r *RestClientImpl) UserAgent() string {
	return r.userAgent
}

func (r *RestClientImpl) HttpClient() *http.Client {
	return r.httpClient
}

func (r *RestClientImpl) Logger() log.Logger {
	return r.logger
}

func (r *RestClientImpl) Do(route *CompiledAPIRoute, rqBody interface{}, rsBody interface{}) RestError {
	return r.DoWithHeaders(route, rqBody, rsBody, r.customHeader)
}

func (r *RestClientImpl) DoWithHeaders(route *CompiledAPIRoute, rqBody interface{}, rsBody interface{}, customHeader http.Header) RestError {
	rqBuffer := &bytes.Buffer{}
	var contentType string

	if rqBody != nil {
		var buffer *bytes.Buffer
		switch v := rqBody.(type) {
		case *MultipartBuffer:
			contentType = v.ContentType
			buffer = v.Buffer

		case url.Values:
			contentType = "application/x-www-form-urlencoded"
			buffer = bytes.NewBufferString(v.Encode())

		default:
			contentType = "application/json"
			buffer = &bytes.Buffer{}
			err := json.NewEncoder(buffer).Encode(rqBody)
			if err != nil {
				return NewRestError(nil, err)
			}
		}
		body, _ := ioutil.ReadAll(io.TeeReader(buffer, rqBuffer))
		r.Logger().Debugf("request to %s, body: %s", route.URL(), string(body))
	}

	rq, err := http.NewRequest(route.Method().String(), route.URL(), rqBuffer)
	if err != nil {
		return NewRestError(nil, err)
	}

	if customHeader != nil {
		rq.Header = customHeader
	}
	rq.Header.Set("User-Agent", r.UserAgent())
	if contentType != "" {
		rq.Header.Set("Content-Type", contentType)
	}

	rs, err := r.httpClient.Do(rq)
	if err != nil {
		return NewRestError(rs, err)
	}

	if rs.Body != nil {
		buffer := &bytes.Buffer{}
		body, _ := ioutil.ReadAll(io.TeeReader(rs.Body, buffer))
		rs.Body = ioutil.NopCloser(buffer)
		r.Logger().Debugf("response from %s, code %d, body: %s", route.URL(), rs.StatusCode, string(body))
	}

	switch rs.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusNoContent:
		if rsBody != nil && rs.Body != nil {
			if err = json.NewDecoder(rs.Body).Decode(rsBody); err != nil {
				r.Logger().Errorf("error unmarshalling response. error: %s", err)
				return NewRestError(rs, err)
			}
		}
		return nil

	case http.StatusTooManyRequests:
		r.Logger().Error(ErrRatelimited)
		return NewRestError(rs, ErrRatelimited)

	case http.StatusBadGateway:
		r.Logger().Error(ErrBadGateway)
		return NewRestError(rs, ErrBadGateway)

	case http.StatusBadRequest:
		r.Logger().Error(ErrBadRequest)
		return NewRestError(rs, ErrBadRequest)

	case http.StatusUnauthorized:
		r.Logger().Error(ErrUnauthorized)
		return NewRestError(rs, ErrUnauthorized)

	default:
		body, _ := ioutil.ReadAll(rq.Body)
		return NewRestError(rs, fmt.Errorf("request to %s failed. statuscode: %d, body: %s", rq.URL, rs.StatusCode, body))
	}
}
