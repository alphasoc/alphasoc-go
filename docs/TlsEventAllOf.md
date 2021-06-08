# TlsEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertHash** | Pointer to **string** | Certificate hash | [optional] 
**Issuer** | Pointer to **string** | Certificate issuer | [optional] 
**Subject** | Pointer to **string** | Certificate subject | [optional] 
**ValidFrom** | Pointer to **time.Time** | From when certificate is valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Certificate expiration date | [optional] 
**DestIP** | Pointer to **string** | Destination IP | [optional] 
**DestPort** | Pointer to **int32** | Destination port | [optional] 
**Ja3** | Pointer to **string** | JA3 fingerprint | [optional] 
**Ja3s** | Pointer to **string** | JA3S fingerprint | [optional] 

## Methods

### NewTlsEventAllOf

`func NewTlsEventAllOf() *TlsEventAllOf`

NewTlsEventAllOf instantiates a new TlsEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsEventAllOfWithDefaults

`func NewTlsEventAllOfWithDefaults() *TlsEventAllOf`

NewTlsEventAllOfWithDefaults instantiates a new TlsEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertHash

`func (o *TlsEventAllOf) GetCertHash() string`

GetCertHash returns the CertHash field if non-nil, zero value otherwise.

### GetCertHashOk

`func (o *TlsEventAllOf) GetCertHashOk() (*string, bool)`

GetCertHashOk returns a tuple with the CertHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertHash

`func (o *TlsEventAllOf) SetCertHash(v string)`

SetCertHash sets CertHash field to given value.

### HasCertHash

`func (o *TlsEventAllOf) HasCertHash() bool`

HasCertHash returns a boolean if a field has been set.

### GetIssuer

`func (o *TlsEventAllOf) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *TlsEventAllOf) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *TlsEventAllOf) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *TlsEventAllOf) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *TlsEventAllOf) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TlsEventAllOf) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TlsEventAllOf) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TlsEventAllOf) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetValidFrom

`func (o *TlsEventAllOf) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *TlsEventAllOf) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *TlsEventAllOf) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *TlsEventAllOf) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *TlsEventAllOf) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *TlsEventAllOf) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *TlsEventAllOf) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *TlsEventAllOf) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### GetDestIP

`func (o *TlsEventAllOf) GetDestIP() string`

GetDestIP returns the DestIP field if non-nil, zero value otherwise.

### GetDestIPOk

`func (o *TlsEventAllOf) GetDestIPOk() (*string, bool)`

GetDestIPOk returns a tuple with the DestIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestIP

`func (o *TlsEventAllOf) SetDestIP(v string)`

SetDestIP sets DestIP field to given value.

### HasDestIP

`func (o *TlsEventAllOf) HasDestIP() bool`

HasDestIP returns a boolean if a field has been set.

### GetDestPort

`func (o *TlsEventAllOf) GetDestPort() int32`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *TlsEventAllOf) GetDestPortOk() (*int32, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *TlsEventAllOf) SetDestPort(v int32)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *TlsEventAllOf) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetJa3

`func (o *TlsEventAllOf) GetJa3() string`

GetJa3 returns the Ja3 field if non-nil, zero value otherwise.

### GetJa3Ok

`func (o *TlsEventAllOf) GetJa3Ok() (*string, bool)`

GetJa3Ok returns a tuple with the Ja3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJa3

`func (o *TlsEventAllOf) SetJa3(v string)`

SetJa3 sets Ja3 field to given value.

### HasJa3

`func (o *TlsEventAllOf) HasJa3() bool`

HasJa3 returns a boolean if a field has been set.

### GetJa3s

`func (o *TlsEventAllOf) GetJa3s() string`

GetJa3s returns the Ja3s field if non-nil, zero value otherwise.

### GetJa3sOk

`func (o *TlsEventAllOf) GetJa3sOk() (*string, bool)`

GetJa3sOk returns a tuple with the Ja3s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJa3s

`func (o *TlsEventAllOf) SetJa3s(v string)`

SetJa3s sets Ja3s field to given value.

### HasJa3s

`func (o *TlsEventAllOf) HasJa3s() bool`

HasJa3s returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


