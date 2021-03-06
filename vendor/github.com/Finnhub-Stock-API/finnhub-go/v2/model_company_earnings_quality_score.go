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

// CompanyEarningsQualityScore struct for CompanyEarningsQualityScore
type CompanyEarningsQualityScore struct {
	// Symbol
	Symbol *string `json:"symbol,omitempty"`
	// Frequency
	Freq *string `json:"freq,omitempty"`
	// Array of earnings quality score.
	Data *[]CompanyEarningsQualityScoreData `json:"data,omitempty"`
}

// NewCompanyEarningsQualityScore instantiates a new CompanyEarningsQualityScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyEarningsQualityScore() *CompanyEarningsQualityScore {
	this := CompanyEarningsQualityScore{}
	return &this
}

// NewCompanyEarningsQualityScoreWithDefaults instantiates a new CompanyEarningsQualityScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyEarningsQualityScoreWithDefaults() *CompanyEarningsQualityScore {
	this := CompanyEarningsQualityScore{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CompanyEarningsQualityScore) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyEarningsQualityScore) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CompanyEarningsQualityScore) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CompanyEarningsQualityScore) SetSymbol(v string) {
	o.Symbol = &v
}

// GetFreq returns the Freq field value if set, zero value otherwise.
func (o *CompanyEarningsQualityScore) GetFreq() string {
	if o == nil || o.Freq == nil {
		var ret string
		return ret
	}
	return *o.Freq
}

// GetFreqOk returns a tuple with the Freq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyEarningsQualityScore) GetFreqOk() (*string, bool) {
	if o == nil || o.Freq == nil {
		return nil, false
	}
	return o.Freq, true
}

// HasFreq returns a boolean if a field has been set.
func (o *CompanyEarningsQualityScore) HasFreq() bool {
	if o != nil && o.Freq != nil {
		return true
	}

	return false
}

// SetFreq gets a reference to the given string and assigns it to the Freq field.
func (o *CompanyEarningsQualityScore) SetFreq(v string) {
	o.Freq = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CompanyEarningsQualityScore) GetData() []CompanyEarningsQualityScoreData {
	if o == nil || o.Data == nil {
		var ret []CompanyEarningsQualityScoreData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyEarningsQualityScore) GetDataOk() (*[]CompanyEarningsQualityScoreData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CompanyEarningsQualityScore) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []CompanyEarningsQualityScoreData and assigns it to the Data field.
func (o *CompanyEarningsQualityScore) SetData(v []CompanyEarningsQualityScoreData) {
	o.Data = &v
}

func (o CompanyEarningsQualityScore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Freq != nil {
		toSerialize["freq"] = o.Freq
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyEarningsQualityScore struct {
	value *CompanyEarningsQualityScore
	isSet bool
}

func (v NullableCompanyEarningsQualityScore) Get() *CompanyEarningsQualityScore {
	return v.value
}

func (v *NullableCompanyEarningsQualityScore) Set(val *CompanyEarningsQualityScore) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyEarningsQualityScore) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyEarningsQualityScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyEarningsQualityScore(val *CompanyEarningsQualityScore) *NullableCompanyEarningsQualityScore {
	return &NullableCompanyEarningsQualityScore{value: val, isSet: true}
}

func (v NullableCompanyEarningsQualityScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyEarningsQualityScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


