# LeaseEvent

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
**Type** | Pointer to **string** |  | [optional] 
**Termination** | Pointer to **bool** |  | [optional] 
**Duration** | Pointer to **int64** | Duration of the event | [optional] 

## Methods

### NewLeaseEvent

`func NewLeaseEvent() *LeaseEvent`

NewLeaseEvent instantiates a new LeaseEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaseEventWithDefaults

`func NewLeaseEventWithDefaults() *LeaseEvent`

NewLeaseEventWithDefaults instantiates a new LeaseEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *LeaseEvent) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *LeaseEvent) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *LeaseEvent) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *LeaseEvent) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetSrcIP

`func (o *LeaseEvent) GetSrcIP() string`

GetSrcIP returns the SrcIP field if non-nil, zero value otherwise.

### GetSrcIPOk

`func (o *LeaseEvent) GetSrcIPOk() (*string, bool)`

GetSrcIPOk returns a tuple with the SrcIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIP

`func (o *LeaseEvent) SetSrcIP(v string)`

SetSrcIP sets SrcIP field to given value.

### HasSrcIP

`func (o *LeaseEvent) HasSrcIP() bool`

HasSrcIP returns a boolean if a field has been set.

### GetSrcPort

`func (o *LeaseEvent) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *LeaseEvent) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *LeaseEvent) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *LeaseEvent) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcHost

`func (o *LeaseEvent) GetSrcHost() string`

GetSrcHost returns the SrcHost field if non-nil, zero value otherwise.

### GetSrcHostOk

`func (o *LeaseEvent) GetSrcHostOk() (*string, bool)`

GetSrcHostOk returns a tuple with the SrcHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcHost

`func (o *LeaseEvent) SetSrcHost(v string)`

SetSrcHost sets SrcHost field to given value.

### HasSrcHost

`func (o *LeaseEvent) HasSrcHost() bool`

HasSrcHost returns a boolean if a field has been set.

### GetSrcMac

`func (o *LeaseEvent) GetSrcMac() string`

GetSrcMac returns the SrcMac field if non-nil, zero value otherwise.

### GetSrcMacOk

`func (o *LeaseEvent) GetSrcMacOk() (*string, bool)`

GetSrcMacOk returns a tuple with the SrcMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcMac

`func (o *LeaseEvent) SetSrcMac(v string)`

SetSrcMac sets SrcMac field to given value.

### HasSrcMac

`func (o *LeaseEvent) HasSrcMac() bool`

HasSrcMac returns a boolean if a field has been set.

### GetSrcUser

`func (o *LeaseEvent) GetSrcUser() string`

GetSrcUser returns the SrcUser field if non-nil, zero value otherwise.

### GetSrcUserOk

`func (o *LeaseEvent) GetSrcUserOk() (*string, bool)`

GetSrcUserOk returns a tuple with the SrcUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcUser

`func (o *LeaseEvent) SetSrcUser(v string)`

SetSrcUser sets SrcUser field to given value.

### HasSrcUser

`func (o *LeaseEvent) HasSrcUser() bool`

HasSrcUser returns a boolean if a field has been set.

### GetSrcID

`func (o *LeaseEvent) GetSrcID() string`

GetSrcID returns the SrcID field if non-nil, zero value otherwise.

### GetSrcIDOk

`func (o *LeaseEvent) GetSrcIDOk() (*string, bool)`

GetSrcIDOk returns a tuple with the SrcID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcID

`func (o *LeaseEvent) SetSrcID(v string)`

SetSrcID sets SrcID field to given value.

### HasSrcID

`func (o *LeaseEvent) HasSrcID() bool`

HasSrcID returns a boolean if a field has been set.

### GetType

`func (o *LeaseEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LeaseEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LeaseEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LeaseEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTermination

`func (o *LeaseEvent) GetTermination() bool`

GetTermination returns the Termination field if non-nil, zero value otherwise.

### GetTerminationOk

`func (o *LeaseEvent) GetTerminationOk() (*bool, bool)`

GetTerminationOk returns a tuple with the Termination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermination

`func (o *LeaseEvent) SetTermination(v bool)`

SetTermination sets Termination field to given value.

### HasTermination

`func (o *LeaseEvent) HasTermination() bool`

HasTermination returns a boolean if a field has been set.

### GetDuration

`func (o *LeaseEvent) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *LeaseEvent) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *LeaseEvent) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *LeaseEvent) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


