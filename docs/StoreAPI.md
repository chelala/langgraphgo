# \StoreAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteItem**](StoreAPI.md#DeleteItem) | **Delete** /store/items | Delete an item.
[**GetItem**](StoreAPI.md#GetItem) | **Get** /store/items | Retrieve a single item.
[**ListNamespaces**](StoreAPI.md#ListNamespaces) | **Post** /store/namespaces | List namespaces with optional match conditions.
[**PutItem**](StoreAPI.md#PutItem) | **Put** /store/items | Store or update an item.
[**SearchItems**](StoreAPI.md#SearchItems) | **Post** /store/items/search | Search for items within a namespace prefix.



## DeleteItem

> DeleteItem(ctx).StoreDeleteRequest(storeDeleteRequest).Execute()

Delete an item.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/chelala/langgraphgo"
)

func main() {
	storeDeleteRequest := *openapiclient.NewStoreDeleteRequest("Key_example") // StoreDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.DeleteItem(context.Background()).StoreDeleteRequest(storeDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.DeleteItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeDeleteRequest** | [**StoreDeleteRequest**](StoreDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItem

> Item GetItem(ctx).Key(key).Namespace(namespace).Execute()

Retrieve a single item.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/chelala/langgraphgo"
)

func main() {
	key := "key_example" // string | 
	namespace := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.GetItem(context.Background()).Key(key).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItem`: Item
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.GetItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 
 **namespace** | **[]string** |  | 

### Return type

[**Item**](Item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespaces

> [][]string ListNamespaces(ctx).StoreListNamespacesRequest(storeListNamespacesRequest).Execute()

List namespaces with optional match conditions.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/chelala/langgraphgo"
)

func main() {
	storeListNamespacesRequest := *openapiclient.NewStoreListNamespacesRequest() // StoreListNamespacesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.ListNamespaces(context.Background()).StoreListNamespacesRequest(storeListNamespacesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.ListNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespaces`: [][]string
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.ListNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeListNamespacesRequest** | [**StoreListNamespacesRequest**](StoreListNamespacesRequest.md) |  | 

### Return type

[**[][]string**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutItem

> PutItem(ctx).StorePutRequest(storePutRequest).Execute()

Store or update an item.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/chelala/langgraphgo"
)

func main() {
	storePutRequest := *openapiclient.NewStorePutRequest([]string{"Namespace_example"}, "Key_example", map[string]interface{}(123)) // StorePutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PutItem(context.Background()).StorePutRequest(storePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PutItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storePutRequest** | [**StorePutRequest**](StorePutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchItems

> SearchItemsResponse SearchItems(ctx).StoreSearchRequest(storeSearchRequest).Execute()

Search for items within a namespace prefix.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/chelala/langgraphgo"
)

func main() {
	storeSearchRequest := *openapiclient.NewStoreSearchRequest() // StoreSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.SearchItems(context.Background()).StoreSearchRequest(storeSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.SearchItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchItems`: SearchItemsResponse
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.SearchItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeSearchRequest** | [**StoreSearchRequest**](StoreSearchRequest.md) |  | 

### Return type

[**SearchItemsResponse**](SearchItemsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

