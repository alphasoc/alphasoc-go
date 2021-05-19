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

// AeFlags struct for AeFlags
type AeFlags struct {
	Flags *map[string]Flag `json:"flags,omitempty"`
}

// NewAeFlags instantiates a new AeFlags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeFlags() *AeFlags {
	this := AeFlags{}
	return &this
}

// NewAeFlagsWithDefaults instantiates a new AeFlags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeFlagsWithDefaults() *AeFlags {
	this := AeFlags{}
	return &this
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *AeFlags) GetFlags() map[string]Flag {
	if o == nil || o.Flags == nil {
		var ret map[string]Flag
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeFlags) GetFlagsOk() (*map[string]Flag, bool) {
	if o == nil || o.Flags == nil {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *AeFlags) HasFlags() bool {
	if o != nil && o.Flags != nil {
		return true
	}

	return false
}

// SetFlags gets a reference to the given map[string]Flag and assigns it to the Flags field.
func (o *AeFlags) SetFlags(v map[string]Flag) {
	o.Flags = &v
}

func (o AeFlags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Flags != nil {
		toSerialize["flags"] = o.Flags
	}
	return json.Marshal(toSerialize)
}

type NullableAeFlags struct {
	value *AeFlags
	isSet bool
}

func (v NullableAeFlags) Get() *AeFlags {
	return v.value
}

func (v *NullableAeFlags) Set(val *AeFlags) {
	v.value = val
	v.isSet = true
}

func (v NullableAeFlags) IsSet() bool {
	return v.isSet
}

func (v *NullableAeFlags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeFlags(val *AeFlags) *NullableAeFlags {
	return &NullableAeFlags{value: val, isSet: true}
}

func (v NullableAeFlags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeFlags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

