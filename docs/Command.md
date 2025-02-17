# Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Update** | Pointer to **NullableMap[string]interface{}** | An update to the state. | [optional] 
**Resume** | Pointer to [**NullableResume**](Resume.md) |  | [optional] 
**Goto** | Pointer to [**NullableGoto**](Goto.md) |  | [optional] 

## Methods

### NewCommand

`func NewCommand() *Command`

NewCommand instantiates a new Command object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandWithDefaults

`func NewCommandWithDefaults() *Command`

NewCommandWithDefaults instantiates a new Command object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdate

`func (o *Command) GetUpdate() map[string]interface{}`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Command) GetUpdateOk() (*map[string]interface{}, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Command) SetUpdate(v map[string]interface{})`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Command) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### SetUpdateNil

`func (o *Command) SetUpdateNil(b bool)`

 SetUpdateNil sets the value for Update to be an explicit nil

### UnsetUpdate
`func (o *Command) UnsetUpdate()`

UnsetUpdate ensures that no value is present for Update, not even an explicit nil
### GetResume

`func (o *Command) GetResume() Resume`

GetResume returns the Resume field if non-nil, zero value otherwise.

### GetResumeOk

`func (o *Command) GetResumeOk() (*Resume, bool)`

GetResumeOk returns a tuple with the Resume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResume

`func (o *Command) SetResume(v Resume)`

SetResume sets Resume field to given value.

### HasResume

`func (o *Command) HasResume() bool`

HasResume returns a boolean if a field has been set.

### SetResumeNil

`func (o *Command) SetResumeNil(b bool)`

 SetResumeNil sets the value for Resume to be an explicit nil

### UnsetResume
`func (o *Command) UnsetResume()`

UnsetResume ensures that no value is present for Resume, not even an explicit nil
### GetGoto

`func (o *Command) GetGoto() Goto`

GetGoto returns the Goto field if non-nil, zero value otherwise.

### GetGotoOk

`func (o *Command) GetGotoOk() (*Goto, bool)`

GetGotoOk returns a tuple with the Goto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoto

`func (o *Command) SetGoto(v Goto)`

SetGoto sets Goto field to given value.

### HasGoto

`func (o *Command) HasGoto() bool`

HasGoto returns a boolean if a field has been set.

### SetGotoNil

`func (o *Command) SetGotoNil(b bool)`

 SetGotoNil sets the value for Goto to be an explicit nil

### UnsetGoto
`func (o *Command) UnsetGoto()`

UnsetGoto ensures that no value is present for Goto, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


