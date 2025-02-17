# Go API client for langgraphgo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: 1.0.0
- Generator version: 7.11.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import langgraphgo "github.com/chelala/langgraphgo"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `langgraphgo.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), langgraphgo.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `langgraphgo.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), langgraphgo.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `langgraphgo.ContextOperationServerIndices` and `langgraphgo.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), langgraphgo.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), langgraphgo.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AssistantsAPI* | [**CreateAssistantAssistantsPost**](docs/AssistantsAPI.md#createassistantassistantspost) | **Post** /assistants | Create Assistant
*AssistantsAPI* | [**DeleteAssistantAssistantsAssistantIdDelete**](docs/AssistantsAPI.md#deleteassistantassistantsassistantiddelete) | **Delete** /assistants/{assistant_id} | Delete Assistant
*AssistantsAPI* | [**GetAssistantAssistantsAssistantIdGet**](docs/AssistantsAPI.md#getassistantassistantsassistantidget) | **Get** /assistants/{assistant_id} | Get Assistant
*AssistantsAPI* | [**GetAssistantGraphAssistantsAssistantIdGraphGet**](docs/AssistantsAPI.md#getassistantgraphassistantsassistantidgraphget) | **Get** /assistants/{assistant_id}/graph | Get Assistant Graph
*AssistantsAPI* | [**GetAssistantSchemasAssistantsAssistantIdSchemasGet**](docs/AssistantsAPI.md#getassistantschemasassistantsassistantidschemasget) | **Get** /assistants/{assistant_id}/schemas | Get Assistant Schemas
*AssistantsAPI* | [**GetAssistantSubgraphsAssistantsAssistantIdSubgraphsGet**](docs/AssistantsAPI.md#getassistantsubgraphsassistantsassistantidsubgraphsget) | **Get** /assistants/{assistant_id}/subgraphs | Get Assistant Subgraphs
*AssistantsAPI* | [**GetAssistantSubgraphsAssistantsAssistantIdSubgraphsNamespaceGet**](docs/AssistantsAPI.md#getassistantsubgraphsassistantsassistantidsubgraphsnamespaceget) | **Get** /assistants/{assistant_id}/subgraphs/{namespace} | Get Assistant Subgraphs by Namespace
*AssistantsAPI* | [**GetAssistantVersionsAssistantsAssistantIdVersionsGet**](docs/AssistantsAPI.md#getassistantversionsassistantsassistantidversionsget) | **Post** /assistants/{assistant_id}/versions | Get Assistant Versions
*AssistantsAPI* | [**PatchAssistantAssistantsAssistantIdPatch**](docs/AssistantsAPI.md#patchassistantassistantsassistantidpatch) | **Patch** /assistants/{assistant_id} | Patch Assistant
*AssistantsAPI* | [**SearchAssistantsAssistantsSearchPost**](docs/AssistantsAPI.md#searchassistantsassistantssearchpost) | **Post** /assistants/search | Search Assistants
*AssistantsAPI* | [**SetLatestAssistantVersionAssistantsAssistantIdVersionsPost**](docs/AssistantsAPI.md#setlatestassistantversionassistantsassistantidversionspost) | **Post** /assistants/{assistant_id}/latest | Set Latest Assistant Version
*CronsPlusTierAPI* | [**CreateCronRunsCronsPost**](docs/CronsPlusTierAPI.md#createcronrunscronspost) | **Post** /runs/crons | Create Cron
*CronsPlusTierAPI* | [**CreateThreadCronThreadsThreadIdRunsCronsPost**](docs/CronsPlusTierAPI.md#createthreadcronthreadsthreadidrunscronspost) | **Post** /threads/{thread_id}/runs/crons | Create Thread Cron
*CronsPlusTierAPI* | [**DeleteCronRunsCronsCronIdDelete**](docs/CronsPlusTierAPI.md#deletecronrunscronscroniddelete) | **Delete** /runs/crons/{cron_id} | Delete Cron
*CronsPlusTierAPI* | [**SearchCronsRunsCronsPost**](docs/CronsPlusTierAPI.md#searchcronsrunscronspost) | **Post** /runs/crons/search | Search Crons
*StatelessRunsAPI* | [**RunBatchStatelessRunsPost**](docs/StatelessRunsAPI.md#runbatchstatelessrunspost) | **Post** /runs/batch | Create Run Batch
*StatelessRunsAPI* | [**RunStatelessRunsPost**](docs/StatelessRunsAPI.md#runstatelessrunspost) | **Post** /runs | Create Background Run
*StatelessRunsAPI* | [**StreamRunStatelessRunsStreamPost**](docs/StatelessRunsAPI.md#streamrunstatelessrunsstreampost) | **Post** /runs/stream | Create Run, Stream Output
*StatelessRunsAPI* | [**WaitRunStatelessRunsWaitPost**](docs/StatelessRunsAPI.md#waitrunstatelessrunswaitpost) | **Post** /runs/wait | Create Run, Wait for Output
*StoreAPI* | [**DeleteItem**](docs/StoreAPI.md#deleteitem) | **Delete** /store/items | Delete an item.
*StoreAPI* | [**GetItem**](docs/StoreAPI.md#getitem) | **Get** /store/items | Retrieve a single item.
*StoreAPI* | [**ListNamespaces**](docs/StoreAPI.md#listnamespaces) | **Post** /store/namespaces | List namespaces with optional match conditions.
*StoreAPI* | [**PutItem**](docs/StoreAPI.md#putitem) | **Put** /store/items | Store or update an item.
*StoreAPI* | [**SearchItems**](docs/StoreAPI.md#searchitems) | **Post** /store/items/search | Search for items within a namespace prefix.
*ThreadRunsAPI* | [**CancelRunHttpThreadsThreadIdRunsRunIdCancelPost**](docs/ThreadRunsAPI.md#cancelrunhttpthreadsthreadidrunsrunidcancelpost) | **Post** /threads/{thread_id}/runs/{run_id}/cancel | Cancel Run
*ThreadRunsAPI* | [**CreateRunThreadsThreadIdRunsPost**](docs/ThreadRunsAPI.md#createrunthreadsthreadidrunspost) | **Post** /threads/{thread_id}/runs | Create Background Run
*ThreadRunsAPI* | [**DeleteRunThreadsThreadIdRunsRunIdDelete**](docs/ThreadRunsAPI.md#deleterunthreadsthreadidrunsruniddelete) | **Delete** /threads/{thread_id}/runs/{run_id} | Delete Run
*ThreadRunsAPI* | [**GetRunHttpThreadsThreadIdRunsRunIdGet**](docs/ThreadRunsAPI.md#getrunhttpthreadsthreadidrunsrunidget) | **Get** /threads/{thread_id}/runs/{run_id} | Get Run
*ThreadRunsAPI* | [**JoinRunHttpThreadsThreadIdRunsRunIdJoinGet**](docs/ThreadRunsAPI.md#joinrunhttpthreadsthreadidrunsrunidjoinget) | **Get** /threads/{thread_id}/runs/{run_id}/join | Join Run
*ThreadRunsAPI* | [**ListRunsHttpThreadsThreadIdRunsGet**](docs/ThreadRunsAPI.md#listrunshttpthreadsthreadidrunsget) | **Get** /threads/{thread_id}/runs | List Runs
*ThreadRunsAPI* | [**StreamRunHttpThreadsThreadIdRunsRunIdJoinGet**](docs/ThreadRunsAPI.md#streamrunhttpthreadsthreadidrunsrunidjoinget) | **Get** /threads/{thread_id}/runs/{run_id}/stream | Join Run Stream
*ThreadRunsAPI* | [**StreamRunThreadsThreadIdRunsStreamPost**](docs/ThreadRunsAPI.md#streamrunthreadsthreadidrunsstreampost) | **Post** /threads/{thread_id}/runs/stream | Create Run, Stream Output
*ThreadRunsAPI* | [**WaitRunThreadsThreadIdRunsWaitPost**](docs/ThreadRunsAPI.md#waitrunthreadsthreadidrunswaitpost) | **Post** /threads/{thread_id}/runs/wait | Create Run, Wait for Output
*ThreadsAPI* | [**CopyThreadPostThreadsThreadIdCopyPost**](docs/ThreadsAPI.md#copythreadpostthreadsthreadidcopypost) | **Post** /threads/{thread_id}/copy | Copy Thread
*ThreadsAPI* | [**CreateThreadThreadsPost**](docs/ThreadsAPI.md#createthreadthreadspost) | **Post** /threads | Create Thread
*ThreadsAPI* | [**DeleteThreadThreadsThreadIdDelete**](docs/ThreadsAPI.md#deletethreadthreadsthreadiddelete) | **Delete** /threads/{thread_id} | Delete Thread
*ThreadsAPI* | [**GetLatestThreadStateThreadsThreadIdStateGet**](docs/ThreadsAPI.md#getlatestthreadstatethreadsthreadidstateget) | **Get** /threads/{thread_id}/state | Get Thread State
*ThreadsAPI* | [**GetThreadHistoryPostThreadsThreadIdHistoryPost**](docs/ThreadsAPI.md#getthreadhistorypostthreadsthreadidhistorypost) | **Post** /threads/{thread_id}/history | Get Thread History Post
*ThreadsAPI* | [**GetThreadHistoryThreadsThreadIdHistoryGet**](docs/ThreadsAPI.md#getthreadhistorythreadsthreadidhistoryget) | **Get** /threads/{thread_id}/history | Get Thread History
*ThreadsAPI* | [**GetThreadThreadsThreadIdGet**](docs/ThreadsAPI.md#getthreadthreadsthreadidget) | **Get** /threads/{thread_id} | Get Thread
*ThreadsAPI* | [**PatchThreadThreadsThreadIdPatch**](docs/ThreadsAPI.md#patchthreadthreadsthreadidpatch) | **Patch** /threads/{thread_id} | Patch Thread
*ThreadsAPI* | [**PostThreadStateAtCheckpointThreadsThreadIdStateCheckpointIdGet**](docs/ThreadsAPI.md#postthreadstateatcheckpointthreadsthreadidstatecheckpointidget) | **Post** /threads/{thread_id}/state/checkpoint | Get Thread State At Checkpoint
*ThreadsAPI* | [**SearchThreadsThreadsSearchPost**](docs/ThreadsAPI.md#searchthreadsthreadssearchpost) | **Post** /threads/search | Search Threads
*ThreadsAPI* | [**UpdateThreadStateThreadsThreadIdStatePost**](docs/ThreadsAPI.md#updatethreadstatethreadsthreadidstatepost) | **Post** /threads/{thread_id}/state | Update Thread State


## Documentation For Models

 - [Assistant](docs/Assistant.md)
 - [AssistantCreate](docs/AssistantCreate.md)
 - [AssistantPatch](docs/AssistantPatch.md)
 - [AssistantSearchRequest](docs/AssistantSearchRequest.md)
 - [AssistantVersionChange](docs/AssistantVersionChange.md)
 - [AssistantVersionsSearchRequest](docs/AssistantVersionsSearchRequest.md)
 - [CheckpointConfig](docs/CheckpointConfig.md)
 - [Command](docs/Command.md)
 - [Config](docs/Config.md)
 - [Config1](docs/Config1.md)
 - [Config2](docs/Config2.md)
 - [Cron](docs/Cron.md)
 - [CronCreate](docs/CronCreate.md)
 - [CronCreateAssistantId](docs/CronCreateAssistantId.md)
 - [CronSearch](docs/CronSearch.md)
 - [GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter](docs/GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter.md)
 - [Goto](docs/Goto.md)
 - [GraphSchema](docs/GraphSchema.md)
 - [GraphSchemaNoId](docs/GraphSchemaNoId.md)
 - [Input](docs/Input.md)
 - [Input1](docs/Input1.md)
 - [InterruptAfter](docs/InterruptAfter.md)
 - [InterruptBefore](docs/InterruptBefore.md)
 - [Item](docs/Item.md)
 - [Resume](docs/Resume.md)
 - [Run](docs/Run.md)
 - [RunCreateStateful](docs/RunCreateStateful.md)
 - [RunCreateStatefulAssistantId](docs/RunCreateStatefulAssistantId.md)
 - [RunCreateStateless](docs/RunCreateStateless.md)
 - [SearchItemsResponse](docs/SearchItemsResponse.md)
 - [Send](docs/Send.md)
 - [StoreDeleteRequest](docs/StoreDeleteRequest.md)
 - [StoreListNamespacesRequest](docs/StoreListNamespacesRequest.md)
 - [StorePutRequest](docs/StorePutRequest.md)
 - [StoreSearchRequest](docs/StoreSearchRequest.md)
 - [StreamMode](docs/StreamMode.md)
 - [Thread](docs/Thread.md)
 - [ThreadCreate](docs/ThreadCreate.md)
 - [ThreadPatch](docs/ThreadPatch.md)
 - [ThreadSearchRequest](docs/ThreadSearchRequest.md)
 - [ThreadState](docs/ThreadState.md)
 - [ThreadStateCheckpointRequest](docs/ThreadStateCheckpointRequest.md)
 - [ThreadStateSearch](docs/ThreadStateSearch.md)
 - [ThreadStateTasksInner](docs/ThreadStateTasksInner.md)
 - [ThreadStateUpdate](docs/ThreadStateUpdate.md)
 - [ThreadStateUpdateResponse](docs/ThreadStateUpdateResponse.md)
 - [Values](docs/Values.md)
 - [Values1](docs/Values1.md)
 - [Xray](docs/Xray.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



