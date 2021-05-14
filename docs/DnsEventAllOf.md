# DnsEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** |  | [optional] 
**Qtype** | Pointer to **string** |  | [optional] 

## Methods

### NewDnsEventAllOf

`func NewDnsEventAllOf() *DnsEventAllOf`

NewDnsEventAllOf instantiates a new DnsEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsEventAllOfWithDefaults

`func NewDnsEventAllOfWithDefaults() *DnsEventAllOf`

NewDnsEventAllOfWithDefaults instantiates a new DnsEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *DnsEventAllOf) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DnsEventAllOf) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DnsEventAllOf) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *DnsEventAllOf) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetQtype

`func (o *DnsEventAllOf) GetQtype() string`

GetQtype returns the Qtype field if non-nil, zero value otherwise.

### GetQtypeOk

`func (o *DnsEventAllOf) GetQtypeOk() (*string, bool)`

GetQtypeOk returns a tuple with the Qtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtype

`func (o *DnsEventAllOf) SetQtype(v string)`

SetQtype sets Qtype field to given value.

### HasQtype

`func (o *DnsEventAllOf) HasQtype() bool`

HasQtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


