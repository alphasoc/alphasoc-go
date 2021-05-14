# HttpEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**App** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**BytesIn** | Pointer to **int64** |  | [optional] 
**BytesOut** | Pointer to **int64** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Referrer** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 

## Methods

### NewHttpEventAllOf

`func NewHttpEventAllOf() *HttpEventAllOf`

NewHttpEventAllOf instantiates a new HttpEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpEventAllOfWithDefaults

`func NewHttpEventAllOfWithDefaults() *HttpEventAllOf`

NewHttpEventAllOfWithDefaults instantiates a new HttpEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *HttpEventAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HttpEventAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HttpEventAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HttpEventAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMethod

`func (o *HttpEventAllOf) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HttpEventAllOf) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HttpEventAllOf) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HttpEventAllOf) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetStatus

`func (o *HttpEventAllOf) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HttpEventAllOf) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HttpEventAllOf) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HttpEventAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApp

`func (o *HttpEventAllOf) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *HttpEventAllOf) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *HttpEventAllOf) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *HttpEventAllOf) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetAction

`func (o *HttpEventAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HttpEventAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HttpEventAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *HttpEventAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBytesIn

`func (o *HttpEventAllOf) GetBytesIn() int64`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *HttpEventAllOf) GetBytesInOk() (*int64, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *HttpEventAllOf) SetBytesIn(v int64)`

SetBytesIn sets BytesIn field to given value.

### HasBytesIn

`func (o *HttpEventAllOf) HasBytesIn() bool`

HasBytesIn returns a boolean if a field has been set.

### GetBytesOut

`func (o *HttpEventAllOf) GetBytesOut() int64`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *HttpEventAllOf) GetBytesOutOk() (*int64, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *HttpEventAllOf) SetBytesOut(v int64)`

SetBytesOut sets BytesOut field to given value.

### HasBytesOut

`func (o *HttpEventAllOf) HasBytesOut() bool`

HasBytesOut returns a boolean if a field has been set.

### GetContentType

`func (o *HttpEventAllOf) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *HttpEventAllOf) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *HttpEventAllOf) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *HttpEventAllOf) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetReferrer

`func (o *HttpEventAllOf) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *HttpEventAllOf) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *HttpEventAllOf) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *HttpEventAllOf) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetUserAgent

`func (o *HttpEventAllOf) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *HttpEventAllOf) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *HttpEventAllOf) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *HttpEventAllOf) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


