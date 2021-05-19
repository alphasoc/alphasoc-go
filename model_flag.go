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

// Flag struct for Flag
type Flag struct {
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewFlag instantiates a new Flag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlag() *Flag {
	this := Flag{}
	return &this
}

// NewFlagWithDefaults instantiates a new Flag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlagWithDefaults() *Flag {
	this := Flag{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Flag) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flag) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Flag) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Flag) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Flag) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flag) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Flag) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Flag) SetType(v string) {
	o.Type = &v
}

func (o Flag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFlag struct {
	value *Flag
	isSet bool
}

func (v NullableFlag) Get() *Flag {
	return v.value
}

func (v *NullableFlag) Set(val *Flag) {
	v.value = val
	v.isSet = true
}

func (v NullableFlag) IsSet() bool {
	return v.isSet
}

func (v *NullableFlag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlag(val *Flag) *NullableFlag {
	return &NullableFlag{value: val, isSet: true}
}

func (v NullableFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

