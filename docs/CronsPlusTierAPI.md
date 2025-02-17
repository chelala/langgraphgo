# \CronsPlusTierAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCronRunsCronsPost**](CronsPlusTierAPI.md#CreateCronRunsCronsPost) | **Post** /runs/crons | Create Cron
[**CreateThreadCronThreadsThreadIdRunsCronsPost**](CronsPlusTierAPI.md#CreateThreadCronThreadsThreadIdRunsCronsPost) | **Post** /threads/{thread_id}/runs/crons | Create Thread Cron
[**DeleteCronRunsCronsCronIdDelete**](CronsPlusTierAPI.md#DeleteCronRunsCronsCronIdDelete) | **Delete** /runs/crons/{cron_id} | Delete Cron
[**SearchCronsRunsCronsPost**](CronsPlusTierAPI.md#SearchCronsRunsCronsPost) | **Post** /runs/crons/search | Search Crons



## CreateCronRunsCronsPost

> Cron CreateCronRunsCronsPost(ctx).CronCreate(cronCreate).Execute()

Create Cron



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
	cronCreate := *openapiclient.NewCronCreate("Schedule_example", *openapiclient.NewCronCreateAssistantId()) // CronCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CronsPlusTierAPI.CreateCronRunsCronsPost(context.Background()).CronCreate(cronCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronsPlusTierAPI.CreateCronRunsCronsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCronRunsCronsPost`: Cron
	fmt.Fprintf(os.Stdout, "Response from `CronsPlusTierAPI.CreateCronRunsCronsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCronRunsCronsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cronCreate** | [**CronCreate**](CronCreate.md) |  | 

### Return type

[**Cron**](Cron.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateThreadCronThreadsThreadIdRunsCronsPost

> Cron CreateThreadCronThreadsThreadIdRunsCronsPost(ctx, threadId).CronCreate(cronCreate).Execute()

Create Thread Cron



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
	cronCreate := *openapiclient.NewCronCreate("Schedule_example", *openapiclient.NewCronCreateAssistantId()) // CronCreate | 
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the thread.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CronsPlusTierAPI.CreateThreadCronThreadsThreadIdRunsCronsPost(context.Background(), threadId).CronCreate(cronCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronsPlusTierAPI.CreateThreadCronThreadsThreadIdRunsCronsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateThreadCronThreadsThreadIdRunsCronsPost`: Cron
	fmt.Fprintf(os.Stdout, "Response from `CronsPlusTierAPI.CreateThreadCronThreadsThreadIdRunsCronsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | The ID of the thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateThreadCronThreadsThreadIdRunsCronsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cronCreate** | [**CronCreate**](CronCreate.md) |  | 


### Return type

[**Cron**](Cron.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCronRunsCronsCronIdDelete

> interface{} DeleteCronRunsCronsCronIdDelete(ctx, cronId).Execute()

Delete Cron



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
	cronId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CronsPlusTierAPI.DeleteCronRunsCronsCronIdDelete(context.Background(), cronId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronsPlusTierAPI.DeleteCronRunsCronsCronIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCronRunsCronsCronIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CronsPlusTierAPI.DeleteCronRunsCronsCronIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cronId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCronRunsCronsCronIdDeleteRequest struct via the builder pattern


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


## SearchCronsRunsCronsPost

> []Cron SearchCronsRunsCronsPost(ctx).CronSearch(cronSearch).Execute()

Search Crons



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
	cronSearch := *openapiclient.NewCronSearch() // CronSearch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CronsPlusTierAPI.SearchCronsRunsCronsPost(context.Background()).CronSearch(cronSearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronsPlusTierAPI.SearchCronsRunsCronsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCronsRunsCronsPost`: []Cron
	fmt.Fprintf(os.Stdout, "Response from `CronsPlusTierAPI.SearchCronsRunsCronsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCronsRunsCronsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cronSearch** | [**CronSearch**](CronSearch.md) |  | 

### Return type

[**[]Cron**](Cron.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

