/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0-SDK.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// GroupShare struct for GroupShare
type GroupShare struct {
	// The resource's unique identifier
	Id *string `json:"id,omitempty"`
	// resource as generic type
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path)
	Href *string `json:"href,omitempty"`
	Properties *GroupShareProperties `json:"properties"`
}


// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupShare) GetId() *string {
	if o == nil {
		return nil
	}


	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupShare) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Id, true
}

// SetId sets field value
func (o *GroupShare) SetId(v string) {


	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *GroupShare) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for Type will be returned
func (o *GroupShare) GetType() *Type {
	if o == nil {
		return nil
	}


	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupShare) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}


	return o.Type, true
}

// SetType sets field value
func (o *GroupShare) SetType(v Type) {


	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *GroupShare) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupShare) GetHref() *string {
	if o == nil {
		return nil
	}


	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupShare) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Href, true
}

// SetHref sets field value
func (o *GroupShare) SetHref(v string) {


	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *GroupShare) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for GroupShareProperties will be returned
func (o *GroupShare) GetProperties() *GroupShareProperties {
	if o == nil {
		return nil
	}


	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupShare) GetPropertiesOk() (*GroupShareProperties, bool) {
	if o == nil {
		return nil, false
	}


	return o.Properties, true
}

// SetProperties sets field value
func (o *GroupShare) SetProperties(v GroupShareProperties) {


	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *GroupShare) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o GroupShare) MarshalJSON() ([]byte, error) {
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

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}
type NullableGroupShare struct {
	value *GroupShare
	isSet bool
}

func (v NullableGroupShare) Get() *GroupShare {
	return v.value
}

func (v *NullableGroupShare) Set(val *GroupShare) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupShare) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupShare(val *GroupShare) *NullableGroupShare {
	return &NullableGroupShare{value: val, isSet: true}
}

func (v NullableGroupShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


