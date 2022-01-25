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

// Indicator struct for Indicator
type Indicator struct {
	// Number of buy signals
	Buy *int64 `json:"buy,omitempty"`
	// Number of neutral signals
	Neutral *int64 `json:"neutral,omitempty"`
	// Number of sell signals
	Sell *int64 `json:"sell,omitempty"`
}

// NewIndicator instantiates a new Indicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicator() *Indicator {
	this := Indicator{}
	return &this
}

// NewIndicatorWithDefaults instantiates a new Indicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicatorWithDefaults() *Indicator {
	this := Indicator{}
	return &this
}

// GetBuy returns the Buy field value if set, zero value otherwise.
func (o *Indicator) GetBuy() int64 {
	if o == nil || o.Buy == nil {
		var ret int64
		return ret
	}
	return *o.Buy
}

// GetBuyOk returns a tuple with the Buy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Indicator) GetBuyOk() (*int64, bool) {
	if o == nil || o.Buy == nil {
		return nil, false
	}
	return o.Buy, true
}

// HasBuy returns a boolean if a field has been set.
func (o *Indicator) HasBuy() bool {
	if o != nil && o.Buy != nil {
		return true
	}

	return false
}

// SetBuy gets a reference to the given int64 and assigns it to the Buy field.
func (o *Indicator) SetBuy(v int64) {
	o.Buy = &v
}

// GetNeutral returns the Neutral field value if set, zero value otherwise.
func (o *Indicator) GetNeutral() int64 {
	if o == nil || o.Neutral == nil {
		var ret int64
		return ret
	}
	return *o.Neutral
}

// GetNeutralOk returns a tuple with the Neutral field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Indicator) GetNeutralOk() (*int64, bool) {
	if o == nil || o.Neutral == nil {
		return nil, false
	}
	return o.Neutral, true
}

// HasNeutral returns a boolean if a field has been set.
func (o *Indicator) HasNeutral() bool {
	if o != nil && o.Neutral != nil {
		return true
	}

	return false
}

// SetNeutral gets a reference to the given int64 and assigns it to the Neutral field.
func (o *Indicator) SetNeutral(v int64) {
	o.Neutral = &v
}

// GetSell returns the Sell field value if set, zero value otherwise.
func (o *Indicator) GetSell() int64 {
	if o == nil || o.Sell == nil {
		var ret int64
		return ret
	}
	return *o.Sell
}

// GetSellOk returns a tuple with the Sell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Indicator) GetSellOk() (*int64, bool) {
	if o == nil || o.Sell == nil {
		return nil, false
	}
	return o.Sell, true
}

// HasSell returns a boolean if a field has been set.
func (o *Indicator) HasSell() bool {
	if o != nil && o.Sell != nil {
		return true
	}

	return false
}

// SetSell gets a reference to the given int64 and assigns it to the Sell field.
func (o *Indicator) SetSell(v int64) {
	o.Sell = &v
}

func (o Indicator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Buy != nil {
		toSerialize["buy"] = o.Buy
	}
	if o.Neutral != nil {
		toSerialize["neutral"] = o.Neutral
	}
	if o.Sell != nil {
		toSerialize["sell"] = o.Sell
	}
	return json.Marshal(toSerialize)
}

type NullableIndicator struct {
	value *Indicator
	isSet bool
}

func (v NullableIndicator) Get() *Indicator {
	return v.value
}

func (v *NullableIndicator) Set(val *Indicator) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicator(val *Indicator) *NullableIndicator {
	return &NullableIndicator{value: val, isSet: true}
}

func (v NullableIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


