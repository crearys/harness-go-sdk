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

type ServiceDiagnostics struct {
	// Error message if success were to be false
	Message string `json:"message,omitempty"`
	// Name of the check
	Name string `json:"name,omitempty"`
	// Flag which specifies if a diagnostic check is successful
	Success bool   `json:"success,omitempty"`
	Type_   string `json:"type,omitempty"`
}
