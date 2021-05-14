/*
 * AlphaSOC API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LeaseEventAllOf struct for LeaseEventAllOf
type LeaseEventAllOf struct {
	Type *string `json:"type,omitempty"`
	Termination *bool `json:"termination,omitempty"`
	Duration *int64 `json:"duration,omitempty"`
}

// NewLeaseEventAllOf instantiates a new LeaseEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLeaseEventAllOf() *LeaseEventAllOf {
	this := LeaseEventAllOf{}
	return &this
}

// NewLeaseEventAllOfWithDefaults instantiates a new LeaseEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeaseEventAllOfWithDefaults() *LeaseEventAllOf {
	this := LeaseEventAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LeaseEventAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeaseEventAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LeaseEventAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LeaseEventAllOf) SetType(v string) {
	o.Type = &v
}

// GetTermination returns the Termination field value if set, zero value otherwise.
func (o *LeaseEventAllOf) GetTermination() bool {
	if o == nil || o.Termination == nil {
		var ret bool
		return ret
	}
	return *o.Termination
}

// GetTerminationOk returns a tuple with the Termination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeaseEventAllOf) GetTerminationOk() (*bool, bool) {
	if o == nil || o.Termination == nil {
		return nil, false
	}
	return o.Termination, true
}

// HasTermination returns a boolean if a field has been set.
func (o *LeaseEventAllOf) HasTermination() bool {
	if o != nil && o.Termination != nil {
		return true
	}

	return false
}

// SetTermination gets a reference to the given bool and assigns it to the Termination field.
func (o *LeaseEventAllOf) SetTermination(v bool) {
	o.Termination = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *LeaseEventAllOf) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeaseEventAllOf) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *LeaseEventAllOf) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *LeaseEventAllOf) SetDuration(v int64) {
	o.Duration = &v
}

func (o LeaseEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Termination != nil {
		toSerialize["termination"] = o.Termination
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	return json.Marshal(toSerialize)
}

type NullableLeaseEventAllOf struct {
	value *LeaseEventAllOf
	isSet bool
}

func (v NullableLeaseEventAllOf) Get() *LeaseEventAllOf {
	return v.value
}

func (v *NullableLeaseEventAllOf) Set(val *LeaseEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLeaseEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLeaseEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeaseEventAllOf(val *LeaseEventAllOf) *NullableLeaseEventAllOf {
	return &NullableLeaseEventAllOf{value: val, isSet: true}
}

func (v NullableLeaseEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLeaseEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


