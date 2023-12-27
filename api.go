// Base API definition
package txtai

import (
    "fmt"
    "github.com/go-resty/resty/v2"
    "os"
)

// Base API definition
type API struct {
    url string
    token string
}

// Index result
type IndexResult struct {
    Id    int `json:"id"`
    Score float64 `json:"score"`
}

// Creates a new API instance
func NewAPI(params ...string) API {
    // Parse url and token
    var url, token string
    if len(params) > 0 {
        url = params[0]
    }
    if len(params) > 1 {
        token = params[1]
    }

    // Fallback to env variables
    if url == "" {
        url = os.Getenv("TXTAI_API_URL")
    }
    if token == "" {
        token = os.Getenv("TXTAI_API_TOKEN")
    }

    return API{url, token}
}

// Creates a new Resty request and returns it.
func (api *API) Request(method string, result interface{}) *resty.Request {
    request := resty.New().R()
    if result != nil {
        request.SetResult(result)
    }

    return request
}

// Executes a GET request. Loads response into result.
func (api *API) Get(method string, params map[string]string, result interface{}) {
    request := api.Request(method, result)
    if params != nil {
        request.SetQueryParams(params)
    }

    // Set headers and run request
    request = api.Headers(request)
    _, err := request.Get(api.url + "/" + method)

    if err != nil {
        fmt.Println(err)
    }
}

// Executes a POST request. Loads response into result.
func (api *API) Post(method string, params interface{}, result interface{}) {
    request := api.Request(method, result)
    if params != nil {
        request.SetBody(params)
    }

    // Set headers and run request
    request = api.Headers(request)
    _, err := request.Post(api.url + "/" + method)

    if err != nil {
        fmt.Println(err)
    }
}

// Adds headers
func (api *API) Headers(request *resty.Request) *resty.Request {
    if api.token != "" {
        request.SetHeader("Authorization", "Bearer " + api.token)
    }

    return request
}
