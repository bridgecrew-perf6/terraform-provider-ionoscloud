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
	"fmt"
)

// TerminationPolicyType The type of the termination policy for the autoscaling group so that a specific pattern is followed for Scaling-In instances. Default termination policy is OLDEST_SERVER_FIRST
type TerminationPolicyType string

// List of TerminationPolicyType
const (
	OLDEST_SERVER_FIRST TerminationPolicyType = "OLDEST_SERVER_FIRST"
	NEWEST_SERVER_FIRST TerminationPolicyType = "NEWEST_SERVER_FIRST"
	RANDOM              TerminationPolicyType = "RANDOM"
)

func (v *TerminationPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TerminationPolicyType(value)
	for _, existing := range []TerminationPolicyType{"OLDEST_SERVER_FIRST", "NEWEST_SERVER_FIRST", "RANDOM"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TerminationPolicyType", value)
}

// Ptr returns reference to TerminationPolicyType value
func (v TerminationPolicyType) Ptr() *TerminationPolicyType {
	return &v
}

type NullableTerminationPolicyType struct {
	value *TerminationPolicyType
	isSet bool
}

func (v NullableTerminationPolicyType) Get() *TerminationPolicyType {
	return v.value
}

func (v *NullableTerminationPolicyType) Set(val *TerminationPolicyType) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminationPolicyType) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminationPolicyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminationPolicyType(val *TerminationPolicyType) *NullableTerminationPolicyType {
	return &NullableTerminationPolicyType{value: val, isSet: true}
}

func (v NullableTerminationPolicyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminationPolicyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
