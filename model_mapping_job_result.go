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

// MappingJobResult struct for MappingJobResult
type MappingJobResult struct {
	Data  []FigiResult `json:"data,omitempty"`
	Error string       `json:"error,omitempty"`
}
