/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// EventPriority The priority of the event. For example, `normal` or `low`.
type EventPriority string

// List of EventPriority
const (
	EVENTPRIORITY_NORMAL EventPriority = "normal"
	EVENTPRIORITY_LOW    EventPriority = "low"
)

func (v *EventPriority) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventPriority(value)
	for _, existing := range []EventPriority{"normal", "low"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventPriority", value)
}

// Ptr returns reference to EventPriority value
func (v EventPriority) Ptr() *EventPriority {
	return &v
}

type NullableEventPriority struct {
	value *EventPriority
	isSet bool
}

func (v NullableEventPriority) Get() *EventPriority {
	return v.value
}

func (v *NullableEventPriority) Set(val *EventPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableEventPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableEventPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventPriority(val *EventPriority) *NullableEventPriority {
	return &NullableEventPriority{value: val, isSet: true}
}

func (v NullableEventPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
