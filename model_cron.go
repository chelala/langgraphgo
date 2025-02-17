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

// checks if the Cron type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cron{}

// Cron Represents a scheduled task.
type Cron struct {
	// The ID of the cron.
	CronId string `json:"cron_id"`
	// The ID of the thread.
	ThreadId string `json:"thread_id"`
	// The end date to stop running the cron.
	EndTime time.Time `json:"end_time"`
	// The schedule to run, cron format.
	Schedule string `json:"schedule"`
	// The time the cron was created.
	CreatedAt time.Time `json:"created_at"`
	// The last time the cron was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The run payload to use for creating new run.
	Payload map[string]interface{} `json:"payload"`
	AdditionalProperties map[string]interface{}
}

type _Cron Cron

// NewCron instantiates a new Cron object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCron(cronId string, threadId string, endTime time.Time, schedule string, createdAt time.Time, updatedAt time.Time, payload map[string]interface{}) *Cron {
	this := Cron{}
	this.CronId = cronId
	this.ThreadId = threadId
	this.EndTime = endTime
	this.Schedule = schedule
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Payload = payload
	return &this
}

// NewCronWithDefaults instantiates a new Cron object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronWithDefaults() *Cron {
	this := Cron{}
	return &this
}

// GetCronId returns the CronId field value
func (o *Cron) GetCronId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CronId
}

// GetCronIdOk returns a tuple with the CronId field value
// and a boolean to check if the value has been set.
func (o *Cron) GetCronIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CronId, true
}

// SetCronId sets field value
func (o *Cron) SetCronId(v string) {
	o.CronId = v
}


// GetThreadId returns the ThreadId field value
func (o *Cron) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *Cron) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *Cron) SetThreadId(v string) {
	o.ThreadId = v
}


// GetEndTime returns the EndTime field value
func (o *Cron) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *Cron) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *Cron) SetEndTime(v time.Time) {
	o.EndTime = v
}


// GetSchedule returns the Schedule field value
func (o *Cron) GetSchedule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *Cron) GetScheduleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *Cron) SetSchedule(v string) {
	o.Schedule = v
}


// GetCreatedAt returns the CreatedAt field value
func (o *Cron) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Cron) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Cron) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}


// GetUpdatedAt returns the UpdatedAt field value
func (o *Cron) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Cron) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Cron) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}


// GetPayload returns the Payload field value
func (o *Cron) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *Cron) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *Cron) SetPayload(v map[string]interface{}) {
	o.Payload = v
}


func (o Cron) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cron) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cron_id"] = o.CronId
	toSerialize["thread_id"] = o.ThreadId
	toSerialize["end_time"] = o.EndTime
	toSerialize["schedule"] = o.Schedule
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["payload"] = o.Payload

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Cron) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cron_id",
		"thread_id",
		"end_time",
		"schedule",
		"created_at",
		"updated_at",
		"payload",
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
	varCron := _Cron{}

	err = json.Unmarshal(data, &varCron)

	if err != nil {
		return err
	}

	*o = Cron(varCron)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cron_id")
		delete(additionalProperties, "thread_id")
		delete(additionalProperties, "end_time")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCron struct {
	value *Cron
	isSet bool
}

func (v NullableCron) Get() *Cron {
	return v.value
}

func (v *NullableCron) Set(val *Cron) {
	v.value = val
	v.isSet = true
}

func (v NullableCron) IsSet() bool {
	return v.isSet
}

func (v *NullableCron) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCron(val *Cron) *NullableCron {
	return &NullableCron{value: val, isSet: true}
}

func (v NullableCron) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCron) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


