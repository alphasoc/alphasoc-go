# DnsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Event timestamp | [optional] 
**SrcIP** | Pointer to **string** | Source IP | [optional] 
**SrcPort** | Pointer to **int32** | Source port | [optional] 
**SrcHost** | Pointer to **string** | Source host | [optional] 
**SrcMac** | Pointer to **string** | Source mac address | [optional] 
**SrcUser** | Pointer to **string** | Source user | [optional] 
**SrcID** | Pointer to **string** | Source ID | [optional] 
**Query** | Pointer to **string** | DNS query | [optional] 
**Qtype** | Pointer to **string** | Query type | [optional] 

## Methods

### NewDnsEvent

`func NewDnsEvent() *DnsEvent`

NewDnsEvent instantiates a new DnsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsEventWithDefaults

`func NewDnsEventWithDefaults() *DnsEvent`

NewDnsEventWithDefaults instantiates a new DnsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *DnsEvent) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *DnsEvent) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *DnsEvent) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *DnsEvent) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetSrcIP

`func (o *DnsEvent) GetSrcIP() string`

GetSrcIP returns the SrcIP field if non-nil, zero value otherwise.

### GetSrcIPOk

`func (o *DnsEvent) GetSrcIPOk() (*string, bool)`

GetSrcIPOk returns a tuple with the SrcIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIP

`func (o *DnsEvent) SetSrcIP(v string)`

SetSrcIP sets SrcIP field to given value.

### HasSrcIP

`func (o *DnsEvent) HasSrcIP() bool`

HasSrcIP returns a boolean if a field has been set.

### GetSrcPort

`func (o *DnsEvent) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *DnsEvent) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *DnsEvent) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *DnsEvent) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcHost

`func (o *DnsEvent) GetSrcHost() string`

GetSrcHost returns the SrcHost field if non-nil, zero value otherwise.

### GetSrcHostOk

`func (o *DnsEvent) GetSrcHostOk() (*string, bool)`

GetSrcHostOk returns a tuple with the SrcHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcHost

`func (o *DnsEvent) SetSrcHost(v string)`

SetSrcHost sets SrcHost field to given value.

### HasSrcHost

`func (o *DnsEvent) HasSrcHost() bool`

HasSrcHost returns a boolean if a field has been set.

### GetSrcMac

`func (o *DnsEvent) GetSrcMac() string`

GetSrcMac returns the SrcMac field if non-nil, zero value otherwise.

### GetSrcMacOk

`func (o *DnsEvent) GetSrcMacOk() (*string, bool)`

GetSrcMacOk returns a tuple with the SrcMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcMac

`func (o *DnsEvent) SetSrcMac(v string)`

SetSrcMac sets SrcMac field to given value.

### HasSrcMac

`func (o *DnsEvent) HasSrcMac() bool`

HasSrcMac returns a boolean if a field has been set.

### GetSrcUser

`func (o *DnsEvent) GetSrcUser() string`

GetSrcUser returns the SrcUser field if non-nil, zero value otherwise.

### GetSrcUserOk

`func (o *DnsEvent) GetSrcUserOk() (*string, bool)`

GetSrcUserOk returns a tuple with the SrcUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcUser

`func (o *DnsEvent) SetSrcUser(v string)`

SetSrcUser sets SrcUser field to given value.

### HasSrcUser

`func (o *DnsEvent) HasSrcUser() bool`

HasSrcUser returns a boolean if a field has been set.

### GetSrcID

`func (o *DnsEvent) GetSrcID() string`

GetSrcID returns the SrcID field if non-nil, zero value otherwise.

### GetSrcIDOk

`func (o *DnsEvent) GetSrcIDOk() (*string, bool)`

GetSrcIDOk returns a tuple with the SrcID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcID

`func (o *DnsEvent) SetSrcID(v string)`

SetSrcID sets SrcID field to given value.

### HasSrcID

`func (o *DnsEvent) HasSrcID() bool`

HasSrcID returns a boolean if a field has been set.

### GetQuery

`func (o *DnsEvent) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DnsEvent) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DnsEvent) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *DnsEvent) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetQtype

`func (o *DnsEvent) GetQtype() string`

GetQtype returns the Qtype field if non-nil, zero value otherwise.

### GetQtypeOk

`func (o *DnsEvent) GetQtypeOk() (*string, bool)`

GetQtypeOk returns a tuple with the Qtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtype

`func (o *DnsEvent) SetQtype(v string)`

SetQtype sets Qtype field to given value.

### HasQtype

`func (o *DnsEvent) HasQtype() bool`

HasQtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


