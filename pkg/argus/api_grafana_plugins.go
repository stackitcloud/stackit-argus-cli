/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: v1
Contact: stackit-argus@mail.schwarz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// GrafanaPluginsApiService GrafanaPluginsApi service
type GrafanaPluginsApiService service

type ApiV1GrafanaPluginsCreateRequest struct {
	ctx           context.Context
	ApiService    *GrafanaPluginsApiService
	authorization *string
	data          *V1GrafanaPluginsCreateRequest
}

// Accepts admin auth.
func (r ApiV1GrafanaPluginsCreateRequest) Authorization(authorization string) ApiV1GrafanaPluginsCreateRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1GrafanaPluginsCreateRequest) Data(data V1GrafanaPluginsCreateRequest) ApiV1GrafanaPluginsCreateRequest {
	r.data = &data
	return r
}

func (r ApiV1GrafanaPluginsCreateRequest) Execute() (*GrafanaPluginSingle, *http.Response, error) {
	return r.ApiService.V1GrafanaPluginsCreateExecute(r)
}

/*
V1GrafanaPluginsCreate Method for V1GrafanaPluginsCreate

Create grafana plugins for instance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1GrafanaPluginsCreateRequest
*/
func (a *GrafanaPluginsApiService) V1GrafanaPluginsCreate(ctx context.Context) ApiV1GrafanaPluginsCreateRequest {
	return ApiV1GrafanaPluginsCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GrafanaPluginSingle
func (a *GrafanaPluginsApiService) V1GrafanaPluginsCreateExecute(r ApiV1GrafanaPluginsCreateRequest) (*GrafanaPluginSingle, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GrafanaPluginSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GrafanaPluginsApiService.V1GrafanaPluginsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/grafana-plugins"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.data == nil {
		return localVarReturnValue, nil, reportError("data is required and must be specified")
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
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	// body params
	localVarPostBody = r.data
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v PermissionDenied
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
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

type ApiV1GrafanaPluginsDeleteRequest struct {
	ctx           context.Context
	ApiService    *GrafanaPluginsApiService
	pluginName    string
	authorization *string
}

// Accepts admin auth.
func (r ApiV1GrafanaPluginsDeleteRequest) Authorization(authorization string) ApiV1GrafanaPluginsDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1GrafanaPluginsDeleteRequest) Execute() (*Message, *http.Response, error) {
	return r.ApiService.V1GrafanaPluginsDeleteExecute(r)
}

/*
V1GrafanaPluginsDelete Method for V1GrafanaPluginsDelete

Delete grafana plugin.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pluginName
 @return ApiV1GrafanaPluginsDeleteRequest
*/
func (a *GrafanaPluginsApiService) V1GrafanaPluginsDelete(ctx context.Context, pluginName string) ApiV1GrafanaPluginsDeleteRequest {
	return ApiV1GrafanaPluginsDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		pluginName: pluginName,
	}
}

// Execute executes the request
//  @return Message
func (a *GrafanaPluginsApiService) V1GrafanaPluginsDeleteExecute(r ApiV1GrafanaPluginsDeleteRequest) (*Message, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Message
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GrafanaPluginsApiService.V1GrafanaPluginsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/grafana-plugins/{pluginName}"
	localVarPath = strings.Replace(localVarPath, "{"+"pluginName"+"}", url.PathEscape(parameterToString(r.pluginName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
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
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v Message
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PermissionDenied
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

type ApiV1GrafanaPluginsListRequest struct {
	ctx           context.Context
	ApiService    *GrafanaPluginsApiService
	authorization *string
}

// Accepts admin auth.
func (r ApiV1GrafanaPluginsListRequest) Authorization(authorization string) ApiV1GrafanaPluginsListRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1GrafanaPluginsListRequest) Execute() (*GrafanaPlugin, *http.Response, error) {
	return r.ApiService.V1GrafanaPluginsListExecute(r)
}

/*
V1GrafanaPluginsList Method for V1GrafanaPluginsList

Get grafana plugins for instance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1GrafanaPluginsListRequest
*/
func (a *GrafanaPluginsApiService) V1GrafanaPluginsList(ctx context.Context) ApiV1GrafanaPluginsListRequest {
	return ApiV1GrafanaPluginsListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GrafanaPlugin
func (a *GrafanaPluginsApiService) V1GrafanaPluginsListExecute(r ApiV1GrafanaPluginsListRequest) (*GrafanaPlugin, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GrafanaPlugin
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GrafanaPluginsApiService.V1GrafanaPluginsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/grafana-plugins"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
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
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v PermissionDenied
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

type ApiV1GrafanaPluginsReadRequest struct {
	ctx           context.Context
	ApiService    *GrafanaPluginsApiService
	pluginName    string
	authorization *string
}

// Accepts admin auth.
func (r ApiV1GrafanaPluginsReadRequest) Authorization(authorization string) ApiV1GrafanaPluginsReadRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1GrafanaPluginsReadRequest) Execute() (*GrafanaPluginSingle, *http.Response, error) {
	return r.ApiService.V1GrafanaPluginsReadExecute(r)
}

/*
V1GrafanaPluginsRead Method for V1GrafanaPluginsRead

Get grafana plugin.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pluginName
 @return ApiV1GrafanaPluginsReadRequest
*/
func (a *GrafanaPluginsApiService) V1GrafanaPluginsRead(ctx context.Context, pluginName string) ApiV1GrafanaPluginsReadRequest {
	return ApiV1GrafanaPluginsReadRequest{
		ApiService: a,
		ctx:        ctx,
		pluginName: pluginName,
	}
}

// Execute executes the request
//  @return GrafanaPluginSingle
func (a *GrafanaPluginsApiService) V1GrafanaPluginsReadExecute(r ApiV1GrafanaPluginsReadRequest) (*GrafanaPluginSingle, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GrafanaPluginSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GrafanaPluginsApiService.V1GrafanaPluginsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/grafana-plugins/{pluginName}"
	localVarPath = strings.Replace(localVarPath, "{"+"pluginName"+"}", url.PathEscape(parameterToString(r.pluginName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
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
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v Message
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PermissionDenied
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

type ApiV1GrafanaPluginsUpdateRequest struct {
	ctx           context.Context
	ApiService    *GrafanaPluginsApiService
	pluginName    string
	authorization *string
	data          *V1GrafanaPluginsUpdateRequest
}

// Accepts admin auth.
func (r ApiV1GrafanaPluginsUpdateRequest) Authorization(authorization string) ApiV1GrafanaPluginsUpdateRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1GrafanaPluginsUpdateRequest) Data(data V1GrafanaPluginsUpdateRequest) ApiV1GrafanaPluginsUpdateRequest {
	r.data = &data
	return r
}

func (r ApiV1GrafanaPluginsUpdateRequest) Execute() (*GrafanaPluginSingle, *http.Response, error) {
	return r.ApiService.V1GrafanaPluginsUpdateExecute(r)
}

/*
V1GrafanaPluginsUpdate Method for V1GrafanaPluginsUpdate

Update grafana plugin.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pluginName
 @return ApiV1GrafanaPluginsUpdateRequest
*/
func (a *GrafanaPluginsApiService) V1GrafanaPluginsUpdate(ctx context.Context, pluginName string) ApiV1GrafanaPluginsUpdateRequest {
	return ApiV1GrafanaPluginsUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		pluginName: pluginName,
	}
}

// Execute executes the request
//  @return GrafanaPluginSingle
func (a *GrafanaPluginsApiService) V1GrafanaPluginsUpdateExecute(r ApiV1GrafanaPluginsUpdateRequest) (*GrafanaPluginSingle, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GrafanaPluginSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GrafanaPluginsApiService.V1GrafanaPluginsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/grafana-plugins/{pluginName}"
	localVarPath = strings.Replace(localVarPath, "{"+"pluginName"+"}", url.PathEscape(parameterToString(r.pluginName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.data == nil {
		return localVarReturnValue, nil, reportError("data is required and must be specified")
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
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	// body params
	localVarPostBody = r.data
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v Message
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PermissionDenied
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
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
