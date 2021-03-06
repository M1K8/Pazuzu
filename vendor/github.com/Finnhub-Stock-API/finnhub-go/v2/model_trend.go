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

// Trend struct for Trend
type Trend struct {
	// ADX reading
	Adx *float32 `json:"adx,omitempty"`
	// Whether market is trending or going sideway
	Trending *bool `json:"trending,omitempty"`
}

// NewTrend instantiates a new Trend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrend() *Trend {
	this := Trend{}
	return &this
}

// NewTrendWithDefaults instantiates a new Trend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrendWithDefaults() *Trend {
	this := Trend{}
	return &this
}

// GetAdx returns the Adx field value if set, zero value otherwise.
func (o *Trend) GetAdx() float32 {
	if o == nil || o.Adx == nil {
		var ret float32
		return ret
	}
	return *o.Adx
}

// GetAdxOk returns a tuple with the Adx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trend) GetAdxOk() (*float32, bool) {
	if o == nil || o.Adx == nil {
		return nil, false
	}
	return o.Adx, true
}

// HasAdx returns a boolean if a field has been set.
func (o *Trend) HasAdx() bool {
	if o != nil && o.Adx != nil {
		return true
	}

	return false
}

// SetAdx gets a reference to the given float32 and assigns it to the Adx field.
func (o *Trend) SetAdx(v float32) {
	o.Adx = &v
}

// GetTrending returns the Trending field value if set, zero value otherwise.
func (o *Trend) GetTrending() bool {
	if o == nil || o.Trending == nil {
		var ret bool
		return ret
	}
	return *o.Trending
}

// GetTrendingOk returns a tuple with the Trending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trend) GetTrendingOk() (*bool, bool) {
	if o == nil || o.Trending == nil {
		return nil, false
	}
	return o.Trending, true
}

// HasTrending returns a boolean if a field has been set.
func (o *Trend) HasTrending() bool {
	if o != nil && o.Trending != nil {
		return true
	}

	return false
}

// SetTrending gets a reference to the given bool and assigns it to the Trending field.
func (o *Trend) SetTrending(v bool) {
	o.Trending = &v
}

func (o Trend) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Adx != nil {
		toSerialize["adx"] = o.Adx
	}
	if o.Trending != nil {
		toSerialize["trending"] = o.Trending
	}
	return json.Marshal(toSerialize)
}

type NullableTrend struct {
	value *Trend
	isSet bool
}

func (v NullableTrend) Get() *Trend {
	return v.value
}

func (v *NullableTrend) Set(val *Trend) {
	v.value = val
	v.isSet = true
}

func (v NullableTrend) IsSet() bool {
	return v.isSet
}

func (v *NullableTrend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrend(val *Trend) *NullableTrend {
	return &NullableTrend{value: val, isSet: true}
}

func (v NullableTrend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


