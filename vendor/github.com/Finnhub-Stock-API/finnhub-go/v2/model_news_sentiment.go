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

// NewsSentiment struct for NewsSentiment
type NewsSentiment struct {
	Buzz *CompanyNewsStatistics `json:"buzz,omitempty"`
	// News score.
	CompanyNewsScore *float32 `json:"companyNewsScore,omitempty"`
	// Sector average bullish percent.
	SectorAverageBullishPercent *float32 `json:"sectorAverageBullishPercent,omitempty"`
	// Sectore average score.
	SectorAverageNewsScore *float32 `json:"sectorAverageNewsScore,omitempty"`
	Sentiment *Sentiment `json:"sentiment,omitempty"`
	// Requested symbol.
	Symbol *string `json:"symbol,omitempty"`
}

// NewNewsSentiment instantiates a new NewsSentiment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewsSentiment() *NewsSentiment {
	this := NewsSentiment{}
	return &this
}

// NewNewsSentimentWithDefaults instantiates a new NewsSentiment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewsSentimentWithDefaults() *NewsSentiment {
	this := NewsSentiment{}
	return &this
}

// GetBuzz returns the Buzz field value if set, zero value otherwise.
func (o *NewsSentiment) GetBuzz() CompanyNewsStatistics {
	if o == nil || o.Buzz == nil {
		var ret CompanyNewsStatistics
		return ret
	}
	return *o.Buzz
}

// GetBuzzOk returns a tuple with the Buzz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsSentiment) GetBuzzOk() (*CompanyNewsStatistics, bool) {
	if o == nil || o.Buzz == nil {
		return nil, false
	}
	return o.Buzz, true
}

// HasBuzz returns a boolean if a field has been set.
func (o *NewsSentiment) HasBuzz() bool {
	if o != nil && o.Buzz != nil {
		return true
	}

	return false
}

// SetBuzz gets a reference to the given CompanyNewsStatistics and assigns it to the Buzz field.
func (o *NewsSentiment) SetBuzz(v CompanyNewsStatistics) {
	o.Buzz = &v
}

// GetCompanyNewsScore returns the CompanyNewsScore field value if set, zero value otherwise.
func (o *NewsSentiment) GetCompanyNewsScore() float32 {
	if o == nil || o.CompanyNewsScore == nil {
		var ret float32
		return ret
	}
	return *o.CompanyNewsScore
}

// GetCompanyNewsScoreOk returns a tuple with the CompanyNewsScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsSentiment) GetCompanyNewsScoreOk() (*float32, bool) {
	if o == nil || o.CompanyNewsScore == nil {
		return nil, false
	}
	return o.CompanyNewsScore, true
}

// HasCompanyNewsScore returns a boolean if a field has been set.
func (o *NewsSentiment) HasCompanyNewsScore() bool {
	if o != nil && o.CompanyNewsScore != nil {
		return true
	}

	return false
}

// SetCompanyNewsScore gets a reference to the given float32 and assigns it to the CompanyNewsScore field.
func (o *NewsSentiment) SetCompanyNewsScore(v float32) {
	o.CompanyNewsScore = &v
}

// GetSectorAverageBullishPercent returns the SectorAverageBullishPercent field value if set, zero value otherwise.
func (o *NewsSentiment) GetSectorAverageBullishPercent() float32 {
	if o == nil || o.SectorAverageBullishPercent == nil {
		var ret float32
		return ret
	}
	return *o.SectorAverageBullishPercent
}

// GetSectorAverageBullishPercentOk returns a tuple with the SectorAverageBullishPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsSentiment) GetSectorAverageBullishPercentOk() (*float32, bool) {
	if o == nil || o.SectorAverageBullishPercent == nil {
		return nil, false
	}
	return o.SectorAverageBullishPercent, true
}

// HasSectorAverageBullishPercent returns a boolean if a field has been set.
func (o *NewsSentiment) HasSectorAverageBullishPercent() bool {
	if o != nil && o.SectorAverageBullishPercent != nil {
		return true
	}

	return false
}

// SetSectorAverageBullishPercent gets a reference to the given float32 and assigns it to the SectorAverageBullishPercent field.
func (o *NewsSentiment) SetSectorAverageBullishPercent(v float32) {
	o.SectorAverageBullishPercent = &v
}

// GetSectorAverageNewsScore returns the SectorAverageNewsScore field value if set, zero value otherwise.
func (o *NewsSentiment) GetSectorAverageNewsScore() float32 {
	if o == nil || o.SectorAverageNewsScore == nil {
		var ret float32
		return ret
	}
	return *o.SectorAverageNewsScore
}

// GetSectorAverageNewsScoreOk returns a tuple with the SectorAverageNewsScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsSentiment) GetSectorAverageNewsScoreOk() (*float32, bool) {
	if o == nil || o.SectorAverageNewsScore == nil {
		return nil, false
	}
	return o.SectorAverageNewsScore, true
}

// HasSectorAverageNewsScore returns a boolean if a field has been set.
func (o *NewsSentiment) HasSectorAverageNewsScore() bool {
	if o != nil && o.SectorAverageNewsScore != nil {
		return true
	}

	return false
}

// SetSectorAverageNewsScore gets a reference to the given float32 and assigns it to the SectorAverageNewsScore field.
func (o *NewsSentiment) SetSectorAverageNewsScore(v float32) {
	o.SectorAverageNewsScore = &v
}

// GetSentiment returns the Sentiment field value if set, zero value otherwise.
func (o *NewsSentiment) GetSentiment() Sentiment {
	if o == nil || o.Sentiment == nil {
		var ret Sentiment
		return ret
	}
	return *o.Sentiment
}

// GetSentimentOk returns a tuple with the Sentiment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsSentiment) GetSentimentOk() (*Sentiment, bool) {
	if o == nil || o.Sentiment == nil {
		return nil, false
	}
	return o.Sentiment, true
}

// HasSentiment returns a boolean if a field has been set.
func (o *NewsSentiment) HasSentiment() bool {
	if o != nil && o.Sentiment != nil {
		return true
	}

	return false
}

// SetSentiment gets a reference to the given Sentiment and assigns it to the Sentiment field.
func (o *NewsSentiment) SetSentiment(v Sentiment) {
	o.Sentiment = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *NewsSentiment) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsSentiment) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *NewsSentiment) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *NewsSentiment) SetSymbol(v string) {
	o.Symbol = &v
}

func (o NewsSentiment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Buzz != nil {
		toSerialize["buzz"] = o.Buzz
	}
	if o.CompanyNewsScore != nil {
		toSerialize["companyNewsScore"] = o.CompanyNewsScore
	}
	if o.SectorAverageBullishPercent != nil {
		toSerialize["sectorAverageBullishPercent"] = o.SectorAverageBullishPercent
	}
	if o.SectorAverageNewsScore != nil {
		toSerialize["sectorAverageNewsScore"] = o.SectorAverageNewsScore
	}
	if o.Sentiment != nil {
		toSerialize["sentiment"] = o.Sentiment
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	return json.Marshal(toSerialize)
}

type NullableNewsSentiment struct {
	value *NewsSentiment
	isSet bool
}

func (v NullableNewsSentiment) Get() *NewsSentiment {
	return v.value
}

func (v *NullableNewsSentiment) Set(val *NewsSentiment) {
	v.value = val
	v.isSet = true
}

func (v NullableNewsSentiment) IsSet() bool {
	return v.isSet
}

func (v *NullableNewsSentiment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewsSentiment(val *NewsSentiment) *NullableNewsSentiment {
	return &NullableNewsSentiment{value: val, isSet: true}
}

func (v NullableNewsSentiment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewsSentiment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


