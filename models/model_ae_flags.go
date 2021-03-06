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

// AeFlags struct for AeFlags
type AeFlags struct {
	// Dictionary that contains flags descriptions
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
