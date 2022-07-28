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

// OsbApiService OsbApi service
type OsbApiService service

type ApiV1ClustersInstancesStatusListRequest struct {
	ctx           context.Context
	ApiService    *OsbApiService
	clusterId     string
	instanceId    string
	authorization *string
}

// Accepts admin auth and broker.
func (r ApiV1ClustersInstancesStatusListRequest) Authorization(authorization string) ApiV1ClustersInstancesStatusListRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1ClustersInstancesStatusListRequest) Execute() (*Status, *http.Response, error) {
	return r.ApiService.V1ClustersInstancesStatusListExecute(r)
}

/*
V1ClustersInstancesStatusList Method for V1ClustersInstancesStatusList

Get deployment status.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterId
 @param instanceId
 @return ApiV1ClustersInstancesStatusListRequest
*/
func (a *OsbApiService) V1ClustersInstancesStatusList(ctx context.Context, clusterId string, instanceId string) ApiV1ClustersInstancesStatusListRequest {
	return ApiV1ClustersInstancesStatusListRequest{
		ApiService: a,
		ctx:        ctx,
		clusterId:  clusterId,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return Status
func (a *OsbApiService) V1ClustersInstancesStatusListExecute(r ApiV1ClustersInstancesStatusListRequest) (*Status, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Status
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsbApiService.V1ClustersInstancesStatusList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/clusters/{clusterId}/instances/{instanceId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", url.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterToString(r.instanceId, "")), -1)

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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Message
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

type ApiV1OsbClustersInstancesCreateRequest struct {
	ctx           context.Context
	ApiService    *OsbApiService
	clusterId     string
	instanceId    string
	authorization *string
	data          *V1OsbClustersInstancesUpdateRequest
}

// Accepts admin auth and broker auth.
func (r ApiV1OsbClustersInstancesCreateRequest) Authorization(authorization string) ApiV1OsbClustersInstancesCreateRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1OsbClustersInstancesCreateRequest) Data(data V1OsbClustersInstancesUpdateRequest) ApiV1OsbClustersInstancesCreateRequest {
	r.data = &data
	return r
}

func (r ApiV1OsbClustersInstancesCreateRequest) Execute() (*Dashboard, *http.Response, error) {
	return r.ApiService.V1OsbClustersInstancesCreateExecute(r)
}

/*
V1OsbClustersInstancesCreate Method for V1OsbClustersInstancesCreate

Create monitoring stack.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterId
 @param instanceId
 @return ApiV1OsbClustersInstancesCreateRequest
*/
func (a *OsbApiService) V1OsbClustersInstancesCreate(ctx context.Context, clusterId string, instanceId string) ApiV1OsbClustersInstancesCreateRequest {
	return ApiV1OsbClustersInstancesCreateRequest{
		ApiService: a,
		ctx:        ctx,
		clusterId:  clusterId,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return Dashboard
func (a *OsbApiService) V1OsbClustersInstancesCreateExecute(r ApiV1OsbClustersInstancesCreateRequest) (*Dashboard, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Dashboard
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsbApiService.V1OsbClustersInstancesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/osb/clusters/{clusterId}/instances/{instanceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", url.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterToString(r.instanceId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 500 {
			var v Message
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

type ApiV1OsbClustersInstancesDeleteRequest struct {
	ctx           context.Context
	ApiService    *OsbApiService
	clusterId     string
	instanceId    string
	authorization *string
}

// Accepts admin auth and broker auth.
func (r ApiV1OsbClustersInstancesDeleteRequest) Authorization(authorization string) ApiV1OsbClustersInstancesDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1OsbClustersInstancesDeleteRequest) Execute() (*Message, *http.Response, error) {
	return r.ApiService.V1OsbClustersInstancesDeleteExecute(r)
}

/*
V1OsbClustersInstancesDelete Method for V1OsbClustersInstancesDelete

Delete a monitoring stack.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterId
 @param instanceId
 @return ApiV1OsbClustersInstancesDeleteRequest
*/
func (a *OsbApiService) V1OsbClustersInstancesDelete(ctx context.Context, clusterId string, instanceId string) ApiV1OsbClustersInstancesDeleteRequest {
	return ApiV1OsbClustersInstancesDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		clusterId:  clusterId,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return Message
func (a *OsbApiService) V1OsbClustersInstancesDeleteExecute(r ApiV1OsbClustersInstancesDeleteRequest) (*Message, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Message
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsbApiService.V1OsbClustersInstancesDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/osb/clusters/{clusterId}/instances/{instanceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", url.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterToString(r.instanceId, "")), -1)

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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Message
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

type ApiV1OsbClustersInstancesUpdateRequest struct {
	ctx           context.Context
	ApiService    *OsbApiService
	clusterId     string
	instanceId    string
	authorization *string
	data          *V1OsbClustersInstancesUpdateRequest
}

// Accepts admin auth and broker auth.
func (r ApiV1OsbClustersInstancesUpdateRequest) Authorization(authorization string) ApiV1OsbClustersInstancesUpdateRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1OsbClustersInstancesUpdateRequest) Data(data V1OsbClustersInstancesUpdateRequest) ApiV1OsbClustersInstancesUpdateRequest {
	r.data = &data
	return r
}

func (r ApiV1OsbClustersInstancesUpdateRequest) Execute() (*Dashboard, *http.Response, error) {
	return r.ApiService.V1OsbClustersInstancesUpdateExecute(r)
}

/*
V1OsbClustersInstancesUpdate Method for V1OsbClustersInstancesUpdate

Update monitoring stack.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterId
 @param instanceId
 @return ApiV1OsbClustersInstancesUpdateRequest
*/
func (a *OsbApiService) V1OsbClustersInstancesUpdate(ctx context.Context, clusterId string, instanceId string) ApiV1OsbClustersInstancesUpdateRequest {
	return ApiV1OsbClustersInstancesUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		clusterId:  clusterId,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return Dashboard
func (a *OsbApiService) V1OsbClustersInstancesUpdateExecute(r ApiV1OsbClustersInstancesUpdateRequest) (*Dashboard, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Dashboard
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsbApiService.V1OsbClustersInstancesUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/osb/clusters/{clusterId}/instances/{instanceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", url.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterToString(r.instanceId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 500 {
			var v Message
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

type ApiV1OsbClustersInstancesUsersCreateRequest struct {
	ctx              context.Context
	ApiService       *OsbApiService
	clusterId        string
	instanceId       string
	serviceBindingId string
	authorization    *string
}

// Accepts system permissions.
func (r ApiV1OsbClustersInstancesUsersCreateRequest) Authorization(authorization string) ApiV1OsbClustersInstancesUsersCreateRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1OsbClustersInstancesUsersCreateRequest) Execute() (*ApiUserCreated, *http.Response, error) {
	return r.ApiService.V1OsbClustersInstancesUsersCreateExecute(r)
}

/*
V1OsbClustersInstancesUsersCreate Method for V1OsbClustersInstancesUsersCreate

Create an api user for an instance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterId
 @param instanceId
 @param serviceBindingId
 @return ApiV1OsbClustersInstancesUsersCreateRequest
*/
func (a *OsbApiService) V1OsbClustersInstancesUsersCreate(ctx context.Context, clusterId string, instanceId string, serviceBindingId string) ApiV1OsbClustersInstancesUsersCreateRequest {
	return ApiV1OsbClustersInstancesUsersCreateRequest{
		ApiService:       a,
		ctx:              ctx,
		clusterId:        clusterId,
		instanceId:       instanceId,
		serviceBindingId: serviceBindingId,
	}
}

// Execute executes the request
//  @return ApiUserCreated
func (a *OsbApiService) V1OsbClustersInstancesUsersCreateExecute(r ApiV1OsbClustersInstancesUsersCreateRequest) (*ApiUserCreated, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiUserCreated
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsbApiService.V1OsbClustersInstancesUsersCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/osb/clusters/{clusterId}/instances/{instanceId}/users/{serviceBindingId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", url.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterToString(r.instanceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceBindingId"+"}", url.PathEscape(parameterToString(r.serviceBindingId, "")), -1)

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

type ApiV1OsbClustersInstancesUsersDeleteRequest struct {
	ctx              context.Context
	ApiService       *OsbApiService
	clusterId        string
	instanceId       string
	serviceBindingId string
	authorization    *string
}

// Accepts system permission.
func (r ApiV1OsbClustersInstancesUsersDeleteRequest) Authorization(authorization string) ApiV1OsbClustersInstancesUsersDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1OsbClustersInstancesUsersDeleteRequest) Execute() (*Message, *http.Response, error) {
	return r.ApiService.V1OsbClustersInstancesUsersDeleteExecute(r)
}

/*
V1OsbClustersInstancesUsersDelete Method for V1OsbClustersInstancesUsersDelete

Delete an api user for an instance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterId
 @param instanceId
 @param serviceBindingId
 @return ApiV1OsbClustersInstancesUsersDeleteRequest
*/
func (a *OsbApiService) V1OsbClustersInstancesUsersDelete(ctx context.Context, clusterId string, instanceId string, serviceBindingId string) ApiV1OsbClustersInstancesUsersDeleteRequest {
	return ApiV1OsbClustersInstancesUsersDeleteRequest{
		ApiService:       a,
		ctx:              ctx,
		clusterId:        clusterId,
		instanceId:       instanceId,
		serviceBindingId: serviceBindingId,
	}
}

// Execute executes the request
//  @return Message
func (a *OsbApiService) V1OsbClustersInstancesUsersDeleteExecute(r ApiV1OsbClustersInstancesUsersDeleteRequest) (*Message, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Message
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsbApiService.V1OsbClustersInstancesUsersDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/osb/clusters/{clusterId}/instances/{instanceId}/users/{serviceBindingId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", url.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterToString(r.instanceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceBindingId"+"}", url.PathEscape(parameterToString(r.serviceBindingId, "")), -1)

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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Message
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

type ApiV1OsbClustersInstancesUsersListRequest struct {
	ctx           context.Context
	ApiService    *OsbApiService
	clusterId     string
	instanceId    string
	authorization *string
}

// Accepts system permissions.
func (r ApiV1OsbClustersInstancesUsersListRequest) Authorization(authorization string) ApiV1OsbClustersInstancesUsersListRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1OsbClustersInstancesUsersListRequest) Execute() (*ApiUser, *http.Response, error) {
	return r.ApiService.V1OsbClustersInstancesUsersListExecute(r)
}

/*
V1OsbClustersInstancesUsersList Method for V1OsbClustersInstancesUsersList

Get all api users.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterId
 @param instanceId
 @return ApiV1OsbClustersInstancesUsersListRequest
*/
func (a *OsbApiService) V1OsbClustersInstancesUsersList(ctx context.Context, clusterId string, instanceId string) ApiV1OsbClustersInstancesUsersListRequest {
	return ApiV1OsbClustersInstancesUsersListRequest{
		ApiService: a,
		ctx:        ctx,
		clusterId:  clusterId,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return ApiUser
func (a *OsbApiService) V1OsbClustersInstancesUsersListExecute(r ApiV1OsbClustersInstancesUsersListRequest) (*ApiUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsbApiService.V1OsbClustersInstancesUsersList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/osb/clusters/{clusterId}/instances/{instanceId}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", url.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterToString(r.instanceId, "")), -1)

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