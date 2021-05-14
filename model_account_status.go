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
	"time"
)

// AccountStatus struct for AccountStatus
type AccountStatus struct {
	Today *time.Time `json:"today,omitempty"`
	Registered *bool `json:"registered,omitempty"`
	Expired *bool `json:"expired,omitempty"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	EndpointsSeenToday *int32 `json:"endpointsSeenToday,omitempty"`
	Messages *[]Message `json:"messages,omitempty"`
}

// NewAccountStatus instantiates a new AccountStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountStatus() *AccountStatus {
	this := AccountStatus{}
	return &this
}

// NewAccountStatusWithDefaults instantiates a new AccountStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountStatusWithDefaults() *AccountStatus {
	this := AccountStatus{}
	return &this
}

// GetToday returns the Today field value if set, zero value otherwise.
func (o *AccountStatus) GetToday() time.Time {
	if o == nil || o.Today == nil {
		var ret time.Time
		return ret
	}
	return *o.Today
}

// GetTodayOk returns a tuple with the Today field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatus) GetTodayOk() (*time.Time, bool) {
	if o == nil || o.Today == nil {
		return nil, false
	}
	return o.Today, true
}

// HasToday returns a boolean if a field has been set.
func (o *AccountStatus) HasToday() bool {
	if o != nil && o.Today != nil {
		return true
	}

	return false
}

// SetToday gets a reference to the given time.Time and assigns it to the Today field.
func (o *AccountStatus) SetToday(v time.Time) {
	o.Today = &v
}

// GetRegistered returns the Registered field value if set, zero value otherwise.
func (o *AccountStatus) GetRegistered() bool {
	if o == nil || o.Registered == nil {
		var ret bool
		return ret
	}
	return *o.Registered
}

// GetRegisteredOk returns a tuple with the Registered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatus) GetRegisteredOk() (*bool, bool) {
	if o == nil || o.Registered == nil {
		return nil, false
	}
	return o.Registered, true
}

// HasRegistered returns a boolean if a field has been set.
func (o *AccountStatus) HasRegistered() bool {
	if o != nil && o.Registered != nil {
		return true
	}

	return false
}

// SetRegistered gets a reference to the given bool and assigns it to the Registered field.
func (o *AccountStatus) SetRegistered(v bool) {
	o.Registered = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *AccountStatus) GetExpired() bool {
	if o == nil || o.Expired == nil {
		var ret bool
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatus) GetExpiredOk() (*bool, bool) {
	if o == nil || o.Expired == nil {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *AccountStatus) HasExpired() bool {
	if o != nil && o.Expired != nil {
		return true
	}

	return false
}

// SetExpired gets a reference to the given bool and assigns it to the Expired field.
func (o *AccountStatus) SetExpired(v bool) {
	o.Expired = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *AccountStatus) GetExpirationDate() time.Time {
	if o == nil || o.ExpirationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatus) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *AccountStatus) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *AccountStatus) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetEndpointsSeenToday returns the EndpointsSeenToday field value if set, zero value otherwise.
func (o *AccountStatus) GetEndpointsSeenToday() int32 {
	if o == nil || o.EndpointsSeenToday == nil {
		var ret int32
		return ret
	}
	return *o.EndpointsSeenToday
}

// GetEndpointsSeenTodayOk returns a tuple with the EndpointsSeenToday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatus) GetEndpointsSeenTodayOk() (*int32, bool) {
	if o == nil || o.EndpointsSeenToday == nil {
		return nil, false
	}
	return o.EndpointsSeenToday, true
}

// HasEndpointsSeenToday returns a boolean if a field has been set.
func (o *AccountStatus) HasEndpointsSeenToday() bool {
	if o != nil && o.EndpointsSeenToday != nil {
		return true
	}

	return false
}

// SetEndpointsSeenToday gets a reference to the given int32 and assigns it to the EndpointsSeenToday field.
func (o *AccountStatus) SetEndpointsSeenToday(v int32) {
	o.EndpointsSeenToday = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *AccountStatus) GetMessages() []Message {
	if o == nil || o.Messages == nil {
		var ret []Message
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatus) GetMessagesOk() (*[]Message, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *AccountStatus) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []Message and assigns it to the Messages field.
func (o *AccountStatus) SetMessages(v []Message) {
	o.Messages = &v
}

func (o AccountStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Today != nil {
		toSerialize["today"] = o.Today
	}
	if o.Registered != nil {
		toSerialize["registered"] = o.Registered
	}
	if o.Expired != nil {
		toSerialize["expired"] = o.Expired
	}
	if o.ExpirationDate != nil {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if o.EndpointsSeenToday != nil {
		toSerialize["endpointsSeenToday"] = o.EndpointsSeenToday
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullableAccountStatus struct {
	value *AccountStatus
	isSet bool
}

func (v NullableAccountStatus) Get() *AccountStatus {
	return v.value
}

func (v *NullableAccountStatus) Set(val *AccountStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountStatus(val *AccountStatus) *NullableAccountStatus {
	return &NullableAccountStatus{value: val, isSet: true}
}

func (v NullableAccountStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


