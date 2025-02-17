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

// checks if the Send type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Send{}

// Send A message to send to a node.
type Send struct {
	// The node to send the message to.
	Node string `json:"node"`
	// The message to send.
	Input map[string]interface{} `json:"input"`
	AdditionalProperties map[string]interface{}
}

type _Send Send

// NewSend instantiates a new Send object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSend(node string, input map[string]interface{}) *Send {
	this := Send{}
	this.Node = node
	this.Input = input
	return &this
}

// NewSendWithDefaults instantiates a new Send object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendWithDefaults() *Send {
	this := Send{}
	return &this
}

// GetNode returns the Node field value
func (o *Send) GetNode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Node
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
func (o *Send) GetNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Node, true
}

// SetNode sets field value
func (o *Send) SetNode(v string) {
	o.Node = v
}


// GetInput returns the Input field value
func (o *Send) GetInput() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *Send) GetInputOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// SetInput sets field value
func (o *Send) SetInput(v map[string]interface{}) {
	o.Input = v
}


func (o Send) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Send) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["node"] = o.Node
	toSerialize["input"] = o.Input

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Send) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node",
		"input",
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
	varSend := _Send{}

	err = json.Unmarshal(data, &varSend)

	if err != nil {
		return err
	}

	*o = Send(varSend)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "node")
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSend struct {
	value *Send
	isSet bool
}

func (v NullableSend) Get() *Send {
	return v.value
}

func (v *NullableSend) Set(val *Send) {
	v.value = val
	v.isSet = true
}

func (v NullableSend) IsSet() bool {
	return v.isSet
}

func (v *NullableSend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSend(val *Send) *NullableSend {
	return &NullableSend{value: val, isSet: true}
}

func (v NullableSend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


