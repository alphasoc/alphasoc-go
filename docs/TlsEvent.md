# TlsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Event timestamp | [optional] 
**SrcIP** | Pointer to **string** | Source IP | [optional] 
**SrcPort** | Pointer to **int32** | Source port | [optional] 
**SrcHost** | Pointer to **string** | Source host | [optional] 
**SrcMac** | Pointer to **string** | Source mac address | [optional] 
**SrcUser** | Pointer to **string** | Source user | [optional] 
**SrcID** | Pointer to **string** | Source ID | [optional] 
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

### NewTlsEvent

`func NewTlsEvent() *TlsEvent`

NewTlsEvent instantiates a new TlsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsEventWithDefaults

`func NewTlsEventWithDefaults() *TlsEvent`

NewTlsEventWithDefaults instantiates a new TlsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *TlsEvent) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *TlsEvent) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *TlsEvent) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *TlsEvent) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetSrcIP

`func (o *TlsEvent) GetSrcIP() string`

GetSrcIP returns the SrcIP field if non-nil, zero value otherwise.

### GetSrcIPOk

`func (o *TlsEvent) GetSrcIPOk() (*string, bool)`

GetSrcIPOk returns a tuple with the SrcIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIP

`func (o *TlsEvent) SetSrcIP(v string)`

SetSrcIP sets SrcIP field to given value.

### HasSrcIP

`func (o *TlsEvent) HasSrcIP() bool`

HasSrcIP returns a boolean if a field has been set.

### GetSrcPort

`func (o *TlsEvent) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *TlsEvent) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *TlsEvent) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *TlsEvent) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcHost

`func (o *TlsEvent) GetSrcHost() string`

GetSrcHost returns the SrcHost field if non-nil, zero value otherwise.

### GetSrcHostOk

`func (o *TlsEvent) GetSrcHostOk() (*string, bool)`

GetSrcHostOk returns a tuple with the SrcHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcHost

`func (o *TlsEvent) SetSrcHost(v string)`

SetSrcHost sets SrcHost field to given value.

### HasSrcHost

`func (o *TlsEvent) HasSrcHost() bool`

HasSrcHost returns a boolean if a field has been set.

### GetSrcMac

`func (o *TlsEvent) GetSrcMac() string`

GetSrcMac returns the SrcMac field if non-nil, zero value otherwise.

### GetSrcMacOk

`func (o *TlsEvent) GetSrcMacOk() (*string, bool)`

GetSrcMacOk returns a tuple with the SrcMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcMac

`func (o *TlsEvent) SetSrcMac(v string)`

SetSrcMac sets SrcMac field to given value.

### HasSrcMac

`func (o *TlsEvent) HasSrcMac() bool`

HasSrcMac returns a boolean if a field has been set.

### GetSrcUser

`func (o *TlsEvent) GetSrcUser() string`

GetSrcUser returns the SrcUser field if non-nil, zero value otherwise.

### GetSrcUserOk

`func (o *TlsEvent) GetSrcUserOk() (*string, bool)`

GetSrcUserOk returns a tuple with the SrcUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcUser

`func (o *TlsEvent) SetSrcUser(v string)`

SetSrcUser sets SrcUser field to given value.

### HasSrcUser

`func (o *TlsEvent) HasSrcUser() bool`

HasSrcUser returns a boolean if a field has been set.

### GetSrcID

`func (o *TlsEvent) GetSrcID() string`

GetSrcID returns the SrcID field if non-nil, zero value otherwise.

### GetSrcIDOk

`func (o *TlsEvent) GetSrcIDOk() (*string, bool)`

GetSrcIDOk returns a tuple with the SrcID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcID

`func (o *TlsEvent) SetSrcID(v string)`

SetSrcID sets SrcID field to given value.

### HasSrcID

`func (o *TlsEvent) HasSrcID() bool`

HasSrcID returns a boolean if a field has been set.

### GetCertHash

`func (o *TlsEvent) GetCertHash() string`

GetCertHash returns the CertHash field if non-nil, zero value otherwise.

### GetCertHashOk

`func (o *TlsEvent) GetCertHashOk() (*string, bool)`

GetCertHashOk returns a tuple with the CertHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertHash

`func (o *TlsEvent) SetCertHash(v string)`

SetCertHash sets CertHash field to given value.

### HasCertHash

`func (o *TlsEvent) HasCertHash() bool`

HasCertHash returns a boolean if a field has been set.

### GetIssuer

`func (o *TlsEvent) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *TlsEvent) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *TlsEvent) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *TlsEvent) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *TlsEvent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TlsEvent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TlsEvent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TlsEvent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetValidFrom

`func (o *TlsEvent) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *TlsEvent) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *TlsEvent) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *TlsEvent) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *TlsEvent) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *TlsEvent) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *TlsEvent) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *TlsEvent) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### GetDestIP

`func (o *TlsEvent) GetDestIP() string`

GetDestIP returns the DestIP field if non-nil, zero value otherwise.

### GetDestIPOk

`func (o *TlsEvent) GetDestIPOk() (*string, bool)`

GetDestIPOk returns a tuple with the DestIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestIP

`func (o *TlsEvent) SetDestIP(v string)`

SetDestIP sets DestIP field to given value.

### HasDestIP

`func (o *TlsEvent) HasDestIP() bool`

HasDestIP returns a boolean if a field has been set.

### GetDestPort

`func (o *TlsEvent) GetDestPort() int32`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *TlsEvent) GetDestPortOk() (*int32, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *TlsEvent) SetDestPort(v int32)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *TlsEvent) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetJa3

`func (o *TlsEvent) GetJa3() string`

GetJa3 returns the Ja3 field if non-nil, zero value otherwise.

### GetJa3Ok

`func (o *TlsEvent) GetJa3Ok() (*string, bool)`

GetJa3Ok returns a tuple with the Ja3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJa3

`func (o *TlsEvent) SetJa3(v string)`

SetJa3 sets Ja3 field to given value.

### HasJa3

`func (o *TlsEvent) HasJa3() bool`

HasJa3 returns a boolean if a field has been set.

### GetJa3s

`func (o *TlsEvent) GetJa3s() string`

GetJa3s returns the Ja3s field if non-nil, zero value otherwise.

### GetJa3sOk

`func (o *TlsEvent) GetJa3sOk() (*string, bool)`

GetJa3sOk returns a tuple with the Ja3s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJa3s

`func (o *TlsEvent) SetJa3s(v string)`

SetJa3s sets Ja3s field to given value.

### HasJa3s

`func (o *TlsEvent) HasJa3s() bool`

HasJa3s returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


