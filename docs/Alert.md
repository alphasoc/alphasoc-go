# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] 
**Event** | Pointer to [**OneOfdnsEventipEventhttpEventtlsEvent**](oneOf&lt;dnsEvent,ipEvent,httpEvent,tlsEvent&gt;.md) |  | [optional] 

## Methods

### NewAlert

`func NewAlert() *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *Alert) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Alert) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Alert) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *Alert) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEvent

`func (o *Alert) GetEvent() OneOfdnsEventipEventhttpEventtlsEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Alert) GetEventOk() (*OneOfdnsEventipEventhttpEventtlsEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Alert) SetEvent(v OneOfdnsEventipEventhttpEventtlsEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *Alert) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


