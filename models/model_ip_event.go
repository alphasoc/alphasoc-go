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
	"time"
)

// IpEvent IP traffic event
type IpEvent struct {
	// Event timestamp
	Ts *time.Time `json:"ts,omitempty"`
	// Source IP
	SrcIP *string `json:"srcIP,omitempty"`
	// Source port
	SrcPort *int32 `json:"srcPort,omitempty"`
	// Source host
	SrcHost *string `json:"srcHost,omitempty"`
	// Source mac address
	SrcMac *string `json:"srcMac,omitempty"`
	// Source user
	SrcUser *string `json:"srcUser,omitempty"`
	// Source ID
	SrcID *string `json:"srcID,omitempty"`
	// Destination IP
	DestIP *string `json:"destIP,omitempty"`
	// Destination port
	DestPort *int32 `json:"destPort,omitempty"`
	// Transport layer protocol
	Proto *string `json:"proto,omitempty"`
	// Number of incoming bytes
	BytesIn *int64 `json:"bytesIn,omitempty"`
	// Number of outgoing bytes
	BytesOut *int64 `json:"bytesOut,omitempty"`
	// Application layer protocol
	App *string `json:"app,omitempty"`
	// Defines if event was allowed or denied
	Action *string `json:"action,omitempty"`
	// Duration of connection
	Duration *float64 `json:"duration,omitempty"`
}

// NewIpEvent instantiates a new IpEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpEvent() *IpEvent {
	this := IpEvent{}
	return &this
}

// NewIpEventWithDefaults instantiates a new IpEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpEventWithDefaults() *IpEvent {
	this := IpEvent{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *IpEvent) GetTs() time.Time {
	if o == nil || o.Ts == nil {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetTsOk() (*time.Time, bool) {
	if o == nil || o.Ts == nil {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *IpEvent) HasTs() bool {
	if o != nil && o.Ts != nil {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *IpEvent) SetTs(v time.Time) {
	o.Ts = &v
}

// GetSrcIP returns the SrcIP field value if set, zero value otherwise.
func (o *IpEvent) GetSrcIP() string {
	if o == nil || o.SrcIP == nil {
		var ret string
		return ret
	}
	return *o.SrcIP
}

// GetSrcIPOk returns a tuple with the SrcIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetSrcIPOk() (*string, bool) {
	if o == nil || o.SrcIP == nil {
		return nil, false
	}
	return o.SrcIP, true
}

// HasSrcIP returns a boolean if a field has been set.
func (o *IpEvent) HasSrcIP() bool {
	if o != nil && o.SrcIP != nil {
		return true
	}

	return false
}

// SetSrcIP gets a reference to the given string and assigns it to the SrcIP field.
func (o *IpEvent) SetSrcIP(v string) {
	o.SrcIP = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *IpEvent) GetSrcPort() int32 {
	if o == nil || o.SrcPort == nil {
		var ret int32
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetSrcPortOk() (*int32, bool) {
	if o == nil || o.SrcPort == nil {
		return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *IpEvent) HasSrcPort() bool {
	if o != nil && o.SrcPort != nil {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given int32 and assigns it to the SrcPort field.
func (o *IpEvent) SetSrcPort(v int32) {
	o.SrcPort = &v
}

// GetSrcHost returns the SrcHost field value if set, zero value otherwise.
func (o *IpEvent) GetSrcHost() string {
	if o == nil || o.SrcHost == nil {
		var ret string
		return ret
	}
	return *o.SrcHost
}

// GetSrcHostOk returns a tuple with the SrcHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetSrcHostOk() (*string, bool) {
	if o == nil || o.SrcHost == nil {
		return nil, false
	}
	return o.SrcHost, true
}

// HasSrcHost returns a boolean if a field has been set.
func (o *IpEvent) HasSrcHost() bool {
	if o != nil && o.SrcHost != nil {
		return true
	}

	return false
}

// SetSrcHost gets a reference to the given string and assigns it to the SrcHost field.
func (o *IpEvent) SetSrcHost(v string) {
	o.SrcHost = &v
}

// GetSrcMac returns the SrcMac field value if set, zero value otherwise.
func (o *IpEvent) GetSrcMac() string {
	if o == nil || o.SrcMac == nil {
		var ret string
		return ret
	}
	return *o.SrcMac
}

// GetSrcMacOk returns a tuple with the SrcMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetSrcMacOk() (*string, bool) {
	if o == nil || o.SrcMac == nil {
		return nil, false
	}
	return o.SrcMac, true
}

// HasSrcMac returns a boolean if a field has been set.
func (o *IpEvent) HasSrcMac() bool {
	if o != nil && o.SrcMac != nil {
		return true
	}

	return false
}

// SetSrcMac gets a reference to the given string and assigns it to the SrcMac field.
func (o *IpEvent) SetSrcMac(v string) {
	o.SrcMac = &v
}

// GetSrcUser returns the SrcUser field value if set, zero value otherwise.
func (o *IpEvent) GetSrcUser() string {
	if o == nil || o.SrcUser == nil {
		var ret string
		return ret
	}
	return *o.SrcUser
}

// GetSrcUserOk returns a tuple with the SrcUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetSrcUserOk() (*string, bool) {
	if o == nil || o.SrcUser == nil {
		return nil, false
	}
	return o.SrcUser, true
}

// HasSrcUser returns a boolean if a field has been set.
func (o *IpEvent) HasSrcUser() bool {
	if o != nil && o.SrcUser != nil {
		return true
	}

	return false
}

// SetSrcUser gets a reference to the given string and assigns it to the SrcUser field.
func (o *IpEvent) SetSrcUser(v string) {
	o.SrcUser = &v
}

// GetSrcID returns the SrcID field value if set, zero value otherwise.
func (o *IpEvent) GetSrcID() string {
	if o == nil || o.SrcID == nil {
		var ret string
		return ret
	}
	return *o.SrcID
}

// GetSrcIDOk returns a tuple with the SrcID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetSrcIDOk() (*string, bool) {
	if o == nil || o.SrcID == nil {
		return nil, false
	}
	return o.SrcID, true
}

// HasSrcID returns a boolean if a field has been set.
func (o *IpEvent) HasSrcID() bool {
	if o != nil && o.SrcID != nil {
		return true
	}

	return false
}

// SetSrcID gets a reference to the given string and assigns it to the SrcID field.
func (o *IpEvent) SetSrcID(v string) {
	o.SrcID = &v
}

// GetDestIP returns the DestIP field value if set, zero value otherwise.
func (o *IpEvent) GetDestIP() string {
	if o == nil || o.DestIP == nil {
		var ret string
		return ret
	}
	return *o.DestIP
}

// GetDestIPOk returns a tuple with the DestIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetDestIPOk() (*string, bool) {
	if o == nil || o.DestIP == nil {
		return nil, false
	}
	return o.DestIP, true
}

// HasDestIP returns a boolean if a field has been set.
func (o *IpEvent) HasDestIP() bool {
	if o != nil && o.DestIP != nil {
		return true
	}

	return false
}

// SetDestIP gets a reference to the given string and assigns it to the DestIP field.
func (o *IpEvent) SetDestIP(v string) {
	o.DestIP = &v
}

// GetDestPort returns the DestPort field value if set, zero value otherwise.
func (o *IpEvent) GetDestPort() int32 {
	if o == nil || o.DestPort == nil {
		var ret int32
		return ret
	}
	return *o.DestPort
}

// GetDestPortOk returns a tuple with the DestPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetDestPortOk() (*int32, bool) {
	if o == nil || o.DestPort == nil {
		return nil, false
	}
	return o.DestPort, true
}

// HasDestPort returns a boolean if a field has been set.
func (o *IpEvent) HasDestPort() bool {
	if o != nil && o.DestPort != nil {
		return true
	}

	return false
}

// SetDestPort gets a reference to the given int32 and assigns it to the DestPort field.
func (o *IpEvent) SetDestPort(v int32) {
	o.DestPort = &v
}

// GetProto returns the Proto field value if set, zero value otherwise.
func (o *IpEvent) GetProto() string {
	if o == nil || o.Proto == nil {
		var ret string
		return ret
	}
	return *o.Proto
}

// GetProtoOk returns a tuple with the Proto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetProtoOk() (*string, bool) {
	if o == nil || o.Proto == nil {
		return nil, false
	}
	return o.Proto, true
}

// HasProto returns a boolean if a field has been set.
func (o *IpEvent) HasProto() bool {
	if o != nil && o.Proto != nil {
		return true
	}

	return false
}

// SetProto gets a reference to the given string and assigns it to the Proto field.
func (o *IpEvent) SetProto(v string) {
	o.Proto = &v
}

// GetBytesIn returns the BytesIn field value if set, zero value otherwise.
func (o *IpEvent) GetBytesIn() int64 {
	if o == nil || o.BytesIn == nil {
		var ret int64
		return ret
	}
	return *o.BytesIn
}

// GetBytesInOk returns a tuple with the BytesIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetBytesInOk() (*int64, bool) {
	if o == nil || o.BytesIn == nil {
		return nil, false
	}
	return o.BytesIn, true
}

// HasBytesIn returns a boolean if a field has been set.
func (o *IpEvent) HasBytesIn() bool {
	if o != nil && o.BytesIn != nil {
		return true
	}

	return false
}

// SetBytesIn gets a reference to the given int64 and assigns it to the BytesIn field.
func (o *IpEvent) SetBytesIn(v int64) {
	o.BytesIn = &v
}

// GetBytesOut returns the BytesOut field value if set, zero value otherwise.
func (o *IpEvent) GetBytesOut() int64 {
	if o == nil || o.BytesOut == nil {
		var ret int64
		return ret
	}
	return *o.BytesOut
}

// GetBytesOutOk returns a tuple with the BytesOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetBytesOutOk() (*int64, bool) {
	if o == nil || o.BytesOut == nil {
		return nil, false
	}
	return o.BytesOut, true
}

// HasBytesOut returns a boolean if a field has been set.
func (o *IpEvent) HasBytesOut() bool {
	if o != nil && o.BytesOut != nil {
		return true
	}

	return false
}

// SetBytesOut gets a reference to the given int64 and assigns it to the BytesOut field.
func (o *IpEvent) SetBytesOut(v int64) {
	o.BytesOut = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *IpEvent) GetApp() string {
	if o == nil || o.App == nil {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetAppOk() (*string, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *IpEvent) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *IpEvent) SetApp(v string) {
	o.App = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *IpEvent) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *IpEvent) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *IpEvent) SetAction(v string) {
	o.Action = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *IpEvent) GetDuration() float64 {
	if o == nil || o.Duration == nil {
		var ret float64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEvent) GetDurationOk() (*float64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *IpEvent) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float64 and assigns it to the Duration field.
func (o *IpEvent) SetDuration(v float64) {
	o.Duration = &v
}

func (o IpEvent) MarshalJSON() ([]byte, error) {
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
	if o.DestIP != nil {
		toSerialize["destIP"] = o.DestIP
	}
	if o.DestPort != nil {
		toSerialize["destPort"] = o.DestPort
	}
	if o.Proto != nil {
		toSerialize["proto"] = o.Proto
	}
	if o.BytesIn != nil {
		toSerialize["bytesIn"] = o.BytesIn
	}
	if o.BytesOut != nil {
		toSerialize["bytesOut"] = o.BytesOut
	}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	return json.Marshal(toSerialize)
}

type NullableIpEvent struct {
	value *IpEvent
	isSet bool
}

func (v NullableIpEvent) Get() *IpEvent {
	return v.value
}

func (v *NullableIpEvent) Set(val *IpEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableIpEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableIpEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpEvent(val *IpEvent) *NullableIpEvent {
	return &NullableIpEvent{value: val, isSet: true}
}

func (v NullableIpEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
