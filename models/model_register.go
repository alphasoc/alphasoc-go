/*
 * REST API
 *
 * This document describes the official AlphaSOC REST API. The primary purpose of the API is to allow a wide variety of clients for sending network telemetry and receiving alerts. API endpoints are complementary to other [data sources](/sources/data-ingestion/) and [alert escalations](/destinations/alert-escalation/) in a way that alerts generated for network telemetry submitted outside of the API are available to download via API and vice versa.  ## Schema  The API can be accessed at `https://api.alphasoc.net` over HTTPS. All requests and responses are encoded in JSON.  ## Compression  As the amount of data transmitted via API can be high, it's advisable to use the compression both ways. Usually HTTP clients transparently support compression when fetching data (by providing `Accept-Encoding` header), but the upload needs to be handled manually. AlphaSOC API supports `gzip` and `deflate` compression algorithms and it's recommended to compress large chunks of data (telemetry) before sending, along with attaching corresponding `Content-Encoding` header.  ## Rate limiting  API counts and limits number of requests from a single API key. The limits are not strictly defined and designed to protect from flooding and accidental errors in client's implementation. In the unlikely case of hitting the limit API returns `429 Too Many Requests` response and expects the client to retry after some time.
 *
 * API version: 1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// Register struct for Register
type Register struct {
	Details *RegisterDetails `json:"details,omitempty"`
}

// NewRegister instantiates a new Register object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegister() *Register {
	this := Register{}
	return &this
}

// NewRegisterWithDefaults instantiates a new Register object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterWithDefaults() *Register {
	this := Register{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Register) GetDetails() RegisterDetails {
	if o == nil || o.Details == nil {
		var ret RegisterDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Register) GetDetailsOk() (*RegisterDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Register) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given RegisterDetails and assigns it to the Details field.
func (o *Register) SetDetails(v RegisterDetails) {
	o.Details = &v
}

func (o Register) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableRegister struct {
	value *Register
	isSet bool
}

func (v NullableRegister) Get() *Register {
	return v.value
}

func (v *NullableRegister) Set(val *Register) {
	v.value = val
	v.isSet = true
}

func (v NullableRegister) IsSet() bool {
	return v.isSet
}

func (v *NullableRegister) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegister(val *Register) *NullableRegister {
	return &NullableRegister{value: val, isSet: true}
}

func (v NullableRegister) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegister) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
