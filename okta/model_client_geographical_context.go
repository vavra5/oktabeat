/*
 * System Log (Okta API)
 *
 * The Okta System Log API (https://developer.okta.com/docs/api/resources/system_log) provides read access to your organization's system log. The API is intended to export event data as a batch job from your organization to another system for reporting or analysis.
 *
 * API version: 1.0
 * Contact: rnd@forter.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oktapi

type ClientGeographicalContext struct {
	Geolocation ClientGeographicalContextGeolocation `json:"geolocation,omitempty"`
	City        string                               `json:"city,omitempty"`
	State       string                               `json:"state,omitempty"`
	Country     string                               `json:"country,omitempty"`
	PostalCode  string                               `json:"postalCode,omitempty"`
}
