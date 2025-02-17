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

// checks if the Config1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Config1{}

// Config1 The assistant config.
type Config1 struct {
	Tags []string `json:"tags,omitempty"`
	RecursionLimit *int32 `json:"recursion_limit,omitempty"`
	Configurable map[string]interface{} `json:"configurable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Config1 Config1

// NewConfig1 instantiates a new Config1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfig1() *Config1 {
	this := Config1{}
	return &this
}

// NewConfig1WithDefaults instantiates a new Config1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfig1WithDefaults() *Config1 {
	this := Config1{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Config1) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Config1) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Config1) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Config1) SetTags(v []string) {
	o.Tags = v
}

// GetRecursionLimit returns the RecursionLimit field value if set, zero value otherwise.
func (o *Config1) GetRecursionLimit() int32 {
	if o == nil || IsNil(o.RecursionLimit) {
		var ret int32
		return ret
	}
	return *o.RecursionLimit
}

// GetRecursionLimitOk returns a tuple with the RecursionLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Config1) GetRecursionLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.RecursionLimit) {
		return nil, false
	}
	return o.RecursionLimit, true
}

// HasRecursionLimit returns a boolean if a field has been set.
func (o *Config1) HasRecursionLimit() bool {
	if o != nil && !IsNil(o.RecursionLimit) {
		return true
	}

	return false
}

// SetRecursionLimit gets a reference to the given int32 and assigns it to the RecursionLimit field.
func (o *Config1) SetRecursionLimit(v int32) {
	o.RecursionLimit = &v
}

// GetConfigurable returns the Configurable field value if set, zero value otherwise.
func (o *Config1) GetConfigurable() map[string]interface{} {
	if o == nil || IsNil(o.Configurable) {
		var ret map[string]interface{}
		return ret
	}
	return o.Configurable
}

// GetConfigurableOk returns a tuple with the Configurable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Config1) GetConfigurableOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Configurable) {
		return map[string]interface{}{}, false
	}
	return o.Configurable, true
}

// HasConfigurable returns a boolean if a field has been set.
func (o *Config1) HasConfigurable() bool {
	if o != nil && !IsNil(o.Configurable) {
		return true
	}

	return false
}

// SetConfigurable gets a reference to the given map[string]interface{} and assigns it to the Configurable field.
func (o *Config1) SetConfigurable(v map[string]interface{}) {
	o.Configurable = v
}

func (o Config1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Config1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.RecursionLimit) {
		toSerialize["recursion_limit"] = o.RecursionLimit
	}
	if !IsNil(o.Configurable) {
		toSerialize["configurable"] = o.Configurable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Config1) UnmarshalJSON(data []byte) (err error) {
	varConfig1 := _Config1{}

	err = json.Unmarshal(data, &varConfig1)

	if err != nil {
		return err
	}

	*o = Config1(varConfig1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tags")
		delete(additionalProperties, "recursion_limit")
		delete(additionalProperties, "configurable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfig1 struct {
	value *Config1
	isSet bool
}

func (v NullableConfig1) Get() *Config1 {
	return v.value
}

func (v *NullableConfig1) Set(val *Config1) {
	v.value = val
	v.isSet = true
}

func (v NullableConfig1) IsSet() bool {
	return v.isSet
}

func (v *NullableConfig1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfig1(val *Config1) *NullableConfig1 {
	return &NullableConfig1{value: val, isSet: true}
}

func (v NullableConfig1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfig1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


