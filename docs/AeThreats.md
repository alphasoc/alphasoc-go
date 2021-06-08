# AeThreats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threats** | Pointer to [**map[string]Threat**](Threat.md) | Dictionary containing definition of threats. | [optional] 

## Methods

### NewAeThreats

`func NewAeThreats() *AeThreats`

NewAeThreats instantiates a new AeThreats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeThreatsWithDefaults

`func NewAeThreatsWithDefaults() *AeThreats`

NewAeThreatsWithDefaults instantiates a new AeThreats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreats

`func (o *AeThreats) GetThreats() map[string]Threat`

GetThreats returns the Threats field if non-nil, zero value otherwise.

### GetThreatsOk

`func (o *AeThreats) GetThreatsOk() (*map[string]Threat, bool)`

GetThreatsOk returns a tuple with the Threats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreats

`func (o *AeThreats) SetThreats(v map[string]Threat)`

SetThreats sets Threats field to given value.

### HasThreats

`func (o *AeThreats) HasThreats() bool`

HasThreats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


