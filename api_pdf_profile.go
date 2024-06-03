/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.27.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
)


// PDFProfileAPIService PDFProfileAPI service
type PDFProfileAPIService service

type PDFProfileAPIGeneratePdfProfileRequest struct {
	ctx context.Context
	ApiService *PDFProfileAPIService
	profileRequest *ProfileRequest
}

func (r PDFProfileAPIGeneratePdfProfileRequest) ProfileRequest(profileRequest ProfileRequest) PDFProfileAPIGeneratePdfProfileRequest {
	r.profileRequest = &profileRequest
	return r
}

func (r PDFProfileAPIGeneratePdfProfileRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.GeneratePdfProfileExecute(r)
}

/*
GeneratePdfProfile Generate a PDF profile from a Person

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PDFProfileAPIGeneratePdfProfileRequest
*/
func (a *PDFProfileAPIService) GeneratePdfProfile(ctx context.Context) PDFProfileAPIGeneratePdfProfileRequest {
	return PDFProfileAPIGeneratePdfProfileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *PDFProfileAPIService) GeneratePdfProfileExecute(r PDFProfileAPIGeneratePdfProfileRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PDFProfileAPIService.GeneratePdfProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pdf-profile"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.profileRequest == nil {
		return localVarReturnValue, nil, reportError("profileRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/octet-stream", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.profileRequest
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
