/*
LangGraph Platform

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package langgraphgo

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Thread type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Thread{}

// Thread struct for Thread
type Thread struct {
	// The ID of the thread.
	ThreadId string `json:"thread_id"`
	// The time the thread was created.
	CreatedAt time.Time `json:"created_at"`
	// The last time the thread was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The thread metadata.
	Metadata map[string]interface{} `json:"metadata"`
	// The status of the thread.
	Status string `json:"status"`
	// The current state of the thread.
	Values map[string]interface{} `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Thread Thread

// NewThread instantiates a new Thread object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThread(threadId string, createdAt time.Time, updatedAt time.Time, metadata map[string]interface{}, status string) *Thread {
	this := Thread{}
	this.ThreadId = threadId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Metadata = metadata
	this.Status = status
	return &this
}

// NewThreadWithDefaults instantiates a new Thread object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadWithDefaults() *Thread {
	this := Thread{}
	return &this
}

// GetThreadId returns the ThreadId field value
func (o *Thread) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *Thread) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *Thread) SetThreadId(v string) {
	o.ThreadId = v
}


// GetCreatedAt returns the CreatedAt field value
func (o *Thread) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Thread) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Thread) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}


// GetUpdatedAt returns the UpdatedAt field value
func (o *Thread) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Thread) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Thread) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}


// GetMetadata returns the Metadata field value
func (o *Thread) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Thread) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Thread) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}


// GetStatus returns the Status field value
func (o *Thread) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Thread) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Thread) SetStatus(v string) {
	o.Status = v
}


// GetValues returns the Values field value if set, zero value otherwise.
func (o *Thread) GetValues() map[string]interface{} {
	if o == nil || IsNil(o.Values) {
		var ret map[string]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetValuesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Values) {
		return map[string]interface{}{}, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Thread) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]interface{} and assigns it to the Values field.
func (o *Thread) SetValues(v map[string]interface{}) {
	o.Values = v
}

func (o Thread) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Thread) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["thread_id"] = o.ThreadId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["metadata"] = o.Metadata
	toSerialize["status"] = o.Status
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Thread) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"thread_id",
		"created_at",
		"updated_at",
		"metadata",
		"status",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{} {
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == ""{
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil{
			return err
		}
	}
	varThread := _Thread{}

	err = json.Unmarshal(data, &varThread)

	if err != nil {
		return err
	}

	*o = Thread(varThread)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "thread_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "status")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThread struct {
	value *Thread
	isSet bool
}

func (v NullableThread) Get() *Thread {
	return v.value
}

func (v *NullableThread) Set(val *Thread) {
	v.value = val
	v.isSet = true
}

func (v NullableThread) IsSet() bool {
	return v.isSet
}

func (v *NullableThread) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThread(val *Thread) *NullableThread {
	return &NullableThread{value: val, isSet: true}
}

func (v NullableThread) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThread) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


