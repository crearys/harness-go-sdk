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

// This contains details of Git Sync Settings
type GitSyncSettings struct {
	// Account Identifier for the Entity
	AccountIdentifier string `json:"accountIdentifier"`
	// Project Identifier for the Entity
	ProjectIdentifier string `json:"projectIdentifier"`
	// Organization Identifier for the Entity
	OrganizationIdentifier string `json:"organizationIdentifier"`
	// Specifies Connectivity Mode for Git Sync. If True, executes through Delegate, else executes through Platform. The default value is True
	ExecuteOnDelegate bool `json:"executeOnDelegate"`
}