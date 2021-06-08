/*
 * REST API
 *
 * This document describes the official AlphaSOC REST API. The primary purpose of the API is to allow a wide variety of clients for sending network telemetry and receiving alerts. API endpoints are complementary to other [data sources](/sources/data-ingestion/) and [alert escalations](/destinations/alert-escalation/) in a way that alerts generated for network telemetry submitted outside of the API are available to download via API and vice versa.  ## Schema  The API can be accessed at `https://api.alphasoc.net` over HTTPS. All requests and responses are encoded in JSON.  ## Compression  As the amount of data transmitted via API can be high, it's advisable to use the compression both ways. Usually HTTP clients transparently support compression when fetching data (by providing `Accept-Encoding` header), but the upload needs to be handled manually. AlphaSOC API supports `gzip` and `deflate` compression algorithms and it's recommended to compress large chunks of data (telemetry) before sending, along with attaching corresponding `Content-Encoding` header.  ## Rate limiting  API counts and limits number of requests from a single API key. The limits are not strictly defined and designed to protect from flooding and accidental errors in client's implementation. In the unlikely case of hitting the limit API returns `429 Too Many Requests` response and expects the client to retry after some time.
 *
 * API version: 1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alphasoc

import (
	"encoding/json"
)

// Platform struct for Platform
type Platform struct {
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewPlatform instantiates a new Platform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatform() *Platform {
	this := Platform{}
	return &this
}

// NewPlatformWithDefaults instantiates a new Platform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformWithDefaults() *Platform {
	this := Platform{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Platform) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Platform) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Platform) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Platform) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Platform) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Platform) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Platform) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Platform) SetVersion(v string) {
	o.Version = &v
}

func (o Platform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullablePlatform struct {
	value *Platform
	isSet bool
}

func (v NullablePlatform) Get() *Platform {
	return v.value
}

func (v *NullablePlatform) Set(val *Platform) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatform) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatform(val *Platform) *NullablePlatform {
	return &NullablePlatform{value: val, isSet: true}
}

func (v NullablePlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
