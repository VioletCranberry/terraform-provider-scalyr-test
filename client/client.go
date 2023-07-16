package client

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type ApiClient struct {
	scalyr_app_url string
	scalyr_api_key string
	api_timeout    int
	httpClient     *resty.Client
}

type ApiResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func NewApiClient(app_url string, api_key string, api_timeout int) *ApiClient {
	timeout := time.Duration(api_timeout) * time.Second
	client := resty.New()

	client.
		SetTimeout(timeout).
		SetBaseURL(app_url).
		SetHeader("Content-Type", "application/json").
		SetAuthToken(api_key)

	return &ApiClient{
		scalyr_app_url: app_url,
		scalyr_api_key: api_key,
		api_timeout:    api_timeout,
		httpClient:     client,
	}
}

func (client *ApiClient) PostRequest(api_endpoint string,
	request_body interface{}, result interface{}) (*resty.Response, error) {

	response, err := client.
		httpClient.R().
		SetBody(request_body).
		SetResult(result).
		SetError(result).
		Post(api_endpoint)

	return response, err
}
