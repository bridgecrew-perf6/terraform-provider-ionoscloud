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

// ErrorGroupValidateAllOf struct for ErrorGroupValidateAllOf
type ErrorGroupValidateAllOf struct {
	HttpStatus *int32                       `json:"httpStatus,omitempty"`
	Messages   *[]ErrorGroupValidateMessage `json:"messages,omitempty"`
}

// GetHttpStatus returns the HttpStatus field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ErrorGroupValidateAllOf) GetHttpStatus() *int32 {
	if o == nil {
		return nil
	}

	return o.HttpStatus

}

// GetHttpStatusOk returns a tuple with the HttpStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorGroupValidateAllOf) GetHttpStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.HttpStatus, true
}

// SetHttpStatus sets field value
func (o *ErrorGroupValidateAllOf) SetHttpStatus(v int32) {

	o.HttpStatus = &v

}

// HasHttpStatus returns a boolean if a field has been set.
func (o *ErrorGroupValidateAllOf) HasHttpStatus() bool {
	if o != nil && o.HttpStatus != nil {
		return true
	}

	return false
}

// GetMessages returns the Messages field value
// If the value is explicit nil, the zero value for []ErrorGroupValidateMessage will be returned
func (o *ErrorGroupValidateAllOf) GetMessages() *[]ErrorGroupValidateMessage {
	if o == nil {
		return nil
	}

	return o.Messages

}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorGroupValidateAllOf) GetMessagesOk() (*[]ErrorGroupValidateMessage, bool) {
	if o == nil {
		return nil, false
	}

	return o.Messages, true
}

// SetMessages sets field value
func (o *ErrorGroupValidateAllOf) SetMessages(v []ErrorGroupValidateMessage) {

	o.Messages = &v

}

// HasMessages returns a boolean if a field has been set.
func (o *ErrorGroupValidateAllOf) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

func (o ErrorGroupValidateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.HttpStatus != nil {
		toSerialize["httpStatus"] = o.HttpStatus
	}

	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}

	return json.Marshal(toSerialize)
}

type NullableErrorGroupValidateAllOf struct {
	value *ErrorGroupValidateAllOf
	isSet bool
}

func (v NullableErrorGroupValidateAllOf) Get() *ErrorGroupValidateAllOf {
	return v.value
}

func (v *NullableErrorGroupValidateAllOf) Set(val *ErrorGroupValidateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorGroupValidateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorGroupValidateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorGroupValidateAllOf(val *ErrorGroupValidateAllOf) *NullableErrorGroupValidateAllOf {
	return &NullableErrorGroupValidateAllOf{value: val, isSet: true}
}

func (v NullableErrorGroupValidateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorGroupValidateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
