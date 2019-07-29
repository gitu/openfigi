/*
 * OpenFIGI API
 *
 * A free & open API for FIGI discovery.
 *
 * API version: 1.3.0
 * Contact: support@openfigi.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openfigi

type FigiResult struct {
	Figi string `json:"figi,omitempty"`
	SecurityType *string `json:"securityType,omitempty"`
	MarketSector *string `json:"marketSector,omitempty"`
	Ticker *string `json:"ticker,omitempty"`
	Name *string `json:"name,omitempty"`
	UniqueID *string `json:"uniqueID,omitempty"`
	ExchCode *string `json:"exchCode,omitempty"`
	ShareClassFIGI *string `json:"shareClassFIGI,omitempty"`
	CompositeFIGI *string `json:"compositeFIGI,omitempty"`
	SecurityType2 *string `json:"securityType2,omitempty"`
	SecurityDescription *string `json:"securityDescription,omitempty"`
	UniqueIDFutOpt *string `json:"uniqueIDFutOpt,omitempty"`
	// Exists when API is unable to show non-FIGI fields.
	Metadata *string `json:"metadata,omitempty"`
}
