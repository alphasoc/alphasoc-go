# Threat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Human readable description of the threat | 
**Severity** | **int32** | Severity of the threat | 
**Policy** | Pointer to **bool** |  | [optional] 

## Methods

### NewThreat

`func NewThreat(title string, severity int32, ) *Threat`

NewThreat instantiates a new Threat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatWithDefaults

`func NewThreatWithDefaults() *Threat`

NewThreatWithDefaults instantiates a new Threat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Threat) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Threat) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Threat) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSeverity

`func (o *Threat) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Threat) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Threat) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.


### GetPolicy

`func (o *Threat) GetPolicy() bool`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *Threat) GetPolicyOk() (*bool, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *Threat) SetPolicy(v bool)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *Threat) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


