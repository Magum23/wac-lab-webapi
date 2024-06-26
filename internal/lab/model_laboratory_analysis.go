/*
 * Laboratory API
 *
 * Laboratory management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: xmagulam@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package lab

type LaboratoryAnalysis struct {

	// Unique id of the entry in this waiting list
	Id string `json:"id"`

	// Name of patient
	Name string `json:"name"`

	Analyses []string `json:"analyses"`

	Results string `json:"results"`

	// Status of the analysis
	Status string `json:"status"`
}
