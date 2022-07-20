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

// RegistryResponse struct for RegistryResponse
type RegistryResponse struct {
	Href       *string              `json:"href,omitempty"`
	Id         *string              `json:"id,omitempty"`
	Metadata   *ApiResourceMetadata `json:"metadata"`
	Properties *RegistryProperties  `json:"properties"`
	Type       *string              `json:"type,omitempty"`
}

// NewRegistryResponse instantiates a new RegistryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryResponse(metadata ApiResourceMetadata, properties RegistryProperties) *RegistryResponse {
	this := RegistryResponse{}

	this.Metadata = &metadata
	this.Properties = &properties

	return &this
}

// NewRegistryResponseWithDefaults instantiates a new RegistryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryResponseWithDefaults() *RegistryResponse {
	this := RegistryResponse{}
	return &this
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RegistryResponse) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegistryResponse) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *RegistryResponse) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *RegistryResponse) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RegistryResponse) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegistryResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *RegistryResponse) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *RegistryResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for ApiResourceMetadata will be returned
func (o *RegistryResponse) GetMetadata() *ApiResourceMetadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegistryResponse) GetMetadataOk() (*ApiResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *RegistryResponse) SetMetadata(v ApiResourceMetadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *RegistryResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for RegistryProperties will be returned
func (o *RegistryResponse) GetProperties() *RegistryProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegistryResponse) GetPropertiesOk() (*RegistryProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *RegistryResponse) SetProperties(v RegistryProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *RegistryResponse) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RegistryResponse) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegistryResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *RegistryResponse) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *RegistryResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o RegistryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	return json.Marshal(toSerialize)
}

type NullableRegistryResponse struct {
	value *RegistryResponse
	isSet bool
}

func (v NullableRegistryResponse) Get() *RegistryResponse {
	return v.value
}

func (v *NullableRegistryResponse) Set(val *RegistryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryResponse(val *RegistryResponse) *NullableRegistryResponse {
	return &NullableRegistryResponse{value: val, isSet: true}
}

func (v NullableRegistryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
