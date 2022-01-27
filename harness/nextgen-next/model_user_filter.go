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

type UserFilter struct {
	// This string will be used to filter the results. Details of all the users having this string in their name or email address will be filtered.
	SearchTerm string `json:"searchTerm,omitempty"`
	// Filter by User Identifiers
	Identifiers  []string `json:"identifiers,omitempty"`
	ParentFilter string   `json:"parentFilter,omitempty"`
}