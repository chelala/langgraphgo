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

// checks if the ThreadStateUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreadStateUpdateResponse{}

// ThreadStateUpdateResponse Response for adding state to a thread.
type ThreadStateUpdateResponse struct {
	Checkpoint map[string]interface{} `json:"checkpoint,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThreadStateUpdateResponse ThreadStateUpdateResponse

// NewThreadStateUpdateResponse instantiates a new ThreadStateUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadStateUpdateResponse() *ThreadStateUpdateResponse {
	this := ThreadStateUpdateResponse{}
	return &this
}

// NewThreadStateUpdateResponseWithDefaults instantiates a new ThreadStateUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadStateUpdateResponseWithDefaults() *ThreadStateUpdateResponse {
	this := ThreadStateUpdateResponse{}
	return &this
}

// GetCheckpoint returns the Checkpoint field value if set, zero value otherwise.
func (o *ThreadStateUpdateResponse) GetCheckpoint() map[string]interface{} {
	if o == nil || IsNil(o.Checkpoint) {
		var ret map[string]interface{}
		return ret
	}
	return o.Checkpoint
}

// GetCheckpointOk returns a tuple with the Checkpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadStateUpdateResponse) GetCheckpointOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Checkpoint) {
		return map[string]interface{}{}, false
	}
	return o.Checkpoint, true
}

// HasCheckpoint returns a boolean if a field has been set.
func (o *ThreadStateUpdateResponse) HasCheckpoint() bool {
	if o != nil && !IsNil(o.Checkpoint) {
		return true
	}

	return false
}

// SetCheckpoint gets a reference to the given map[string]interface{} and assigns it to the Checkpoint field.
func (o *ThreadStateUpdateResponse) SetCheckpoint(v map[string]interface{}) {
	o.Checkpoint = v
}

func (o ThreadStateUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreadStateUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Checkpoint) {
		toSerialize["checkpoint"] = o.Checkpoint
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThreadStateUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	varThreadStateUpdateResponse := _ThreadStateUpdateResponse{}

	err = json.Unmarshal(data, &varThreadStateUpdateResponse)

	if err != nil {
		return err
	}

	*o = ThreadStateUpdateResponse(varThreadStateUpdateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "checkpoint")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThreadStateUpdateResponse struct {
	value *ThreadStateUpdateResponse
	isSet bool
}

func (v NullableThreadStateUpdateResponse) Get() *ThreadStateUpdateResponse {
	return v.value
}

func (v *NullableThreadStateUpdateResponse) Set(val *ThreadStateUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadStateUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadStateUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadStateUpdateResponse(val *ThreadStateUpdateResponse) *NullableThreadStateUpdateResponse {
	return &NullableThreadStateUpdateResponse{value: val, isSet: true}
}

func (v NullableThreadStateUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadStateUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


