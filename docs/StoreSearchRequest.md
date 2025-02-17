# StoreSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamespacePrefix** | Pointer to **[]string** | List of strings representing the namespace prefix. | [optional] 
**Filter** | Pointer to **map[string]interface{}** | Optional dictionary of key-value pairs to filter results. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of items to return (default is 10). | [optional] [default to 10]
**Offset** | Pointer to **int32** | Number of items to skip before returning results (default is 0). | [optional] [default to 0]

## Methods

### NewStoreSearchRequest

`func NewStoreSearchRequest() *StoreSearchRequest`

NewStoreSearchRequest instantiates a new StoreSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreSearchRequestWithDefaults

`func NewStoreSearchRequestWithDefaults() *StoreSearchRequest`

NewStoreSearchRequestWithDefaults instantiates a new StoreSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespacePrefix

`func (o *StoreSearchRequest) GetNamespacePrefix() []string`

GetNamespacePrefix returns the NamespacePrefix field if non-nil, zero value otherwise.

### GetNamespacePrefixOk

`func (o *StoreSearchRequest) GetNamespacePrefixOk() (*[]string, bool)`

GetNamespacePrefixOk returns a tuple with the NamespacePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespacePrefix

`func (o *StoreSearchRequest) SetNamespacePrefix(v []string)`

SetNamespacePrefix sets NamespacePrefix field to given value.

### HasNamespacePrefix

`func (o *StoreSearchRequest) HasNamespacePrefix() bool`

HasNamespacePrefix returns a boolean if a field has been set.

### SetNamespacePrefixNil

`func (o *StoreSearchRequest) SetNamespacePrefixNil(b bool)`

 SetNamespacePrefixNil sets the value for NamespacePrefix to be an explicit nil

### UnsetNamespacePrefix
`func (o *StoreSearchRequest) UnsetNamespacePrefix()`

UnsetNamespacePrefix ensures that no value is present for NamespacePrefix, not even an explicit nil
### GetFilter

`func (o *StoreSearchRequest) GetFilter() map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *StoreSearchRequest) GetFilterOk() (*map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *StoreSearchRequest) SetFilter(v map[string]interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *StoreSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLimit

`func (o *StoreSearchRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *StoreSearchRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *StoreSearchRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *StoreSearchRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *StoreSearchRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *StoreSearchRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *StoreSearchRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *StoreSearchRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


