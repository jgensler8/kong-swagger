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

type Certificate struct {

	Id string `json:"id,omitempty"`

	Cert string `json:"cert,omitempty"`

	Key string `json:"key,omitempty"`

	Snis []string `json:"snis,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`
}