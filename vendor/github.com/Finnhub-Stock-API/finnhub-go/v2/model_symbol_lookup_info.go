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

// SymbolLookupInfo struct for SymbolLookupInfo
type SymbolLookupInfo struct {
	// Symbol description
	Description *string `json:"description,omitempty"`
	// Display symbol name.
	DisplaySymbol *string `json:"displaySymbol,omitempty"`
	// Unique symbol used to identify this symbol used in <code>/stock/candle</code> endpoint.
	Symbol *string `json:"symbol,omitempty"`
	// Security type.
	Type *string `json:"type,omitempty"`
}

// NewSymbolLookupInfo instantiates a new SymbolLookupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolLookupInfo() *SymbolLookupInfo {
	this := SymbolLookupInfo{}
	return &this
}

// NewSymbolLookupInfoWithDefaults instantiates a new SymbolLookupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolLookupInfoWithDefaults() *SymbolLookupInfo {
	this := SymbolLookupInfo{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SymbolLookupInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolLookupInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SymbolLookupInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SymbolLookupInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDisplaySymbol returns the DisplaySymbol field value if set, zero value otherwise.
func (o *SymbolLookupInfo) GetDisplaySymbol() string {
	if o == nil || o.DisplaySymbol == nil {
		var ret string
		return ret
	}
	return *o.DisplaySymbol
}

// GetDisplaySymbolOk returns a tuple with the DisplaySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolLookupInfo) GetDisplaySymbolOk() (*string, bool) {
	if o == nil || o.DisplaySymbol == nil {
		return nil, false
	}
	return o.DisplaySymbol, true
}

// HasDisplaySymbol returns a boolean if a field has been set.
func (o *SymbolLookupInfo) HasDisplaySymbol() bool {
	if o != nil && o.DisplaySymbol != nil {
		return true
	}

	return false
}

// SetDisplaySymbol gets a reference to the given string and assigns it to the DisplaySymbol field.
func (o *SymbolLookupInfo) SetDisplaySymbol(v string) {
	o.DisplaySymbol = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SymbolLookupInfo) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolLookupInfo) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SymbolLookupInfo) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SymbolLookupInfo) SetSymbol(v string) {
	o.Symbol = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SymbolLookupInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolLookupInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SymbolLookupInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SymbolLookupInfo) SetType(v string) {
	o.Type = &v
}

func (o SymbolLookupInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplaySymbol != nil {
		toSerialize["displaySymbol"] = o.DisplaySymbol
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSymbolLookupInfo struct {
	value *SymbolLookupInfo
	isSet bool
}

func (v NullableSymbolLookupInfo) Get() *SymbolLookupInfo {
	return v.value
}

func (v *NullableSymbolLookupInfo) Set(val *SymbolLookupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolLookupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolLookupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolLookupInfo(val *SymbolLookupInfo) *NullableSymbolLookupInfo {
	return &NullableSymbolLookupInfo{value: val, isSet: true}
}

func (v NullableSymbolLookupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolLookupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


