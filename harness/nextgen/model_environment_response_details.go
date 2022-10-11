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

// This is the Environment entity defined in Harness
type EnvironmentResponseDetails struct {
	AccountId         string            `json:"accountId,omitempty"`
	OrgIdentifier     string            `json:"orgIdentifier,omitempty"`
	ProjectIdentifier string            `json:"projectIdentifier,omitempty"`
	Identifier        string            `json:"identifier,omitempty"`
	Name              string            `json:"name,omitempty"`
	Description       string            `json:"description,omitempty"`
	Color             string            `json:"color,omitempty"`
	Type_             EnvironmentType   `json:"type,omitempty"`
	Deleted           bool              `json:"deleted,omitempty"`
	Tags              map[string]string `json:"tags,omitempty"`
	Version           int64             `json:"version,omitempty"`
	Yaml              string            `json:"yaml,omitempty"`
}
