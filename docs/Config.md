# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**RecursionLimit** | Pointer to **int32** |  | [optional] 
**Configurable** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConfig

`func NewConfig() *Config`

NewConfig instantiates a new Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWithDefaults

`func NewConfigWithDefaults() *Config`

NewConfigWithDefaults instantiates a new Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *Config) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Config) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Config) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Config) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRecursionLimit

`func (o *Config) GetRecursionLimit() int32`

GetRecursionLimit returns the RecursionLimit field if non-nil, zero value otherwise.

### GetRecursionLimitOk

`func (o *Config) GetRecursionLimitOk() (*int32, bool)`

GetRecursionLimitOk returns a tuple with the RecursionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionLimit

`func (o *Config) SetRecursionLimit(v int32)`

SetRecursionLimit sets RecursionLimit field to given value.

### HasRecursionLimit

`func (o *Config) HasRecursionLimit() bool`

HasRecursionLimit returns a boolean if a field has been set.

### GetConfigurable

`func (o *Config) GetConfigurable() map[string]interface{}`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *Config) GetConfigurableOk() (*map[string]interface{}, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *Config) SetConfigurable(v map[string]interface{})`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *Config) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


