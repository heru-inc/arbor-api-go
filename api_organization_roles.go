/*
arborxr-api-v2

This API provides a RESTful interface to interact with your organization's data.

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arborapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// OrganizationRolesAPIService OrganizationRolesAPI service
type OrganizationRolesAPIService service

type ApiOrganizationRolesOrganizationRoleRequest struct {
	ctx context.Context
	ApiService *OrganizationRolesAPIService
	organizationRoleId string
}

func (r ApiOrganizationRolesOrganizationRoleRequest) Execute() (*OrganizationRole, *http.Response, error) {
	return r.ApiService.OrganizationRolesOrganizationRoleExecute(r)
}

/*
OrganizationRolesOrganizationRole Method for OrganizationRolesOrganizationRole

Get a organization role by id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationRoleId
 @return ApiOrganizationRolesOrganizationRoleRequest
*/
func (a *OrganizationRolesAPIService) OrganizationRolesOrganizationRole(ctx context.Context, organizationRoleId string) ApiOrganizationRolesOrganizationRoleRequest {
	return ApiOrganizationRolesOrganizationRoleRequest{
		ApiService: a,
		ctx: ctx,
		organizationRoleId: organizationRoleId,
	}
}

// Execute executes the request
//  @return OrganizationRole
func (a *OrganizationRolesAPIService) OrganizationRolesOrganizationRoleExecute(r ApiOrganizationRolesOrganizationRoleRequest) (*OrganizationRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrganizationRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationRolesAPIService.OrganizationRolesOrganizationRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization-roles/{organizationRoleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationRoleId"+"}", url.PathEscape(parameterValueToString(r.organizationRoleId, "organizationRoleId")), -1)

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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ForbiddenErrorBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v NotFoundErrorBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiOrganizationRolesOrganizationRolesRequest struct {
	ctx context.Context
	ApiService *OrganizationRolesAPIService
	perPage *int32
	page *int32
}

func (r ApiOrganizationRolesOrganizationRolesRequest) PerPage(perPage int32) ApiOrganizationRolesOrganizationRolesRequest {
	r.perPage = &perPage
	return r
}

func (r ApiOrganizationRolesOrganizationRolesRequest) Page(page int32) ApiOrganizationRolesOrganizationRolesRequest {
	r.page = &page
	return r
}

func (r ApiOrganizationRolesOrganizationRolesRequest) Execute() (*OrganizationRolesResponse, *http.Response, error) {
	return r.ApiService.OrganizationRolesOrganizationRolesExecute(r)
}

/*
OrganizationRolesOrganizationRoles Method for OrganizationRolesOrganizationRoles

Get a paginated list of organization roles.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOrganizationRolesOrganizationRolesRequest
*/
func (a *OrganizationRolesAPIService) OrganizationRolesOrganizationRoles(ctx context.Context) ApiOrganizationRolesOrganizationRolesRequest {
	return ApiOrganizationRolesOrganizationRolesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OrganizationRolesResponse
func (a *OrganizationRolesAPIService) OrganizationRolesOrganizationRolesExecute(r ApiOrganizationRolesOrganizationRolesRequest) (*OrganizationRolesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrganizationRolesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationRolesAPIService.OrganizationRolesOrganizationRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization-roles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
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
