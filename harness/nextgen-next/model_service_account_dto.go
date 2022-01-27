/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type ServiceAccountDto struct {
	Identifier        string            `json:"identifier,omitempty"`
	Name              string            `json:"name,omitempty"`
	Email             string            `json:"email,omitempty"`
	Description       string            `json:"description,omitempty"`
	Tags              map[string]string `json:"tags,omitempty"`
	AccountIdentifier string            `json:"accountIdentifier,omitempty"`
	OrgIdentifier     string            `json:"orgIdentifier,omitempty"`
	ProjectIdentifier string            `json:"projectIdentifier,omitempty"`
}