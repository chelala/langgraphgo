# \AssistantsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssistantAssistantsPost**](AssistantsAPI.md#CreateAssistantAssistantsPost) | **Post** /assistants | Create Assistant
[**DeleteAssistantAssistantsAssistantIdDelete**](AssistantsAPI.md#DeleteAssistantAssistantsAssistantIdDelete) | **Delete** /assistants/{assistant_id} | Delete Assistant
[**GetAssistantAssistantsAssistantIdGet**](AssistantsAPI.md#GetAssistantAssistantsAssistantIdGet) | **Get** /assistants/{assistant_id} | Get Assistant
[**GetAssistantGraphAssistantsAssistantIdGraphGet**](AssistantsAPI.md#GetAssistantGraphAssistantsAssistantIdGraphGet) | **Get** /assistants/{assistant_id}/graph | Get Assistant Graph
[**GetAssistantSchemasAssistantsAssistantIdSchemasGet**](AssistantsAPI.md#GetAssistantSchemasAssistantsAssistantIdSchemasGet) | **Get** /assistants/{assistant_id}/schemas | Get Assistant Schemas
[**GetAssistantSubgraphsAssistantsAssistantIdSubgraphsGet**](AssistantsAPI.md#GetAssistantSubgraphsAssistantsAssistantIdSubgraphsGet) | **Get** /assistants/{assistant_id}/subgraphs | Get Assistant Subgraphs
[**GetAssistantSubgraphsAssistantsAssistantIdSubgraphsNamespaceGet**](AssistantsAPI.md#GetAssistantSubgraphsAssistantsAssistantIdSubgraphsNamespaceGet) | **Get** /assistants/{assistant_id}/subgraphs/{namespace} | Get Assistant Subgraphs by Namespace
[**GetAssistantVersionsAssistantsAssistantIdVersionsGet**](AssistantsAPI.md#GetAssistantVersionsAssistantsAssistantIdVersionsGet) | **Post** /assistants/{assistant_id}/versions | Get Assistant Versions
[**PatchAssistantAssistantsAssistantIdPatch**](AssistantsAPI.md#PatchAssistantAssistantsAssistantIdPatch) | **Patch** /assistants/{assistant_id} | Patch Assistant
[**SearchAssistantsAssistantsSearchPost**](AssistantsAPI.md#SearchAssistantsAssistantsSearchPost) | **Post** /assistants/search | Search Assistants
[**SetLatestAssistantVersionAssistantsAssistantIdVersionsPost**](AssistantsAPI.md#SetLatestAssistantVersionAssistantsAssistantIdVersionsPost) | **Post** /assistants/{assistant_id}/latest | Set Latest Assistant Version



## CreateAssistantAssistantsPost

> Assistant CreateAssistantAssistantsPost(ctx).AssistantCreate(assistantCreate).Execute()

Create Assistant



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
	assistantCreate := *openapiclient.NewAssistantCreate("GraphId_example") // AssistantCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.CreateAssistantAssistantsPost(context.Background()).AssistantCreate(assistantCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.CreateAssistantAssistantsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssistantAssistantsPost`: Assistant
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.CreateAssistantAssistantsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssistantAssistantsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assistantCreate** | [**AssistantCreate**](AssistantCreate.md) |  | 

### Return type

[**Assistant**](Assistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssistantAssistantsAssistantIdDelete

> interface{} DeleteAssistantAssistantsAssistantIdDelete(ctx, assistantId).Execute()

Delete Assistant



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
	assistantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the assistant.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.DeleteAssistantAssistantsAssistantIdDelete(context.Background(), assistantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.DeleteAssistantAssistantsAssistantIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAssistantAssistantsAssistantIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.DeleteAssistantAssistantsAssistantIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** | The ID of the assistant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssistantAssistantsAssistantIdDeleteRequest struct via the builder pattern


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


## GetAssistantAssistantsAssistantIdGet

> Assistant GetAssistantAssistantsAssistantIdGet(ctx, assistantId).Execute()

Get Assistant



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
	assistantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the assistant.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.GetAssistantAssistantsAssistantIdGet(context.Background(), assistantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.GetAssistantAssistantsAssistantIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssistantAssistantsAssistantIdGet`: Assistant
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.GetAssistantAssistantsAssistantIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** | The ID of the assistant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssistantAssistantsAssistantIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Assistant**](Assistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssistantGraphAssistantsAssistantIdGraphGet

> map[string][]map[string]interface{} GetAssistantGraphAssistantsAssistantIdGraphGet(ctx, assistantId).Xray(xray).Execute()

Get Assistant Graph



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
	assistantId := *openapiclient.NewGetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter() // GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter | The ID of the assistant.
	xray := openapiclient.Xray{Bool: new(bool)} // Xray | Include graph representation of subgraphs. If an integer value is provided, only subgraphs with a depth less than or equal to the value will be included. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.GetAssistantGraphAssistantsAssistantIdGraphGet(context.Background(), assistantId).Xray(xray).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.GetAssistantGraphAssistantsAssistantIdGraphGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssistantGraphAssistantsAssistantIdGraphGet`: map[string][]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.GetAssistantGraphAssistantsAssistantIdGraphGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | [**GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter**](.md) | The ID of the assistant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssistantGraphAssistantsAssistantIdGraphGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xray** | [**Xray**](Xray.md) | Include graph representation of subgraphs. If an integer value is provided, only subgraphs with a depth less than or equal to the value will be included. | [default to false]

### Return type

[**map[string][]map[string]interface{}**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssistantSchemasAssistantsAssistantIdSchemasGet

> GraphSchema GetAssistantSchemasAssistantsAssistantIdSchemasGet(ctx, assistantId).Execute()

Get Assistant Schemas



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
	assistantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the assistant.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.GetAssistantSchemasAssistantsAssistantIdSchemasGet(context.Background(), assistantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.GetAssistantSchemasAssistantsAssistantIdSchemasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssistantSchemasAssistantsAssistantIdSchemasGet`: GraphSchema
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.GetAssistantSchemasAssistantsAssistantIdSchemasGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** | The ID of the assistant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssistantSchemasAssistantsAssistantIdSchemasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GraphSchema**](GraphSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssistantSubgraphsAssistantsAssistantIdSubgraphsGet

> map[string]GraphSchemaNoId GetAssistantSubgraphsAssistantsAssistantIdSubgraphsGet(ctx, assistantId).Recurse(recurse).Execute()

Get Assistant Subgraphs



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
	assistantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the assistant.
	recurse := true // bool | Recursively retrieve subgraphs of subgraphs. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.GetAssistantSubgraphsAssistantsAssistantIdSubgraphsGet(context.Background(), assistantId).Recurse(recurse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.GetAssistantSubgraphsAssistantsAssistantIdSubgraphsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssistantSubgraphsAssistantsAssistantIdSubgraphsGet`: map[string]GraphSchemaNoId
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.GetAssistantSubgraphsAssistantsAssistantIdSubgraphsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** | The ID of the assistant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssistantSubgraphsAssistantsAssistantIdSubgraphsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recurse** | **bool** | Recursively retrieve subgraphs of subgraphs. | [default to false]

### Return type

[**map[string]GraphSchemaNoId**](GraphSchemaNoId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssistantSubgraphsAssistantsAssistantIdSubgraphsNamespaceGet

> map[string]GraphSchemaNoId GetAssistantSubgraphsAssistantsAssistantIdSubgraphsNamespaceGet(ctx, assistantId, namespace).Recurse(recurse).Execute()

Get Assistant Subgraphs by Namespace



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
	assistantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the assistant.
	namespace := "namespace_example" // string | Namespace of the subgraph to filter by.
	recurse := true // bool | Recursively retrieve subgraphs of subgraphs. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.GetAssistantSubgraphsAssistantsAssistantIdSubgraphsNamespaceGet(context.Background(), assistantId, namespace).Recurse(recurse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.GetAssistantSubgraphsAssistantsAssistantIdSubgraphsNamespaceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssistantSubgraphsAssistantsAssistantIdSubgraphsNamespaceGet`: map[string]GraphSchemaNoId
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.GetAssistantSubgraphsAssistantsAssistantIdSubgraphsNamespaceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** | The ID of the assistant. | 
**namespace** | **string** | Namespace of the subgraph to filter by. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssistantSubgraphsAssistantsAssistantIdSubgraphsNamespaceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recurse** | **bool** | Recursively retrieve subgraphs of subgraphs. | [default to false]

### Return type

[**map[string]GraphSchemaNoId**](GraphSchemaNoId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssistantVersionsAssistantsAssistantIdVersionsGet

> []Assistant GetAssistantVersionsAssistantsAssistantIdVersionsGet(ctx, assistantId).Execute()

Get Assistant Versions



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
	assistantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the assistant.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.GetAssistantVersionsAssistantsAssistantIdVersionsGet(context.Background(), assistantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.GetAssistantVersionsAssistantsAssistantIdVersionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssistantVersionsAssistantsAssistantIdVersionsGet`: []Assistant
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.GetAssistantVersionsAssistantsAssistantIdVersionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** | The ID of the assistant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssistantVersionsAssistantsAssistantIdVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Assistant**](Assistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAssistantAssistantsAssistantIdPatch

> Assistant PatchAssistantAssistantsAssistantIdPatch(ctx, assistantId).AssistantPatch(assistantPatch).Execute()

Patch Assistant



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
	assistantPatch := *openapiclient.NewAssistantPatch() // AssistantPatch | 
	assistantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the assistant.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.PatchAssistantAssistantsAssistantIdPatch(context.Background(), assistantId).AssistantPatch(assistantPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.PatchAssistantAssistantsAssistantIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAssistantAssistantsAssistantIdPatch`: Assistant
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.PatchAssistantAssistantsAssistantIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** | The ID of the assistant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAssistantAssistantsAssistantIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assistantPatch** | [**AssistantPatch**](AssistantPatch.md) |  | 


### Return type

[**Assistant**](Assistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAssistantsAssistantsSearchPost

> []Assistant SearchAssistantsAssistantsSearchPost(ctx).AssistantSearchRequest(assistantSearchRequest).Execute()

Search Assistants



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
	assistantSearchRequest := *openapiclient.NewAssistantSearchRequest() // AssistantSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.SearchAssistantsAssistantsSearchPost(context.Background()).AssistantSearchRequest(assistantSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.SearchAssistantsAssistantsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAssistantsAssistantsSearchPost`: []Assistant
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.SearchAssistantsAssistantsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAssistantsAssistantsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assistantSearchRequest** | [**AssistantSearchRequest**](AssistantSearchRequest.md) |  | 

### Return type

[**[]Assistant**](Assistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLatestAssistantVersionAssistantsAssistantIdVersionsPost

> Assistant SetLatestAssistantVersionAssistantsAssistantIdVersionsPost(ctx, assistantId).Version(version).Execute()

Set Latest Assistant Version



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
	assistantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the assistant.
	version := int32(56) // int32 | The version to change to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.SetLatestAssistantVersionAssistantsAssistantIdVersionsPost(context.Background(), assistantId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.SetLatestAssistantVersionAssistantsAssistantIdVersionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetLatestAssistantVersionAssistantsAssistantIdVersionsPost`: Assistant
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.SetLatestAssistantVersionAssistantsAssistantIdVersionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** | The ID of the assistant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLatestAssistantVersionAssistantsAssistantIdVersionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | The version to change to. | 

### Return type

[**Assistant**](Assistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

