# \ShoppingCartAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPersonToShoppingCart**](ShoppingCartAPI.md#AddPersonToShoppingCart) | **Post** /shopping-carts/{shoppingCartId}/persons/{personId} | Add a Person to a ShoppingCart
[**CreateShoppingCart**](ShoppingCartAPI.md#CreateShoppingCart) | **Post** /shopping-carts | Create a ShoppingCart
[**DeleteShoppingCart**](ShoppingCartAPI.md#DeleteShoppingCart) | **Delete** /shopping-carts/{shoppingCartId} | Delete a ShoppingCart
[**GetShoppingCart**](ShoppingCartAPI.md#GetShoppingCart) | **Get** /shopping-carts/{shoppingCartId} | Get details about a ShoppingCart
[**RemovePersonFromShoppingCart**](ShoppingCartAPI.md#RemovePersonFromShoppingCart) | **Delete** /shopping-carts/{shoppingCartId}/persons/{personId} | Remove a Person from a ShoppingCart
[**SearchShoppingCarts**](ShoppingCartAPI.md#SearchShoppingCarts) | **Post** /shopping-carts/search | Complex search over shopping cart entities
[**ShareShoppingCart**](ShoppingCartAPI.md#ShareShoppingCart) | **Post** /shopping-carts/{shoppingCartId}/shared-with/{personId} | Share a ShoppingCart with a Person
[**UnshareShoppingCart**](ShoppingCartAPI.md#UnshareShoppingCart) | **Delete** /shopping-carts/{shoppingCartId}/shared-with/{personId} | Unshare a ShoppingCart with a Person
[**UpdateShoppingCart**](ShoppingCartAPI.md#UpdateShoppingCart) | **Put** /shopping-carts/{shoppingCartId} | Update a ShoppingCart



## AddPersonToShoppingCart

> ShoppingCartDetail AddPersonToShoppingCart(ctx, shoppingCartId, personId).Execute()

Add a Person to a ShoppingCart

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    shoppingCartId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShoppingCartAPI.AddPersonToShoppingCart(context.Background(), shoppingCartId, personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingCartAPI.AddPersonToShoppingCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonToShoppingCart`: ShoppingCartDetail
    fmt.Fprintf(os.Stdout, "Response from `ShoppingCartAPI.AddPersonToShoppingCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shoppingCartId** | **string** |  | 
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonToShoppingCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ShoppingCartDetail**](ShoppingCartDetail.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShoppingCart

> ShoppingCartDetail CreateShoppingCart(ctx).ShoppingCart(shoppingCart).Execute()

Create a ShoppingCart

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    shoppingCart := *openapiclient.NewShoppingCart("Id_example", "Name_example") // ShoppingCart | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShoppingCartAPI.CreateShoppingCart(context.Background()).ShoppingCart(shoppingCart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingCartAPI.CreateShoppingCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateShoppingCart`: ShoppingCartDetail
    fmt.Fprintf(os.Stdout, "Response from `ShoppingCartAPI.CreateShoppingCart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShoppingCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shoppingCart** | [**ShoppingCart**](ShoppingCart.md) |  | 

### Return type

[**ShoppingCartDetail**](ShoppingCartDetail.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShoppingCart

> Status DeleteShoppingCart(ctx, shoppingCartId).Execute()

Delete a ShoppingCart

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    shoppingCartId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShoppingCartAPI.DeleteShoppingCart(context.Background(), shoppingCartId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingCartAPI.DeleteShoppingCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteShoppingCart`: Status
    fmt.Fprintf(os.Stdout, "Response from `ShoppingCartAPI.DeleteShoppingCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shoppingCartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteShoppingCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShoppingCart

> ShoppingCartDetail GetShoppingCart(ctx, shoppingCartId).Execute()

Get details about a ShoppingCart

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    shoppingCartId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShoppingCartAPI.GetShoppingCart(context.Background(), shoppingCartId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingCartAPI.GetShoppingCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShoppingCart`: ShoppingCartDetail
    fmt.Fprintf(os.Stdout, "Response from `ShoppingCartAPI.GetShoppingCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shoppingCartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShoppingCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShoppingCartDetail**](ShoppingCartDetail.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePersonFromShoppingCart

> ShoppingCartDetail RemovePersonFromShoppingCart(ctx, shoppingCartId, personId).Execute()

Remove a Person from a ShoppingCart

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    shoppingCartId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShoppingCartAPI.RemovePersonFromShoppingCart(context.Background(), shoppingCartId, personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingCartAPI.RemovePersonFromShoppingCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemovePersonFromShoppingCart`: ShoppingCartDetail
    fmt.Fprintf(os.Stdout, "Response from `ShoppingCartAPI.RemovePersonFromShoppingCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shoppingCartId** | **string** |  | 
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePersonFromShoppingCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ShoppingCartDetail**](ShoppingCartDetail.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchShoppingCarts

> PagedShoppingCarts SearchShoppingCarts(ctx).ShoppingCartSearch(shoppingCartSearch).Execute()

Complex search over shopping cart entities

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    shoppingCartSearch := *openapiclient.NewShoppingCartSearch(int32(123), int32(123)) // ShoppingCartSearch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShoppingCartAPI.SearchShoppingCarts(context.Background()).ShoppingCartSearch(shoppingCartSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingCartAPI.SearchShoppingCarts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchShoppingCarts`: PagedShoppingCarts
    fmt.Fprintf(os.Stdout, "Response from `ShoppingCartAPI.SearchShoppingCarts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchShoppingCartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shoppingCartSearch** | [**ShoppingCartSearch**](ShoppingCartSearch.md) |  | 

### Return type

[**PagedShoppingCarts**](PagedShoppingCarts.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareShoppingCart

> ShoppingCartDetail ShareShoppingCart(ctx, shoppingCartId, personId).Execute()

Share a ShoppingCart with a Person

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    shoppingCartId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShoppingCartAPI.ShareShoppingCart(context.Background(), shoppingCartId, personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingCartAPI.ShareShoppingCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShareShoppingCart`: ShoppingCartDetail
    fmt.Fprintf(os.Stdout, "Response from `ShoppingCartAPI.ShareShoppingCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shoppingCartId** | **string** |  | 
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareShoppingCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ShoppingCartDetail**](ShoppingCartDetail.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnshareShoppingCart

> ShoppingCartDetail UnshareShoppingCart(ctx, shoppingCartId, personId).Execute()

Unshare a ShoppingCart with a Person

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    shoppingCartId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShoppingCartAPI.UnshareShoppingCart(context.Background(), shoppingCartId, personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingCartAPI.UnshareShoppingCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnshareShoppingCart`: ShoppingCartDetail
    fmt.Fprintf(os.Stdout, "Response from `ShoppingCartAPI.UnshareShoppingCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shoppingCartId** | **string** |  | 
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnshareShoppingCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ShoppingCartDetail**](ShoppingCartDetail.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShoppingCart

> ShoppingCartDetail UpdateShoppingCart(ctx, shoppingCartId).ShoppingCart(shoppingCart).Execute()

Update a ShoppingCart

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    shoppingCartId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    shoppingCart := *openapiclient.NewShoppingCart("Id_example", "Name_example") // ShoppingCart | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShoppingCartAPI.UpdateShoppingCart(context.Background(), shoppingCartId).ShoppingCart(shoppingCart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingCartAPI.UpdateShoppingCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateShoppingCart`: ShoppingCartDetail
    fmt.Fprintf(os.Stdout, "Response from `ShoppingCartAPI.UpdateShoppingCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shoppingCartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShoppingCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shoppingCart** | [**ShoppingCart**](ShoppingCart.md) |  | 

### Return type

[**ShoppingCartDetail**](ShoppingCartDetail.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

