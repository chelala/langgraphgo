# StorePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **[]string** | A list of strings representing the namespace path. | 
**Key** | **string** | The unique identifier for the item within the namespace. | 
**Value** | **map[string]interface{}** | A dictionary containing the item&#39;s data. | 

## Methods

### NewStorePutRequest

`func NewStorePutRequest(namespace []string, key string, value map[string]interface{}, ) *StorePutRequest`

NewStorePutRequest instantiates a new StorePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorePutRequestWithDefaults

`func NewStorePutRequestWithDefaults() *StorePutRequest`

NewStorePutRequestWithDefaults instantiates a new StorePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *StorePutRequest) GetNamespace() []string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *StorePutRequest) GetNamespaceOk() (*[]string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *StorePutRequest) SetNamespace(v []string)`

SetNamespace sets Namespace field to given value.


### GetKey

`func (o *StorePutRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StorePutRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StorePutRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *StorePutRequest) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StorePutRequest) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StorePutRequest) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


