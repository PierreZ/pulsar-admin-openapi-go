/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type BacklogQuota struct {
	Limit int64 `json:"limit,omitempty"`
	Policy string `json:"policy,omitempty"`
}