/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.67.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// SkillsProfileAPIService SkillsProfileAPI service
type SkillsProfileAPIService service

type SkillsProfileAPICreateSkillsProfileForPersonRequest struct {
	ctx context.Context
	ApiService *SkillsProfileAPIService
	personId string
	skillsProfileRequest *SkillsProfileRequest
}

func (r SkillsProfileAPICreateSkillsProfileForPersonRequest) SkillsProfileRequest(skillsProfileRequest SkillsProfileRequest) SkillsProfileAPICreateSkillsProfileForPersonRequest {
	r.skillsProfileRequest = &skillsProfileRequest
	return r
}

func (r SkillsProfileAPICreateSkillsProfileForPersonRequest) Execute() (*SkillsProfileDetails, *http.Response, error) {
	return r.ApiService.CreateSkillsProfileForPersonExecute(r)
}

/*
CreateSkillsProfileForPerson Create a SkillsProfile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param personId
 @return SkillsProfileAPICreateSkillsProfileForPersonRequest
*/
func (a *SkillsProfileAPIService) CreateSkillsProfileForPerson(ctx context.Context, personId string) SkillsProfileAPICreateSkillsProfileForPersonRequest {
	return SkillsProfileAPICreateSkillsProfileForPersonRequest{
		ApiService: a,
		ctx: ctx,
		personId: personId,
	}
}

// Execute executes the request
//  @return SkillsProfileDetails
func (a *SkillsProfileAPIService) CreateSkillsProfileForPersonExecute(r SkillsProfileAPICreateSkillsProfileForPersonRequest) (*SkillsProfileDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SkillsProfileDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SkillsProfileAPIService.CreateSkillsProfileForPerson")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/persons/{personId}/skills-profiles"
	localVarPath = strings.Replace(localVarPath, "{"+"personId"+"}", url.PathEscape(parameterValueToString(r.personId, "personId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.skillsProfileRequest == nil {
		return localVarReturnValue, nil, reportError("skillsProfileRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.skillsProfileRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type SkillsProfileAPIDeleteSkillsProfileRequest struct {
	ctx context.Context
	ApiService *SkillsProfileAPIService
	skillsProfileId string
}

func (r SkillsProfileAPIDeleteSkillsProfileRequest) Execute() (*Status, *http.Response, error) {
	return r.ApiService.DeleteSkillsProfileExecute(r)
}

/*
DeleteSkillsProfile Delete a skills profile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skillsProfileId
 @return SkillsProfileAPIDeleteSkillsProfileRequest
*/
func (a *SkillsProfileAPIService) DeleteSkillsProfile(ctx context.Context, skillsProfileId string) SkillsProfileAPIDeleteSkillsProfileRequest {
	return SkillsProfileAPIDeleteSkillsProfileRequest{
		ApiService: a,
		ctx: ctx,
		skillsProfileId: skillsProfileId,
	}
}

// Execute executes the request
//  @return Status
func (a *SkillsProfileAPIService) DeleteSkillsProfileExecute(r SkillsProfileAPIDeleteSkillsProfileRequest) (*Status, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Status
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SkillsProfileAPIService.DeleteSkillsProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/skillsProfiles/{skillsProfileId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skillsProfileId"+"}", url.PathEscape(parameterValueToString(r.skillsProfileId, "skillsProfileId")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type SkillsProfileAPIGetPersonSkillsProfilesRequest struct {
	ctx context.Context
	ApiService *SkillsProfileAPIService
	personId string
}

func (r SkillsProfileAPIGetPersonSkillsProfilesRequest) Execute() (*PagedSkillsProfiles, *http.Response, error) {
	return r.ApiService.GetPersonSkillsProfilesExecute(r)
}

/*
GetPersonSkillsProfiles Get all SkillsProfiles of a single person

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param personId
 @return SkillsProfileAPIGetPersonSkillsProfilesRequest
*/
func (a *SkillsProfileAPIService) GetPersonSkillsProfiles(ctx context.Context, personId string) SkillsProfileAPIGetPersonSkillsProfilesRequest {
	return SkillsProfileAPIGetPersonSkillsProfilesRequest{
		ApiService: a,
		ctx: ctx,
		personId: personId,
	}
}

// Execute executes the request
//  @return PagedSkillsProfiles
func (a *SkillsProfileAPIService) GetPersonSkillsProfilesExecute(r SkillsProfileAPIGetPersonSkillsProfilesRequest) (*PagedSkillsProfiles, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PagedSkillsProfiles
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SkillsProfileAPIService.GetPersonSkillsProfiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/persons/{personId}/skills-profiles"
	localVarPath = strings.Replace(localVarPath, "{"+"personId"+"}", url.PathEscape(parameterValueToString(r.personId, "personId")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type SkillsProfileAPIGetSkillsProfileRequest struct {
	ctx context.Context
	ApiService *SkillsProfileAPIService
	skillsProfileId string
}

func (r SkillsProfileAPIGetSkillsProfileRequest) Execute() (*SkillsProfileDetails, *http.Response, error) {
	return r.ApiService.GetSkillsProfileExecute(r)
}

/*
GetSkillsProfile Get a SkillsProfile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skillsProfileId
 @return SkillsProfileAPIGetSkillsProfileRequest
*/
func (a *SkillsProfileAPIService) GetSkillsProfile(ctx context.Context, skillsProfileId string) SkillsProfileAPIGetSkillsProfileRequest {
	return SkillsProfileAPIGetSkillsProfileRequest{
		ApiService: a,
		ctx: ctx,
		skillsProfileId: skillsProfileId,
	}
}

// Execute executes the request
//  @return SkillsProfileDetails
func (a *SkillsProfileAPIService) GetSkillsProfileExecute(r SkillsProfileAPIGetSkillsProfileRequest) (*SkillsProfileDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SkillsProfileDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SkillsProfileAPIService.GetSkillsProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/skillsProfiles/{skillsProfileId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skillsProfileId"+"}", url.PathEscape(parameterValueToString(r.skillsProfileId, "skillsProfileId")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type SkillsProfileAPIUpdateSkillsProfileRequest struct {
	ctx context.Context
	ApiService *SkillsProfileAPIService
	skillsProfileId string
	skillsProfileRequest *SkillsProfileRequest
}

func (r SkillsProfileAPIUpdateSkillsProfileRequest) SkillsProfileRequest(skillsProfileRequest SkillsProfileRequest) SkillsProfileAPIUpdateSkillsProfileRequest {
	r.skillsProfileRequest = &skillsProfileRequest
	return r
}

func (r SkillsProfileAPIUpdateSkillsProfileRequest) Execute() (*SkillsProfileDetails, *http.Response, error) {
	return r.ApiService.UpdateSkillsProfileExecute(r)
}

/*
UpdateSkillsProfile Update a SkillsProfile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skillsProfileId
 @return SkillsProfileAPIUpdateSkillsProfileRequest
*/
func (a *SkillsProfileAPIService) UpdateSkillsProfile(ctx context.Context, skillsProfileId string) SkillsProfileAPIUpdateSkillsProfileRequest {
	return SkillsProfileAPIUpdateSkillsProfileRequest{
		ApiService: a,
		ctx: ctx,
		skillsProfileId: skillsProfileId,
	}
}

// Execute executes the request
//  @return SkillsProfileDetails
func (a *SkillsProfileAPIService) UpdateSkillsProfileExecute(r SkillsProfileAPIUpdateSkillsProfileRequest) (*SkillsProfileDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SkillsProfileDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SkillsProfileAPIService.UpdateSkillsProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/skillsProfiles/{skillsProfileId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skillsProfileId"+"}", url.PathEscape(parameterValueToString(r.skillsProfileId, "skillsProfileId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.skillsProfileRequest == nil {
		return localVarReturnValue, nil, reportError("skillsProfileRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.skillsProfileRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
