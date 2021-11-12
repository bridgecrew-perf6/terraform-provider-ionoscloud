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

// GroupCollection struct for GroupCollection
type GroupCollection struct {
	Id    *string                 `json:"id"`
	Type  *map[string]interface{} `json:"type,omitempty"`
	Href  *string                 `json:"href,omitempty"`
	Items *[]GroupResource        `json:"items,omitempty"`
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupCollection) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupCollection) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *GroupCollection) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *GroupCollection) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *GroupCollection) GetType() *map[string]interface{} {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupCollection) GetTypeOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *GroupCollection) SetType(v map[string]interface{}) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *GroupCollection) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupCollection) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupCollection) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *GroupCollection) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *GroupCollection) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []GroupResource will be returned
func (o *GroupCollection) GetItems() *[]GroupResource {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupCollection) GetItemsOk() (*[]GroupResource, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *GroupCollection) SetItems(v []GroupResource) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *GroupCollection) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o GroupCollection) MarshalJSON() ([]byte, error) {
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

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	return json.Marshal(toSerialize)
}

type NullableGroupCollection struct {
	value *GroupCollection
	isSet bool
}

func (v NullableGroupCollection) Get() *GroupCollection {
	return v.value
}

func (v *NullableGroupCollection) Set(val *GroupCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupCollection(val *GroupCollection) *NullableGroupCollection {
	return &NullableGroupCollection{value: val, isSet: true}
}

func (v NullableGroupCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
