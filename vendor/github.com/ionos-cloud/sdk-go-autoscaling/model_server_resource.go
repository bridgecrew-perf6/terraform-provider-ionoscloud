/*
 * VM Auto Scaling service (CloudAPI)
 *
 * VM Auto Scaling service enables IONOS clients to horizontally scale the number of VM instances, based on configured rules. Use Auto Scaling to ensure you will have a sufficient number of instances to handle your application loads at all times.  Create an Auto Scaling group that contains the server instances; Auto Scaling service will ensure that the number of instances in the group is always within these limits.  When target replica count is specified, Auto Scaling will maintain the set number on instances.  When scaling policies are specified, Auto Scaling will create or delete instances based on the demands of your applications. For each policy, specified scale-in and scale-out actions are performed whenever the corresponding thresholds are met.
 *
 * API version: 1-SDK.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ServerResource struct for ServerResource
type ServerResource struct {
	// The resource's unique identifier
	Id *string `json:"id"`
	// The type of object that has been created
	Type *string `json:"type,omitempty"`
	// URL to the object representation (absolute path)
	Href *string `json:"href,omitempty"`
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerResource) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *ServerResource) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *ServerResource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerResource) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *ServerResource) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *ServerResource) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerResource) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerResource) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *ServerResource) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *ServerResource) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

func (o ServerResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	return json.Marshal(toSerialize)
}

type NullableServerResource struct {
	value *ServerResource
	isSet bool
}

func (v NullableServerResource) Get() *ServerResource {
	return v.value
}

func (v *NullableServerResource) Set(val *ServerResource) {
	v.value = val
	v.isSet = true
}

func (v NullableServerResource) IsSet() bool {
	return v.isSet
}

func (v *NullableServerResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerResource(val *ServerResource) *NullableServerResource {
	return &NullableServerResource{value: val, isSet: true}
}

func (v NullableServerResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
