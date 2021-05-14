# \DataProcessingApi

All URIs are relative to *https://api.alphasoc.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AlertsGet**](DataProcessingApi.md#V1AlertsGet) | **Get** /v1/alerts | Retrieve alerts
[**V1EventsDnsPost**](DataProcessingApi.md#V1EventsDnsPost) | **Post** /v1/events/dns | Submit DNS query events
[**V1EventsHttpPost**](DataProcessingApi.md#V1EventsHttpPost) | **Post** /v1/events/http | Submit HTTP events
[**V1EventsIpPost**](DataProcessingApi.md#V1EventsIpPost) | **Post** /v1/events/ip | Submit IP events
[**V1EventsLeasePost**](DataProcessingApi.md#V1EventsLeasePost) | **Post** /v1/events/lease | Submit lease (DHCP) events
[**V1EventsTlsPost**](DataProcessingApi.md#V1EventsTlsPost) | **Post** /v1/events/tls | Submit TLS events



## V1AlertsGet

> Alerts V1AlertsGet(ctx).Follow(follow).Threats(threats).Execute()

Retrieve alerts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    follow := "follow_example" // string |  (optional)
    threats := "threats_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProcessingApi.V1AlertsGet(context.Background()).Follow(follow).Threats(threats).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingApi.V1AlertsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AlertsGet`: Alerts
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingApi.V1AlertsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AlertsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **follow** | **string** |  | 
 **threats** | **string** |  | 

### Return type

[**Alerts**](Alerts.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EventsDnsPost

> EventsResponseBody V1EventsDnsPost(ctx).DnsEvent(dnsEvent).Execute()

Submit DNS query events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dnsEvent := *openapiclient.NewDnsEvent() // DnsEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProcessingApi.V1EventsDnsPost(context.Background()).DnsEvent(dnsEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingApi.V1EventsDnsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EventsDnsPost`: EventsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingApi.V1EventsDnsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EventsDnsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsEvent** | [**DnsEvent**](DnsEvent.md) |  | 

### Return type

[**EventsResponseBody**](EventsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EventsHttpPost

> EventsResponseBody V1EventsHttpPost(ctx).HttpEvent(httpEvent).Execute()

Submit HTTP events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    httpEvent := *openapiclient.NewHttpEvent() // HttpEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProcessingApi.V1EventsHttpPost(context.Background()).HttpEvent(httpEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingApi.V1EventsHttpPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EventsHttpPost`: EventsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingApi.V1EventsHttpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EventsHttpPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **httpEvent** | [**HttpEvent**](HttpEvent.md) |  | 

### Return type

[**EventsResponseBody**](EventsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EventsIpPost

> EventsResponseBody V1EventsIpPost(ctx).IpEvent(ipEvent).Execute()

Submit IP events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ipEvent := *openapiclient.NewIpEvent() // IpEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProcessingApi.V1EventsIpPost(context.Background()).IpEvent(ipEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingApi.V1EventsIpPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EventsIpPost`: EventsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingApi.V1EventsIpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EventsIpPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipEvent** | [**IpEvent**](IpEvent.md) |  | 

### Return type

[**EventsResponseBody**](EventsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EventsLeasePost

> EventsResponseBody V1EventsLeasePost(ctx).LeaseEvent(leaseEvent).Execute()

Submit lease (DHCP) events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    leaseEvent := *openapiclient.NewLeaseEvent() // LeaseEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProcessingApi.V1EventsLeasePost(context.Background()).LeaseEvent(leaseEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingApi.V1EventsLeasePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EventsLeasePost`: EventsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingApi.V1EventsLeasePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EventsLeasePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **leaseEvent** | [**LeaseEvent**](LeaseEvent.md) |  | 

### Return type

[**EventsResponseBody**](EventsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EventsTlsPost

> EventsResponseBody V1EventsTlsPost(ctx).TlsEvent(tlsEvent).Execute()

Submit TLS events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tlsEvent := *openapiclient.NewTlsEvent() // TlsEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProcessingApi.V1EventsTlsPost(context.Background()).TlsEvent(tlsEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingApi.V1EventsTlsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EventsTlsPost`: EventsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingApi.V1EventsTlsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EventsTlsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsEvent** | [**TlsEvent**](TlsEvent.md) |  | 

### Return type

[**EventsResponseBody**](EventsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

