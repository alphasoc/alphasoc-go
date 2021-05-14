# \AccountManagementApi

All URIs are relative to *https://api.alphasoc.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountRegisterPost**](AccountManagementApi.md#V1AccountRegisterPost) | **Post** /v1/account/register | Register new account
[**V1AccountStatusGet**](AccountManagementApi.md#V1AccountStatusGet) | **Get** /v1/account/status | Account status
[**V1AccountVerifyGet**](AccountManagementApi.md#V1AccountVerifyGet) | **Get** /v1/account/verify | Account verification



## V1AccountRegisterPost

> map[string]interface{} V1AccountRegisterPost(ctx).Register(register).Execute()

Register new account



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
    register := *openapiclient.NewRegister() // Register |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountManagementApi.V1AccountRegisterPost(context.Background()).Register(register).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.V1AccountRegisterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AccountRegisterPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.V1AccountRegisterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **register** | [**Register**](Register.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountStatusGet

> AccountStatus V1AccountStatusGet(ctx).Execute()

Account status



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
    resp, r, err := api_client.AccountManagementApi.V1AccountStatusGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.V1AccountStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AccountStatusGet`: AccountStatus
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.V1AccountStatusGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountStatusGetRequest struct via the builder pattern


### Return type

[**AccountStatus**](AccountStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountVerifyGet

> V1AccountVerifyGet(ctx, token).Execute()

Account verification



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
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountManagementApi.V1AccountVerifyGet(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.V1AccountVerifyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountVerifyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

