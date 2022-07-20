/*
 * Container Registry service (CloudAPI)
 *
 * Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their manage Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls.
 *
 * API version: 1.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	_context "context"
	"fmt"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// TokensApiService TokensApi service
type TokensApiService service

type ApiRegistriesTokensDeleteRequest struct {
	ctx        _context.Context
	ApiService *TokensApiService
	registryId string
	tokenId    string
}

func (r ApiRegistriesTokensDeleteRequest) Execute() (*APIResponse, error) {
	return r.ApiService.RegistriesTokensDeleteExecute(r)
}

/*
 * RegistriesTokensDelete Delete token
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param registryId
 * @param tokenId
 * @return ApiRegistriesTokensDeleteRequest
 */
func (a *TokensApiService) RegistriesTokensDelete(ctx _context.Context, registryId string, tokenId string) ApiRegistriesTokensDeleteRequest {
	return ApiRegistriesTokensDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		registryId: registryId,
		tokenId:    tokenId,
	}
}

/*
 * Execute executes the request
 */
func (a *TokensApiService) RegistriesTokensDeleteExecute(r ApiRegistriesTokensDeleteRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokensApiService.RegistriesTokensDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registries/{registryId}/tokens/{tokenId}"
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", _neturl.PathEscape(parameterToString(r.registryId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tokenId"+"}", _neturl.PathEscape(parameterToString(r.tokenId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:   localVarHTTPResponse,
		Method:     localVarHTTPMethod,
		RequestURL: localVarPath,
		Operation:  "RegistriesTokensDelete",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		return localVarAPIResponse, newErr
	}

	return localVarAPIResponse, nil
}

type ApiRegistriesTokensFindByIdRequest struct {
	ctx        _context.Context
	ApiService *TokensApiService
	registryId string
	tokenId    string
}

func (r ApiRegistriesTokensFindByIdRequest) Execute() (TokenResponse, *APIResponse, error) {
	return r.ApiService.RegistriesTokensFindByIdExecute(r)
}

/*
 * RegistriesTokensFindById Get Token Information
 * Gets all information for a specific token used to access a container registry
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param registryId
 * @param tokenId
 * @return ApiRegistriesTokensFindByIdRequest
 */
func (a *TokensApiService) RegistriesTokensFindById(ctx _context.Context, registryId string, tokenId string) ApiRegistriesTokensFindByIdRequest {
	return ApiRegistriesTokensFindByIdRequest{
		ApiService: a,
		ctx:        ctx,
		registryId: registryId,
		tokenId:    tokenId,
	}
}

/*
 * Execute executes the request
 * @return TokenResponse
 */
func (a *TokensApiService) RegistriesTokensFindByIdExecute(r ApiRegistriesTokensFindByIdRequest) (TokenResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokensApiService.RegistriesTokensFindById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registries/{registryId}/tokens/{tokenId}"
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", _neturl.PathEscape(parameterToString(r.registryId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tokenId"+"}", _neturl.PathEscape(parameterToString(r.tokenId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:   localVarHTTPResponse,
		Method:     localVarHTTPMethod,
		RequestURL: localVarPath,
		Operation:  "RegistriesTokensFindById",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiRegistriesTokensGetRequest struct {
	ctx        _context.Context
	ApiService *TokensApiService
	registryId string
	offset     *string
	limit      *string
}

func (r ApiRegistriesTokensGetRequest) Offset(offset string) ApiRegistriesTokensGetRequest {
	r.offset = &offset
	return r
}
func (r ApiRegistriesTokensGetRequest) Limit(limit string) ApiRegistriesTokensGetRequest {
	r.limit = &limit
	return r
}

func (r ApiRegistriesTokensGetRequest) Execute() (TokensResponse, *APIResponse, error) {
	return r.ApiService.RegistriesTokensGetExecute(r)
}

/*
 * RegistriesTokensGet List all tokens for the container registry
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param registryId
 * @return ApiRegistriesTokensGetRequest
 */
func (a *TokensApiService) RegistriesTokensGet(ctx _context.Context, registryId string) ApiRegistriesTokensGetRequest {
	return ApiRegistriesTokensGetRequest{
		ApiService: a,
		ctx:        ctx,
		registryId: registryId,
	}
}

/*
 * Execute executes the request
 * @return TokensResponse
 */
func (a *TokensApiService) RegistriesTokensGetExecute(r ApiRegistriesTokensGetRequest) (TokensResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TokensResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokensApiService.RegistriesTokensGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registries/{registryId}/tokens"
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", _neturl.PathEscape(parameterToString(r.registryId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:   localVarHTTPResponse,
		Method:     localVarHTTPMethod,
		RequestURL: localVarPath,
		Operation:  "RegistriesTokensGet",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiRegistriesTokensPatchRequest struct {
	ctx             _context.Context
	ApiService      *TokensApiService
	registryId      string
	tokenId         string
	patchTokenInput *PatchTokenInput
}

func (r ApiRegistriesTokensPatchRequest) PatchTokenInput(patchTokenInput PatchTokenInput) ApiRegistriesTokensPatchRequest {
	r.patchTokenInput = &patchTokenInput
	return r
}

func (r ApiRegistriesTokensPatchRequest) Execute() (TokenResponse, *APIResponse, error) {
	return r.ApiService.RegistriesTokensPatchExecute(r)
}

/*
 * RegistriesTokensPatch Update token
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param registryId
 * @param tokenId
 * @return ApiRegistriesTokensPatchRequest
 */
func (a *TokensApiService) RegistriesTokensPatch(ctx _context.Context, registryId string, tokenId string) ApiRegistriesTokensPatchRequest {
	return ApiRegistriesTokensPatchRequest{
		ApiService: a,
		ctx:        ctx,
		registryId: registryId,
		tokenId:    tokenId,
	}
}

/*
 * Execute executes the request
 * @return TokenResponse
 */
func (a *TokensApiService) RegistriesTokensPatchExecute(r ApiRegistriesTokensPatchRequest) (TokenResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokensApiService.RegistriesTokensPatch")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registries/{registryId}/tokens/{tokenId}"
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", _neturl.PathEscape(parameterToString(r.registryId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tokenId"+"}", _neturl.PathEscape(parameterToString(r.tokenId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = r.patchTokenInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:   localVarHTTPResponse,
		Method:     localVarHTTPMethod,
		RequestURL: localVarPath,
		Operation:  "RegistriesTokensPatch",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiRegistriesTokensPostRequest struct {
	ctx            _context.Context
	ApiService     *TokensApiService
	registryId     string
	postTokenInput *PostTokenInput
}

func (r ApiRegistriesTokensPostRequest) PostTokenInput(postTokenInput PostTokenInput) ApiRegistriesTokensPostRequest {
	r.postTokenInput = &postTokenInput
	return r
}

func (r ApiRegistriesTokensPostRequest) Execute() (PostTokenOutput, *APIResponse, error) {
	return r.ApiService.RegistriesTokensPostExecute(r)
}

/*
 * RegistriesTokensPost Create token
 * Create a token
- password is only available once in the POST response
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param registryId
 * @return ApiRegistriesTokensPostRequest
*/
func (a *TokensApiService) RegistriesTokensPost(ctx _context.Context, registryId string) ApiRegistriesTokensPostRequest {
	return ApiRegistriesTokensPostRequest{
		ApiService: a,
		ctx:        ctx,
		registryId: registryId,
	}
}

/*
 * Execute executes the request
 * @return PostTokenOutput
 */
func (a *TokensApiService) RegistriesTokensPostExecute(r ApiRegistriesTokensPostRequest) (PostTokenOutput, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PostTokenOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokensApiService.RegistriesTokensPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registries/{registryId}/tokens"
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", _neturl.PathEscape(parameterToString(r.registryId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = r.postTokenInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:   localVarHTTPResponse,
		Method:     localVarHTTPMethod,
		RequestURL: localVarPath,
		Operation:  "RegistriesTokensPost",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiRegistriesTokensPutRequest struct {
	ctx           _context.Context
	ApiService    *TokensApiService
	registryId    string
	tokenId       string
	putTokenInput *PutTokenInput
}

func (r ApiRegistriesTokensPutRequest) PutTokenInput(putTokenInput PutTokenInput) ApiRegistriesTokensPutRequest {
	r.putTokenInput = &putTokenInput
	return r
}

func (r ApiRegistriesTokensPutRequest) Execute() (PutTokenOutput, *APIResponse, error) {
	return r.ApiService.RegistriesTokensPutExecute(r)
}

/*
 * RegistriesTokensPut Create or replace token
 * Create/replace a token
- password is only available once in the create response
- "name" cannot be changed
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param registryId
 * @param tokenId
 * @return ApiRegistriesTokensPutRequest
*/
func (a *TokensApiService) RegistriesTokensPut(ctx _context.Context, registryId string, tokenId string) ApiRegistriesTokensPutRequest {
	return ApiRegistriesTokensPutRequest{
		ApiService: a,
		ctx:        ctx,
		registryId: registryId,
		tokenId:    tokenId,
	}
}

/*
 * Execute executes the request
 * @return PutTokenOutput
 */
func (a *TokensApiService) RegistriesTokensPutExecute(r ApiRegistriesTokensPutRequest) (PutTokenOutput, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PutTokenOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokensApiService.RegistriesTokensPut")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registries/{registryId}/tokens/{tokenId}"
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", _neturl.PathEscape(parameterToString(r.registryId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tokenId"+"}", _neturl.PathEscape(parameterToString(r.tokenId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = r.putTokenInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:   localVarHTTPResponse,
		Method:     localVarHTTPMethod,
		RequestURL: localVarPath,
		Operation:  "RegistriesTokensPut",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}
