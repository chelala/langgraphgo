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

// checks if the AssistantVersionsSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssistantVersionsSearchRequest{}

// AssistantVersionsSearchRequest Payload for listing assistant versions.
type AssistantVersionsSearchRequest struct {
	// Metadata to filter versions by. Exact match filter for each KV pair.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The maximum number of versions to return.
	Limit *int32 `json:"limit,omitempty"`
	// The number of versions to skip.
	Offset *int32 `json:"offset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssistantVersionsSearchRequest AssistantVersionsSearchRequest

// NewAssistantVersionsSearchRequest instantiates a new AssistantVersionsSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssistantVersionsSearchRequest() *AssistantVersionsSearchRequest {
	this := AssistantVersionsSearchRequest{}
	var limit int32 = 10
	this.Limit = &limit
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// NewAssistantVersionsSearchRequestWithDefaults instantiates a new AssistantVersionsSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssistantVersionsSearchRequestWithDefaults() *AssistantVersionsSearchRequest {
	this := AssistantVersionsSearchRequest{}
	var limit int32 = 10
	this.Limit = &limit
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AssistantVersionsSearchRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssistantVersionsSearchRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AssistantVersionsSearchRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AssistantVersionsSearchRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *AssistantVersionsSearchRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssistantVersionsSearchRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *AssistantVersionsSearchRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *AssistantVersionsSearchRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *AssistantVersionsSearchRequest) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssistantVersionsSearchRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *AssistantVersionsSearchRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *AssistantVersionsSearchRequest) SetOffset(v int32) {
	o.Offset = &v
}

func (o AssistantVersionsSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssistantVersionsSearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssistantVersionsSearchRequest) UnmarshalJSON(data []byte) (err error) {
	varAssistantVersionsSearchRequest := _AssistantVersionsSearchRequest{}

	err = json.Unmarshal(data, &varAssistantVersionsSearchRequest)

	if err != nil {
		return err
	}

	*o = AssistantVersionsSearchRequest(varAssistantVersionsSearchRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "offset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssistantVersionsSearchRequest struct {
	value *AssistantVersionsSearchRequest
	isSet bool
}

func (v NullableAssistantVersionsSearchRequest) Get() *AssistantVersionsSearchRequest {
	return v.value
}

func (v *NullableAssistantVersionsSearchRequest) Set(val *AssistantVersionsSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssistantVersionsSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssistantVersionsSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssistantVersionsSearchRequest(val *AssistantVersionsSearchRequest) *NullableAssistantVersionsSearchRequest {
	return &NullableAssistantVersionsSearchRequest{value: val, isSet: true}
}

func (v NullableAssistantVersionsSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssistantVersionsSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


