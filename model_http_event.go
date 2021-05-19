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

// HttpEvent struct for HttpEvent
type HttpEvent struct {
	Ts *time.Time `json:"ts,omitempty"`
	SrcIP *string `json:"srcIP,omitempty"`
	SrcPort *int32 `json:"srcPort,omitempty"`
	SrcHost *string `json:"srcHost,omitempty"`
	SrcMac *string `json:"srcMac,omitempty"`
	SrcUser *string `json:"srcUser,omitempty"`
	SrcID *string `json:"srcID,omitempty"`
	Url *string `json:"url,omitempty"`
	Method *string `json:"method,omitempty"`
	Status *int64 `json:"status,omitempty"`
	App *string `json:"app,omitempty"`
	Action *string `json:"action,omitempty"`
	BytesIn *int64 `json:"bytesIn,omitempty"`
	BytesOut *int64 `json:"bytesOut,omitempty"`
	ContentType *string `json:"contentType,omitempty"`
	Referrer *string `json:"referrer,omitempty"`
	UserAgent *string `json:"userAgent,omitempty"`
}

// NewHttpEvent instantiates a new HttpEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpEvent() *HttpEvent {
	this := HttpEvent{}
	return &this
}

// NewHttpEventWithDefaults instantiates a new HttpEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpEventWithDefaults() *HttpEvent {
	this := HttpEvent{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *HttpEvent) GetTs() time.Time {
	if o == nil || o.Ts == nil {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetTsOk() (*time.Time, bool) {
	if o == nil || o.Ts == nil {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *HttpEvent) HasTs() bool {
	if o != nil && o.Ts != nil {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *HttpEvent) SetTs(v time.Time) {
	o.Ts = &v
}

// GetSrcIP returns the SrcIP field value if set, zero value otherwise.
func (o *HttpEvent) GetSrcIP() string {
	if o == nil || o.SrcIP == nil {
		var ret string
		return ret
	}
	return *o.SrcIP
}

// GetSrcIPOk returns a tuple with the SrcIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetSrcIPOk() (*string, bool) {
	if o == nil || o.SrcIP == nil {
		return nil, false
	}
	return o.SrcIP, true
}

// HasSrcIP returns a boolean if a field has been set.
func (o *HttpEvent) HasSrcIP() bool {
	if o != nil && o.SrcIP != nil {
		return true
	}

	return false
}

// SetSrcIP gets a reference to the given string and assigns it to the SrcIP field.
func (o *HttpEvent) SetSrcIP(v string) {
	o.SrcIP = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *HttpEvent) GetSrcPort() int32 {
	if o == nil || o.SrcPort == nil {
		var ret int32
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetSrcPortOk() (*int32, bool) {
	if o == nil || o.SrcPort == nil {
		return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *HttpEvent) HasSrcPort() bool {
	if o != nil && o.SrcPort != nil {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given int32 and assigns it to the SrcPort field.
func (o *HttpEvent) SetSrcPort(v int32) {
	o.SrcPort = &v
}

// GetSrcHost returns the SrcHost field value if set, zero value otherwise.
func (o *HttpEvent) GetSrcHost() string {
	if o == nil || o.SrcHost == nil {
		var ret string
		return ret
	}
	return *o.SrcHost
}

// GetSrcHostOk returns a tuple with the SrcHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetSrcHostOk() (*string, bool) {
	if o == nil || o.SrcHost == nil {
		return nil, false
	}
	return o.SrcHost, true
}

// HasSrcHost returns a boolean if a field has been set.
func (o *HttpEvent) HasSrcHost() bool {
	if o != nil && o.SrcHost != nil {
		return true
	}

	return false
}

// SetSrcHost gets a reference to the given string and assigns it to the SrcHost field.
func (o *HttpEvent) SetSrcHost(v string) {
	o.SrcHost = &v
}

// GetSrcMac returns the SrcMac field value if set, zero value otherwise.
func (o *HttpEvent) GetSrcMac() string {
	if o == nil || o.SrcMac == nil {
		var ret string
		return ret
	}
	return *o.SrcMac
}

// GetSrcMacOk returns a tuple with the SrcMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetSrcMacOk() (*string, bool) {
	if o == nil || o.SrcMac == nil {
		return nil, false
	}
	return o.SrcMac, true
}

// HasSrcMac returns a boolean if a field has been set.
func (o *HttpEvent) HasSrcMac() bool {
	if o != nil && o.SrcMac != nil {
		return true
	}

	return false
}

// SetSrcMac gets a reference to the given string and assigns it to the SrcMac field.
func (o *HttpEvent) SetSrcMac(v string) {
	o.SrcMac = &v
}

// GetSrcUser returns the SrcUser field value if set, zero value otherwise.
func (o *HttpEvent) GetSrcUser() string {
	if o == nil || o.SrcUser == nil {
		var ret string
		return ret
	}
	return *o.SrcUser
}

// GetSrcUserOk returns a tuple with the SrcUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetSrcUserOk() (*string, bool) {
	if o == nil || o.SrcUser == nil {
		return nil, false
	}
	return o.SrcUser, true
}

// HasSrcUser returns a boolean if a field has been set.
func (o *HttpEvent) HasSrcUser() bool {
	if o != nil && o.SrcUser != nil {
		return true
	}

	return false
}

// SetSrcUser gets a reference to the given string and assigns it to the SrcUser field.
func (o *HttpEvent) SetSrcUser(v string) {
	o.SrcUser = &v
}

// GetSrcID returns the SrcID field value if set, zero value otherwise.
func (o *HttpEvent) GetSrcID() string {
	if o == nil || o.SrcID == nil {
		var ret string
		return ret
	}
	return *o.SrcID
}

// GetSrcIDOk returns a tuple with the SrcID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetSrcIDOk() (*string, bool) {
	if o == nil || o.SrcID == nil {
		return nil, false
	}
	return o.SrcID, true
}

// HasSrcID returns a boolean if a field has been set.
func (o *HttpEvent) HasSrcID() bool {
	if o != nil && o.SrcID != nil {
		return true
	}

	return false
}

// SetSrcID gets a reference to the given string and assigns it to the SrcID field.
func (o *HttpEvent) SetSrcID(v string) {
	o.SrcID = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *HttpEvent) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *HttpEvent) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *HttpEvent) SetUrl(v string) {
	o.Url = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *HttpEvent) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *HttpEvent) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *HttpEvent) SetMethod(v string) {
	o.Method = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HttpEvent) GetStatus() int64 {
	if o == nil || o.Status == nil {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetStatusOk() (*int64, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HttpEvent) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *HttpEvent) SetStatus(v int64) {
	o.Status = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *HttpEvent) GetApp() string {
	if o == nil || o.App == nil {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetAppOk() (*string, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *HttpEvent) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *HttpEvent) SetApp(v string) {
	o.App = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *HttpEvent) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *HttpEvent) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *HttpEvent) SetAction(v string) {
	o.Action = &v
}

// GetBytesIn returns the BytesIn field value if set, zero value otherwise.
func (o *HttpEvent) GetBytesIn() int64 {
	if o == nil || o.BytesIn == nil {
		var ret int64
		return ret
	}
	return *o.BytesIn
}

// GetBytesInOk returns a tuple with the BytesIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetBytesInOk() (*int64, bool) {
	if o == nil || o.BytesIn == nil {
		return nil, false
	}
	return o.BytesIn, true
}

// HasBytesIn returns a boolean if a field has been set.
func (o *HttpEvent) HasBytesIn() bool {
	if o != nil && o.BytesIn != nil {
		return true
	}

	return false
}

// SetBytesIn gets a reference to the given int64 and assigns it to the BytesIn field.
func (o *HttpEvent) SetBytesIn(v int64) {
	o.BytesIn = &v
}

// GetBytesOut returns the BytesOut field value if set, zero value otherwise.
func (o *HttpEvent) GetBytesOut() int64 {
	if o == nil || o.BytesOut == nil {
		var ret int64
		return ret
	}
	return *o.BytesOut
}

// GetBytesOutOk returns a tuple with the BytesOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetBytesOutOk() (*int64, bool) {
	if o == nil || o.BytesOut == nil {
		return nil, false
	}
	return o.BytesOut, true
}

// HasBytesOut returns a boolean if a field has been set.
func (o *HttpEvent) HasBytesOut() bool {
	if o != nil && o.BytesOut != nil {
		return true
	}

	return false
}

// SetBytesOut gets a reference to the given int64 and assigns it to the BytesOut field.
func (o *HttpEvent) SetBytesOut(v int64) {
	o.BytesOut = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *HttpEvent) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *HttpEvent) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *HttpEvent) SetContentType(v string) {
	o.ContentType = &v
}

// GetReferrer returns the Referrer field value if set, zero value otherwise.
func (o *HttpEvent) GetReferrer() string {
	if o == nil || o.Referrer == nil {
		var ret string
		return ret
	}
	return *o.Referrer
}

// GetReferrerOk returns a tuple with the Referrer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetReferrerOk() (*string, bool) {
	if o == nil || o.Referrer == nil {
		return nil, false
	}
	return o.Referrer, true
}

// HasReferrer returns a boolean if a field has been set.
func (o *HttpEvent) HasReferrer() bool {
	if o != nil && o.Referrer != nil {
		return true
	}

	return false
}

// SetReferrer gets a reference to the given string and assigns it to the Referrer field.
func (o *HttpEvent) SetReferrer(v string) {
	o.Referrer = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *HttpEvent) GetUserAgent() string {
	if o == nil || o.UserAgent == nil {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEvent) GetUserAgentOk() (*string, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *HttpEvent) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *HttpEvent) SetUserAgent(v string) {
	o.UserAgent = &v
}

func (o HttpEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ts != nil {
		toSerialize["ts"] = o.Ts
	}
	if o.SrcIP != nil {
		toSerialize["srcIP"] = o.SrcIP
	}
	if o.SrcPort != nil {
		toSerialize["srcPort"] = o.SrcPort
	}
	if o.SrcHost != nil {
		toSerialize["srcHost"] = o.SrcHost
	}
	if o.SrcMac != nil {
		toSerialize["srcMac"] = o.SrcMac
	}
	if o.SrcUser != nil {
		toSerialize["srcUser"] = o.SrcUser
	}
	if o.SrcID != nil {
		toSerialize["srcID"] = o.SrcID
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.BytesIn != nil {
		toSerialize["bytesIn"] = o.BytesIn
	}
	if o.BytesOut != nil {
		toSerialize["bytesOut"] = o.BytesOut
	}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	if o.Referrer != nil {
		toSerialize["referrer"] = o.Referrer
	}
	if o.UserAgent != nil {
		toSerialize["userAgent"] = o.UserAgent
	}
	return json.Marshal(toSerialize)
}

type NullableHttpEvent struct {
	value *HttpEvent
	isSet bool
}

func (v NullableHttpEvent) Get() *HttpEvent {
	return v.value
}

func (v *NullableHttpEvent) Set(val *HttpEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpEvent(val *HttpEvent) *NullableHttpEvent {
	return &NullableHttpEvent{value: val, isSet: true}
}

func (v NullableHttpEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

