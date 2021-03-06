/*
Finnhub API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finnhub

import (
	"encoding/json"
)

// PatternRecognition struct for PatternRecognition
type PatternRecognition struct {
	// Array of patterns.
	Points *[]map[string]interface{} `json:"points,omitempty"`
}

// NewPatternRecognition instantiates a new PatternRecognition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatternRecognition() *PatternRecognition {
	this := PatternRecognition{}
	return &this
}

// NewPatternRecognitionWithDefaults instantiates a new PatternRecognition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatternRecognitionWithDefaults() *PatternRecognition {
	this := PatternRecognition{}
	return &this
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *PatternRecognition) GetPoints() []map[string]interface{} {
	if o == nil || o.Points == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatternRecognition) GetPointsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *PatternRecognition) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []map[string]interface{} and assigns it to the Points field.
func (o *PatternRecognition) SetPoints(v []map[string]interface{}) {
	o.Points = &v
}

func (o PatternRecognition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	return json.Marshal(toSerialize)
}

type NullablePatternRecognition struct {
	value *PatternRecognition
	isSet bool
}

func (v NullablePatternRecognition) Get() *PatternRecognition {
	return v.value
}

func (v *NullablePatternRecognition) Set(val *PatternRecognition) {
	v.value = val
	v.isSet = true
}

func (v NullablePatternRecognition) IsSet() bool {
	return v.isSet
}

func (v *NullablePatternRecognition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatternRecognition(val *PatternRecognition) *NullablePatternRecognition {
	return &NullablePatternRecognition{value: val, isSet: true}
}

func (v NullablePatternRecognition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatternRecognition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


