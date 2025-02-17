/*
LangGraph Platform

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package langgraphgo

import (
	"encoding/json"
	"fmt"
)


// CronCreateAssistantId The assistant ID or graph name to run. If using graph name, will default to the assistant automatically created from that graph by the server.
type CronCreateAssistantId struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CronCreateAssistantId) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CronCreateAssistantId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CronCreateAssistantId) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableCronCreateAssistantId struct {
	value *CronCreateAssistantId
	isSet bool
}

func (v NullableCronCreateAssistantId) Get() *CronCreateAssistantId {
	return v.value
}

func (v *NullableCronCreateAssistantId) Set(val *CronCreateAssistantId) {
	v.value = val
	v.isSet = true
}

func (v NullableCronCreateAssistantId) IsSet() bool {
	return v.isSet
}

func (v *NullableCronCreateAssistantId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCronCreateAssistantId(val *CronCreateAssistantId) *NullableCronCreateAssistantId {
	return &NullableCronCreateAssistantId{value: val, isSet: true}
}

func (v NullableCronCreateAssistantId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCronCreateAssistantId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


