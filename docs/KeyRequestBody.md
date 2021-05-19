# KeyRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to [**Platform**](Platform.md) |  | [optional] 
**Uname** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewKeyRequestBody

`func NewKeyRequestBody() *KeyRequestBody`

NewKeyRequestBody instantiates a new KeyRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyRequestBodyWithDefaults

`func NewKeyRequestBodyWithDefaults() *KeyRequestBody`

NewKeyRequestBodyWithDefaults instantiates a new KeyRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *KeyRequestBody) GetPlatform() Platform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *KeyRequestBody) GetPlatformOk() (*Platform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *KeyRequestBody) SetPlatform(v Platform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *KeyRequestBody) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetUname

`func (o *KeyRequestBody) GetUname() string`

GetUname returns the Uname field if non-nil, zero value otherwise.

### GetUnameOk

`func (o *KeyRequestBody) GetUnameOk() (*string, bool)`

GetUnameOk returns a tuple with the Uname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUname

`func (o *KeyRequestBody) SetUname(v string)`

SetUname sets Uname field to given value.

### HasUname

`func (o *KeyRequestBody) HasUname() bool`

HasUname returns a boolean if a field has been set.

### GetToken

`func (o *KeyRequestBody) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KeyRequestBody) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KeyRequestBody) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KeyRequestBody) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


