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


// GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter struct for GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter
type GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableGetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter struct {
	value *GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter
	isSet bool
}

func (v NullableGetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter) Get() *GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter {
	return v.value
}

func (v *NullableGetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter) Set(val *GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter(val *GetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter) *NullableGetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter {
	return &NullableGetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter{value: val, isSet: true}
}

func (v NullableGetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssistantGraphAssistantsAssistantIdGraphGetAssistantIdParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


