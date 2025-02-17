# GraphSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraphId** | **string** | The ID of the graph. | 
**InputSchema** | Pointer to **map[string]interface{}** | The schema for the graph input. Missing if unable to generate JSON schema from graph. | [optional] 
**OutputSchema** | Pointer to **map[string]interface{}** | The schema for the graph output. Missing if unable to generate JSON schema from graph. | [optional] 
**StateSchema** | **map[string]interface{}** | The schema for the graph state. Missing if unable to generate JSON schema from graph. | 
**ConfigSchema** | **map[string]interface{}** | The schema for the graph config. Missing if unable to generate JSON schema from graph. | 

## Methods

### NewGraphSchema

`func NewGraphSchema(graphId string, stateSchema map[string]interface{}, configSchema map[string]interface{}, ) *GraphSchema`

NewGraphSchema instantiates a new GraphSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphSchemaWithDefaults

`func NewGraphSchemaWithDefaults() *GraphSchema`

NewGraphSchemaWithDefaults instantiates a new GraphSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraphId

`func (o *GraphSchema) GetGraphId() string`

GetGraphId returns the GraphId field if non-nil, zero value otherwise.

### GetGraphIdOk

`func (o *GraphSchema) GetGraphIdOk() (*string, bool)`

GetGraphIdOk returns a tuple with the GraphId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphId

`func (o *GraphSchema) SetGraphId(v string)`

SetGraphId sets GraphId field to given value.


### GetInputSchema

`func (o *GraphSchema) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *GraphSchema) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *GraphSchema) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *GraphSchema) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *GraphSchema) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *GraphSchema) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *GraphSchema) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *GraphSchema) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetStateSchema

`func (o *GraphSchema) GetStateSchema() map[string]interface{}`

GetStateSchema returns the StateSchema field if non-nil, zero value otherwise.

### GetStateSchemaOk

`func (o *GraphSchema) GetStateSchemaOk() (*map[string]interface{}, bool)`

GetStateSchemaOk returns a tuple with the StateSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateSchema

`func (o *GraphSchema) SetStateSchema(v map[string]interface{})`

SetStateSchema sets StateSchema field to given value.


### GetConfigSchema

`func (o *GraphSchema) GetConfigSchema() map[string]interface{}`

GetConfigSchema returns the ConfigSchema field if non-nil, zero value otherwise.

### GetConfigSchemaOk

`func (o *GraphSchema) GetConfigSchemaOk() (*map[string]interface{}, bool)`

GetConfigSchemaOk returns a tuple with the ConfigSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSchema

`func (o *GraphSchema) SetConfigSchema(v map[string]interface{})`

SetConfigSchema sets ConfigSchema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


