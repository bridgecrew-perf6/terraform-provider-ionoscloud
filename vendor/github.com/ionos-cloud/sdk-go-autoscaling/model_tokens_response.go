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
	"encoding/json"
)

// TokensResponse struct for TokensResponse
type TokensResponse struct {
	Links  *PaginationLinks `json:"_links"`
	Count  *int32           `json:"count"`
	Href   *string          `json:"href,omitempty"`
	Id     *string          `json:"id,omitempty"`
	Items  *[]TokenResponse `json:"items,omitempty"`
	Limit  *int32           `json:"limit"`
	Offset *int32           `json:"offset"`
	Total  *int32           `json:"total"`
	Type   *string          `json:"type,omitempty"`
}

// NewTokensResponse instantiates a new TokensResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokensResponse(links PaginationLinks, count int32, limit int32, offset int32, total int32) *TokensResponse {
	this := TokensResponse{}

	this.Links = &links
	this.Count = &count
	this.Limit = &limit
	this.Offset = &offset
	this.Total = &total

	return &this
}

// NewTokensResponseWithDefaults instantiates a new TokensResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokensResponseWithDefaults() *TokensResponse {
	this := TokensResponse{}
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for PaginationLinks will be returned
func (o *TokensResponse) GetLinks() *PaginationLinks {
	if o == nil {
		return nil
	}

	return o.Links

}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokensResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil {
		return nil, false
	}

	return o.Links, true
}

// SetLinks sets field value
func (o *TokensResponse) SetLinks(v PaginationLinks) {

	o.Links = &v

}

// HasLinks returns a boolean if a field has been set.
func (o *TokensResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// GetCount returns the Count field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TokensResponse) GetCount() *int32 {
	if o == nil {
		return nil
	}

	return o.Count

}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokensResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Count, true
}

// SetCount sets field value
func (o *TokensResponse) SetCount(v int32) {

	o.Count = &v

}

// HasCount returns a boolean if a field has been set.
func (o *TokensResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TokensResponse) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokensResponse) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *TokensResponse) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *TokensResponse) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TokensResponse) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokensResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *TokensResponse) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *TokensResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []TokenResponse will be returned
func (o *TokensResponse) GetItems() *[]TokenResponse {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokensResponse) GetItemsOk() (*[]TokenResponse, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *TokensResponse) SetItems(v []TokenResponse) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *TokensResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// GetLimit returns the Limit field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TokensResponse) GetLimit() *int32 {
	if o == nil {
		return nil
	}

	return o.Limit

}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokensResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Limit, true
}

// SetLimit sets field value
func (o *TokensResponse) SetLimit(v int32) {

	o.Limit = &v

}

// HasLimit returns a boolean if a field has been set.
func (o *TokensResponse) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// GetOffset returns the Offset field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TokensResponse) GetOffset() *int32 {
	if o == nil {
		return nil
	}

	return o.Offset

}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokensResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Offset, true
}

// SetOffset sets field value
func (o *TokensResponse) SetOffset(v int32) {

	o.Offset = &v

}

// HasOffset returns a boolean if a field has been set.
func (o *TokensResponse) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// GetTotal returns the Total field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TokensResponse) GetTotal() *int32 {
	if o == nil {
		return nil
	}

	return o.Total

}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokensResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Total, true
}

// SetTotal sets field value
func (o *TokensResponse) SetTotal(v int32) {

	o.Total = &v

}

// HasTotal returns a boolean if a field has been set.
func (o *TokensResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TokensResponse) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokensResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *TokensResponse) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *TokensResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o TokensResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	if o.Count != nil {
		toSerialize["count"] = o.Count
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}

	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	return json.Marshal(toSerialize)
}

type NullableTokensResponse struct {
	value *TokensResponse
	isSet bool
}

func (v NullableTokensResponse) Get() *TokensResponse {
	return v.value
}

func (v *NullableTokensResponse) Set(val *TokensResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokensResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokensResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokensResponse(val *TokensResponse) *NullableTokensResponse {
	return &NullableTokensResponse{value: val, isSet: true}
}

func (v NullableTokensResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokensResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
