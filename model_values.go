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


// Values struct for Values
type Values struct {
	ArrayOfMapmapOfStringAny *[]map[string]interface{}
	MapmapOfStringAny *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Values) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ArrayOfMapmapOfStringAny
	err = json.Unmarshal(data, &dst.ArrayOfMapmapOfStringAny);
	if err == nil {
		jsonArrayOfMapmapOfStringAny, _ := json.Marshal(dst.ArrayOfMapmapOfStringAny)
		if string(jsonArrayOfMapmapOfStringAny) == "{}" { // empty struct
			dst.ArrayOfMapmapOfStringAny = nil
		} else {
			return nil // data stored in dst.ArrayOfMapmapOfStringAny, return on the first match
		}
	} else {
		dst.ArrayOfMapmapOfStringAny = nil
	}

	// try to unmarshal JSON data into MapmapOfStringAny
	err = json.Unmarshal(data, &dst.MapmapOfStringAny);
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			return nil // data stored in dst.MapmapOfStringAny, return on the first match
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Values)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Values) MarshalJSON() ([]byte, error) {
	if src.ArrayOfMapmapOfStringAny != nil {
		return json.Marshal(&src.ArrayOfMapmapOfStringAny)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableValues struct {
	value *Values
	isSet bool
}

func (v NullableValues) Get() *Values {
	return v.value
}

func (v *NullableValues) Set(val *Values) {
	v.value = val
	v.isSet = true
}

func (v NullableValues) IsSet() bool {
	return v.isSet
}

func (v *NullableValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValues(val *Values) *NullableValues {
	return &NullableValues{value: val, isSet: true}
}

func (v NullableValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


