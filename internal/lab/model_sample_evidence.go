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

type SampleEvidence struct {

	// Unique id of the entry in this waiting list
	Id string `json:"id"`

	// Name of patient
	Name string `json:"name"`

	// Volume of the sample
	Volume float32 `json:"volume"`

	// Type of the test
	TestType string `json:"testType"`

	// Result of the test
	Result string `json:"result"`

	// Status of the test
	Status string `json:"status"`
}
