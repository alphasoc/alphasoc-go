# AccountStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Today** | Pointer to **time.Time** |  | [optional] 
**Registered** | Pointer to **bool** |  | [optional] 
**Expired** | Pointer to **bool** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**EndpointsSeenToday** | Pointer to **int32** |  | [optional] 
**Messages** | Pointer to [**[]Message**](Message.md) |  | [optional] 

## Methods

### NewAccountStatus

`func NewAccountStatus() *AccountStatus`

NewAccountStatus instantiates a new AccountStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountStatusWithDefaults

`func NewAccountStatusWithDefaults() *AccountStatus`

NewAccountStatusWithDefaults instantiates a new AccountStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToday

`func (o *AccountStatus) GetToday() time.Time`

GetToday returns the Today field if non-nil, zero value otherwise.

### GetTodayOk

`func (o *AccountStatus) GetTodayOk() (*time.Time, bool)`

GetTodayOk returns a tuple with the Today field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToday

`func (o *AccountStatus) SetToday(v time.Time)`

SetToday sets Today field to given value.

### HasToday

`func (o *AccountStatus) HasToday() bool`

HasToday returns a boolean if a field has been set.

### GetRegistered

`func (o *AccountStatus) GetRegistered() bool`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *AccountStatus) GetRegisteredOk() (*bool, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *AccountStatus) SetRegistered(v bool)`

SetRegistered sets Registered field to given value.

### HasRegistered

`func (o *AccountStatus) HasRegistered() bool`

HasRegistered returns a boolean if a field has been set.

### GetExpired

`func (o *AccountStatus) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *AccountStatus) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *AccountStatus) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *AccountStatus) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetExpirationDate

`func (o *AccountStatus) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *AccountStatus) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *AccountStatus) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *AccountStatus) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetEndpointsSeenToday

`func (o *AccountStatus) GetEndpointsSeenToday() int32`

GetEndpointsSeenToday returns the EndpointsSeenToday field if non-nil, zero value otherwise.

### GetEndpointsSeenTodayOk

`func (o *AccountStatus) GetEndpointsSeenTodayOk() (*int32, bool)`

GetEndpointsSeenTodayOk returns a tuple with the EndpointsSeenToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointsSeenToday

`func (o *AccountStatus) SetEndpointsSeenToday(v int32)`

SetEndpointsSeenToday sets EndpointsSeenToday field to given value.

### HasEndpointsSeenToday

`func (o *AccountStatus) HasEndpointsSeenToday() bool`

HasEndpointsSeenToday returns a boolean if a field has been set.

### GetMessages

`func (o *AccountStatus) GetMessages() []Message`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AccountStatus) GetMessagesOk() (*[]Message, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AccountStatus) SetMessages(v []Message)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *AccountStatus) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


