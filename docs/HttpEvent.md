# HttpEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** |  | [optional] 
**SrcIP** | Pointer to **string** |  | [optional] 
**SrcPort** | Pointer to **int32** |  | [optional] 
**SrcHost** | Pointer to **string** |  | [optional] 
**SrcMac** | Pointer to **string** |  | [optional] 
**SrcUser** | Pointer to **string** |  | [optional] 
**SrcID** | Pointer to **string** |  | [optional] 
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

### NewHttpEvent

`func NewHttpEvent() *HttpEvent`

NewHttpEvent instantiates a new HttpEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpEventWithDefaults

`func NewHttpEventWithDefaults() *HttpEvent`

NewHttpEventWithDefaults instantiates a new HttpEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *HttpEvent) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *HttpEvent) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *HttpEvent) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *HttpEvent) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetSrcIP

`func (o *HttpEvent) GetSrcIP() string`

GetSrcIP returns the SrcIP field if non-nil, zero value otherwise.

### GetSrcIPOk

`func (o *HttpEvent) GetSrcIPOk() (*string, bool)`

GetSrcIPOk returns a tuple with the SrcIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIP

`func (o *HttpEvent) SetSrcIP(v string)`

SetSrcIP sets SrcIP field to given value.

### HasSrcIP

`func (o *HttpEvent) HasSrcIP() bool`

HasSrcIP returns a boolean if a field has been set.

### GetSrcPort

`func (o *HttpEvent) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *HttpEvent) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *HttpEvent) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *HttpEvent) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcHost

`func (o *HttpEvent) GetSrcHost() string`

GetSrcHost returns the SrcHost field if non-nil, zero value otherwise.

### GetSrcHostOk

`func (o *HttpEvent) GetSrcHostOk() (*string, bool)`

GetSrcHostOk returns a tuple with the SrcHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcHost

`func (o *HttpEvent) SetSrcHost(v string)`

SetSrcHost sets SrcHost field to given value.

### HasSrcHost

`func (o *HttpEvent) HasSrcHost() bool`

HasSrcHost returns a boolean if a field has been set.

### GetSrcMac

`func (o *HttpEvent) GetSrcMac() string`

GetSrcMac returns the SrcMac field if non-nil, zero value otherwise.

### GetSrcMacOk

`func (o *HttpEvent) GetSrcMacOk() (*string, bool)`

GetSrcMacOk returns a tuple with the SrcMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcMac

`func (o *HttpEvent) SetSrcMac(v string)`

SetSrcMac sets SrcMac field to given value.

### HasSrcMac

`func (o *HttpEvent) HasSrcMac() bool`

HasSrcMac returns a boolean if a field has been set.

### GetSrcUser

`func (o *HttpEvent) GetSrcUser() string`

GetSrcUser returns the SrcUser field if non-nil, zero value otherwise.

### GetSrcUserOk

`func (o *HttpEvent) GetSrcUserOk() (*string, bool)`

GetSrcUserOk returns a tuple with the SrcUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcUser

`func (o *HttpEvent) SetSrcUser(v string)`

SetSrcUser sets SrcUser field to given value.

### HasSrcUser

`func (o *HttpEvent) HasSrcUser() bool`

HasSrcUser returns a boolean if a field has been set.

### GetSrcID

`func (o *HttpEvent) GetSrcID() string`

GetSrcID returns the SrcID field if non-nil, zero value otherwise.

### GetSrcIDOk

`func (o *HttpEvent) GetSrcIDOk() (*string, bool)`

GetSrcIDOk returns a tuple with the SrcID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcID

`func (o *HttpEvent) SetSrcID(v string)`

SetSrcID sets SrcID field to given value.

### HasSrcID

`func (o *HttpEvent) HasSrcID() bool`

HasSrcID returns a boolean if a field has been set.

### GetUrl

`func (o *HttpEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HttpEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HttpEvent) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HttpEvent) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMethod

`func (o *HttpEvent) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HttpEvent) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HttpEvent) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HttpEvent) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetStatus

`func (o *HttpEvent) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HttpEvent) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HttpEvent) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HttpEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApp

`func (o *HttpEvent) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *HttpEvent) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *HttpEvent) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *HttpEvent) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetAction

`func (o *HttpEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HttpEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HttpEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *HttpEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBytesIn

`func (o *HttpEvent) GetBytesIn() int64`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *HttpEvent) GetBytesInOk() (*int64, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *HttpEvent) SetBytesIn(v int64)`

SetBytesIn sets BytesIn field to given value.

### HasBytesIn

`func (o *HttpEvent) HasBytesIn() bool`

HasBytesIn returns a boolean if a field has been set.

### GetBytesOut

`func (o *HttpEvent) GetBytesOut() int64`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *HttpEvent) GetBytesOutOk() (*int64, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *HttpEvent) SetBytesOut(v int64)`

SetBytesOut sets BytesOut field to given value.

### HasBytesOut

`func (o *HttpEvent) HasBytesOut() bool`

HasBytesOut returns a boolean if a field has been set.

### GetContentType

`func (o *HttpEvent) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *HttpEvent) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *HttpEvent) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *HttpEvent) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetReferrer

`func (o *HttpEvent) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *HttpEvent) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *HttpEvent) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *HttpEvent) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetUserAgent

`func (o *HttpEvent) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *HttpEvent) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *HttpEvent) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *HttpEvent) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


