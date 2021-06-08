# EventsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Received** | Pointer to **int32** | Number of received events | [optional] 
**Accepted** | Pointer to **int32** | Number of accepted events | [optional] 

## Methods

### NewEventsResponseBody

`func NewEventsResponseBody() *EventsResponseBody`

NewEventsResponseBody instantiates a new EventsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsResponseBodyWithDefaults

`func NewEventsResponseBodyWithDefaults() *EventsResponseBody`

NewEventsResponseBodyWithDefaults instantiates a new EventsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceived

`func (o *EventsResponseBody) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *EventsResponseBody) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *EventsResponseBody) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *EventsResponseBody) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetAccepted

`func (o *EventsResponseBody) GetAccepted() int32`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *EventsResponseBody) GetAcceptedOk() (*int32, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *EventsResponseBody) SetAccepted(v int32)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *EventsResponseBody) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


