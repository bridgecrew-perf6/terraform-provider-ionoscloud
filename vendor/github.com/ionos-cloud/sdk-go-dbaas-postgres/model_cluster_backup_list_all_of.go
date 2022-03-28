/*
 * IONOS DBaaS REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ClusterBackupListAllOf struct for ClusterBackupListAllOf
type ClusterBackupListAllOf struct {
	Type *ResourceType `json:"type,omitempty"`
	// The unique ID of the resource.
	Id    *string           `json:"id,omitempty"`
	Items *[]BackupResponse `json:"items,omitempty"`
}

// NewClusterBackupListAllOf instantiates a new ClusterBackupListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterBackupListAllOf() *ClusterBackupListAllOf {
	this := ClusterBackupListAllOf{}

	return &this
}

// NewClusterBackupListAllOfWithDefaults instantiates a new ClusterBackupListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterBackupListAllOfWithDefaults() *ClusterBackupListAllOf {
	this := ClusterBackupListAllOf{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for ResourceType will be returned
func (o *ClusterBackupListAllOf) GetType() *ResourceType {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterBackupListAllOf) GetTypeOk() (*ResourceType, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *ClusterBackupListAllOf) SetType(v ResourceType) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *ClusterBackupListAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ClusterBackupListAllOf) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterBackupListAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *ClusterBackupListAllOf) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *ClusterBackupListAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []BackupResponse will be returned
func (o *ClusterBackupListAllOf) GetItems() *[]BackupResponse {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterBackupListAllOf) GetItemsOk() (*[]BackupResponse, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *ClusterBackupListAllOf) SetItems(v []BackupResponse) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *ClusterBackupListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o ClusterBackupListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	return json.Marshal(toSerialize)
}

type NullableClusterBackupListAllOf struct {
	value *ClusterBackupListAllOf
	isSet bool
}

func (v NullableClusterBackupListAllOf) Get() *ClusterBackupListAllOf {
	return v.value
}

func (v *NullableClusterBackupListAllOf) Set(val *ClusterBackupListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterBackupListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterBackupListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterBackupListAllOf(val *ClusterBackupListAllOf) *NullableClusterBackupListAllOf {
	return &NullableClusterBackupListAllOf{value: val, isSet: true}
}

func (v NullableClusterBackupListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterBackupListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}