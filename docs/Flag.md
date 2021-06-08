# Flag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Flag description | [optional] 
**Type** | Pointer to **string** | Flag type | [optional] 

## Methods

### NewFlag

`func NewFlag() *Flag`

NewFlag instantiates a new Flag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagWithDefaults

`func NewFlagWithDefaults() *Flag`

NewFlagWithDefaults instantiates a new Flag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Flag) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Flag) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Flag) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Flag) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *Flag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Flag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Flag) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Flag) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


