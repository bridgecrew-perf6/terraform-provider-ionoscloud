/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
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

// DataCenterApiService DataCenterApi service
type DataCenterApiService service

type ApiDatacentersDeleteRequest struct {
	ctx             _context.Context
	ApiService      *DataCenterApiService
	datacenterId    string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiDatacentersDeleteRequest) Pretty(pretty bool) ApiDatacentersDeleteRequest {
	r.pretty = &pretty
	return r
}
func (r ApiDatacentersDeleteRequest) Depth(depth int32) ApiDatacentersDeleteRequest {
	r.depth = &depth
	return r
}
func (r ApiDatacentersDeleteRequest) XContractNumber(xContractNumber int32) ApiDatacentersDeleteRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiDatacentersDeleteRequest) Execute() (map[string]interface{}, *APIResponse, error) {
	return r.ApiService.DatacentersDeleteExecute(r)
}

/*
 * DatacentersDelete Delete a Data Center
 * Will remove all objects within the datacenter and remove the datacenter object itself, too. This is a highly destructive method which should be used with caution
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param datacenterId The unique ID of the datacenter
 * @return ApiDatacentersDeleteRequest
 */
func (a *DataCenterApiService) DatacentersDelete(ctx _context.Context, datacenterId string) ApiDatacentersDeleteRequest {
	return ApiDatacentersDeleteRequest{
		ApiService:   a,
		ctx:          ctx,
		datacenterId: datacenterId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *DataCenterApiService) DatacentersDeleteExecute(r ApiDatacentersDeleteRequest) (map[string]interface{}, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataCenterApiService.DatacentersDelete")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/datacenters/{datacenterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"datacenterId"+"}", _neturl.PathEscape(parameterToString(r.datacenterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
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
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "DatacentersDelete",
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
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
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

type ApiDatacentersFindByIdRequest struct {
	ctx             _context.Context
	ApiService      *DataCenterApiService
	datacenterId    string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiDatacentersFindByIdRequest) Pretty(pretty bool) ApiDatacentersFindByIdRequest {
	r.pretty = &pretty
	return r
}
func (r ApiDatacentersFindByIdRequest) Depth(depth int32) ApiDatacentersFindByIdRequest {
	r.depth = &depth
	return r
}
func (r ApiDatacentersFindByIdRequest) XContractNumber(xContractNumber int32) ApiDatacentersFindByIdRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiDatacentersFindByIdRequest) Execute() (Datacenter, *APIResponse, error) {
	return r.ApiService.DatacentersFindByIdExecute(r)
}

/*
 * DatacentersFindById Retrieve a Data Center
 * You can retrieve a data center by using the resource's ID. This value can be found in the response body when a datacenter is created or when you GET a list of datacenters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param datacenterId The unique ID of the datacenter
 * @return ApiDatacentersFindByIdRequest
 */
func (a *DataCenterApiService) DatacentersFindById(ctx _context.Context, datacenterId string) ApiDatacentersFindByIdRequest {
	return ApiDatacentersFindByIdRequest{
		ApiService:   a,
		ctx:          ctx,
		datacenterId: datacenterId,
	}
}

/*
 * Execute executes the request
 * @return Datacenter
 */
func (a *DataCenterApiService) DatacentersFindByIdExecute(r ApiDatacentersFindByIdRequest) (Datacenter, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Datacenter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataCenterApiService.DatacentersFindById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/datacenters/{datacenterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"datacenterId"+"}", _neturl.PathEscape(parameterToString(r.datacenterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
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
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "DatacentersFindById",
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
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
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

type ApiDatacentersGetRequest struct {
	ctx             _context.Context
	ApiService      *DataCenterApiService
	filters         _neturl.Values
	orderBy         *string
	maxResults      *int32
	pretty          *bool
	depth           *int32
	xContractNumber *int32
	offset          *int32
	limit           *int32
}

func (r ApiDatacentersGetRequest) Pretty(pretty bool) ApiDatacentersGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiDatacentersGetRequest) Depth(depth int32) ApiDatacentersGetRequest {
	r.depth = &depth
	return r
}
func (r ApiDatacentersGetRequest) XContractNumber(xContractNumber int32) ApiDatacentersGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}
func (r ApiDatacentersGetRequest) Offset(offset int32) ApiDatacentersGetRequest {
	r.offset = &offset
	return r
}
func (r ApiDatacentersGetRequest) Limit(limit int32) ApiDatacentersGetRequest {
	r.limit = &limit
	return r
}

// Filters query parameters limit results to those containing a matching value for a specific property.
func (r ApiDatacentersGetRequest) Filter(key string, value string) ApiDatacentersGetRequest {
	filterKey := fmt.Sprintf(FilterQueryParam, key)
	r.filters[filterKey] = []string{value}
	return r
}

// OrderBy query param sorts the results alphanumerically in ascending order based on the specified property.
func (r ApiDatacentersGetRequest) OrderBy(orderBy string) ApiDatacentersGetRequest {
	r.orderBy = &orderBy
	return r
}

// MaxResults query param limits the number of results returned.
func (r ApiDatacentersGetRequest) MaxResults(maxResults int32) ApiDatacentersGetRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiDatacentersGetRequest) Execute() (Datacenters, *APIResponse, error) {
	return r.ApiService.DatacentersGetExecute(r)
}

/*
 * DatacentersGet List Data Centers under your account
 * You can retrieve a list of data centers provisioned under your account. Default list will contain first 100 items. For more items use pagination query parameters
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiDatacentersGetRequest
 */
func (a *DataCenterApiService) DatacentersGet(ctx _context.Context) ApiDatacentersGetRequest {
	return ApiDatacentersGetRequest{
		ApiService: a,
		ctx:        ctx,
		filters:    _neturl.Values{},
	}
}

/*
 * Execute executes the request
 * @return Datacenters
 */
func (a *DataCenterApiService) DatacentersGetExecute(r ApiDatacentersGetRequest) (Datacenters, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Datacenters
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataCenterApiService.DatacentersGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/datacenters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("offset")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("offset", parameterToString(0, ""))
		}
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("limit")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("limit", parameterToString(1000, ""))
		}
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.maxResults != nil {
		localVarQueryParams.Add("maxResults", parameterToString(*r.maxResults, ""))
	}
	if len(r.filters) > 0 {
		for k, v := range r.filters {
			for _, iv := range v {
				localVarQueryParams.Add(k, iv)
			}
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
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "DatacentersGet",
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
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
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

type ApiDatacentersPatchRequest struct {
	ctx             _context.Context
	ApiService      *DataCenterApiService
	datacenterId    string
	datacenter      *DatacenterProperties
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiDatacentersPatchRequest) Datacenter(datacenter DatacenterProperties) ApiDatacentersPatchRequest {
	r.datacenter = &datacenter
	return r
}
func (r ApiDatacentersPatchRequest) Pretty(pretty bool) ApiDatacentersPatchRequest {
	r.pretty = &pretty
	return r
}
func (r ApiDatacentersPatchRequest) Depth(depth int32) ApiDatacentersPatchRequest {
	r.depth = &depth
	return r
}
func (r ApiDatacentersPatchRequest) XContractNumber(xContractNumber int32) ApiDatacentersPatchRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiDatacentersPatchRequest) Execute() (Datacenter, *APIResponse, error) {
	return r.ApiService.DatacentersPatchExecute(r)
}

/*
 * DatacentersPatch Partially modify a Data Center
 * You can use update datacenter to re-name the datacenter or update its description
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param datacenterId The unique ID of the datacenter
 * @return ApiDatacentersPatchRequest
 */
func (a *DataCenterApiService) DatacentersPatch(ctx _context.Context, datacenterId string) ApiDatacentersPatchRequest {
	return ApiDatacentersPatchRequest{
		ApiService:   a,
		ctx:          ctx,
		datacenterId: datacenterId,
	}
}

/*
 * Execute executes the request
 * @return Datacenter
 */
func (a *DataCenterApiService) DatacentersPatchExecute(r ApiDatacentersPatchRequest) (Datacenter, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Datacenter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataCenterApiService.DatacentersPatch")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/datacenters/{datacenterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"datacenterId"+"}", _neturl.PathEscape(parameterToString(r.datacenterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.datacenter == nil {
		return localVarReturnValue, nil, reportError("datacenter is required and must be specified")
	}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
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
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	// body params
	localVarPostBody = r.datacenter
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "DatacentersPatch",
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
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
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

type ApiDatacentersPostRequest struct {
	ctx             _context.Context
	ApiService      *DataCenterApiService
	datacenter      *Datacenter
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiDatacentersPostRequest) Datacenter(datacenter Datacenter) ApiDatacentersPostRequest {
	r.datacenter = &datacenter
	return r
}
func (r ApiDatacentersPostRequest) Pretty(pretty bool) ApiDatacentersPostRequest {
	r.pretty = &pretty
	return r
}
func (r ApiDatacentersPostRequest) Depth(depth int32) ApiDatacentersPostRequest {
	r.depth = &depth
	return r
}
func (r ApiDatacentersPostRequest) XContractNumber(xContractNumber int32) ApiDatacentersPostRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiDatacentersPostRequest) Execute() (Datacenter, *APIResponse, error) {
	return r.ApiService.DatacentersPostExecute(r)
}

/*
 * DatacentersPost Create a Data Center
 * Virtual data centers are the foundation of the platform. They act as logical containers for all other objects you will be creating, e.g. servers. You can provision as many data centers as you want. Datacenters have their own private network and are logically segmented from each other to create isolation. You can use this POST method to create a simple datacenter or to create a datacenter with multiple objects under it such as servers and storage volumes.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiDatacentersPostRequest
 */
func (a *DataCenterApiService) DatacentersPost(ctx _context.Context) ApiDatacentersPostRequest {
	return ApiDatacentersPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Datacenter
 */
func (a *DataCenterApiService) DatacentersPostExecute(r ApiDatacentersPostRequest) (Datacenter, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Datacenter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataCenterApiService.DatacentersPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/datacenters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.datacenter == nil {
		return localVarReturnValue, nil, reportError("datacenter is required and must be specified")
	}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
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
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	// body params
	localVarPostBody = r.datacenter
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "DatacentersPost",
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
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
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

type ApiDatacentersPutRequest struct {
	ctx             _context.Context
	ApiService      *DataCenterApiService
	datacenterId    string
	datacenter      *Datacenter
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiDatacentersPutRequest) Datacenter(datacenter Datacenter) ApiDatacentersPutRequest {
	r.datacenter = &datacenter
	return r
}
func (r ApiDatacentersPutRequest) Pretty(pretty bool) ApiDatacentersPutRequest {
	r.pretty = &pretty
	return r
}
func (r ApiDatacentersPutRequest) Depth(depth int32) ApiDatacentersPutRequest {
	r.depth = &depth
	return r
}
func (r ApiDatacentersPutRequest) XContractNumber(xContractNumber int32) ApiDatacentersPutRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiDatacentersPutRequest) Execute() (Datacenter, *APIResponse, error) {
	return r.ApiService.DatacentersPutExecute(r)
}

/*
 * DatacentersPut Modify a Data Center
 * You can use update datacenter to re-name the datacenter or update its description
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param datacenterId The unique ID of the datacenter
 * @return ApiDatacentersPutRequest
 */
func (a *DataCenterApiService) DatacentersPut(ctx _context.Context, datacenterId string) ApiDatacentersPutRequest {
	return ApiDatacentersPutRequest{
		ApiService:   a,
		ctx:          ctx,
		datacenterId: datacenterId,
	}
}

/*
 * Execute executes the request
 * @return Datacenter
 */
func (a *DataCenterApiService) DatacentersPutExecute(r ApiDatacentersPutRequest) (Datacenter, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Datacenter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataCenterApiService.DatacentersPut")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/datacenters/{datacenterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"datacenterId"+"}", _neturl.PathEscape(parameterToString(r.datacenterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.datacenter == nil {
		return localVarReturnValue, nil, reportError("datacenter is required and must be specified")
	}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
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
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	// body params
	localVarPostBody = r.datacenter
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "DatacentersPut",
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
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
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