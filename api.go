// Base API definition
package txtai

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

// Base API definition
type API struct {
	url string
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

	_, err := request.Post(api.url + "/" + method)

	if err != nil {
		fmt.Println(err)
	}
}
