# IpEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestIP** | Pointer to **string** | Destination IP | [optional] 
**DestPort** | Pointer to **int32** | Destination port | [optional] 
**Proto** | Pointer to **string** | Transport layer protocol | [optional] 
**BytesIn** | Pointer to **int64** | Number of incoming bytes | [optional] 
**BytesOut** | Pointer to **int64** | Number of outgoing bytes | [optional] 
**App** | Pointer to **string** | Application layer protocol | [optional] 
**Action** | Pointer to **string** | Defines if event was allowed or denied | [optional] 
**Duration** | Pointer to **float64** | Duration of connection | [optional] 

## Methods

### NewIpEventAllOf

`func NewIpEventAllOf() *IpEventAllOf`

NewIpEventAllOf instantiates a new IpEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpEventAllOfWithDefaults

`func NewIpEventAllOfWithDefaults() *IpEventAllOf`

NewIpEventAllOfWithDefaults instantiates a new IpEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestIP

`func (o *IpEventAllOf) GetDestIP() string`

GetDestIP returns the DestIP field if non-nil, zero value otherwise.

### GetDestIPOk

`func (o *IpEventAllOf) GetDestIPOk() (*string, bool)`

GetDestIPOk returns a tuple with the DestIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestIP

`func (o *IpEventAllOf) SetDestIP(v string)`

SetDestIP sets DestIP field to given value.

### HasDestIP

`func (o *IpEventAllOf) HasDestIP() bool`

HasDestIP returns a boolean if a field has been set.

### GetDestPort

`func (o *IpEventAllOf) GetDestPort() int32`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *IpEventAllOf) GetDestPortOk() (*int32, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *IpEventAllOf) SetDestPort(v int32)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *IpEventAllOf) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetProto

`func (o *IpEventAllOf) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *IpEventAllOf) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *IpEventAllOf) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *IpEventAllOf) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetBytesIn

`func (o *IpEventAllOf) GetBytesIn() int64`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *IpEventAllOf) GetBytesInOk() (*int64, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *IpEventAllOf) SetBytesIn(v int64)`

SetBytesIn sets BytesIn field to given value.

### HasBytesIn

`func (o *IpEventAllOf) HasBytesIn() bool`

HasBytesIn returns a boolean if a field has been set.

### GetBytesOut

`func (o *IpEventAllOf) GetBytesOut() int64`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *IpEventAllOf) GetBytesOutOk() (*int64, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *IpEventAllOf) SetBytesOut(v int64)`

SetBytesOut sets BytesOut field to given value.

### HasBytesOut

`func (o *IpEventAllOf) HasBytesOut() bool`

HasBytesOut returns a boolean if a field has been set.

### GetApp

`func (o *IpEventAllOf) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *IpEventAllOf) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *IpEventAllOf) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *IpEventAllOf) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetAction

`func (o *IpEventAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IpEventAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IpEventAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IpEventAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDuration

`func (o *IpEventAllOf) GetDuration() float64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *IpEventAllOf) GetDurationOk() (*float64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *IpEventAllOf) SetDuration(v float64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *IpEventAllOf) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


