/*
LangGraph Platform

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package langgraphgo

import (
	"encoding/json"
)


// NullableMap is a helper abstraction for handling nullable map[string]interface{} types.
type NullableMap struct {
    value *map[string]interface{}
    isSet bool
}

func (v NullableMap) Get() *map[string]interface{} {
    return v.value
}

func (v *NullableMap) Set(val *map[string]interface{}) {
    v.value = val
    v.isSet = true
}

func (v NullableMap) IsSet() bool {
    return v.isSet
}

func (v *NullableMap) Unset() {
    v.value = nil
    v.isSet = false
}

func NewNullableMap(val *map[string]interface{}) *NullableMap {
    return &NullableMap{value: val, isSet: true}
}

func (v NullableMap) MarshalJSON() ([]byte, error) {
    return json.Marshal(v.value)
}

func (v *NullableMap) UnmarshalJSON(src []byte) error {
    v.isSet = true
    return json.Unmarshal(src, &v.value)
}

// checks if the Command type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Command{}

// Command The command to run.
type Command struct {
	// An update to the state.
	Update NullableMap `json:"update,omitempty"`
	Resume NullableResume `json:"resume,omitempty"`
	Goto NullableGoto `json:"goto,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Command Command

// NewCommand instantiates a new Command object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommand() *Command {
	this := Command{}
	return &this
}

// NewCommandWithDefaults instantiates a new Command object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandWithDefaults() *Command {
	this := Command{}
	return &this
}

// GetUpdate returns the Update field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Command) GetUpdate() map[string]interface{} {
	if o == nil || IsNil(o.Update.Get()) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Update.Get()
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Command) GetUpdateOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Update.Get(), o.Update.IsSet()
}

// HasUpdate returns a boolean if a field has been set.
func (o *Command) HasUpdate() bool {
	if o != nil && o.Update.IsSet() {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given NullableMap[string]interface{} and assigns it to the Update field.
func (o *Command) SetUpdate(v map[string]interface{}) {
	o.Update.Set(&v)
}
// SetUpdateNil sets the value for Update to be an explicit nil
func (o *Command) SetUpdateNil() {
	o.Update.Set(nil)
}

// UnsetUpdate ensures that no value is present for Update, not even an explicit nil
func (o *Command) UnsetUpdate() {
	o.Update.Unset()
}

// GetResume returns the Resume field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Command) GetResume() Resume {
	if o == nil || IsNil(o.Resume.Get()) {
		var ret Resume
		return ret
	}
	return *o.Resume.Get()
}

// GetResumeOk returns a tuple with the Resume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Command) GetResumeOk() (*Resume, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resume.Get(), o.Resume.IsSet()
}

// HasResume returns a boolean if a field has been set.
func (o *Command) HasResume() bool {
	if o != nil && o.Resume.IsSet() {
		return true
	}

	return false
}

// SetResume gets a reference to the given NullableResume and assigns it to the Resume field.
func (o *Command) SetResume(v Resume) {
	o.Resume.Set(&v)
}
// SetResumeNil sets the value for Resume to be an explicit nil
func (o *Command) SetResumeNil() {
	o.Resume.Set(nil)
}

// UnsetResume ensures that no value is present for Resume, not even an explicit nil
func (o *Command) UnsetResume() {
	o.Resume.Unset()
}

// GetGoto returns the Goto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Command) GetGoto() Goto {
	if o == nil || IsNil(o.Goto.Get()) {
		var ret Goto
		return ret
	}
	return *o.Goto.Get()
}

// GetGotoOk returns a tuple with the Goto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Command) GetGotoOk() (*Goto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Goto.Get(), o.Goto.IsSet()
}

// HasGoto returns a boolean if a field has been set.
func (o *Command) HasGoto() bool {
	if o != nil && o.Goto.IsSet() {
		return true
	}

	return false
}

// SetGoto gets a reference to the given NullableGoto and assigns it to the Goto field.
func (o *Command) SetGoto(v Goto) {
	o.Goto.Set(&v)
}
// SetGotoNil sets the value for Goto to be an explicit nil
func (o *Command) SetGotoNil() {
	o.Goto.Set(nil)
}

// UnsetGoto ensures that no value is present for Goto, not even an explicit nil
func (o *Command) UnsetGoto() {
	o.Goto.Unset()
}

func (o Command) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Command) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Update.IsSet() {
		toSerialize["update"] = o.Update.Get()
	}
	if o.Resume.IsSet() {
		toSerialize["resume"] = o.Resume.Get()
	}
	if o.Goto.IsSet() {
		toSerialize["goto"] = o.Goto.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Command) UnmarshalJSON(data []byte) (err error) {
	varCommand := _Command{}

	err = json.Unmarshal(data, &varCommand)

	if err != nil {
		return err
	}

	*o = Command(varCommand)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "update")
		delete(additionalProperties, "resume")
		delete(additionalProperties, "goto")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommand struct {
	value *Command
	isSet bool
}

func (v NullableCommand) Get() *Command {
	return v.value
}

func (v *NullableCommand) Set(val *Command) {
	v.value = val
	v.isSet = true
}

func (v NullableCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommand(val *Command) *NullableCommand {
	return &NullableCommand{value: val, isSet: true}
}

func (v NullableCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


