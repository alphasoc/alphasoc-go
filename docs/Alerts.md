# Alerts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Follow** | Pointer to **string** | Page bookmark. Can be passed to consecutive request to retrieve only new alerts since the last query. | [optional] 
**More** | Pointer to **bool** | Indicates if there are more alerts to retrieve. | [optional] 
**Alerts** | Pointer to [**[]Alert**](Alert.md) | Array of alerts. | [optional] 
**Threats** | Pointer to [**map[string]Threat**](Threat.md) | Dictionary containing definition of threats. | [optional] 

## Methods

### NewAlerts

`func NewAlerts() *Alerts`

NewAlerts instantiates a new Alerts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsWithDefaults

`func NewAlertsWithDefaults() *Alerts`

NewAlertsWithDefaults instantiates a new Alerts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollow

`func (o *Alerts) GetFollow() string`

GetFollow returns the Follow field if non-nil, zero value otherwise.

### GetFollowOk

`func (o *Alerts) GetFollowOk() (*string, bool)`

GetFollowOk returns a tuple with the Follow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollow

`func (o *Alerts) SetFollow(v string)`

SetFollow sets Follow field to given value.

### HasFollow

`func (o *Alerts) HasFollow() bool`

HasFollow returns a boolean if a field has been set.

### GetMore

`func (o *Alerts) GetMore() bool`

GetMore returns the More field if non-nil, zero value otherwise.

### GetMoreOk

`func (o *Alerts) GetMoreOk() (*bool, bool)`

GetMoreOk returns a tuple with the More field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMore

`func (o *Alerts) SetMore(v bool)`

SetMore sets More field to given value.

### HasMore

`func (o *Alerts) HasMore() bool`

HasMore returns a boolean if a field has been set.

### GetAlerts

`func (o *Alerts) GetAlerts() []Alert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *Alerts) GetAlertsOk() (*[]Alert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *Alerts) SetAlerts(v []Alert)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *Alerts) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetThreats

`func (o *Alerts) GetThreats() map[string]Threat`

GetThreats returns the Threats field if non-nil, zero value otherwise.

### GetThreatsOk

`func (o *Alerts) GetThreatsOk() (*map[string]Threat, bool)`

GetThreatsOk returns a tuple with the Threats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreats

`func (o *Alerts) SetThreats(v map[string]Threat)`

SetThreats sets Threats field to given value.

### HasThreats

`func (o *Alerts) HasThreats() bool`

HasThreats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


