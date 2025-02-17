# GraphSchemaNoId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputSchema** | **map[string]interface{}** | The schema for the graph input. Missing if unable to generate JSON schema from graph. | 
**OutputSchema** | **map[string]interface{}** | The schema for the graph output. Missing if unable to generate JSON schema from graph. | 
**StateSchema** | **map[string]interface{}** | The schema for the graph state. Missing if unable to generate JSON schema from graph. | 
**ConfigSchema** | **map[string]interface{}** | The schema for the graph config. Missing if unable to generate JSON schema from graph. | 

## Methods

### NewGraphSchemaNoId

`func NewGraphSchemaNoId(inputSchema map[string]interface{}, outputSchema map[string]interface{}, stateSchema map[string]interface{}, configSchema map[string]interface{}, ) *GraphSchemaNoId`

NewGraphSchemaNoId instantiates a new GraphSchemaNoId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphSchemaNoIdWithDefaults

`func NewGraphSchemaNoIdWithDefaults() *GraphSchemaNoId`

NewGraphSchemaNoIdWithDefaults instantiates a new GraphSchemaNoId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputSchema

`func (o *GraphSchemaNoId) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *GraphSchemaNoId) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *GraphSchemaNoId) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.


### GetOutputSchema

`func (o *GraphSchemaNoId) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *GraphSchemaNoId) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *GraphSchemaNoId) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.


### GetStateSchema

`func (o *GraphSchemaNoId) GetStateSchema() map[string]interface{}`

GetStateSchema returns the StateSchema field if non-nil, zero value otherwise.

### GetStateSchemaOk

`func (o *GraphSchemaNoId) GetStateSchemaOk() (*map[string]interface{}, bool)`

GetStateSchemaOk returns a tuple with the StateSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateSchema

`func (o *GraphSchemaNoId) SetStateSchema(v map[string]interface{})`

SetStateSchema sets StateSchema field to given value.


### GetConfigSchema

`func (o *GraphSchemaNoId) GetConfigSchema() map[string]interface{}`

GetConfigSchema returns the ConfigSchema field if non-nil, zero value otherwise.

### GetConfigSchemaOk

`func (o *GraphSchemaNoId) GetConfigSchemaOk() (*map[string]interface{}, bool)`

GetConfigSchemaOk returns a tuple with the ConfigSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSchema

`func (o *GraphSchemaNoId) SetConfigSchema(v map[string]interface{})`

SetConfigSchema sets ConfigSchema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


