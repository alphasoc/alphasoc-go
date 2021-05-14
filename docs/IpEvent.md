# IpEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** |  | [optional] 
**SrcIP** | Pointer to **string** |  | [optional] 
**SrcPort** | Pointer to **int32** |  | [optional] 
**SrcHost** | Pointer to **string** |  | [optional] 
**SrcMac** | Pointer to **string** |  | [optional] 
**SrcUser** | Pointer to **string** |  | [optional] 
**SrcID** | Pointer to **string** |  | [optional] 
**DestIP** | Pointer to **string** |  | [optional] 
**DestPort** | Pointer to **int32** |  | [optional] 
**Proto** | Pointer to **string** |  | [optional] 
**BytesIn** | Pointer to **int64** |  | [optional] 
**BytesOut** | Pointer to **int64** |  | [optional] 
**App** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **float64** |  | [optional] 

## Methods

### NewIpEvent

`func NewIpEvent() *IpEvent`

NewIpEvent instantiates a new IpEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpEventWithDefaults

`func NewIpEventWithDefaults() *IpEvent`

NewIpEventWithDefaults instantiates a new IpEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *IpEvent) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *IpEvent) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *IpEvent) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *IpEvent) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetSrcIP

`func (o *IpEvent) GetSrcIP() string`

GetSrcIP returns the SrcIP field if non-nil, zero value otherwise.

### GetSrcIPOk

`func (o *IpEvent) GetSrcIPOk() (*string, bool)`

GetSrcIPOk returns a tuple with the SrcIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIP

`func (o *IpEvent) SetSrcIP(v string)`

SetSrcIP sets SrcIP field to given value.

### HasSrcIP

`func (o *IpEvent) HasSrcIP() bool`

HasSrcIP returns a boolean if a field has been set.

### GetSrcPort

`func (o *IpEvent) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *IpEvent) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *IpEvent) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *IpEvent) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcHost

`func (o *IpEvent) GetSrcHost() string`

GetSrcHost returns the SrcHost field if non-nil, zero value otherwise.

### GetSrcHostOk

`func (o *IpEvent) GetSrcHostOk() (*string, bool)`

GetSrcHostOk returns a tuple with the SrcHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcHost

`func (o *IpEvent) SetSrcHost(v string)`

SetSrcHost sets SrcHost field to given value.

### HasSrcHost

`func (o *IpEvent) HasSrcHost() bool`

HasSrcHost returns a boolean if a field has been set.

### GetSrcMac

`func (o *IpEvent) GetSrcMac() string`

GetSrcMac returns the SrcMac field if non-nil, zero value otherwise.

### GetSrcMacOk

`func (o *IpEvent) GetSrcMacOk() (*string, bool)`

GetSrcMacOk returns a tuple with the SrcMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcMac

`func (o *IpEvent) SetSrcMac(v string)`

SetSrcMac sets SrcMac field to given value.

### HasSrcMac

`func (o *IpEvent) HasSrcMac() bool`

HasSrcMac returns a boolean if a field has been set.

### GetSrcUser

`func (o *IpEvent) GetSrcUser() string`

GetSrcUser returns the SrcUser field if non-nil, zero value otherwise.

### GetSrcUserOk

`func (o *IpEvent) GetSrcUserOk() (*string, bool)`

GetSrcUserOk returns a tuple with the SrcUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcUser

`func (o *IpEvent) SetSrcUser(v string)`

SetSrcUser sets SrcUser field to given value.

### HasSrcUser

`func (o *IpEvent) HasSrcUser() bool`

HasSrcUser returns a boolean if a field has been set.

### GetSrcID

`func (o *IpEvent) GetSrcID() string`

GetSrcID returns the SrcID field if non-nil, zero value otherwise.

### GetSrcIDOk

`func (o *IpEvent) GetSrcIDOk() (*string, bool)`

GetSrcIDOk returns a tuple with the SrcID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcID

`func (o *IpEvent) SetSrcID(v string)`

SetSrcID sets SrcID field to given value.

### HasSrcID

`func (o *IpEvent) HasSrcID() bool`

HasSrcID returns a boolean if a field has been set.

### GetDestIP

`func (o *IpEvent) GetDestIP() string`

GetDestIP returns the DestIP field if non-nil, zero value otherwise.

### GetDestIPOk

`func (o *IpEvent) GetDestIPOk() (*string, bool)`

GetDestIPOk returns a tuple with the DestIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestIP

`func (o *IpEvent) SetDestIP(v string)`

SetDestIP sets DestIP field to given value.

### HasDestIP

`func (o *IpEvent) HasDestIP() bool`

HasDestIP returns a boolean if a field has been set.

### GetDestPort

`func (o *IpEvent) GetDestPort() int32`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *IpEvent) GetDestPortOk() (*int32, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *IpEvent) SetDestPort(v int32)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *IpEvent) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetProto

`func (o *IpEvent) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *IpEvent) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *IpEvent) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *IpEvent) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetBytesIn

`func (o *IpEvent) GetBytesIn() int64`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *IpEvent) GetBytesInOk() (*int64, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *IpEvent) SetBytesIn(v int64)`

SetBytesIn sets BytesIn field to given value.

### HasBytesIn

`func (o *IpEvent) HasBytesIn() bool`

HasBytesIn returns a boolean if a field has been set.

### GetBytesOut

`func (o *IpEvent) GetBytesOut() int64`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *IpEvent) GetBytesOutOk() (*int64, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *IpEvent) SetBytesOut(v int64)`

SetBytesOut sets BytesOut field to given value.

### HasBytesOut

`func (o *IpEvent) HasBytesOut() bool`

HasBytesOut returns a boolean if a field has been set.

### GetApp

`func (o *IpEvent) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *IpEvent) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *IpEvent) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *IpEvent) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetAction

`func (o *IpEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IpEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IpEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IpEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDuration

`func (o *IpEvent) GetDuration() float64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *IpEvent) GetDurationOk() (*float64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *IpEvent) SetDuration(v float64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *IpEvent) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


