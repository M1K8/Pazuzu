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

// CompanyNewsStatistics struct for CompanyNewsStatistics
type CompanyNewsStatistics struct {
	// 
	ArticlesInLastWeek *int64 `json:"articlesInLastWeek,omitempty"`
	// 
	Buzz *float32 `json:"buzz,omitempty"`
	// 
	WeeklyAverage *float32 `json:"weeklyAverage,omitempty"`
}

// NewCompanyNewsStatistics instantiates a new CompanyNewsStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyNewsStatistics() *CompanyNewsStatistics {
	this := CompanyNewsStatistics{}
	return &this
}

// NewCompanyNewsStatisticsWithDefaults instantiates a new CompanyNewsStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyNewsStatisticsWithDefaults() *CompanyNewsStatistics {
	this := CompanyNewsStatistics{}
	return &this
}

// GetArticlesInLastWeek returns the ArticlesInLastWeek field value if set, zero value otherwise.
func (o *CompanyNewsStatistics) GetArticlesInLastWeek() int64 {
	if o == nil || o.ArticlesInLastWeek == nil {
		var ret int64
		return ret
	}
	return *o.ArticlesInLastWeek
}

// GetArticlesInLastWeekOk returns a tuple with the ArticlesInLastWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyNewsStatistics) GetArticlesInLastWeekOk() (*int64, bool) {
	if o == nil || o.ArticlesInLastWeek == nil {
		return nil, false
	}
	return o.ArticlesInLastWeek, true
}

// HasArticlesInLastWeek returns a boolean if a field has been set.
func (o *CompanyNewsStatistics) HasArticlesInLastWeek() bool {
	if o != nil && o.ArticlesInLastWeek != nil {
		return true
	}

	return false
}

// SetArticlesInLastWeek gets a reference to the given int64 and assigns it to the ArticlesInLastWeek field.
func (o *CompanyNewsStatistics) SetArticlesInLastWeek(v int64) {
	o.ArticlesInLastWeek = &v
}

// GetBuzz returns the Buzz field value if set, zero value otherwise.
func (o *CompanyNewsStatistics) GetBuzz() float32 {
	if o == nil || o.Buzz == nil {
		var ret float32
		return ret
	}
	return *o.Buzz
}

// GetBuzzOk returns a tuple with the Buzz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyNewsStatistics) GetBuzzOk() (*float32, bool) {
	if o == nil || o.Buzz == nil {
		return nil, false
	}
	return o.Buzz, true
}

// HasBuzz returns a boolean if a field has been set.
func (o *CompanyNewsStatistics) HasBuzz() bool {
	if o != nil && o.Buzz != nil {
		return true
	}

	return false
}

// SetBuzz gets a reference to the given float32 and assigns it to the Buzz field.
func (o *CompanyNewsStatistics) SetBuzz(v float32) {
	o.Buzz = &v
}

// GetWeeklyAverage returns the WeeklyAverage field value if set, zero value otherwise.
func (o *CompanyNewsStatistics) GetWeeklyAverage() float32 {
	if o == nil || o.WeeklyAverage == nil {
		var ret float32
		return ret
	}
	return *o.WeeklyAverage
}

// GetWeeklyAverageOk returns a tuple with the WeeklyAverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyNewsStatistics) GetWeeklyAverageOk() (*float32, bool) {
	if o == nil || o.WeeklyAverage == nil {
		return nil, false
	}
	return o.WeeklyAverage, true
}

// HasWeeklyAverage returns a boolean if a field has been set.
func (o *CompanyNewsStatistics) HasWeeklyAverage() bool {
	if o != nil && o.WeeklyAverage != nil {
		return true
	}

	return false
}

// SetWeeklyAverage gets a reference to the given float32 and assigns it to the WeeklyAverage field.
func (o *CompanyNewsStatistics) SetWeeklyAverage(v float32) {
	o.WeeklyAverage = &v
}

func (o CompanyNewsStatistics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArticlesInLastWeek != nil {
		toSerialize["articlesInLastWeek"] = o.ArticlesInLastWeek
	}
	if o.Buzz != nil {
		toSerialize["buzz"] = o.Buzz
	}
	if o.WeeklyAverage != nil {
		toSerialize["weeklyAverage"] = o.WeeklyAverage
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyNewsStatistics struct {
	value *CompanyNewsStatistics
	isSet bool
}

func (v NullableCompanyNewsStatistics) Get() *CompanyNewsStatistics {
	return v.value
}

func (v *NullableCompanyNewsStatistics) Set(val *CompanyNewsStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyNewsStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyNewsStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyNewsStatistics(val *CompanyNewsStatistics) *NullableCompanyNewsStatistics {
	return &NullableCompanyNewsStatistics{value: val, isSet: true}
}

func (v NullableCompanyNewsStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyNewsStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


