package client

import (
	"context"
	"net/http"
	"net/url"
)

//*************************************************************************************
// NOTE: this file is not generated.
// It makes some of the low level methods accessible
//*************************************************************************************

// PrepareRequest exposes the client `prepareRequest`
func (c *APIClient) PrepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile) (localVarRequest *http.Request, err error) {

	return c.prepareRequest(ctx, path, method, postBody, headerParams, queryParams, formParams, formFiles)
}

// CallAPI exposes the client `callAPI`
func (c *APIClient) CallAPI(request *http.Request) (*http.Response, error) {
	return c.callAPI(request)
}
