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

// FinancialsAsReported struct for FinancialsAsReported
type FinancialsAsReported struct {
	// Symbol
	Symbol *string `json:"symbol,omitempty"`
	// CIK
	Cik *string `json:"cik,omitempty"`
	// Array of filings.
	Data *[]Report `json:"data,omitempty"`
}

// NewFinancialsAsReported instantiates a new FinancialsAsReported object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialsAsReported() *FinancialsAsReported {
	this := FinancialsAsReported{}
	return &this
}

// NewFinancialsAsReportedWithDefaults instantiates a new FinancialsAsReported object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialsAsReportedWithDefaults() *FinancialsAsReported {
	this := FinancialsAsReported{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *FinancialsAsReported) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialsAsReported) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *FinancialsAsReported) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *FinancialsAsReported) SetSymbol(v string) {
	o.Symbol = &v
}

// GetCik returns the Cik field value if set, zero value otherwise.
func (o *FinancialsAsReported) GetCik() string {
	if o == nil || o.Cik == nil {
		var ret string
		return ret
	}
	return *o.Cik
}

// GetCikOk returns a tuple with the Cik field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialsAsReported) GetCikOk() (*string, bool) {
	if o == nil || o.Cik == nil {
		return nil, false
	}
	return o.Cik, true
}

// HasCik returns a boolean if a field has been set.
func (o *FinancialsAsReported) HasCik() bool {
	if o != nil && o.Cik != nil {
		return true
	}

	return false
}

// SetCik gets a reference to the given string and assigns it to the Cik field.
func (o *FinancialsAsReported) SetCik(v string) {
	o.Cik = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *FinancialsAsReported) GetData() []Report {
	if o == nil || o.Data == nil {
		var ret []Report
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialsAsReported) GetDataOk() (*[]Report, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FinancialsAsReported) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Report and assigns it to the Data field.
func (o *FinancialsAsReported) SetData(v []Report) {
	o.Data = &v
}

func (o FinancialsAsReported) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Cik != nil {
		toSerialize["cik"] = o.Cik
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableFinancialsAsReported struct {
	value *FinancialsAsReported
	isSet bool
}

func (v NullableFinancialsAsReported) Get() *FinancialsAsReported {
	return v.value
}

func (v *NullableFinancialsAsReported) Set(val *FinancialsAsReported) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialsAsReported) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialsAsReported) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialsAsReported(val *FinancialsAsReported) *NullableFinancialsAsReported {
	return &NullableFinancialsAsReported{value: val, isSet: true}
}

func (v NullableFinancialsAsReported) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialsAsReported) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


