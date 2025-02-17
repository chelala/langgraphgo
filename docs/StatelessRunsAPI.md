# \StatelessRunsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunBatchStatelessRunsPost**](StatelessRunsAPI.md#RunBatchStatelessRunsPost) | **Post** /runs/batch | Create Run Batch
[**RunStatelessRunsPost**](StatelessRunsAPI.md#RunStatelessRunsPost) | **Post** /runs | Create Background Run
[**StreamRunStatelessRunsStreamPost**](StatelessRunsAPI.md#StreamRunStatelessRunsStreamPost) | **Post** /runs/stream | Create Run, Stream Output
[**WaitRunStatelessRunsWaitPost**](StatelessRunsAPI.md#WaitRunStatelessRunsWaitPost) | **Post** /runs/wait | Create Run, Wait for Output



## RunBatchStatelessRunsPost

> interface{} RunBatchStatelessRunsPost(ctx).RunCreateStateless(runCreateStateless).Execute()

Create Run Batch



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
	runCreateStateless := []openapiclient.RunCreateStateless{*openapiclient.NewRunCreateStateless(*openapiclient.NewRunCreateStatefulAssistantId())} // []RunCreateStateless | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatelessRunsAPI.RunBatchStatelessRunsPost(context.Background()).RunCreateStateless(runCreateStateless).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatelessRunsAPI.RunBatchStatelessRunsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunBatchStatelessRunsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StatelessRunsAPI.RunBatchStatelessRunsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunBatchStatelessRunsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runCreateStateless** | [**[]RunCreateStateless**](RunCreateStateless.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunStatelessRunsPost

> interface{} RunStatelessRunsPost(ctx).RunCreateStateless(runCreateStateless).Execute()

Create Background Run



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
	runCreateStateless := *openapiclient.NewRunCreateStateless(*openapiclient.NewRunCreateStatefulAssistantId()) // RunCreateStateless | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatelessRunsAPI.RunStatelessRunsPost(context.Background()).RunCreateStateless(runCreateStateless).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatelessRunsAPI.RunStatelessRunsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunStatelessRunsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StatelessRunsAPI.RunStatelessRunsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunStatelessRunsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runCreateStateless** | [**RunCreateStateless**](RunCreateStateless.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamRunStatelessRunsStreamPost

> string StreamRunStatelessRunsStreamPost(ctx).RunCreateStateless(runCreateStateless).Execute()

Create Run, Stream Output



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
	runCreateStateless := *openapiclient.NewRunCreateStateless(*openapiclient.NewRunCreateStatefulAssistantId()) // RunCreateStateless | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatelessRunsAPI.StreamRunStatelessRunsStreamPost(context.Background()).RunCreateStateless(runCreateStateless).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatelessRunsAPI.StreamRunStatelessRunsStreamPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamRunStatelessRunsStreamPost`: string
	fmt.Fprintf(os.Stdout, "Response from `StatelessRunsAPI.StreamRunStatelessRunsStreamPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStreamRunStatelessRunsStreamPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runCreateStateless** | [**RunCreateStateless**](RunCreateStateless.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WaitRunStatelessRunsWaitPost

> interface{} WaitRunStatelessRunsWaitPost(ctx).RunCreateStateless(runCreateStateless).Execute()

Create Run, Wait for Output



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
	runCreateStateless := *openapiclient.NewRunCreateStateless(*openapiclient.NewRunCreateStatefulAssistantId()) // RunCreateStateless | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatelessRunsAPI.WaitRunStatelessRunsWaitPost(context.Background()).RunCreateStateless(runCreateStateless).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatelessRunsAPI.WaitRunStatelessRunsWaitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WaitRunStatelessRunsWaitPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StatelessRunsAPI.WaitRunStatelessRunsWaitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWaitRunStatelessRunsWaitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runCreateStateless** | [**RunCreateStateless**](RunCreateStateless.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

