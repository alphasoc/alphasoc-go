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

// KeyRequestBody struct for KeyRequestBody
type KeyRequestBody struct {
	Platform *Platform `json:"platform,omitempty"`
	Uname *string `json:"uname,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NewKeyRequestBody instantiates a new KeyRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyRequestBody() *KeyRequestBody {
	this := KeyRequestBody{}
	return &this
}

// NewKeyRequestBodyWithDefaults instantiates a new KeyRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyRequestBodyWithDefaults() *KeyRequestBody {
	this := KeyRequestBody{}
	return &this
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *KeyRequestBody) GetPlatform() Platform {
	if o == nil || o.Platform == nil {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRequestBody) GetPlatformOk() (*Platform, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *KeyRequestBody) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *KeyRequestBody) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetUname returns the Uname field value if set, zero value otherwise.
func (o *KeyRequestBody) GetUname() string {
	if o == nil || o.Uname == nil {
		var ret string
		return ret
	}
	return *o.Uname
}

// GetUnameOk returns a tuple with the Uname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRequestBody) GetUnameOk() (*string, bool) {
	if o == nil || o.Uname == nil {
		return nil, false
	}
	return o.Uname, true
}

// HasUname returns a boolean if a field has been set.
func (o *KeyRequestBody) HasUname() bool {
	if o != nil && o.Uname != nil {
		return true
	}

	return false
}

// SetUname gets a reference to the given string and assigns it to the Uname field.
func (o *KeyRequestBody) SetUname(v string) {
	o.Uname = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *KeyRequestBody) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRequestBody) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *KeyRequestBody) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *KeyRequestBody) SetToken(v string) {
	o.Token = &v
}

func (o KeyRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.Uname != nil {
		toSerialize["uname"] = o.Uname
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableKeyRequestBody struct {
	value *KeyRequestBody
	isSet bool
}

func (v NullableKeyRequestBody) Get() *KeyRequestBody {
	return v.value
}

func (v *NullableKeyRequestBody) Set(val *KeyRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyRequestBody(val *KeyRequestBody) *NullableKeyRequestBody {
	return &NullableKeyRequestBody{value: val, isSet: true}
}

func (v NullableKeyRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

