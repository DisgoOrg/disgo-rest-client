package endpoints

import (
	"github.com/DisgoOrg/log"
	"net/http"
)

func NewRestClient(httpClient *http.Client, logger log.Logger, userAgent string) *RestClient {
	return &RestClient{httpClient: httpClient, logger: logger}
}

type RestClient struct {
	userAgent  string
	httpClient *http.Client
	logger     log.Logger
}
