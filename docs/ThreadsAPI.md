# \ThreadsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyThreadPostThreadsThreadIdCopyPost**](ThreadsAPI.md#CopyThreadPostThreadsThreadIdCopyPost) | **Post** /threads/{thread_id}/copy | Copy Thread
[**CreateThreadThreadsPost**](ThreadsAPI.md#CreateThreadThreadsPost) | **Post** /threads | Create Thread
[**DeleteThreadThreadsThreadIdDelete**](ThreadsAPI.md#DeleteThreadThreadsThreadIdDelete) | **Delete** /threads/{thread_id} | Delete Thread
[**GetLatestThreadStateThreadsThreadIdStateGet**](ThreadsAPI.md#GetLatestThreadStateThreadsThreadIdStateGet) | **Get** /threads/{thread_id}/state | Get Thread State
[**GetThreadHistoryPostThreadsThreadIdHistoryPost**](ThreadsAPI.md#GetThreadHistoryPostThreadsThreadIdHistoryPost) | **Post** /threads/{thread_id}/history | Get Thread History Post
[**GetThreadHistoryThreadsThreadIdHistoryGet**](ThreadsAPI.md#GetThreadHistoryThreadsThreadIdHistoryGet) | **Get** /threads/{thread_id}/history | Get Thread History
[**GetThreadThreadsThreadIdGet**](ThreadsAPI.md#GetThreadThreadsThreadIdGet) | **Get** /threads/{thread_id} | Get Thread
[**PatchThreadThreadsThreadIdPatch**](ThreadsAPI.md#PatchThreadThreadsThreadIdPatch) | **Patch** /threads/{thread_id} | Patch Thread
[**PostThreadStateAtCheckpointThreadsThreadIdStateCheckpointIdGet**](ThreadsAPI.md#PostThreadStateAtCheckpointThreadsThreadIdStateCheckpointIdGet) | **Post** /threads/{thread_id}/state/checkpoint | Get Thread State At Checkpoint
[**SearchThreadsThreadsSearchPost**](ThreadsAPI.md#SearchThreadsThreadsSearchPost) | **Post** /threads/search | Search Threads
[**UpdateThreadStateThreadsThreadIdStatePost**](ThreadsAPI.md#UpdateThreadStateThreadsThreadIdStatePost) | **Post** /threads/{thread_id}/state | Update Thread State



## CopyThreadPostThreadsThreadIdCopyPost

> Thread CopyThreadPostThreadsThreadIdCopyPost(ctx, threadId).Execute()

Copy Thread



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
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the thread.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.CopyThreadPostThreadsThreadIdCopyPost(context.Background(), threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.CopyThreadPostThreadsThreadIdCopyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyThreadPostThreadsThreadIdCopyPost`: Thread
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.CopyThreadPostThreadsThreadIdCopyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyThreadPostThreadsThreadIdCopyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateThreadThreadsPost

> Thread CreateThreadThreadsPost(ctx).ThreadCreate(threadCreate).Execute()

Create Thread



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
	threadCreate := *openapiclient.NewThreadCreate() // ThreadCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.CreateThreadThreadsPost(context.Background()).ThreadCreate(threadCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.CreateThreadThreadsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateThreadThreadsPost`: Thread
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.CreateThreadThreadsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateThreadThreadsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadCreate** | [**ThreadCreate**](ThreadCreate.md) |  | 

### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteThreadThreadsThreadIdDelete

> interface{} DeleteThreadThreadsThreadIdDelete(ctx, threadId).Execute()

Delete Thread



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
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the thread.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.DeleteThreadThreadsThreadIdDelete(context.Background(), threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.DeleteThreadThreadsThreadIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteThreadThreadsThreadIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.DeleteThreadThreadsThreadIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteThreadThreadsThreadIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestThreadStateThreadsThreadIdStateGet

> ThreadState GetLatestThreadStateThreadsThreadIdStateGet(ctx, threadId).Execute()

Get Thread State



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
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the thread.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.GetLatestThreadStateThreadsThreadIdStateGet(context.Background(), threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.GetLatestThreadStateThreadsThreadIdStateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestThreadStateThreadsThreadIdStateGet`: ThreadState
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.GetLatestThreadStateThreadsThreadIdStateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestThreadStateThreadsThreadIdStateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThreadState**](ThreadState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreadHistoryPostThreadsThreadIdHistoryPost

> []ThreadState GetThreadHistoryPostThreadsThreadIdHistoryPost(ctx, threadId).ThreadStateSearch(threadStateSearch).Execute()

Get Thread History Post



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
	threadStateSearch := *openapiclient.NewThreadStateSearch() // ThreadStateSearch | 
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the thread.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.GetThreadHistoryPostThreadsThreadIdHistoryPost(context.Background(), threadId).ThreadStateSearch(threadStateSearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.GetThreadHistoryPostThreadsThreadIdHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreadHistoryPostThreadsThreadIdHistoryPost`: []ThreadState
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.GetThreadHistoryPostThreadsThreadIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreadHistoryPostThreadsThreadIdHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadStateSearch** | [**ThreadStateSearch**](ThreadStateSearch.md) |  | 


### Return type

[**[]ThreadState**](ThreadState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreadHistoryThreadsThreadIdHistoryGet

> []ThreadState GetThreadHistoryThreadsThreadIdHistoryGet(ctx, threadId).Limit(limit).Before(before).Execute()

Get Thread History



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
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the thread.
	limit := int32(56) // int32 |  (optional) (default to 10)
	before := "before_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.GetThreadHistoryThreadsThreadIdHistoryGet(context.Background(), threadId).Limit(limit).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.GetThreadHistoryThreadsThreadIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreadHistoryThreadsThreadIdHistoryGet`: []ThreadState
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.GetThreadHistoryThreadsThreadIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreadHistoryThreadsThreadIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 10]
 **before** | **string** |  | 

### Return type

[**[]ThreadState**](ThreadState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreadThreadsThreadIdGet

> Thread GetThreadThreadsThreadIdGet(ctx, threadId).Execute()

Get Thread



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
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the thread.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.GetThreadThreadsThreadIdGet(context.Background(), threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.GetThreadThreadsThreadIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreadThreadsThreadIdGet`: Thread
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.GetThreadThreadsThreadIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreadThreadsThreadIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchThreadThreadsThreadIdPatch

> Thread PatchThreadThreadsThreadIdPatch(ctx, threadId).ThreadPatch(threadPatch).Execute()

Patch Thread



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
	threadPatch := *openapiclient.NewThreadPatch() // ThreadPatch | 
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the thread.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.PatchThreadThreadsThreadIdPatch(context.Background(), threadId).ThreadPatch(threadPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.PatchThreadThreadsThreadIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchThreadThreadsThreadIdPatch`: Thread
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.PatchThreadThreadsThreadIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchThreadThreadsThreadIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadPatch** | [**ThreadPatch**](ThreadPatch.md) |  | 


### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostThreadStateAtCheckpointThreadsThreadIdStateCheckpointIdGet

> ThreadState PostThreadStateAtCheckpointThreadsThreadIdStateCheckpointIdGet(ctx).ThreadStateCheckpointRequest(threadStateCheckpointRequest).Execute()

Get Thread State At Checkpoint



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
	threadStateCheckpointRequest := *openapiclient.NewThreadStateCheckpointRequest(*openapiclient.NewCheckpointConfig()) // ThreadStateCheckpointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.PostThreadStateAtCheckpointThreadsThreadIdStateCheckpointIdGet(context.Background()).ThreadStateCheckpointRequest(threadStateCheckpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.PostThreadStateAtCheckpointThreadsThreadIdStateCheckpointIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostThreadStateAtCheckpointThreadsThreadIdStateCheckpointIdGet`: ThreadState
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.PostThreadStateAtCheckpointThreadsThreadIdStateCheckpointIdGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostThreadStateAtCheckpointThreadsThreadIdStateCheckpointIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadStateCheckpointRequest** | [**ThreadStateCheckpointRequest**](ThreadStateCheckpointRequest.md) |  | 

### Return type

[**ThreadState**](ThreadState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchThreadsThreadsSearchPost

> []Thread SearchThreadsThreadsSearchPost(ctx).ThreadSearchRequest(threadSearchRequest).Execute()

Search Threads



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
	threadSearchRequest := *openapiclient.NewThreadSearchRequest() // ThreadSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.SearchThreadsThreadsSearchPost(context.Background()).ThreadSearchRequest(threadSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.SearchThreadsThreadsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchThreadsThreadsSearchPost`: []Thread
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.SearchThreadsThreadsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchThreadsThreadsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadSearchRequest** | [**ThreadSearchRequest**](ThreadSearchRequest.md) |  | 

### Return type

[**[]Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateThreadStateThreadsThreadIdStatePost

> ThreadStateUpdateResponse UpdateThreadStateThreadsThreadIdStatePost(ctx, threadId).ThreadStateUpdate(threadStateUpdate).Execute()

Update Thread State



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
	threadStateUpdate := *openapiclient.NewThreadStateUpdate() // ThreadStateUpdate | 
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the thread.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.UpdateThreadStateThreadsThreadIdStatePost(context.Background(), threadId).ThreadStateUpdate(threadStateUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.UpdateThreadStateThreadsThreadIdStatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateThreadStateThreadsThreadIdStatePost`: ThreadStateUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.UpdateThreadStateThreadsThreadIdStatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateThreadStateThreadsThreadIdStatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadStateUpdate** | [**ThreadStateUpdate**](ThreadStateUpdate.md) |  | 


### Return type

[**ThreadStateUpdateResponse**](ThreadStateUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

