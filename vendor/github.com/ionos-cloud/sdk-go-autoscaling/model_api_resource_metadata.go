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
	"time"
)

// ApiResourceMetadata struct for ApiResourceMetadata
type ApiResourceMetadata struct {
	CreatedByUserId      *string    `json:"createdByUserId"`
	CreatedDate          *IonosTime `json:"createdDate"`
	LastModifiedBy       *string    `json:"lastModifiedBy,omitempty"`
	LastModifiedByUserId *string    `json:"lastModifiedByUserId,omitempty"`
	LastModifiedDate     *IonosTime `json:"lastModifiedDate,omitempty"`
	State                *string    `json:"state"`
}

// NewApiResourceMetadata instantiates a new ApiResourceMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResourceMetadata(createdByUserId string, createdDate time.Time, state string) *ApiResourceMetadata {
	this := ApiResourceMetadata{}

	this.CreatedByUserId = &createdByUserId
	this.CreatedDate = &IonosTime{createdDate}
	this.State = &state

	return &this
}

// NewApiResourceMetadataWithDefaults instantiates a new ApiResourceMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResourceMetadataWithDefaults() *ApiResourceMetadata {
	this := ApiResourceMetadata{}
	return &this
}

// GetCreatedByUserId returns the CreatedByUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApiResourceMetadata) GetCreatedByUserId() *string {
	if o == nil {
		return nil
	}

	return o.CreatedByUserId

}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiResourceMetadata) GetCreatedByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.CreatedByUserId, true
}

// SetCreatedByUserId sets field value
func (o *ApiResourceMetadata) SetCreatedByUserId(v string) {

	o.CreatedByUserId = &v

}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *ApiResourceMetadata) HasCreatedByUserId() bool {
	if o != nil && o.CreatedByUserId != nil {
		return true
	}

	return false
}

// GetCreatedDate returns the CreatedDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ApiResourceMetadata) GetCreatedDate() *time.Time {
	if o == nil {
		return nil
	}

	if o.CreatedDate == nil {
		return nil
	}
	return &o.CreatedDate.Time

}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiResourceMetadata) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.CreatedDate == nil {
		return nil, false
	}
	return &o.CreatedDate.Time, true

}

// SetCreatedDate sets field value
func (o *ApiResourceMetadata) SetCreatedDate(v time.Time) {

	o.CreatedDate = &IonosTime{v}

}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ApiResourceMetadata) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// GetLastModifiedBy returns the LastModifiedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApiResourceMetadata) GetLastModifiedBy() *string {
	if o == nil {
		return nil
	}

	return o.LastModifiedBy

}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiResourceMetadata) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.LastModifiedBy, true
}

// SetLastModifiedBy sets field value
func (o *ApiResourceMetadata) SetLastModifiedBy(v string) {

	o.LastModifiedBy = &v

}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *ApiResourceMetadata) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// GetLastModifiedByUserId returns the LastModifiedByUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApiResourceMetadata) GetLastModifiedByUserId() *string {
	if o == nil {
		return nil
	}

	return o.LastModifiedByUserId

}

// GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiResourceMetadata) GetLastModifiedByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.LastModifiedByUserId, true
}

// SetLastModifiedByUserId sets field value
func (o *ApiResourceMetadata) SetLastModifiedByUserId(v string) {

	o.LastModifiedByUserId = &v

}

// HasLastModifiedByUserId returns a boolean if a field has been set.
func (o *ApiResourceMetadata) HasLastModifiedByUserId() bool {
	if o != nil && o.LastModifiedByUserId != nil {
		return true
	}

	return false
}

// GetLastModifiedDate returns the LastModifiedDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ApiResourceMetadata) GetLastModifiedDate() *time.Time {
	if o == nil {
		return nil
	}

	if o.LastModifiedDate == nil {
		return nil
	}
	return &o.LastModifiedDate.Time

}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiResourceMetadata) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.LastModifiedDate == nil {
		return nil, false
	}
	return &o.LastModifiedDate.Time, true

}

// SetLastModifiedDate sets field value
func (o *ApiResourceMetadata) SetLastModifiedDate(v time.Time) {

	o.LastModifiedDate = &IonosTime{v}

}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *ApiResourceMetadata) HasLastModifiedDate() bool {
	if o != nil && o.LastModifiedDate != nil {
		return true
	}

	return false
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApiResourceMetadata) GetState() *string {
	if o == nil {
		return nil
	}

	return o.State

}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiResourceMetadata) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.State, true
}

// SetState sets field value
func (o *ApiResourceMetadata) SetState(v string) {

	o.State = &v

}

// HasState returns a boolean if a field has been set.
func (o *ApiResourceMetadata) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

func (o ApiResourceMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedByUserId != nil {
		toSerialize["createdByUserId"] = o.CreatedByUserId
	}

	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}

	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}

	if o.LastModifiedByUserId != nil {
		toSerialize["lastModifiedByUserId"] = o.LastModifiedByUserId
	}

	if o.LastModifiedDate != nil {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}

	if o.State != nil {
		toSerialize["state"] = o.State
	}

	return json.Marshal(toSerialize)
}

type NullableApiResourceMetadata struct {
	value *ApiResourceMetadata
	isSet bool
}

func (v NullableApiResourceMetadata) Get() *ApiResourceMetadata {
	return v.value
}

func (v *NullableApiResourceMetadata) Set(val *ApiResourceMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResourceMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResourceMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResourceMetadata(val *ApiResourceMetadata) *NullableApiResourceMetadata {
	return &NullableApiResourceMetadata{value: val, isSet: true}
}

func (v NullableApiResourceMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResourceMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
