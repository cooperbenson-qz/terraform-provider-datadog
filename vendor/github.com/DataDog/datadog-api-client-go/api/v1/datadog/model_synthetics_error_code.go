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

// SyntheticsErrorCode Error code that can be returned by a Synthetic test.
type SyntheticsErrorCode string

// List of SyntheticsErrorCode
const (
	SYNTHETICSERRORCODE_NO_ERROR            SyntheticsErrorCode = "NO_ERROR"
	SYNTHETICSERRORCODE_UNKNOWN             SyntheticsErrorCode = "UNKNOWN"
	SYNTHETICSERRORCODE_DNS                 SyntheticsErrorCode = "DNS"
	SYNTHETICSERRORCODE_SSL                 SyntheticsErrorCode = "SSL"
	SYNTHETICSERRORCODE_TIMEOUT             SyntheticsErrorCode = "TIMEOUT"
	SYNTHETICSERRORCODE_DENIED              SyntheticsErrorCode = "DENIED"
	SYNTHETICSERRORCODE_INCORRECT_ASSERTION SyntheticsErrorCode = "INCORRECT_ASSERTION"
)

func (v *SyntheticsErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SyntheticsErrorCode(value)
	for _, existing := range []SyntheticsErrorCode{"NO_ERROR", "UNKNOWN", "DNS", "SSL", "TIMEOUT", "DENIED", "INCORRECT_ASSERTION"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SyntheticsErrorCode", value)
}

// Ptr returns reference to SyntheticsErrorCode value
func (v SyntheticsErrorCode) Ptr() *SyntheticsErrorCode {
	return &v
}

type NullableSyntheticsErrorCode struct {
	value *SyntheticsErrorCode
	isSet bool
}

func (v NullableSyntheticsErrorCode) Get() *SyntheticsErrorCode {
	return v.value
}

func (v *NullableSyntheticsErrorCode) Set(val *SyntheticsErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsErrorCode(val *SyntheticsErrorCode) *NullableSyntheticsErrorCode {
	return &NullableSyntheticsErrorCode{value: val, isSet: true}
}

func (v NullableSyntheticsErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
