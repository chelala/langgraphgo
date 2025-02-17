# StoreDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **[]string** | A list of strings representing the namespace path. | [optional] 
**Key** | **string** | The unique identifier for the item. | 

## Methods

### NewStoreDeleteRequest

`func NewStoreDeleteRequest(key string, ) *StoreDeleteRequest`

NewStoreDeleteRequest instantiates a new StoreDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDeleteRequestWithDefaults

`func NewStoreDeleteRequestWithDefaults() *StoreDeleteRequest`

NewStoreDeleteRequestWithDefaults instantiates a new StoreDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *StoreDeleteRequest) GetNamespace() []string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *StoreDeleteRequest) GetNamespaceOk() (*[]string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *StoreDeleteRequest) SetNamespace(v []string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *StoreDeleteRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetKey

`func (o *StoreDeleteRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StoreDeleteRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StoreDeleteRequest) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


