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

// UpgradeDowngrade struct for UpgradeDowngrade
type UpgradeDowngrade struct {
	// Company symbol.
	Symbol *string `json:"symbol,omitempty"`
	// Upgrade/downgrade time in UNIX timestamp.
	GradeTime *int64 `json:"gradeTime,omitempty"`
	// From grade.
	FromGrade *string `json:"fromGrade,omitempty"`
	// To grade.
	ToGrade *string `json:"toGrade,omitempty"`
	// Company/analyst who did the upgrade/downgrade.
	Company *string `json:"company,omitempty"`
	// Action can take any of the following values: <code>up(upgrade), down(downgrade), main(maintains), init(initiate), reit(reiterate)</code>.
	Action *string `json:"action,omitempty"`
}

// NewUpgradeDowngrade instantiates a new UpgradeDowngrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradeDowngrade() *UpgradeDowngrade {
	this := UpgradeDowngrade{}
	return &this
}

// NewUpgradeDowngradeWithDefaults instantiates a new UpgradeDowngrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradeDowngradeWithDefaults() *UpgradeDowngrade {
	this := UpgradeDowngrade{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UpgradeDowngrade) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeDowngrade) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UpgradeDowngrade) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UpgradeDowngrade) SetSymbol(v string) {
	o.Symbol = &v
}

// GetGradeTime returns the GradeTime field value if set, zero value otherwise.
func (o *UpgradeDowngrade) GetGradeTime() int64 {
	if o == nil || o.GradeTime == nil {
		var ret int64
		return ret
	}
	return *o.GradeTime
}

// GetGradeTimeOk returns a tuple with the GradeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeDowngrade) GetGradeTimeOk() (*int64, bool) {
	if o == nil || o.GradeTime == nil {
		return nil, false
	}
	return o.GradeTime, true
}

// HasGradeTime returns a boolean if a field has been set.
func (o *UpgradeDowngrade) HasGradeTime() bool {
	if o != nil && o.GradeTime != nil {
		return true
	}

	return false
}

// SetGradeTime gets a reference to the given int64 and assigns it to the GradeTime field.
func (o *UpgradeDowngrade) SetGradeTime(v int64) {
	o.GradeTime = &v
}

// GetFromGrade returns the FromGrade field value if set, zero value otherwise.
func (o *UpgradeDowngrade) GetFromGrade() string {
	if o == nil || o.FromGrade == nil {
		var ret string
		return ret
	}
	return *o.FromGrade
}

// GetFromGradeOk returns a tuple with the FromGrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeDowngrade) GetFromGradeOk() (*string, bool) {
	if o == nil || o.FromGrade == nil {
		return nil, false
	}
	return o.FromGrade, true
}

// HasFromGrade returns a boolean if a field has been set.
func (o *UpgradeDowngrade) HasFromGrade() bool {
	if o != nil && o.FromGrade != nil {
		return true
	}

	return false
}

// SetFromGrade gets a reference to the given string and assigns it to the FromGrade field.
func (o *UpgradeDowngrade) SetFromGrade(v string) {
	o.FromGrade = &v
}

// GetToGrade returns the ToGrade field value if set, zero value otherwise.
func (o *UpgradeDowngrade) GetToGrade() string {
	if o == nil || o.ToGrade == nil {
		var ret string
		return ret
	}
	return *o.ToGrade
}

// GetToGradeOk returns a tuple with the ToGrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeDowngrade) GetToGradeOk() (*string, bool) {
	if o == nil || o.ToGrade == nil {
		return nil, false
	}
	return o.ToGrade, true
}

// HasToGrade returns a boolean if a field has been set.
func (o *UpgradeDowngrade) HasToGrade() bool {
	if o != nil && o.ToGrade != nil {
		return true
	}

	return false
}

// SetToGrade gets a reference to the given string and assigns it to the ToGrade field.
func (o *UpgradeDowngrade) SetToGrade(v string) {
	o.ToGrade = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *UpgradeDowngrade) GetCompany() string {
	if o == nil || o.Company == nil {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeDowngrade) GetCompanyOk() (*string, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *UpgradeDowngrade) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *UpgradeDowngrade) SetCompany(v string) {
	o.Company = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *UpgradeDowngrade) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeDowngrade) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *UpgradeDowngrade) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *UpgradeDowngrade) SetAction(v string) {
	o.Action = &v
}

func (o UpgradeDowngrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.GradeTime != nil {
		toSerialize["gradeTime"] = o.GradeTime
	}
	if o.FromGrade != nil {
		toSerialize["fromGrade"] = o.FromGrade
	}
	if o.ToGrade != nil {
		toSerialize["toGrade"] = o.ToGrade
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	return json.Marshal(toSerialize)
}

type NullableUpgradeDowngrade struct {
	value *UpgradeDowngrade
	isSet bool
}

func (v NullableUpgradeDowngrade) Get() *UpgradeDowngrade {
	return v.value
}

func (v *NullableUpgradeDowngrade) Set(val *UpgradeDowngrade) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradeDowngrade) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradeDowngrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradeDowngrade(val *UpgradeDowngrade) *NullableUpgradeDowngrade {
	return &NullableUpgradeDowngrade{value: val, isSet: true}
}

func (v NullableUpgradeDowngrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradeDowngrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


