# \KeyManagementApi

All URIs are relative to *https://api.alphasoc.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1KeyRequestPost**](KeyManagementApi.md#V1KeyRequestPost) | **Post** /v1/key/request | Key request
[**V1KeyResetPost**](KeyManagementApi.md#V1KeyResetPost) | **Post** /v1/key/reset | Key reset



## V1KeyRequestPost

> KeyRequestResponseBody V1KeyRequestPost(ctx).KeyRequestBody(keyRequestBody).Execute()

Key request



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
    keyRequestBody := *openapiclient.NewKeyRequestBody() // KeyRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyManagementApi.V1KeyRequestPost(context.Background()).KeyRequestBody(keyRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.V1KeyRequestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1KeyRequestPost`: KeyRequestResponseBody
    fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.V1KeyRequestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1KeyRequestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyRequestBody** | [**KeyRequestBody**](KeyRequestBody.md) |  | 

### Return type

[**KeyRequestResponseBody**](KeyRequestResponseBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1KeyResetPost

> map[string]interface{} V1KeyResetPost(ctx).KeyResetRequestBody(keyResetRequestBody).Execute()

Key reset



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
    keyResetRequestBody := *openapiclient.NewKeyResetRequestBody() // KeyResetRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyManagementApi.V1KeyResetPost(context.Background()).KeyResetRequestBody(keyResetRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.V1KeyResetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1KeyResetPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.V1KeyResetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1KeyResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyResetRequestBody** | [**KeyResetRequestBody**](KeyResetRequestBody.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

