# Config2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**RecursionLimit** | Pointer to **int32** |  | [optional] 
**Configurable** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConfig2

`func NewConfig2() *Config2`

NewConfig2 instantiates a new Config2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfig2WithDefaults

`func NewConfig2WithDefaults() *Config2`

NewConfig2WithDefaults instantiates a new Config2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *Config2) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Config2) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Config2) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Config2) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRecursionLimit

`func (o *Config2) GetRecursionLimit() int32`

GetRecursionLimit returns the RecursionLimit field if non-nil, zero value otherwise.

### GetRecursionLimitOk

`func (o *Config2) GetRecursionLimitOk() (*int32, bool)`

GetRecursionLimitOk returns a tuple with the RecursionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionLimit

`func (o *Config2) SetRecursionLimit(v int32)`

SetRecursionLimit sets RecursionLimit field to given value.

### HasRecursionLimit

`func (o *Config2) HasRecursionLimit() bool`

HasRecursionLimit returns a boolean if a field has been set.

### GetConfigurable

`func (o *Config2) GetConfigurable() map[string]interface{}`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *Config2) GetConfigurableOk() (*map[string]interface{}, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *Config2) SetConfigurable(v map[string]interface{})`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *Config2) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


