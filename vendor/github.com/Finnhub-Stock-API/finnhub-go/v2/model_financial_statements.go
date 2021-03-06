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

// FinancialStatements struct for FinancialStatements
type FinancialStatements struct {
	// Symbol of the company.
	Symbol *string `json:"symbol,omitempty"`
	// An array of map of key, value pairs containing the data for each period.
	Financials *[]map[string]interface{} `json:"financials,omitempty"`
}

// NewFinancialStatements instantiates a new FinancialStatements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialStatements() *FinancialStatements {
	this := FinancialStatements{}
	return &this
}

// NewFinancialStatementsWithDefaults instantiates a new FinancialStatements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialStatementsWithDefaults() *FinancialStatements {
	this := FinancialStatements{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *FinancialStatements) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialStatements) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *FinancialStatements) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *FinancialStatements) SetSymbol(v string) {
	o.Symbol = &v
}

// GetFinancials returns the Financials field value if set, zero value otherwise.
func (o *FinancialStatements) GetFinancials() []map[string]interface{} {
	if o == nil || o.Financials == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Financials
}

// GetFinancialsOk returns a tuple with the Financials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialStatements) GetFinancialsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Financials == nil {
		return nil, false
	}
	return o.Financials, true
}

// HasFinancials returns a boolean if a field has been set.
func (o *FinancialStatements) HasFinancials() bool {
	if o != nil && o.Financials != nil {
		return true
	}

	return false
}

// SetFinancials gets a reference to the given []map[string]interface{} and assigns it to the Financials field.
func (o *FinancialStatements) SetFinancials(v []map[string]interface{}) {
	o.Financials = &v
}

func (o FinancialStatements) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Financials != nil {
		toSerialize["financials"] = o.Financials
	}
	return json.Marshal(toSerialize)
}

type NullableFinancialStatements struct {
	value *FinancialStatements
	isSet bool
}

func (v NullableFinancialStatements) Get() *FinancialStatements {
	return v.value
}

func (v *NullableFinancialStatements) Set(val *FinancialStatements) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialStatements) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialStatements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialStatements(val *FinancialStatements) *NullableFinancialStatements {
	return &NullableFinancialStatements{value: val, isSet: true}
}

func (v NullableFinancialStatements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialStatements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


