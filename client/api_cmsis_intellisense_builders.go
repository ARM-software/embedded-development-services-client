/*
Solar API

This API provides a RESTful interface to all the Solar services e.g. looking for boards, building projects, etc. - This API uses Hypermedia as the Engine of Application State (HATEOAS) to drive the discovery and provide   affordances. - Discovery is possible by following links from the well known root resource. While this specification lists   all supported endpoints, it is only recommended that these are hard coded into a client if code generation is   being used. Otherwise, it is recommended that the discovery mechanisms present in the resources (affordances)   are used exclusively. - Affordances are links which indicate whether an action is currently possible, this is significantly different from   whether the service supports an action in general. This specification defines what actions could be possible,   but only by checking the affordances returned by the API in the returned resources, can a client determine whether   this action is currently possible or available for the current user. For example:   - An operation to modify a resource could be defined in this specification, but the user may lack the appropriate     privileges. In that situation, the affordance link would not be present in the resource when read. Therefore,     the client can infer that it is not possible to edit this resource and present appropriate information to the     user.   - An operation to delete a resource could be defined and be possible in some circumstances. The specification     describes that the delete is supported and how to use it, but the affordance describes whether it is currently     possible. The logic in the API may dictate that if the resource was in use (perhaps it is a running job or used     by another resource), then it will not be possible to delete that resource as it would result in a conflicted     state. - It is strongly encouraged that affordances are used by all clients, even those using code generation. This has the   ability to both improve robustness and the user experience by decoupling the client and server. For example, if for   some reason the criteria for deleting a resource changes, the logic is only implemented in the server and there is   no need to update the logic in the client as it is driven by the affordances. - The format used for the resources is the Hypertext Application Language (HAL), which includes the definition   of links and embedded resources. 

API version: 1.0.0
Contact: support@arm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// CMSISIntellisenseBuildersApiService CMSISIntellisenseBuildersApi service
type CMSISIntellisenseBuildersApiService service

type ApiGetCmsisIntellisenseRequest struct {
	ctx context.Context
	ApiService *CMSISIntellisenseBuildersApiService
	builderName string
	acceptVersion *string
	ifNoneMatch *string
}

// Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning.
func (r ApiGetCmsisIntellisenseRequest) AcceptVersion(acceptVersion string) ApiGetCmsisIntellisenseRequest {
	r.acceptVersion = &acceptVersion
	return r
}

// Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content.
func (r ApiGetCmsisIntellisenseRequest) IfNoneMatch(ifNoneMatch string) ApiGetCmsisIntellisenseRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

func (r ApiGetCmsisIntellisenseRequest) Execute() (*CmsisBuilderItem, *http.Response, error) {
	return r.ApiService.GetCmsisIntellisenseExecute(r)
}

/*
GetCmsisIntellisense Return details of specific CMSIS Intellisense Builders.

A CMSIS Intellisense Builder is a combination of a specific set of CMSIS-Build tools and a specific toolchain, which can be used to build a CMSIS project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param builderName The ID of the CMSIS Intellisense Builder.
 @return ApiGetCmsisIntellisenseRequest
*/
func (a *CMSISIntellisenseBuildersApiService) GetCmsisIntellisense(ctx context.Context, builderName string) ApiGetCmsisIntellisenseRequest {
	return ApiGetCmsisIntellisenseRequest{
		ApiService: a,
		ctx: ctx,
		builderName: builderName,
	}
}

// Execute executes the request
//  @return CmsisBuilderItem
func (a *CMSISIntellisenseBuildersApiService) GetCmsisIntellisenseExecute(r ApiGetCmsisIntellisenseRequest) (*CmsisBuilderItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CmsisBuilderItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CMSISIntellisenseBuildersApiService.GetCmsisIntellisense")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cmsis-intellisense/{builderName}"
	localVarPath = strings.Replace(localVarPath, "{"+"builderName"+"}", url.PathEscape(parameterToString(r.builderName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptVersion != nil {
		localVarHeaderParams["Accept-Version"] = parameterToString(*r.acceptVersion, "")
	}
	if r.ifNoneMatch != nil {
		localVarHeaderParams["if-none-match"] = parameterToString(*r.ifNoneMatch, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 406 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListCmsisIntellisenseRequest struct {
	ctx context.Context
	ApiService *CMSISIntellisenseBuildersApiService
	acceptVersion *string
	embed *bool
	ifNoneMatch *string
	limit *int32
	offset *int32
}

// Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning.
func (r ApiListCmsisIntellisenseRequest) AcceptVersion(acceptVersion string) ApiListCmsisIntellisenseRequest {
	r.acceptVersion = &acceptVersion
	return r
}

// Embedding: The whether or not to embed resources into the collection (rather than return links).
func (r ApiListCmsisIntellisenseRequest) Embed(embed bool) ApiListCmsisIntellisenseRequest {
	r.embed = &embed
	return r
}

// Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content.
func (r ApiListCmsisIntellisenseRequest) IfNoneMatch(ifNoneMatch string) ApiListCmsisIntellisenseRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

// Paging: The maximum number of items to return in a resource.
func (r ApiListCmsisIntellisenseRequest) Limit(limit int32) ApiListCmsisIntellisenseRequest {
	r.limit = &limit
	return r
}

// Paging:  The index of the first item to return in the resource.
func (r ApiListCmsisIntellisenseRequest) Offset(offset int32) ApiListCmsisIntellisenseRequest {
	r.offset = &offset
	return r
}

func (r ApiListCmsisIntellisenseRequest) Execute() (*CmsisIntellisenseCollection, *http.Response, error) {
	return r.ApiService.ListCmsisIntellisenseExecute(r)
}

/*
ListCmsisIntellisense List available CMSIS Intellisense Builders.

This returns a collection resource that lists all of the CMSIS Intellisense Builders supported by the build service. Each builder listed represents a specific combination of the following:
  - CMSIS-Build tool version
  - Toolchain type e.g. GCC/AC6
  - Toolchain version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListCmsisIntellisenseRequest
*/
func (a *CMSISIntellisenseBuildersApiService) ListCmsisIntellisense(ctx context.Context) ApiListCmsisIntellisenseRequest {
	return ApiListCmsisIntellisenseRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CmsisIntellisenseCollection
func (a *CMSISIntellisenseBuildersApiService) ListCmsisIntellisenseExecute(r ApiListCmsisIntellisenseRequest) (*CmsisIntellisenseCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CmsisIntellisenseCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CMSISIntellisenseBuildersApiService.ListCmsisIntellisense")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cmsis-intellisense/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.embed != nil {
		localVarQueryParams.Add("embed", parameterToString(*r.embed, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptVersion != nil {
		localVarHeaderParams["Accept-Version"] = parameterToString(*r.acceptVersion, "")
	}
	if r.ifNoneMatch != nil {
		localVarHeaderParams["if-none-match"] = parameterToString(*r.ifNoneMatch, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 406 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStartCmsisIntellisenseRequest struct {
	ctx context.Context
	ApiService *CMSISIntellisenseBuildersApiService
	builderName string
	intellisenseJobItem *IntellisenseJobItem
	acceptVersion *string
}

// A name of the CMSIS project to generate compilation database.
func (r ApiStartCmsisIntellisenseRequest) IntellisenseJobItem(intellisenseJobItem IntellisenseJobItem) ApiStartCmsisIntellisenseRequest {
	r.intellisenseJobItem = &intellisenseJobItem
	return r
}

// Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning.
func (r ApiStartCmsisIntellisenseRequest) AcceptVersion(acceptVersion string) ApiStartCmsisIntellisenseRequest {
	r.acceptVersion = &acceptVersion
	return r
}

func (r ApiStartCmsisIntellisenseRequest) Execute() (*IntellisenseJobItem, *http.Response, error) {
	return r.ApiService.StartCmsisIntellisenseExecute(r)
}

/*
StartCmsisIntellisense Initiate a compilation database generation using the specified CMSIS Intellisense Builder.

Initiate a compilation database generation using the specified CMSIS Intellisense Builder.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param builderName The ID of the CMSIS Intellisense Builder.
 @return ApiStartCmsisIntellisenseRequest
*/
func (a *CMSISIntellisenseBuildersApiService) StartCmsisIntellisense(ctx context.Context, builderName string) ApiStartCmsisIntellisenseRequest {
	return ApiStartCmsisIntellisenseRequest{
		ApiService: a,
		ctx: ctx,
		builderName: builderName,
	}
}

// Execute executes the request
//  @return IntellisenseJobItem
func (a *CMSISIntellisenseBuildersApiService) StartCmsisIntellisenseExecute(r ApiStartCmsisIntellisenseRequest) (*IntellisenseJobItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IntellisenseJobItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CMSISIntellisenseBuildersApiService.StartCmsisIntellisense")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cmsis-intellisense/{builderName}"
	localVarPath = strings.Replace(localVarPath, "{"+"builderName"+"}", url.PathEscape(parameterToString(r.builderName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.intellisenseJobItem == nil {
		return localVarReturnValue, nil, reportError("intellisenseJobItem is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptVersion != nil {
		localVarHeaderParams["Accept-Version"] = parameterToString(*r.acceptVersion, "")
	}
	// body params
	localVarPostBody = r.intellisenseJobItem
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 406 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
