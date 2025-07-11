/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.74.0
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
	"reflect"
)


// ProjectParticipationAPIService ProjectParticipationAPI service
type ProjectParticipationAPIService service

type ProjectParticipationAPIReadPersonProjectParticipationRequest struct {
	ctx context.Context
	ApiService *ProjectParticipationAPIService
	personId string
	projectIds *[]string
}

// Filter by project ids, if not set, all projects are returned
func (r ProjectParticipationAPIReadPersonProjectParticipationRequest) ProjectIds(projectIds []string) ProjectParticipationAPIReadPersonProjectParticipationRequest {
	r.projectIds = &projectIds
	return r
}

func (r ProjectParticipationAPIReadPersonProjectParticipationRequest) Execute() ([]PersonProjectParticipationDetails, *http.Response, error) {
	return r.ApiService.ReadPersonProjectParticipationExecute(r)
}

/*
ReadPersonProjectParticipation Get a Project Participation of a Person

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param personId
 @return ProjectParticipationAPIReadPersonProjectParticipationRequest
*/
func (a *ProjectParticipationAPIService) ReadPersonProjectParticipation(ctx context.Context, personId string) ProjectParticipationAPIReadPersonProjectParticipationRequest {
	return ProjectParticipationAPIReadPersonProjectParticipationRequest{
		ApiService: a,
		ctx: ctx,
		personId: personId,
	}
}

// Execute executes the request
//  @return []PersonProjectParticipationDetails
func (a *ProjectParticipationAPIService) ReadPersonProjectParticipationExecute(r ProjectParticipationAPIReadPersonProjectParticipationRequest) ([]PersonProjectParticipationDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PersonProjectParticipationDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectParticipationAPIService.ReadPersonProjectParticipation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/persons/{personId}/project-participation"
	localVarPath = strings.Replace(localVarPath, "{"+"personId"+"}", url.PathEscape(parameterValueToString(r.personId, "personId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.projectIds != nil {
		t := *r.projectIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "projectIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "projectIds", t, "multi")
		}
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
