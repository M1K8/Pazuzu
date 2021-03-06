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

// Forexrates struct for Forexrates
type Forexrates struct {
	// Base currency.
	Base *string `json:"base,omitempty"`
	Quote *map[string]interface{} `json:"quote,omitempty"`
}

// NewForexrates instantiates a new Forexrates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForexrates() *Forexrates {
	this := Forexrates{}
	return &this
}

// NewForexratesWithDefaults instantiates a new Forexrates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForexratesWithDefaults() *Forexrates {
	this := Forexrates{}
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *Forexrates) GetBase() string {
	if o == nil || o.Base == nil {
		var ret string
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Forexrates) GetBaseOk() (*string, bool) {
	if o == nil || o.Base == nil {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *Forexrates) HasBase() bool {
	if o != nil && o.Base != nil {
		return true
	}

	return false
}

// SetBase gets a reference to the given string and assigns it to the Base field.
func (o *Forexrates) SetBase(v string) {
	o.Base = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *Forexrates) GetQuote() map[string]interface{} {
	if o == nil || o.Quote == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Forexrates) GetQuoteOk() (*map[string]interface{}, bool) {
	if o == nil || o.Quote == nil {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *Forexrates) HasQuote() bool {
	if o != nil && o.Quote != nil {
		return true
	}

	return false
}

// SetQuote gets a reference to the given map[string]interface{} and assigns it to the Quote field.
func (o *Forexrates) SetQuote(v map[string]interface{}) {
	o.Quote = &v
}

func (o Forexrates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Base != nil {
		toSerialize["base"] = o.Base
	}
	if o.Quote != nil {
		toSerialize["quote"] = o.Quote
	}
	return json.Marshal(toSerialize)
}

type NullableForexrates struct {
	value *Forexrates
	isSet bool
}

func (v NullableForexrates) Get() *Forexrates {
	return v.value
}

func (v *NullableForexrates) Set(val *Forexrates) {
	v.value = val
	v.isSet = true
}

func (v NullableForexrates) IsSet() bool {
	return v.isSet
}

func (v *NullableForexrates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForexrates(val *Forexrates) *NullableForexrates {
	return &NullableForexrates{value: val, isSet: true}
}

func (v NullableForexrates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForexrates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


