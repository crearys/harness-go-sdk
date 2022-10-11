/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type Permission struct {
	Identifier         string   `json:"identifier,omitempty"`
	Name               string   `json:"name,omitempty"`
	Status             string   `json:"status,omitempty"`
	IncludeInAllRoles  bool     `json:"includeInAllRoles,omitempty"`
	AllowedScopeLevels []string `json:"allowedScopeLevels,omitempty"`
	ResourceType       string   `json:"resourceType,omitempty"`
	Action             string   `json:"action,omitempty"`
}
