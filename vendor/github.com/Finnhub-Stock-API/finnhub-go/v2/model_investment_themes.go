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

// InvestmentThemes struct for InvestmentThemes
type InvestmentThemes struct {
	// Investment theme
	Theme *string `json:"theme,omitempty"`
	// Investment theme portfolio.
	Data *[]InvestmentThemePortfolio `json:"data,omitempty"`
}

// NewInvestmentThemes instantiates a new InvestmentThemes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentThemes() *InvestmentThemes {
	this := InvestmentThemes{}
	return &this
}

// NewInvestmentThemesWithDefaults instantiates a new InvestmentThemes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentThemesWithDefaults() *InvestmentThemes {
	this := InvestmentThemes{}
	return &this
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *InvestmentThemes) GetTheme() string {
	if o == nil || o.Theme == nil {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentThemes) GetThemeOk() (*string, bool) {
	if o == nil || o.Theme == nil {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *InvestmentThemes) HasTheme() bool {
	if o != nil && o.Theme != nil {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *InvestmentThemes) SetTheme(v string) {
	o.Theme = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InvestmentThemes) GetData() []InvestmentThemePortfolio {
	if o == nil || o.Data == nil {
		var ret []InvestmentThemePortfolio
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentThemes) GetDataOk() (*[]InvestmentThemePortfolio, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InvestmentThemes) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []InvestmentThemePortfolio and assigns it to the Data field.
func (o *InvestmentThemes) SetData(v []InvestmentThemePortfolio) {
	o.Data = &v
}

func (o InvestmentThemes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Theme != nil {
		toSerialize["theme"] = o.Theme
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInvestmentThemes struct {
	value *InvestmentThemes
	isSet bool
}

func (v NullableInvestmentThemes) Get() *InvestmentThemes {
	return v.value
}

func (v *NullableInvestmentThemes) Set(val *InvestmentThemes) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentThemes) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentThemes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentThemes(val *InvestmentThemes) *NullableInvestmentThemes {
	return &NullableInvestmentThemes{value: val, isSet: true}
}

func (v NullableInvestmentThemes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentThemes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


