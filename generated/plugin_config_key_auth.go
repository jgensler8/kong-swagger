/* 
 * Kong-Admin
 *
 * A Swagger definition for the Kong Admin API
 *
 * OpenAPI spec version: 1.1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type PluginConfigKeyAuth struct {

	KeyNames string `json:"key_names,omitempty"`

	KeyInBody bool `json:"key_in_body,omitempty"`

	HideCredentials bool `json:"hide_credentials,omitempty"`

	Anonymous string `json:"anonymous,omitempty"`

	RunOnPreflight bool `json:"run_on_preflight,omitempty"`
}