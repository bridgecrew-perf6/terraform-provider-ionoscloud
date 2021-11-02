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
	"encoding/json"
)

// KubernetesMaintenanceWindow struct for KubernetesMaintenanceWindow
type KubernetesMaintenanceWindow struct {
	// The day of the week for a maintenance window.
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty"`
	// The time to use for a maintenance window. Accepted formats are: HH:mm:ss; HH:mm:ss\"Z\"; HH:mm:ssZ. This time may varies by 15 minutes.
	Time *string `json:"time,omitempty"`
}



// GetDayOfTheWeek returns the DayOfTheWeek field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesMaintenanceWindow) GetDayOfTheWeek() *string {
	if o == nil {
		return nil
	}


	return o.DayOfTheWeek

}

// GetDayOfTheWeekOk returns a tuple with the DayOfTheWeek field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesMaintenanceWindow) GetDayOfTheWeekOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.DayOfTheWeek, true
}

// SetDayOfTheWeek sets field value
func (o *KubernetesMaintenanceWindow) SetDayOfTheWeek(v string) {


	o.DayOfTheWeek = &v

}

// HasDayOfTheWeek returns a boolean if a field has been set.
func (o *KubernetesMaintenanceWindow) HasDayOfTheWeek() bool {
	if o != nil && o.DayOfTheWeek != nil {
		return true
	}

	return false
}


// GetTime returns the Time field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesMaintenanceWindow) GetTime() *string {
	if o == nil {
		return nil
	}


	return o.Time

}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesMaintenanceWindow) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Time, true
}

// SetTime sets field value
func (o *KubernetesMaintenanceWindow) SetTime(v string) {


	o.Time = &v

}

// HasTime returns a boolean if a field has been set.
func (o *KubernetesMaintenanceWindow) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

func (o KubernetesMaintenanceWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.DayOfTheWeek != nil {
		toSerialize["dayOfTheWeek"] = o.DayOfTheWeek
	}

	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesMaintenanceWindow struct {
	value *KubernetesMaintenanceWindow
	isSet bool
}

func (v NullableKubernetesMaintenanceWindow) Get() *KubernetesMaintenanceWindow {
	return v.value
}

func (v *NullableKubernetesMaintenanceWindow) Set(val *KubernetesMaintenanceWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesMaintenanceWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesMaintenanceWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesMaintenanceWindow(val *KubernetesMaintenanceWindow) *NullableKubernetesMaintenanceWindow {
	return &NullableKubernetesMaintenanceWindow{value: val, isSet: true}
}

func (v NullableKubernetesMaintenanceWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesMaintenanceWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

