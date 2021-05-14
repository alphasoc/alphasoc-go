# EventHeader

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

## Methods

### NewEventHeader

`func NewEventHeader() *EventHeader`

NewEventHeader instantiates a new EventHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHeaderWithDefaults

`func NewEventHeaderWithDefaults() *EventHeader`

NewEventHeaderWithDefaults instantiates a new EventHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *EventHeader) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *EventHeader) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *EventHeader) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *EventHeader) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetSrcIP

`func (o *EventHeader) GetSrcIP() string`

GetSrcIP returns the SrcIP field if non-nil, zero value otherwise.

### GetSrcIPOk

`func (o *EventHeader) GetSrcIPOk() (*string, bool)`

GetSrcIPOk returns a tuple with the SrcIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIP

`func (o *EventHeader) SetSrcIP(v string)`

SetSrcIP sets SrcIP field to given value.

### HasSrcIP

`func (o *EventHeader) HasSrcIP() bool`

HasSrcIP returns a boolean if a field has been set.

### GetSrcPort

`func (o *EventHeader) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *EventHeader) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *EventHeader) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *EventHeader) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcHost

`func (o *EventHeader) GetSrcHost() string`

GetSrcHost returns the SrcHost field if non-nil, zero value otherwise.

### GetSrcHostOk

`func (o *EventHeader) GetSrcHostOk() (*string, bool)`

GetSrcHostOk returns a tuple with the SrcHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcHost

`func (o *EventHeader) SetSrcHost(v string)`

SetSrcHost sets SrcHost field to given value.

### HasSrcHost

`func (o *EventHeader) HasSrcHost() bool`

HasSrcHost returns a boolean if a field has been set.

### GetSrcMac

`func (o *EventHeader) GetSrcMac() string`

GetSrcMac returns the SrcMac field if non-nil, zero value otherwise.

### GetSrcMacOk

`func (o *EventHeader) GetSrcMacOk() (*string, bool)`

GetSrcMacOk returns a tuple with the SrcMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcMac

`func (o *EventHeader) SetSrcMac(v string)`

SetSrcMac sets SrcMac field to given value.

### HasSrcMac

`func (o *EventHeader) HasSrcMac() bool`

HasSrcMac returns a boolean if a field has been set.

### GetSrcUser

`func (o *EventHeader) GetSrcUser() string`

GetSrcUser returns the SrcUser field if non-nil, zero value otherwise.

### GetSrcUserOk

`func (o *EventHeader) GetSrcUserOk() (*string, bool)`

GetSrcUserOk returns a tuple with the SrcUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcUser

`func (o *EventHeader) SetSrcUser(v string)`

SetSrcUser sets SrcUser field to given value.

### HasSrcUser

`func (o *EventHeader) HasSrcUser() bool`

HasSrcUser returns a boolean if a field has been set.

### GetSrcID

`func (o *EventHeader) GetSrcID() string`

GetSrcID returns the SrcID field if non-nil, zero value otherwise.

### GetSrcIDOk

`func (o *EventHeader) GetSrcIDOk() (*string, bool)`

GetSrcIDOk returns a tuple with the SrcID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcID

`func (o *EventHeader) SetSrcID(v string)`

SetSrcID sets SrcID field to given value.

### HasSrcID

`func (o *EventHeader) HasSrcID() bool`

HasSrcID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


