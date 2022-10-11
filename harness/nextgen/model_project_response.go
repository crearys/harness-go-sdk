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

// This has Project details along with its metadata as defined in Harness .
type ProjectResponse struct {
	Project *Project `json:"project"`
	// This specifies the time at which project was created.
	CreatedAt int64 `json:"createdAt,omitempty"`
	// This specifies the time at which project was last modified.
	LastModifiedAt int64 `json:"lastModifiedAt,omitempty"`
}
