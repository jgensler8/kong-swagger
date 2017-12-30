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

// Someone that interacts with the API
type Consumer struct {

	Id string `json:"id,omitempty"`

	CustomId string `json:"custom_id,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`
}
