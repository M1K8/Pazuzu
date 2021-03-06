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

// Split struct for Split
type Split struct {
	// Symbol.
	Symbol *string `json:"symbol,omitempty"`
	// Split date.
	Date *string `json:"date,omitempty"`
	// From factor.
	FromFactor *float32 `json:"fromFactor,omitempty"`
	// To factor.
	ToFactor *float32 `json:"toFactor,omitempty"`
}

// NewSplit instantiates a new Split object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplit() *Split {
	this := Split{}
	return &this
}

// NewSplitWithDefaults instantiates a new Split object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitWithDefaults() *Split {
	this := Split{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Split) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Split) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Split) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Split) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Split) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Split) SetDate(v string) {
	o.Date = &v
}

// GetFromFactor returns the FromFactor field value if set, zero value otherwise.
func (o *Split) GetFromFactor() float32 {
	if o == nil || o.FromFactor == nil {
		var ret float32
		return ret
	}
	return *o.FromFactor
}

// GetFromFactorOk returns a tuple with the FromFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetFromFactorOk() (*float32, bool) {
	if o == nil || o.FromFactor == nil {
		return nil, false
	}
	return o.FromFactor, true
}

// HasFromFactor returns a boolean if a field has been set.
func (o *Split) HasFromFactor() bool {
	if o != nil && o.FromFactor != nil {
		return true
	}

	return false
}

// SetFromFactor gets a reference to the given float32 and assigns it to the FromFactor field.
func (o *Split) SetFromFactor(v float32) {
	o.FromFactor = &v
}

// GetToFactor returns the ToFactor field value if set, zero value otherwise.
func (o *Split) GetToFactor() float32 {
	if o == nil || o.ToFactor == nil {
		var ret float32
		return ret
	}
	return *o.ToFactor
}

// GetToFactorOk returns a tuple with the ToFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetToFactorOk() (*float32, bool) {
	if o == nil || o.ToFactor == nil {
		return nil, false
	}
	return o.ToFactor, true
}

// HasToFactor returns a boolean if a field has been set.
func (o *Split) HasToFactor() bool {
	if o != nil && o.ToFactor != nil {
		return true
	}

	return false
}

// SetToFactor gets a reference to the given float32 and assigns it to the ToFactor field.
func (o *Split) SetToFactor(v float32) {
	o.ToFactor = &v
}

func (o Split) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.FromFactor != nil {
		toSerialize["fromFactor"] = o.FromFactor
	}
	if o.ToFactor != nil {
		toSerialize["toFactor"] = o.ToFactor
	}
	return json.Marshal(toSerialize)
}

type NullableSplit struct {
	value *Split
	isSet bool
}

func (v NullableSplit) Get() *Split {
	return v.value
}

func (v *NullableSplit) Set(val *Split) {
	v.value = val
	v.isSet = true
}

func (v NullableSplit) IsSet() bool {
	return v.isSet
}

func (v *NullableSplit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplit(val *Split) *NullableSplit {
	return &NullableSplit{value: val, isSet: true}
}

func (v NullableSplit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


