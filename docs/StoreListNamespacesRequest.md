# StoreListNamespacesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **[]string** | Optional list of strings representing the prefix to filter namespaces. | [optional] 
**Suffix** | Pointer to **[]string** | Optional list of strings representing the suffix to filter namespaces. | [optional] 
**MaxDepth** | Pointer to **int32** | Optional integer specifying the maximum depth of namespaces to return. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of namespaces to return (default is 100). | [optional] [default to 100]
**Offset** | Pointer to **int32** | Number of namespaces to skip before returning results (default is 0). | [optional] [default to 0]

## Methods

### NewStoreListNamespacesRequest

`func NewStoreListNamespacesRequest() *StoreListNamespacesRequest`

NewStoreListNamespacesRequest instantiates a new StoreListNamespacesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreListNamespacesRequestWithDefaults

`func NewStoreListNamespacesRequestWithDefaults() *StoreListNamespacesRequest`

NewStoreListNamespacesRequestWithDefaults instantiates a new StoreListNamespacesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *StoreListNamespacesRequest) GetPrefix() []string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *StoreListNamespacesRequest) GetPrefixOk() (*[]string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *StoreListNamespacesRequest) SetPrefix(v []string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *StoreListNamespacesRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSuffix

`func (o *StoreListNamespacesRequest) GetSuffix() []string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *StoreListNamespacesRequest) GetSuffixOk() (*[]string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *StoreListNamespacesRequest) SetSuffix(v []string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *StoreListNamespacesRequest) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetMaxDepth

`func (o *StoreListNamespacesRequest) GetMaxDepth() int32`

GetMaxDepth returns the MaxDepth field if non-nil, zero value otherwise.

### GetMaxDepthOk

`func (o *StoreListNamespacesRequest) GetMaxDepthOk() (*int32, bool)`

GetMaxDepthOk returns a tuple with the MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDepth

`func (o *StoreListNamespacesRequest) SetMaxDepth(v int32)`

SetMaxDepth sets MaxDepth field to given value.

### HasMaxDepth

`func (o *StoreListNamespacesRequest) HasMaxDepth() bool`

HasMaxDepth returns a boolean if a field has been set.

### GetLimit

`func (o *StoreListNamespacesRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *StoreListNamespacesRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *StoreListNamespacesRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *StoreListNamespacesRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *StoreListNamespacesRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *StoreListNamespacesRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *StoreListNamespacesRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *StoreListNamespacesRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


