# \InventoryApi

All URIs are relative to *https://api.alphasoc.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AeInventoryFlagsGet**](InventoryApi.md#V1AeInventoryFlagsGet) | **Get** /v1/ae/inventory/flags | Analytics Engine flags
[**V1AeInventoryThreatsGet**](InventoryApi.md#V1AeInventoryThreatsGet) | **Get** /v1/ae/inventory/threats | Analytics Engine threats



## V1AeInventoryFlagsGet

> AeFlags V1AeInventoryFlagsGet(ctx).Execute()

Analytics Engine flags

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.V1AeInventoryFlagsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.V1AeInventoryFlagsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AeInventoryFlagsGet`: AeFlags
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.V1AeInventoryFlagsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AeInventoryFlagsGetRequest struct via the builder pattern


### Return type

[**AeFlags**](AeFlags.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AeInventoryThreatsGet

> AeThreats V1AeInventoryThreatsGet(ctx).Execute()

Analytics Engine threats

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.V1AeInventoryThreatsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.V1AeInventoryThreatsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AeInventoryThreatsGet`: AeThreats
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.V1AeInventoryThreatsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AeInventoryThreatsGetRequest struct via the builder pattern


### Return type

[**AeThreats**](AeThreats.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

