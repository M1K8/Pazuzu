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

// FundOwnershipInfo struct for FundOwnershipInfo
type FundOwnershipInfo struct {
	// Investor's name.
	Name *string `json:"name,omitempty"`
	// Number of shares held by the investor.
	Share *int64 `json:"share,omitempty"`
	// Number of share changed (net buy or sell) from the last period.
	Change *int64 `json:"change,omitempty"`
	// Filing date.
	FilingDate *string `json:"filingDate,omitempty"`
	// Percent of the fund's portfolio comprised of the company's share.
	PortfolioPercent *float32 `json:"portfolioPercent,omitempty"`
}

// NewFundOwnershipInfo instantiates a new FundOwnershipInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundOwnershipInfo() *FundOwnershipInfo {
	this := FundOwnershipInfo{}
	return &this
}

// NewFundOwnershipInfoWithDefaults instantiates a new FundOwnershipInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundOwnershipInfoWithDefaults() *FundOwnershipInfo {
	this := FundOwnershipInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FundOwnershipInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOwnershipInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FundOwnershipInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FundOwnershipInfo) SetName(v string) {
	o.Name = &v
}

// GetShare returns the Share field value if set, zero value otherwise.
func (o *FundOwnershipInfo) GetShare() int64 {
	if o == nil || o.Share == nil {
		var ret int64
		return ret
	}
	return *o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOwnershipInfo) GetShareOk() (*int64, bool) {
	if o == nil || o.Share == nil {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *FundOwnershipInfo) HasShare() bool {
	if o != nil && o.Share != nil {
		return true
	}

	return false
}

// SetShare gets a reference to the given int64 and assigns it to the Share field.
func (o *FundOwnershipInfo) SetShare(v int64) {
	o.Share = &v
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *FundOwnershipInfo) GetChange() int64 {
	if o == nil || o.Change == nil {
		var ret int64
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOwnershipInfo) GetChangeOk() (*int64, bool) {
	if o == nil || o.Change == nil {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *FundOwnershipInfo) HasChange() bool {
	if o != nil && o.Change != nil {
		return true
	}

	return false
}

// SetChange gets a reference to the given int64 and assigns it to the Change field.
func (o *FundOwnershipInfo) SetChange(v int64) {
	o.Change = &v
}

// GetFilingDate returns the FilingDate field value if set, zero value otherwise.
func (o *FundOwnershipInfo) GetFilingDate() string {
	if o == nil || o.FilingDate == nil {
		var ret string
		return ret
	}
	return *o.FilingDate
}

// GetFilingDateOk returns a tuple with the FilingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOwnershipInfo) GetFilingDateOk() (*string, bool) {
	if o == nil || o.FilingDate == nil {
		return nil, false
	}
	return o.FilingDate, true
}

// HasFilingDate returns a boolean if a field has been set.
func (o *FundOwnershipInfo) HasFilingDate() bool {
	if o != nil && o.FilingDate != nil {
		return true
	}

	return false
}

// SetFilingDate gets a reference to the given string and assigns it to the FilingDate field.
func (o *FundOwnershipInfo) SetFilingDate(v string) {
	o.FilingDate = &v
}

// GetPortfolioPercent returns the PortfolioPercent field value if set, zero value otherwise.
func (o *FundOwnershipInfo) GetPortfolioPercent() float32 {
	if o == nil || o.PortfolioPercent == nil {
		var ret float32
		return ret
	}
	return *o.PortfolioPercent
}

// GetPortfolioPercentOk returns a tuple with the PortfolioPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOwnershipInfo) GetPortfolioPercentOk() (*float32, bool) {
	if o == nil || o.PortfolioPercent == nil {
		return nil, false
	}
	return o.PortfolioPercent, true
}

// HasPortfolioPercent returns a boolean if a field has been set.
func (o *FundOwnershipInfo) HasPortfolioPercent() bool {
	if o != nil && o.PortfolioPercent != nil {
		return true
	}

	return false
}

// SetPortfolioPercent gets a reference to the given float32 and assigns it to the PortfolioPercent field.
func (o *FundOwnershipInfo) SetPortfolioPercent(v float32) {
	o.PortfolioPercent = &v
}

func (o FundOwnershipInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Share != nil {
		toSerialize["share"] = o.Share
	}
	if o.Change != nil {
		toSerialize["change"] = o.Change
	}
	if o.FilingDate != nil {
		toSerialize["filingDate"] = o.FilingDate
	}
	if o.PortfolioPercent != nil {
		toSerialize["portfolioPercent"] = o.PortfolioPercent
	}
	return json.Marshal(toSerialize)
}

type NullableFundOwnershipInfo struct {
	value *FundOwnershipInfo
	isSet bool
}

func (v NullableFundOwnershipInfo) Get() *FundOwnershipInfo {
	return v.value
}

func (v *NullableFundOwnershipInfo) Set(val *FundOwnershipInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFundOwnershipInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFundOwnershipInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundOwnershipInfo(val *FundOwnershipInfo) *NullableFundOwnershipInfo {
	return &NullableFundOwnershipInfo{value: val, isSet: true}
}

func (v NullableFundOwnershipInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundOwnershipInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


