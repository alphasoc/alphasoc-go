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

// KeyRequestResponseBody struct for KeyRequestResponseBody
type KeyRequestResponseBody struct {
	Key *string `json:"key,omitempty"`
}

// NewKeyRequestResponseBody instantiates a new KeyRequestResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyRequestResponseBody() *KeyRequestResponseBody {
	this := KeyRequestResponseBody{}
	return &this
}

// NewKeyRequestResponseBodyWithDefaults instantiates a new KeyRequestResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyRequestResponseBodyWithDefaults() *KeyRequestResponseBody {
	this := KeyRequestResponseBody{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *KeyRequestResponseBody) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRequestResponseBody) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *KeyRequestResponseBody) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *KeyRequestResponseBody) SetKey(v string) {
	o.Key = &v
}

func (o KeyRequestResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableKeyRequestResponseBody struct {
	value *KeyRequestResponseBody
	isSet bool
}

func (v NullableKeyRequestResponseBody) Get() *KeyRequestResponseBody {
	return v.value
}

func (v *NullableKeyRequestResponseBody) Set(val *KeyRequestResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyRequestResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyRequestResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyRequestResponseBody(val *KeyRequestResponseBody) *NullableKeyRequestResponseBody {
	return &NullableKeyRequestResponseBody{value: val, isSet: true}
}

func (v NullableKeyRequestResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyRequestResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


