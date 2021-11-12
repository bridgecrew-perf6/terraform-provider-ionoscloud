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

// Group A group of virtual servers where replica count can be scaled automatically.
type Group struct {
	Id         *string                 `json:"id"`
	Type       *string                 `json:"type,omitempty"`
	Href       *string                 `json:"href,omitempty"`
	Metadata   *Metadata               `json:"metadata,omitempty"`
	Properties *GroupProperties        `json:"properties,omitempty"`
	Entities   *map[string]interface{} `json:"entities,omitempty"`
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Group) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *Group) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *Group) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Group) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *Group) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *Group) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Group) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *Group) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *Group) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for Metadata will be returned
func (o *Group) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Group) SetMetadata(v Metadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *Group) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for GroupProperties will be returned
func (o *Group) GetProperties() *GroupProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetPropertiesOk() (*GroupProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *Group) SetProperties(v GroupProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *Group) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// GetEntities returns the Entities field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *Group) GetEntities() *map[string]interface{} {
	if o == nil {
		return nil
	}

	return o.Entities

}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetEntitiesOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}

	return o.Entities, true
}

// SetEntities sets field value
func (o *Group) SetEntities(v map[string]interface{}) {

	o.Entities = &v

}

// HasEntities returns a boolean if a field has been set.
func (o *Group) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

func (o Group) MarshalJSON() ([]byte, error) {
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

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}

	return json.Marshal(toSerialize)
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
