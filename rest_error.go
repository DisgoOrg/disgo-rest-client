package restclient

import (
	"net/http"
)

var _ RestError = (*restErrorImpl)(nil)

type RestError interface {
	error
	Response() *http.Response
}

func NewRestError(response *http.Response, err error) *restErrorImpl {
	return &restErrorImpl{
		err:      err,
		response: response,
	}
}

type restErrorImpl struct {
	err      error
	response *http.Response
}

// Error returns the specific error message
func (r *restErrorImpl) Error() string {
	return r.err.Error()
}

// Error returns the specific error message
func (r *restErrorImpl) String() string {
	return r.err.Error()
}

// Response returns the http.Response. May be null depending on what broke during the request
func (r *restErrorImpl) Response() *http.Response {
	return r.response
}
