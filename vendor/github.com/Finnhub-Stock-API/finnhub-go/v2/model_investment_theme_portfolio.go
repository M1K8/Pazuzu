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

// InvestmentThemePortfolio struct for InvestmentThemePortfolio
type InvestmentThemePortfolio struct {
	// Symbol
	Symbol *string `json:"symbol,omitempty"`
}

// NewInvestmentThemePortfolio instantiates a new InvestmentThemePortfolio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentThemePortfolio() *InvestmentThemePortfolio {
	this := InvestmentThemePortfolio{}
	return &this
}

// NewInvestmentThemePortfolioWithDefaults instantiates a new InvestmentThemePortfolio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentThemePortfolioWithDefaults() *InvestmentThemePortfolio {
	this := InvestmentThemePortfolio{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *InvestmentThemePortfolio) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentThemePortfolio) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *InvestmentThemePortfolio) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *InvestmentThemePortfolio) SetSymbol(v string) {
	o.Symbol = &v
}

func (o InvestmentThemePortfolio) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	return json.Marshal(toSerialize)
}

type NullableInvestmentThemePortfolio struct {
	value *InvestmentThemePortfolio
	isSet bool
}

func (v NullableInvestmentThemePortfolio) Get() *InvestmentThemePortfolio {
	return v.value
}

func (v *NullableInvestmentThemePortfolio) Set(val *InvestmentThemePortfolio) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentThemePortfolio) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentThemePortfolio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentThemePortfolio(val *InvestmentThemePortfolio) *NullableInvestmentThemePortfolio {
	return &NullableInvestmentThemePortfolio{value: val, isSet: true}
}

func (v NullableInvestmentThemePortfolio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentThemePortfolio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


