/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.0
Contact: patrick.koss@mail.schwarz
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

// AlertConfigApiService AlertConfigApi service
type AlertConfigApiService service

type ApiV1InstancesAlertconfigsListRequest struct {
	ctx           context.Context
	ApiService    *AlertConfigApiService
	instanceId    string
	authorization *string
}

// Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user
func (r ApiV1InstancesAlertconfigsListRequest) Authorization(authorization string) ApiV1InstancesAlertconfigsListRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1InstancesAlertconfigsListRequest) Execute() (*GetAlert, *http.Response, error) {
	return r.ApiService.V1InstancesAlertconfigsListExecute(r)
}

/*
V1InstancesAlertconfigsList Method for V1InstancesAlertconfigsList

get alert rule config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param instanceId
	@return ApiV1InstancesAlertconfigsListRequest
*/
func (a *AlertConfigApiService) V1InstancesAlertconfigsList(ctx context.Context, instanceId string) ApiV1InstancesAlertconfigsListRequest {
	return ApiV1InstancesAlertconfigsListRequest{
		ApiService: a,
		ctx:        ctx,
		instanceId: instanceId,
	}
}

// Execute executes the request
//
//	@return GetAlert
func (a *AlertConfigApiService) V1InstancesAlertconfigsListExecute(r ApiV1InstancesAlertconfigsListRequest) (*GetAlert, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetAlert
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertConfigApiService.V1InstancesAlertconfigsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instances/{instanceId}/alertconfigs"
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
		if localVarHTTPResponse.StatusCode == 400 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiV1InstancesAlertconfigsUpdateRequest struct {
	ctx                                  context.Context
	ApiService                           *AlertConfigApiService
	instanceId                           string
	authorization                        *string
	v1InstancesAlertconfigsUpdateRequest *V1InstancesAlertconfigsUpdateRequest
}

// Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user
func (r ApiV1InstancesAlertconfigsUpdateRequest) Authorization(authorization string) ApiV1InstancesAlertconfigsUpdateRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1InstancesAlertconfigsUpdateRequest) V1InstancesAlertconfigsUpdateRequest(v1InstancesAlertconfigsUpdateRequest V1InstancesAlertconfigsUpdateRequest) ApiV1InstancesAlertconfigsUpdateRequest {
	r.v1InstancesAlertconfigsUpdateRequest = &v1InstancesAlertconfigsUpdateRequest
	return r
}

func (r ApiV1InstancesAlertconfigsUpdateRequest) Execute() (*PutAlert, *http.Response, error) {
	return r.ApiService.V1InstancesAlertconfigsUpdateExecute(r)
}

/*
V1InstancesAlertconfigsUpdate Method for V1InstancesAlertconfigsUpdate

route for alert config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param instanceId
	@return ApiV1InstancesAlertconfigsUpdateRequest
*/
func (a *AlertConfigApiService) V1InstancesAlertconfigsUpdate(ctx context.Context, instanceId string) ApiV1InstancesAlertconfigsUpdateRequest {
	return ApiV1InstancesAlertconfigsUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		instanceId: instanceId,
	}
}

// Execute executes the request
//
//	@return PutAlert
func (a *AlertConfigApiService) V1InstancesAlertconfigsUpdateExecute(r ApiV1InstancesAlertconfigsUpdateRequest) (*PutAlert, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PutAlert
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertConfigApiService.V1InstancesAlertconfigsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instances/{instanceId}/alertconfigs"
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterToString(r.instanceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.v1InstancesAlertconfigsUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("v1InstancesAlertconfigsUpdateRequest is required and must be specified")
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
	localVarPostBody = r.v1InstancesAlertconfigsUpdateRequest
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiV1ProjectsInstancesAlertconfigsListRequest struct {
	ctx           context.Context
	ApiService    *AlertConfigApiService
	instanceId    string
	projectId     string
	authorization *string
}

// Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user
func (r ApiV1ProjectsInstancesAlertconfigsListRequest) Authorization(authorization string) ApiV1ProjectsInstancesAlertconfigsListRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1ProjectsInstancesAlertconfigsListRequest) Execute() (*GetAlert, *http.Response, error) {
	return r.ApiService.V1ProjectsInstancesAlertconfigsListExecute(r)
}

/*
V1ProjectsInstancesAlertconfigsList Method for V1ProjectsInstancesAlertconfigsList

get alert rule config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param instanceId
	@param projectId
	@return ApiV1ProjectsInstancesAlertconfigsListRequest
*/
func (a *AlertConfigApiService) V1ProjectsInstancesAlertconfigsList(ctx context.Context, instanceId string, projectId string) ApiV1ProjectsInstancesAlertconfigsListRequest {
	return ApiV1ProjectsInstancesAlertconfigsListRequest{
		ApiService: a,
		ctx:        ctx,
		instanceId: instanceId,
		projectId:  projectId,
	}
}

// Execute executes the request
//
//	@return GetAlert
func (a *AlertConfigApiService) V1ProjectsInstancesAlertconfigsListExecute(r ApiV1ProjectsInstancesAlertconfigsListRequest) (*GetAlert, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetAlert
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertConfigApiService.V1ProjectsInstancesAlertconfigsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/instances/{instanceId}/alertconfigs"
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterToString(r.instanceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterToString(r.projectId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiV1ProjectsInstancesAlertconfigsUpdateRequest struct {
	ctx                                  context.Context
	ApiService                           *AlertConfigApiService
	instanceId                           string
	projectId                            string
	authorization                        *string
	v1InstancesAlertconfigsUpdateRequest *V1InstancesAlertconfigsUpdateRequest
}

// Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user
func (r ApiV1ProjectsInstancesAlertconfigsUpdateRequest) Authorization(authorization string) ApiV1ProjectsInstancesAlertconfigsUpdateRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1ProjectsInstancesAlertconfigsUpdateRequest) V1InstancesAlertconfigsUpdateRequest(v1InstancesAlertconfigsUpdateRequest V1InstancesAlertconfigsUpdateRequest) ApiV1ProjectsInstancesAlertconfigsUpdateRequest {
	r.v1InstancesAlertconfigsUpdateRequest = &v1InstancesAlertconfigsUpdateRequest
	return r
}

func (r ApiV1ProjectsInstancesAlertconfigsUpdateRequest) Execute() (*PutAlert, *http.Response, error) {
	return r.ApiService.V1ProjectsInstancesAlertconfigsUpdateExecute(r)
}

/*
V1ProjectsInstancesAlertconfigsUpdate Method for V1ProjectsInstancesAlertconfigsUpdate

route for alert config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param instanceId
	@param projectId
	@return ApiV1ProjectsInstancesAlertconfigsUpdateRequest
*/
func (a *AlertConfigApiService) V1ProjectsInstancesAlertconfigsUpdate(ctx context.Context, instanceId string, projectId string) ApiV1ProjectsInstancesAlertconfigsUpdateRequest {
	return ApiV1ProjectsInstancesAlertconfigsUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		instanceId: instanceId,
		projectId:  projectId,
	}
}

// Execute executes the request
//
//	@return PutAlert
func (a *AlertConfigApiService) V1ProjectsInstancesAlertconfigsUpdateExecute(r ApiV1ProjectsInstancesAlertconfigsUpdateRequest) (*PutAlert, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PutAlert
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertConfigApiService.V1ProjectsInstancesAlertconfigsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/instances/{instanceId}/alertconfigs"
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterToString(r.instanceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterToString(r.projectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.v1InstancesAlertconfigsUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("v1InstancesAlertconfigsUpdateRequest is required and must be specified")
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
	localVarPostBody = r.v1InstancesAlertconfigsUpdateRequest
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
		if localVarHTTPResponse.StatusCode == 404 {
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