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

// RegisterDetails struct for RegisterDetails
type RegisterDetails struct {
	Name         *string   `json:"name,omitempty"`
	Organization *string   `json:"organization,omitempty"`
	Email        *string   `json:"email,omitempty"`
	Phone        *string   `json:"phone,omitempty"`
	Address      *[]string `json:"address,omitempty"`
}

// NewRegisterDetails instantiates a new RegisterDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterDetails() *RegisterDetails {
	this := RegisterDetails{}
	return &this
}

// NewRegisterDetailsWithDefaults instantiates a new RegisterDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterDetailsWithDefaults() *RegisterDetails {
	this := RegisterDetails{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RegisterDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegisterDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegisterDetails) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *RegisterDetails) GetOrganization() string {
	if o == nil || o.Organization == nil {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDetails) GetOrganizationOk() (*string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *RegisterDetails) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *RegisterDetails) SetOrganization(v string) {
	o.Organization = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *RegisterDetails) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDetails) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *RegisterDetails) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *RegisterDetails) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *RegisterDetails) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDetails) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *RegisterDetails) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *RegisterDetails) SetPhone(v string) {
	o.Phone = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *RegisterDetails) GetAddress() []string {
	if o == nil || o.Address == nil {
		var ret []string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDetails) GetAddressOk() (*[]string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *RegisterDetails) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given []string and assigns it to the Address field.
func (o *RegisterDetails) SetAddress(v []string) {
	o.Address = &v
}

func (o RegisterDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableRegisterDetails struct {
	value *RegisterDetails
	isSet bool
}

func (v NullableRegisterDetails) Get() *RegisterDetails {
	return v.value
}

func (v *NullableRegisterDetails) Set(val *RegisterDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterDetails(val *RegisterDetails) *NullableRegisterDetails {
	return &NullableRegisterDetails{value: val, isSet: true}
}

func (v NullableRegisterDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
