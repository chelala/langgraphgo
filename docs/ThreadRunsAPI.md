# \ThreadRunsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRunHttpThreadsThreadIdRunsRunIdCancelPost**](ThreadRunsAPI.md#CancelRunHttpThreadsThreadIdRunsRunIdCancelPost) | **Post** /threads/{thread_id}/runs/{run_id}/cancel | Cancel Run
[**CreateRunThreadsThreadIdRunsPost**](ThreadRunsAPI.md#CreateRunThreadsThreadIdRunsPost) | **Post** /threads/{thread_id}/runs | Create Background Run
[**DeleteRunThreadsThreadIdRunsRunIdDelete**](ThreadRunsAPI.md#DeleteRunThreadsThreadIdRunsRunIdDelete) | **Delete** /threads/{thread_id}/runs/{run_id} | Delete Run
[**GetRunHttpThreadsThreadIdRunsRunIdGet**](ThreadRunsAPI.md#GetRunHttpThreadsThreadIdRunsRunIdGet) | **Get** /threads/{thread_id}/runs/{run_id} | Get Run
[**JoinRunHttpThreadsThreadIdRunsRunIdJoinGet**](ThreadRunsAPI.md#JoinRunHttpThreadsThreadIdRunsRunIdJoinGet) | **Get** /threads/{thread_id}/runs/{run_id}/join | Join Run
[**ListRunsHttpThreadsThreadIdRunsGet**](ThreadRunsAPI.md#ListRunsHttpThreadsThreadIdRunsGet) | **Get** /threads/{thread_id}/runs | List Runs
[**StreamRunHttpThreadsThreadIdRunsRunIdJoinGet**](ThreadRunsAPI.md#StreamRunHttpThreadsThreadIdRunsRunIdJoinGet) | **Get** /threads/{thread_id}/runs/{run_id}/stream | Join Run Stream
[**StreamRunThreadsThreadIdRunsStreamPost**](ThreadRunsAPI.md#StreamRunThreadsThreadIdRunsStreamPost) | **Post** /threads/{thread_id}/runs/stream | Create Run, Stream Output
[**WaitRunThreadsThreadIdRunsWaitPost**](ThreadRunsAPI.md#WaitRunThreadsThreadIdRunsWaitPost) | **Post** /threads/{thread_id}/runs/wait | Create Run, Wait for Output



## CancelRunHttpThreadsThreadIdRunsRunIdCancelPost

> interface{} CancelRunHttpThreadsThreadIdRunsRunIdCancelPost(ctx, threadId, runId).Wait(wait).Action(action).Execute()

Cancel Run

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
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the run.
	wait := true // bool |  (optional) (default to false)
	action := "action_example" // string | Action to take when cancelling the run. Possible values are `interrupt` or `rollback`. `interrupt` will simply cancel the run. `rollback` will cancel the run and delete the run and associated checkpoints afterwards. (optional) (default to "interrupt")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadRunsAPI.CancelRunHttpThreadsThreadIdRunsRunIdCancelPost(context.Background(), threadId, runId).Wait(wait).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadRunsAPI.CancelRunHttpThreadsThreadIdRunsRunIdCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelRunHttpThreadsThreadIdRunsRunIdCancelPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ThreadRunsAPI.CancelRunHttpThreadsThreadIdRunsRunIdCancelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 
**runId** | **string** | The ID of the run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRunHttpThreadsThreadIdRunsRunIdCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wait** | **bool** |  | [default to false]
 **action** | **string** | Action to take when cancelling the run. Possible values are &#x60;interrupt&#x60; or &#x60;rollback&#x60;. &#x60;interrupt&#x60; will simply cancel the run. &#x60;rollback&#x60; will cancel the run and delete the run and associated checkpoints afterwards. | [default to &quot;interrupt&quot;]

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


## CreateRunThreadsThreadIdRunsPost

> Run CreateRunThreadsThreadIdRunsPost(ctx, threadId).RunCreateStateful(runCreateStateful).Execute()

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
	runCreateStateful := *openapiclient.NewRunCreateStateful(*openapiclient.NewRunCreateStatefulAssistantId()) // RunCreateStateful | 
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the thread.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadRunsAPI.CreateRunThreadsThreadIdRunsPost(context.Background(), threadId).RunCreateStateful(runCreateStateful).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadRunsAPI.CreateRunThreadsThreadIdRunsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRunThreadsThreadIdRunsPost`: Run
	fmt.Fprintf(os.Stdout, "Response from `ThreadRunsAPI.CreateRunThreadsThreadIdRunsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRunThreadsThreadIdRunsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runCreateStateful** | [**RunCreateStateful**](RunCreateStateful.md) |  | 


### Return type

[**Run**](Run.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRunThreadsThreadIdRunsRunIdDelete

> interface{} DeleteRunThreadsThreadIdRunsRunIdDelete(ctx, threadId, runId).Execute()

Delete Run



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
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadRunsAPI.DeleteRunThreadsThreadIdRunsRunIdDelete(context.Background(), threadId, runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadRunsAPI.DeleteRunThreadsThreadIdRunsRunIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRunThreadsThreadIdRunsRunIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ThreadRunsAPI.DeleteRunThreadsThreadIdRunsRunIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 
**runId** | **string** | The ID of the run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRunThreadsThreadIdRunsRunIdDeleteRequest struct via the builder pattern


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


## GetRunHttpThreadsThreadIdRunsRunIdGet

> Run GetRunHttpThreadsThreadIdRunsRunIdGet(ctx, threadId, runId).Execute()

Get Run



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
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadRunsAPI.GetRunHttpThreadsThreadIdRunsRunIdGet(context.Background(), threadId, runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadRunsAPI.GetRunHttpThreadsThreadIdRunsRunIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunHttpThreadsThreadIdRunsRunIdGet`: Run
	fmt.Fprintf(os.Stdout, "Response from `ThreadRunsAPI.GetRunHttpThreadsThreadIdRunsRunIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 
**runId** | **string** | The ID of the run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunHttpThreadsThreadIdRunsRunIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Run**](Run.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinRunHttpThreadsThreadIdRunsRunIdJoinGet

> interface{} JoinRunHttpThreadsThreadIdRunsRunIdJoinGet(ctx, threadId, runId).CancelOnDisconnect(cancelOnDisconnect).Execute()

Join Run



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
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the run.
	cancelOnDisconnect := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadRunsAPI.JoinRunHttpThreadsThreadIdRunsRunIdJoinGet(context.Background(), threadId, runId).CancelOnDisconnect(cancelOnDisconnect).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadRunsAPI.JoinRunHttpThreadsThreadIdRunsRunIdJoinGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JoinRunHttpThreadsThreadIdRunsRunIdJoinGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ThreadRunsAPI.JoinRunHttpThreadsThreadIdRunsRunIdJoinGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 
**runId** | **string** | The ID of the run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoinRunHttpThreadsThreadIdRunsRunIdJoinGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cancelOnDisconnect** | **bool** |  | [default to false]

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


## ListRunsHttpThreadsThreadIdRunsGet

> []Run ListRunsHttpThreadsThreadIdRunsGet(ctx, threadId).Limit(limit).Offset(offset).Status(status).Execute()

List Runs



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
	offset := int32(56) // int32 |  (optional) (default to 0)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadRunsAPI.ListRunsHttpThreadsThreadIdRunsGet(context.Background(), threadId).Limit(limit).Offset(offset).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadRunsAPI.ListRunsHttpThreadsThreadIdRunsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRunsHttpThreadsThreadIdRunsGet`: []Run
	fmt.Fprintf(os.Stdout, "Response from `ThreadRunsAPI.ListRunsHttpThreadsThreadIdRunsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRunsHttpThreadsThreadIdRunsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 10]
 **offset** | **int32** |  | [default to 0]
 **status** | **string** |  | 

### Return type

[**[]Run**](Run.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamRunHttpThreadsThreadIdRunsRunIdJoinGet

> string StreamRunHttpThreadsThreadIdRunsRunIdJoinGet(ctx, threadId, runId).Execute()

Join Run Stream



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
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadRunsAPI.StreamRunHttpThreadsThreadIdRunsRunIdJoinGet(context.Background(), threadId, runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadRunsAPI.StreamRunHttpThreadsThreadIdRunsRunIdJoinGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamRunHttpThreadsThreadIdRunsRunIdJoinGet`: string
	fmt.Fprintf(os.Stdout, "Response from `ThreadRunsAPI.StreamRunHttpThreadsThreadIdRunsRunIdJoinGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 
**runId** | **string** | The ID of the run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamRunHttpThreadsThreadIdRunsRunIdJoinGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamRunThreadsThreadIdRunsStreamPost

> string StreamRunThreadsThreadIdRunsStreamPost(ctx, threadId).RunCreateStateful(runCreateStateful).Execute()

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
	runCreateStateful := *openapiclient.NewRunCreateStateful(*openapiclient.NewRunCreateStatefulAssistantId()) // RunCreateStateful | 
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the thread.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadRunsAPI.StreamRunThreadsThreadIdRunsStreamPost(context.Background(), threadId).RunCreateStateful(runCreateStateful).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadRunsAPI.StreamRunThreadsThreadIdRunsStreamPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamRunThreadsThreadIdRunsStreamPost`: string
	fmt.Fprintf(os.Stdout, "Response from `ThreadRunsAPI.StreamRunThreadsThreadIdRunsStreamPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamRunThreadsThreadIdRunsStreamPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runCreateStateful** | [**RunCreateStateful**](RunCreateStateful.md) |  | 


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


## WaitRunThreadsThreadIdRunsWaitPost

> interface{} WaitRunThreadsThreadIdRunsWaitPost(ctx, threadId).RunCreateStateful(runCreateStateful).Execute()

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
	runCreateStateful := *openapiclient.NewRunCreateStateful(*openapiclient.NewRunCreateStatefulAssistantId()) // RunCreateStateful | 
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the thread.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadRunsAPI.WaitRunThreadsThreadIdRunsWaitPost(context.Background(), threadId).RunCreateStateful(runCreateStateful).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadRunsAPI.WaitRunThreadsThreadIdRunsWaitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WaitRunThreadsThreadIdRunsWaitPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ThreadRunsAPI.WaitRunThreadsThreadIdRunsWaitPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWaitRunThreadsThreadIdRunsWaitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runCreateStateful** | [**RunCreateStateful**](RunCreateStateful.md) |  | 


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

