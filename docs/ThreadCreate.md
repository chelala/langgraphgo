# ThreadCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | Pointer to **string** | The ID of the thread. If not provided, a random UUID will be generated. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata to add to thread. | [optional] 
**IfExists** | Pointer to **string** | How to handle duplicate creation. Must be either &#39;raise&#39; (raise error if duplicate), or &#39;do_nothing&#39; (return existing thread). | [optional] [default to "raise"]

## Methods

### NewThreadCreate

`func NewThreadCreate() *ThreadCreate`

NewThreadCreate instantiates a new ThreadCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadCreateWithDefaults

`func NewThreadCreateWithDefaults() *ThreadCreate`

NewThreadCreateWithDefaults instantiates a new ThreadCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *ThreadCreate) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *ThreadCreate) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *ThreadCreate) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *ThreadCreate) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetMetadata

`func (o *ThreadCreate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ThreadCreate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ThreadCreate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ThreadCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetIfExists

`func (o *ThreadCreate) GetIfExists() string`

GetIfExists returns the IfExists field if non-nil, zero value otherwise.

### GetIfExistsOk

`func (o *ThreadCreate) GetIfExistsOk() (*string, bool)`

GetIfExistsOk returns a tuple with the IfExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfExists

`func (o *ThreadCreate) SetIfExists(v string)`

SetIfExists sets IfExists field to given value.

### HasIfExists

`func (o *ThreadCreate) HasIfExists() bool`

HasIfExists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


